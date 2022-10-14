package models

func GetAllGolies() ([]Goly, error){
	var golies []Goly

	tx := DB.Find(&golies)

	if tx.Error != nil {
		return []Goly{}, tx.Error
	}

	return golies, nil
}


func GetGoly(id uint64) (Goly, error) {
	var goly Goly
	txt := DB.Where("id = ?", id).First(&goly)
	if txt.Error != nil {
		return Goly{}, txt.Error
	}
	return goly, nil
}

func CreateGoly(goly Goly) error{
	tx := DB.Create(&goly)
	return tx.Error
}

func UpdateGoly(goly Goly) error {
	tx := DB.Save(&goly)
	return tx.Error
}

func DeleteGoly(id uint64) error {
	txt := DB.Unscoped().Delete(&Goly{}, id)
	return txt.Error
}

func FindByGolyUrl(url string) (Goly, error) {
	var goly Goly
	tx := DB.Where("goly = ?", url).First(&goly)
	return goly, tx.Error
}