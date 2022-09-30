package models

import (
	"errors"
	"html"
	"strings"

	"github.com/fahad-md-kamal/go-jwt/initializers"
	"github.com/fahad-md-kamal/go-jwt/utils/token"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"passwords"`
}

func GetUserByID(uid uint) (User, error) {
	var u User
	if err := initializers.DB.First(&u, uid).Error; err != nil {
		return u, errors.New("User not found!")
	}

	u.PrepareGive()

	return u, nil
}

func (u *User) PrepareGive() {
	u.Password = ""
}


func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(username string, password string) (string, error)  {
	var err error

	user := User{}

	err = initializers.DB.Model(User{}).Where("username = ?", username).Take(&user).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, user.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword{
		return "", err
	}

	token, err := token.GenerateToken(user.ID)

	if err != nil {
		return "", err
	}


	return token, nil
}

func (user *User) SaveUser() (*User, error){
	var err error
	user.BeforeSave()
	err = initializers.DB.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func (user *User) BeforeSave() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	//remove spaces in username 
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))

	return nil
}