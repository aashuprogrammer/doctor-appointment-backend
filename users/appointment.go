package users

import (
	"log"
	"mylearning/db/pgdb"

	"github.com/gofiber/fiber/v2"
)

type CreateAppointmentRequeset struct {
	DoctorId int64  `json:"doctor_id"`
	UsersId  int64  `json:"users_id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Gender   string `json:"gender"`
	Age      string `json:"age"`
	Date     string `json:"date"`
	Time     string `json:"time"`
	Address  string `json:"address"`
}

type CreateAppointmentResponse struct {
	Msg  string           `json:"msg"`
	User pgdb.Appointment `json:"user"`
}

func CreateAppointment(p pgdb.Store) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var req CreateAppointmentRequeset
		err := c.BodyParser(&req)
		if err != nil {
			log.Fatal("Unable to Appointment", err)
		}

		user, err := p.CreateAppointment(c.Context(), pgdb.CreateAppointmentParams{
			DoctorID: req.DoctorId,
			UsersID:  req.UsersId,
			Name:     req.Name,
			Phone:    req.Phone,
			Gender:   pgdb.GenderT(req.Gender),
			Age:      req.Age,
			Date:     req.Date,
			Time:     req.Time,
			Address:  req.Address,
		})
		if err != nil {
			return err
		}
		return c.JSON(CreateAppointmentResponse{Msg: "Create Appointment", User: user})
	}
}

type DoctorAppointmentResponse struct {
	Msg      string             `json:"msg"`
	DoctorID int64              `json:"doctor_id"`
	User     []pgdb.Appointment `json:"user"`
}

func DoctorAppointment(d pgdb.Store) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var req DoctorAppointmentResponse
		err := c.BodyParser(&req)
		if err != nil {
			log.Fatal("Not Unable Appointment")
		}
		user, err := d.GetDoctorAppointment(c.Context(), req.DoctorID)
		if err != nil {
			return err
		}
		return c.JSON(DoctorAppointmentResponse{Msg: "Get Appointment", User: user})
	}
}

// ////////////// Get All Appointments  ////////////////////////////////
type GetAllAppointmentResponse struct {
	AppointmentList []pgdb.Appointment `json:"appointment"`
}

func GetAllAppointments(a pgdb.Store) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		arr, err := a.GetAllAppointment(c.Context())
		if err != nil {
			return err
		}
		return c.JSON(GetAllAppointmentResponse{AppointmentList: arr})
	}

}

// /////////  My Appointment  ////////////////////////////
type GetMyAppointmentRequest struct {
	UserId        int64              `json:"user_id"`
	AppointmentId []pgdb.Appointment `json:"appointment"`
}

func MyAppointment(m pgdb.Store) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var req GetMyAppointmentRequest
		err := c.BodyParser(&req)
		if err != nil {
			log.Fatal(" Unale to Your Appointment ")
		}
		user, err := m.MyAppointment(c.Context(), req.UserId)
		if err != nil {
			return err
		}
		return c.JSON(GetMyAppointmentRequest{AppointmentId: user})
	}
}

// ///////////// Delete Appointment  ////////////////////////
type DeleteAppointmentResponse struct {
	Msg string `json:"msg"`
}
type DeleteAppointmentRequest struct {
	ID int64 `json:"id"`
}

func DeleteAppointment(f pgdb.Store) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var req DeleteAppointmentRequest
		err := c.BodyParser(&req)
		if err != nil {
			log.Fatal("Appointment not Deleted")
		}
		id, _ := c.ParamsInt("user_id")
		err = f.DeleteAppointment(c.Context(), pgdb.DeleteAppointmentParams{
			ID:      req.ID,
			UsersID: int64(id),
		})
		if err != nil {
			return err
		}
		return c.JSON(DeleteAppointmentResponse{Msg: "Appointment Delete"})
	}
}

// ///////// See All Appointment Doctor id ////////////////////////////

type SeeAllAppointmentDoctorIdResponse struct {
	Msg  string             `json:"msg"`
	User []pgdb.Appointment `json:"user"`
}
type SeeAllAppointmentDoctorIdRequest struct {
	DoctorId int64 `json:"doctor_id"`
}

func SeeAllDoctorAppointment(z pgdb.Store) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var req SeeAllAppointmentDoctorIdRequest
		err := c.BodyParser(&req)
		if err != nil {
			return err
		}
		user, err := z.SeeAllAppointmentByDoctorId(c.Context(), req.DoctorId)
		if err != nil {
			return err
		}
		return c.JSON(SeeAllAppointmentDoctorIdResponse{Msg: "Doctor Id Seccess", User: user})
	}
}

///////// See All Appointment User id //////////////////////////////

type SeeAllAppointmentUserIdResponse struct {
	Msg  string             `json:"msg"`
	User []pgdb.Appointment `json:"user"`
}
type SeeAllAppointmentUserIdRequest struct {
	UserId int64 `json:"user_id"`
}

func SeeAllUserAppointment(z pgdb.Store) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var req SeeAllAppointmentUserIdRequest
		err := c.BodyParser(&req)
		if err != nil {
			return err
		}
		user, err := z.SeeAllAppointmentByUserId(c.Context(), req.UserId)
		if err != nil {
			return err
		}
		return c.JSON(SeeAllAppointmentUserIdResponse{Msg: "Doctor Id Seccess", User: user})
	}
}

// //////////////// Appointment Details ///////////////

type AppointmentDetialsResponse struct {
	Msg  string                        `json:"msg"`
	User pgdb.GetAppointmentDetailsRow `json:"user"`
}

func AppointmentDetials(w pgdb.Store) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("appointment_id")
		user, err := w.GetAppointmentDetails(c.Context(), int64(id))
		if err != nil {
			return err
		}
		return c.JSON(AppointmentDetialsResponse{Msg: "Get Appointment Details", User: user})
	}
}
