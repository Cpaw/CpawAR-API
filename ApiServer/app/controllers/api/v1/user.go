package controllers

import (
	"ApiServer/app/db"
	"ApiServer/app/models"
	"github.com/revel/revel"
	"gopkg.in/validator.v2"
	"crypto/rand"
	"fmt"
)

type ApiUser struct {
	ApiV1Controller
}

type ReseposeUser struct {
	User *models.User `json:"user"`
}

// User Create
func (c ApiUser) Create() revel.Result {
	user := &models.User{ Token: TokenGenerator(32) }

	if err := validator.Validate(user); err != nil {
		return c.HandleBadRequestError(err.Error())
	}
	if err := db.DB.Create(user).Error; err != nil {
		return c.HandleBadRequestError(err.Error())
	}
	if err := db.DB.First(&user).Error; err != nil {
		return c.HandleNotFoundError(err.Error())
	}

	r := Response{ReseposeUser{user}}
	return c.RenderJSON(r)
}

func TokenGenerator(len int) string {
	b := make([]byte, len)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
