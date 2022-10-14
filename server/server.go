package server

import (
	"fmt"
	"strconv"

	"github.com/fahad-md-kamal/go-bitly/models"
	"github.com/fahad-md-kamal/go-bitly/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func redirect(c *fiber.Ctx) error {
	golyUrl := c.Params("redirect")
	goly, err := models.FindByGolyUrl(golyUrl)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error getting all goly links" + err.Error(),
		})
	}
	// grab any stats you want ...
	goly.Clicked += 1
	err = models.UpdateGoly(goly)

	if err != nil {
		fmt.Printf("error updating: %v\n", err)
	}
	return c.Redirect(goly.Redirect, fiber.StatusTemporaryRedirect)
}

func getAllGolies(c *fiber.Ctx) error {
	golies, err := models.GetAllGolies()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error getting all goly links" + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(golies)
}

func getGoly(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error could not parse id" + err.Error(),
		})
	}

	goly, err := models.GetGoly(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error could not retrieve goly from db" + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(goly)
}

func createGoly(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var goly models.Goly
	err := c.BodyParser(&goly)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error parsing JSON" + err.Error(),
		})
	}

	if goly.Random {
		goly.Goly = utils.RandomURL(8)
	}

	err = models.CreateGoly(goly)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error parsing JSON" + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(goly)
}

func updateGoly(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var goly models.Goly

	err := c.BodyParser(&goly)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not parse json " + err.Error(),
		})
	}

	err = models.UpdateGoly(goly)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not update goly link in DB " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(goly)
}

func deleteGoly(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not parsing JSON" + err.Error(),
		})
	}

	err = models.DeleteGoly(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not parsing JSON" + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "goly deleted.",
	})
}

func SetupAndListen() {

	router := fiber.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	router.Get("/r/:redirect", redirect)

	router.Post("/goly", createGoly)
	router.Patch("/goly", updateGoly)
	router.Get("/goly", getAllGolies)
	router.Get("/goly/:id", getGoly)
	router.Delete("/goly/:id", deleteGoly)

	router.Listen(":3000")
}
