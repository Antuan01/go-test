package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/Antuan01/go-test/database"
	"github.com/Antuan01/go-test/models"
)

func CreatePostReport(c *fiber.Ctx) error {
	var post models.Post

	report := new(models.PostReport)

	c.BodyParser(report)
	
	fmt.Println(report.Reason)
	fmt.Println(report.PostId)

	if len(report.Reason) < 10 {
		c.JSON(fiber.Map{"error":"reason must be at least 10 characters long"})
		return c.SendStatus(422)
	}

	db := database.Database

	err := db.First(&post, report.PostId).Error

	if err != nil {
		fmt.Println(err)
		c.JSON(fiber.Map{"error":"post not found"})
		return c.SendStatus(422)
	}
	
	db.Create(&report)

	c.JSON(report)
	return c.SendStatus(201)
}

func CreateCommentReport(c *fiber.Ctx) error {
	var comment models.Comment

	report := new(models.CommentReport)

	c.BodyParser(report)
	
	fmt.Println(report.Reason)
	fmt.Println(report.CommentId)

	if len(report.Reason) < 10 {
		c.JSON(fiber.Map{"error":"reason must be at least 10 characters long"})
		return c.SendStatus(422)
	}

	db := database.Database

	err := db.First(&comment, report.CommentId).Error

	if err != nil {
		fmt.Println(err)
		c.JSON(fiber.Map{"error":"comment not found"})
		return c.SendStatus(422)
	}
	
	db.Create(&report)

	c.JSON(report)
	return c.SendStatus(201)
}
