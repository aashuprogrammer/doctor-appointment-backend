package users

import (
	"log"
	"mylearning/db/pgdb"

	"github.com/gofiber/fiber/v2"
)

type DoctorCategoryRequeset struct {
	CategoryName string `json:"categoryName"`
}

type DoctorCategoryResponse struct {
	Msg  string        `json:"msg"`
	User pgdb.Category `json:"user"`
}

func DoctorCategory(o pgdb.Store) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var req DoctorCategoryRequeset
		err := c.BodyParser(&req)
		if err != nil {
			log.Fatal("Unable to doctor category", err)
		}
		user, err := o.CreateCategory(c.Context(), req.CategoryName)
		if err != nil {
			return err
		}
		return c.JSON(DoctorCategoryResponse{Msg: "Doctor Category Create", User: user})
	}
}
