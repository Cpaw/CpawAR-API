package controllers

import (
	"AiCompServer/app/db"
	"AiCompServer/app/models"
	"github.com/revel/revel"
	"gopkg.in/validator.v2"
)

type ApiNotification struct {
	ApiV1Controller
}

type ResponseNotification struct {
	Notification *models.Notification `json:"notification"`
}

type ResponseNotifications struct {
	Notifications []models.Notification `json:"notifications"`
}

func (c ApiNotification) Index() revel.Result {
	if err := CheckToken(c.ApiV1Controller); err != nil {
		return err
	}
	notifications := []models.Notification{}
	if err := db.DB.Order("id desc").Find(&notifications).Error; err != nil {
		return c.HandleNotFoundError("Record Find Failure")
	}
	r := Response{ResponseNotifications{notifications}}
	return c.RenderJSON(r)
}

func (c ApiNotification) Show(id int) revel.Result {
	if err := CheckToken(c.ApiV1Controller); err != nil {
		return err
	}
	notification := &models.Notification{}
	if err := db.DB.First(&notification, id).Error; err != nil {
		return c.HandleNotFoundError(err.Error())
	}
	r := Response{ResponseNotification{notification}}
	return c.RenderJSON(r)
}

func (c ApiNotification) Create() revel.Result {
	if err := CheckRole(c.ApiV1Controller, []string{"admin"}); err != nil {
		return err
	}
	if err := CheckToken(c.ApiV1Controller); err != nil {
		return err
	}
	notification := &models.Notification{}
	if err := c.BindParams(notification); err != nil {
		return c.HandleBadRequestError(err.Error())
	}
	if err := validator.Validate(notification); err != nil {
		return c.HandleBadRequestError(err.Error())
	}
	if err := db.DB.Create(notification).Error; err != nil {
		return c.HandleBadRequestError(err.Error())
	}
	r := Response{ResponseNotification{notification}}
	return c.RenderJSON(r)
}

func (c ApiNotification) Update(id int) revel.Result {
	if err := CheckRole(c.ApiV1Controller, []string{"admin"}); err != nil {
		return err
	}
	if err := CheckToken(c.ApiV1Controller); err != nil {
		return err
	}
	notificationOld := &models.Notification{}
	if err := db.DB.First(&notificationOld, id).Error; err != nil {
		return c.HandleNotFoundError(err.Error())
	}
	notificationNew := &models.Notification{}
	if err := c.BindParams(notificationNew); err != nil {
		return c.HandleBadRequestError(err.Error())
	}
	if err := validator.Validate(notificationNew); err != nil {
		return c.HandleBadRequestError(err.Error())
	}
	if err := db.DB.Model(&notificationOld).Update(&notificationNew).Error; err != nil {
		return c.HandleNotFoundError(err.Error())
	}
	r := Response{ResponseNotification{notificationNew}}
	return c.RenderJSON(r)
}

func (c ApiNotification) Delete(id int) revel.Result {
	if err := CheckRole(c.ApiV1Controller, []string{"admin"}); err != nil {
		return err
	}
	if err := CheckToken(c.ApiV1Controller); err != nil {
		return err
	}
	notification := &models.Notification{}
	if err := db.DB.First(&notification, id).Error; err != nil {
		return c.HandleNotFoundError(err.Error())
	}
	if err := db.DB.Delete(&notification).Error; err != nil {
		return c.HandleInternalServerError("Record Delete Failure")
	}
	r := Response{"Success Delete"}
	return c.RenderJSON(r)
}
