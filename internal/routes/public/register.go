package routes

import (
	"0day-backend/internal/helpers"
	"0day-backend/pkg/logging"
	"0day-backend/pkg/mongodb"
	"strings"

	"github.com/gofiber/fiber/v3"
	"go.mongodb.org/mongo-driver/bson"
)

type Student struct {
	Name            string `json:"name"`
	Phone           string `json:"phone"`
	Roll            string `json:"roll"`
	DiscordUsername string `json:"discord_username"`
	GitHubProfile   string `json:"github_profile"`
	StudentEmailID  string `json:"student_email_id"`
}

func Register(c fiber.Ctx) error {
	var student Student

	if err := c.Bind().JSON(&student); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}
	if student.Name == "" || student.Phone == "" || student.Roll == "" || student.DiscordUsername == "" || student.StudentEmailID == "" {
		return helpers.Response(c, fiber.ErrBadRequest.Code, `{"sucess":"false","message":"Missing or Empty Required Fields"}`)
	}

	if strings.Contains(student.Phone, `+91`) {

		student.Phone = strings.Replace(student.Phone, `+91`, ``, 1)
	}

	payload := bson.M{"roll": student.Roll}
	_, err := mongodb.DB.FindOne(payload, mongodb.DB.Collections.Registrations)
	if err == nil {
		return helpers.Response(c, fiber.ErrBadRequest.Code, `{"sucess":"false","message":"User is Already Registered"}`)

	}

	payload = bson.M{
		"name":             student.Name,
		"phone":            student.Phone,
		"roll":             student.Roll,
		"discord_username": student.DiscordUsername,
		"github_profile":   student.GitHubProfile,
		"student_email_id": student.StudentEmailID,
	}

	_, err = mongodb.DB.InsertOne(payload, mongodb.DB.Collections.Registrations)
	if err != nil {
		logging.Logger.Error().Msgf("Failed to insert Registrations into database: %v", err)
	}
	//whatsapp.SendDM(student.Phone)
	return helpers.Response(c, fiber.StatusOK, `{"sucess":"true","message":"User Has Been Registered Successfully"}`)

}
