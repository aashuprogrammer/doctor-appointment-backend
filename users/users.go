package users

import (
	"log"
	"mylearning/db/pgdb"
	"mylearning/token"
	"mylearning/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgtype"
)

type CreateUserRequeset struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"pass"`
}

type CreateUserResponse struct {
	Msg  string    `json:"msg"`
	User pgdb.User `json:"user"`
}

func CreateUser(a pgdb.Store) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var req CreateUserRequeset
		err := c.BodyParser(&req)
		if err != nil {
			log.Fatal("unable to parse", err)
		}
		HashPassword, err := utils.HashPassword(req.Password)
		if err != nil {
			return err
		}
		user, err := a.CreateUser(c.Context(), pgdb.CreateUserParams{
			Name:     req.Name,
			Email:    req.Email,
			Password: HashPassword,
		})
		if err != nil {
			return err
		}
		return c.JSON(CreateUserResponse{Msg: "user Created", User: user})
	}
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"pass"`
}

type LoginResponse struct {
	Name                string    `json:"name"`
	Email               string    `json:"email"`
	UserId              int64     `json:"user_id"`
	AccessToken         string    `json:"access_token"`
	Role                string    `json:"role"`
	AccessTokenExpireAt time.Time `json:"access_token_expireAt"`
}

////////  Users login ////////////////////

func UserLogin(a pgdb.Store, config utils.Config, token token.Maker) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var req LoginRequest
		err := c.BodyParser(&req)
		if err != nil {
			log.Fatal("unable to parse", err)
		}
		user, err := a.LoginUser(c.Context(), req.Email)
		if err != nil {
			return err
		}
		err = utils.VerifyPassword(user.Password, req.Password)
		if err != nil {
			return err
		}
		role := ""
		s, p, errr := token.CreateAccessToken(user.ID, user.Name, user.Email, role, config.AccessTokenDuration)
		if errr != nil {
			return err
		}

		return c.JSON(LoginResponse{
			Name:                p.Name,
			Email:               p.Email,
			UserId:              p.Id,
			AccessToken:         s,
			AccessTokenExpireAt: p.ExpireAt,
		})
	}
}

type AllUserResponse struct {
	User []pgdb.User `json:"user"`
}

func AllGetUser(g pgdb.Store) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		user, err := g.GetAllusers(c.Context())
		if err != nil {
			return err
		}
		return c.JSON(AllUserResponse{User: user})
	}
}

// type AllUserGetByNameResponse struct {
// 	Msg  string      `json:"msg"`
// 	User []pgdb.User `json:"user"`
// }

// func AllUserName(n pgdb.Store) func(c *fiber.Ctx) error {
// 	return func(c *fiber.Ctx) error {
// 		user, err := n.GetAllusers(c.Context())
// 		if err != nil {
// 			return err
// 		}
// 		return c.JSON(AllUserGetByNameResponse{Msg: "get all user Success", User: user})
// 	}
// }

type AllGetUserEmailRequest struct {
	Email string `json:"email"`
}
type AllGetUserEmailResponse struct {
	User pgdb.User `json:"user"`
}

func AllGetByEmail(e pgdb.Store) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var req AllGetUserEmailRequest
		err := c.BodyParser(&req)
		if err != nil {
			log.Fatal("unable to email", err)
		}
		user, err := e.GetUserByEmail(c.Context(), req.Email)
		if err != nil {
			return err
		}
		return c.JSON(AllGetUserEmailResponse{User: user})
	}

}

type Doctor struct {
	User  pgdb.User `json:"user"`
	Field string    `json:"field"`
}

type GetUserByIdRequest struct {
	User pgdb.User `json:"user"`
	Id   int64     `json:"id"`
}

type GetUserByIdResponse struct {
	Msg  string    `json:"msg"`
	User pgdb.User `json:"user"`
}

func GetUserId(i pgdb.Store) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var req GetUserByIdRequest
		err := c.BodyParser(&req)
		if err != nil {
			log.Fatal("Unable to id", err)
		}
		user, err := i.GetUserById(c.Context(), int64(req.Id))
		if err != nil {
			return err
		}

		return c.JSON(GetUserByIdResponse{Msg: "User id", User: user})
	}
}

///////////  User Delete  /////////////////////

type DeleteUserResponse struct {
	Msg string `json:"msg"`
}

func DeleteUser(d pgdb.Store) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("user_id")
		err := d.DeleteUser(c.Context(), int64(id))
		if err != nil {
			return err
		}
		return c.JSON(DeleteUserResponse{Msg: "User Delete"})
	}
}

// ///////// User Update  //////////////////////
type UpdateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserResponse struct {
	Msg string `json:"msg"`
}

func UpdateUser(u pgdb.Store) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var req UpdateUserRequest
		err := c.BodyParser(&req)
		if err != nil {
			log.Fatal("Not Update", err)
		}
		id, _ := c.ParamsInt("user_id")
		HashPassword, err := utils.HashPassword(req.Password)
		if err != nil {
			return err
		}
		err = u.UpdateUser(c.Context(), pgdb.UpdateUserParams{
			ID:       int64(id),
			Name:     pgtype.Text{String: req.Name, Valid: req.Name != ""},
			Email:    pgtype.Text{String: req.Email, Valid: req.Email != ""},
			Password: pgtype.Text{String: HashPassword, Valid: HashPassword != ""},
		})
		if err != nil {
			return err
		}
		return c.JSON(UpdateUserResponse{Msg: "User Update"})
	}
}
