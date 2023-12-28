package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

func createBook(db *gorm.DB, c *fiber.Ctx) error {
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	result := db.Create(book)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(result.Error.Error())
	}
	return c.JSON(book)
}

func getBook(db *gorm.DB, c *fiber.Ctx) error {
	id := c.Params("id")
	var book Book
	result := db.First(&book, id)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(result.Error.Error())
	}
	return c.JSON(book)
}

func getBooks(db *gorm.DB, c *fiber.Ctx) error {
	var books []Book
	result := db.Find(&books)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(result.Error.Error())
	}
	return c.JSON(books)
}

func updateBook(db *gorm.DB, c *fiber.Ctx) error {
	id := c.Params("id")
	book := new(Book)
	db.First(&book, id)
	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	db.Save(&book)
	return c.Status(fiber.StatusOK).SendString("Book successfully updated!")
}

func deleteBook(db *gorm.DB, c *fiber.Ctx) error {
	id := c.Params("id")
	result := db.Delete(&Book{}, id)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(result.Error.Error())
	}
	return c.Status(fiber.StatusOK).SendString(fmt.Sprintf("Book with id %s successfully deleted!", id))
}
