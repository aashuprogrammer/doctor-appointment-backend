package users

import (
	"log"
	"mylearning/db/pgdb"
	"mylearning/utils"

	"github.com/gofiber/fiber/v2"
)

type CreateDoctorRequeset struct {
	CategoryId int64  `json:"category_id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Password   string `json:"password"`
	Gender     string `json:"gender"`
	Dob        string `json:"dob"`
	Img        string `json:"img"`
	Shedule    string `json:"shedule"`
	Degree     string `json:"degree"`
	Address    string `json:"address"`
	Roles      string `json:"roles"`
}
type CreateDoctorResponse struct {
	Msg  string      `json:"msg"`
	User pgdb.Doctor `json:"user"`
}

func CreateDoctor(d pgdb.Store) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var req CreateDoctorRequeset
		err := c.BodyParser(&req)
		if err != nil {
			log.Fatal("Unable to Doctor parse", err)
		}
		HashPassword, err := utils.HashPassword(req.Password)
		if err != nil {
			return err
		}
		user, err := d.CreateDoctor(c.Context(), pgdb.CreateDoctorParams{
			CategoryID: req.CategoryId,
			Name:       req.Name,
			Email:      req.Email,
			Phone:      req.Phone,
			Password:   HashPassword,
			Gender:     pgdb.GenderT(req.Gender),
			Dob:        req.Dob,
			Shedule:    req.Shedule,
			Degree:     req.Degree,
			Address:    req.Address,
			Roles:      pgdb.RolesT(req.Roles),
		})
		if err != nil {
			return err
		}
		return c.JSON(CreateDoctorResponse{Msg: "User Created", User: user})
	}
}
