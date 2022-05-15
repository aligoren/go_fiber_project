package routes

import (
	"github.com/aligoren/fiber_project/models"
	"github.com/gofiber/fiber/v2"
	"time"
)

func checkUserIsValid(u *models.UserLoginModel) bool {
	return u.Username == "admin" && u.Password == "admin"
}

func Login(c *fiber.Ctx) error {

	userLogin := new(models.UserLoginModel)

	if err := c.BodyParser(userLogin); err != nil {
		return err
	}

	isUserValid := checkUserIsValid(userLogin)
	responseModel := models.ResponseModel{
		Success:    true,
		StatusCode: 200,
		Message:    "User logged in successfully",
	}

	if !isUserValid {
		responseModel.Success = false
		responseModel.StatusCode = 404
		responseModel.Message = "User couldn't login successfully"
	}

	return c.Status(responseModel.StatusCode).JSON(responseModel)

}

func GetUsers(c *fiber.Ctx) error {
	users := []models.User{
		{
			ID:        1,
			FirstName: "Ali",
			Surname:   "GÖREN",
			CreatedAt: time.Time{},
		},
		{
			ID:        2,
			FirstName: "Eren",
			Surname:   "Jeager",
			CreatedAt: time.Time{},
		},
	}

	return c.JSON(users)
}
