package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"golang-clean-architecture/entity"
	"gorm.io/gorm"
)

func NewAuth(db *gorm.DB, log *logrus.Logger) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		token := ctx.Get("Authorization", "NOT_FOUND")
		log.Debugf("Authorization : %s", token)

		user := new(entity.User)
		err := db.Take(user, "token = ?", token).Error
		if err != nil {
			log.Warnf("Failed find user by token : %+v", err)
			return err
		}

		log.Debugf("User : %+v", user)
		ctx.Locals("user", user)
		return ctx.Next()
	}
}
