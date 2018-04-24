package controllers

import (
	"AiCompServer/app/db"
	"AiCompServer/app/models"
	"github.com/revel/revel"
	"github.com/revel/revel/cache"
	"gopkg.in/validator.v2"
	"time"
)

type ApiSetting struct {
	ApiV1Controller
}

type ResponseSetting struct {
	Setting *models.Setting `json:"setting"`
}

type ResponseSettings struct {
	Settings []models.Setting `json:"settings"`
}

func (c ApiSetting) Index() revel.Result {
	if err := CheckRole(c.ApiV1Controller, []string{"admin"}); err != nil {
		return err
	}
	if err := CheckToken(c.ApiV1Controller); err != nil {
		return err
	}
	settings := []models.Setting{}
	if err := db.DB.Order("id desc").Find(&settings).Error; err != nil {
		return c.HandleNotFoundError("Record Find Failure")
	}
	r := Response{ResponseSettings{settings}}
	return c.RenderJSON(r)
}

func (c ApiSetting) Show(id int) revel.Result {
	if err := CheckRole(c.ApiV1Controller, []string{"admin"}); err != nil {
		return err
	}
	if err := CheckToken(c.ApiV1Controller); err != nil {
		return err
	}
	setting := &models.Setting{}
	if err := db.DB.First(&setting, id).Error; err != nil {
		return c.HandleNotFoundError(err.Error())
	}
	r := Response{ResponseSetting{setting}}
	return c.RenderJSON(r)
}

func (c ApiSetting) Create() revel.Result {
	if err := CheckRole(c.ApiV1Controller, []string{"admin"}); err != nil {
		return err
	}
	if err := CheckToken(c.ApiV1Controller); err != nil {
		return err
	}
	setting := &models.Setting{}
	if err := c.BindParams(setting); err != nil {
		return c.HandleBadRequestError(err.Error())
	}
	if err := validator.Validate(setting); err != nil {
		return c.HandleBadRequestError(err.Error())
	}
	if err := db.DB.Create(setting).Error; err != nil {
		return c.HandleBadRequestError(err.Error())
	}
	r := Response{ResponseSetting{setting}}
	return c.RenderJSON(r)
}

func (c ApiSetting) Update(id int) revel.Result {
	if err := CheckRole(c.ApiV1Controller, []string{"admin"}); err != nil {
		return err
	}
	if err := CheckToken(c.ApiV1Controller); err != nil {
		return err
	}
	settingOld := &models.Setting{}
	if err := db.DB.First(&settingOld, id).Error; err != nil {
		return c.HandleNotFoundError(err.Error())
	}
	settingNew := &models.Setting{}
	if err := c.BindParams(settingNew); err != nil {
		return c.HandleBadRequestError(err.Error())
	}
	if err := validator.Validate(settingNew); err != nil {
		return c.HandleBadRequestError(err.Error())
	}
	if err := db.DB.Model(&settingOld).Update(&settingNew).Error; err != nil {
		return c.HandleNotFoundError(err.Error())
	}
	r := Response{ResponseSetting{settingNew}}
	return c.RenderJSON(r)
}

func (c ApiSetting) Delete(id int) revel.Result {
	if err := CheckRole(c.ApiV1Controller, []string{"admin"}); err != nil {
		return err
	}
	if err := CheckToken(c.ApiV1Controller); err != nil {
		return err
	}
	setting := &models.Setting{}
	if err := db.DB.First(&setting, id).Error; err != nil {
		return c.HandleNotFoundError(err.Error())
	}
	if err := db.DB.Delete(&setting).Error; err != nil {
		return c.HandleInternalServerError("Record Delete Failure")
	}
	r := Response{"Success Delete"}
	return c.RenderJSON(r)
}

func (c ApiSetting) Adapt(id int) revel.Result {
	if err := CheckRole(c.ApiV1Controller, []string{"admin"}); err != nil {
		return err
	}
	if err := CheckToken(c.ApiV1Controller); err != nil {
		return err
	}
	setting := &models.Setting{}
	if err := db.DB.First(&setting, id).Error; err != nil {
		return c.HandleNotFoundError(err.Error())
	}
	// キャッシュする。または更新する
	go cache.Set("setting_start_at", setting.StartAt, cache.ForEverNeverExpiry)
	go cache.Set("setting_end_at", setting.EndAt, cache.ForEverNeverExpiry)
	go cache.Set("setting_challenge_public", setting.ChallangePublic, cache.ForEverNeverExpiry)
	go cache.Set("setting_submit", setting.Submit, cache.ForEverNeverExpiry)
	go cache.Set("setting_preranking", setting.PreRanking, cache.ForEverNeverExpiry)
	r := Response{"Success Adapt"}
	return c.RenderJSON(r)
}

func Holding() bool {
	var start time.Time
	if err := cache.Get("setting_start_at", &start); err != nil {
		return false
	}
	var end time.Time
	if err := cache.Get("setting_end_at", &end); err != nil {
		return false
	}
	// compare start < now < end
	now := time.Now()
	if start.Before(now) && end.After(now) {
		return true
	}
	return false
}

func ChallengePublic() bool {
	var challengepublic bool
	if err := cache.Get("setting_challenge_public", &challengepublic); err != nil {
		return false
	}
	return challengepublic
}

func SubmitPublic() bool {
	var submitpublic bool
	if err := cache.Get("setting_challenge_public", &submitpublic); err != nil {
		return false
	}
	return submitpublic
}

func PreRanking() bool {
	var preranking bool
	if err := cache.Get("setting_preranking", &preranking); err != nil {
		return false
	}
	return preranking
}
