package main

import (
	"context"
	"log"
	"mylearning/db/pgdb"
	"mylearning/token"
	"mylearning/users"
	"mylearning/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("unable to load config", err)
		return
	}
	pgxConfig, err := pgxpool.ParseConfig(config.DbUrl)
	if err != nil {
		log.Fatal("unable to parse", err)
		return
	}
	tokenMaker, err := token.NewPastroMaker(config.SecretKey)
	if err != nil {
		log.Fatal("failed to create token", err)
	}

	conn, err := pgxpool.NewWithConfig(context.Background(), pgxConfig)
	if err != nil {
		log.Fatal("unable to create connection pool", err)
	}
	store := pgdb.NewStore(conn)
	defer conn.Close()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello Users ")
	})
	// //////////////  User Create //////////////
	app.Post("/user/signup", users.CreateUser(store))
	app.Get("/user/get_all_users", users.AllGetUser(store))
	app.Post("/user/login", users.UserLogin(store, config, tokenMaker))
	app.Get("/user/allGetByEmail", users.AllGetByEmail(store))
	app.Get("/user/getUserId", users.GetUserId(store))
	app.Delete("/user/update_my_profile/:user_id", users.DeleteUser(store))
	app.Patch("/user/delete_user/:user_id", users.UpdateUser(store))
	// //////////  Create Doctor  ///////////////////////
	app.Get("/doctor/signup", users.CreateDoctor(store))
	app.Get("user/doctorAppointment", users.DoctorAppointment(store))
	app.Get("category/create_category", users.DoctorCategory(store))
	// ////////  Appointment  /////////////////////////
	app.Post("appointment/create_appointment", users.CreateAppointment(store))
	app.Get("/appointment/all_appointment", users.GetAllAppointments(store))
	app.Get("/appointment/my_appointment", users.MyAppointment(store))
	app.Delete("/appointment/delete_appointment/:user_id", users.DeleteAppointment(store))
	app.Get("/user/appointmentDetials/:appointment_id", users.AppointmentDetials(store))
	app.Get("/user/seeAllDoctorAppointment", users.SeeAllDoctorAppointment(store))
	app.Get("/user/seeAllUserAppointment", users.SeeAllUserAppointment(store))
	app.Use("*", func(c *fiber.Ctx) error {
		return c.JSON("Route not exites")
	})

	app.Listen(":3000")
}
