package controllers

import (
	"AiCompServer/app/db"
	"AiCompServer/app/models"
	"bufio"
	"github.com/revel/revel"
	"gopkg.in/validator.v2"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type ApiAnswer struct {
	ApiV1Controller
}

type ResponseAnswer struct {
	Answer *models.Answer `json:"answer"`
}

type ResponseAnswers struct {
	Answers []models.Answer `json:"answers"`
}

type ResponseAccuracy struct {
	Accuracy float64 `json:"accuracy"`
}

// Answer Index
func (c ApiAnswer) Index() revel.Result {
	if err := CheckToken(c.ApiV1Controller); err != nil {
		return err
	}
	answers := []models.Answer{}
	if err := db.DB.Order("id desc").Find(&answers).Error; err != nil {
		return c.HandleNotFoundError("Record Find Failure")
	}
	r := Response{ResponseAnswers{answers}}
	return c.RenderJSON(r)
}

// Answer Show
func (c ApiAnswer) Show(id int) revel.Result {
	if err := CheckToken(c.ApiV1Controller); err != nil {
		return err
	}
	answer := &models.Answer{}
	if err := db.DB.First(&answer, id).Error; err != nil {
		return c.HandleNotFoundError(err.Error())
	}
	r := Response{ResponseAnswer{answer}}
	return c.RenderJSON(r)
}

// Answer User's Answer of Some Challenge
func (c ApiAnswer) UserChallengeAnswer(id int) revel.Result {
	if err := CheckToken(c.ApiV1Controller); err != nil {
		return err
	}
	token := c.Request.Header.Get("Authorization")
	user := &models.User{}
	if err := db.DB.Find(&user, models.User{Token: token}).Error; err != nil {
		return c.HandleNotFoundError(err.Error())
	}
	answer := &models.Answer{}
	if err := db.DB.Where("user_id = ? AND challenge_id = ?", user.ID, id).First(&answer).Error; err != nil {
		return c.HandleNotFoundError(err.Error())
	}
	r := Response{ResponseAnswer{answer}}
	return c.RenderJSON(r)
}

// Answer Create
func (c ApiAnswer) Create() revel.Result {
	if err := CheckRole(c.ApiV1Controller, []string{"admin"}); err != nil {
		return err
	}
	if err := CheckToken(c.ApiV1Controller); err != nil {
		return err
	}
	answer := &models.Answer{}
	if err := c.BindParams(answer); err != nil {
		return c.HandleBadRequestError(err.Error())
	}
	token := c.Request.Header.Get("Authorization")
	user := &models.User{}
	if err := db.DB.Find(&user, models.User{Token: token}).Error; err != nil {
		return c.HandleNotFoundError(err.Error())
	}
	if err := validator.Validate(answer); err != nil {
		return c.HandleBadRequestError(err.Error())
	}
	if err := db.DB.Create(answer).Error; err != nil {
		return c.HandleBadRequestError(err.Error())
	}
	r := Response{ResponseAnswer{answer}}
	return c.RenderJSON(r)
}

// Answer Update
func (c ApiAnswer) Update(id int) revel.Result {
	if err := CheckRole(c.ApiV1Controller, []string{"admin"}); err != nil {
		return err
	}
	if err := CheckToken(c.ApiV1Controller); err != nil {
		return err
	}
	answerOld := &models.Answer{}
	if err := db.DB.First(&answerOld, id).Error; err != nil {
		return c.HandleNotFoundError(err.Error())
	}
	answerNew := &models.Answer{}
	if err := c.BindParams(answerNew); err != nil {
		return c.HandleBadRequestError(err.Error())
	}
	if err := validator.Validate(answerNew); err != nil {
		return c.HandleBadRequestError(err.Error())
	}
	if err := db.DB.Model(&answerOld).Update(&answerNew).Error; err != nil {
		return c.HandleNotFoundError(err.Error())
	}
	r := Response{ResponseAnswer{answerNew}}
	return c.RenderJSON(r)
}

// Answer Delete
func (c ApiAnswer) Delete(id int) revel.Result {
	if err := CheckRole(c.ApiV1Controller, []string{"admin"}); err != nil {
		return err
	}
	if err := CheckToken(c.ApiV1Controller); err != nil {
		return err
	}
	answer := &models.Answer{}
	if err := db.DB.First(&answer, id).Error; err != nil {
		return c.HandleNotFoundError(err.Error())
	}
	if err := db.DB.Delete(&answer).Error; err != nil {
		return c.HandleInternalServerError("Record Delete Failure")
	}
	r := Response{"Success Delete"}
	return c.RenderJSON(r)
}

func (c ApiAnswer) Submit(ChallengeID uint64, ansFP *os.File) revel.Result {
	if err := CheckToken(c.ApiV1Controller); err != nil {
		return err
	}
	ansFile := revel.Config.StringDefault("answer.file", "/correct/answer") + strconv.Itoa(int(ChallengeID)) + ".csv"
	fp, err := os.Open(ansFile)
	if err != nil {
		return c.HandleBadRequestError("正解ファイルが開けませんでした")
	}
	scanner1 := bufio.NewScanner(ansFP)
	scanner2 := bufio.NewScanner(fp)
	var acc float64
	acc = 0.0
	a1 := map[string]string{}
	a2 := map[string]string{}
	for scanner1.Scan() {
		st1 := scanner1.Text()
		l1 := strings.Split(st1, ",")
		if len(l1) > 1 {
			a1[l1[0]] = l1[1]
		}
	}
	for scanner2.Scan() {
		st2 := scanner2.Text()
		l2 := strings.Split(st2, ",")
		if len(l2) > 1 {
			a2[l2[0]] = l2[1]
		}
	}
	PreCount := 0
	pacc := 0.0
	for asKey, asVal := range a1 {
		if aaVal, err := a2[asKey]; err == true {
			if aaVal == asVal {
				acc = acc + 1.0
				if PreCount < 100 {
					pacc = pacc + 1.0
				}
			}
			PreCount = PreCount + 1
		}
	}
	if err := scanner1.Err(); err != nil {
		return c.HandleBadRequestError("採点中に解答ファイルにエラーが起きました")
	}
	if err := scanner2.Err(); err != nil {
		return c.HandleBadRequestError("採点中に正解ファイルにエラーが起きました")
	}
	challenge := &models.Challenge{}
	if err := db.DB.First(&challenge, ChallengeID).Error; err != nil {
		return c.HandleNotFoundError(err.Error())
	}
	acc = acc * 100.0 * challenge.Weight / float64(len(a2))
	// Submitしたユーザーを特定する
	token := c.Request.Header.Get("Authorization")
	user := &models.User{}
	if err := db.DB.Find(&user, models.User{Token: token}).Error; err != nil {
		return c.HandleNotFoundError(err.Error())
	}
	log.Println(user.Username)
	log.Println(acc)
	log.Println(pacc)
	// まずはDBを探してあったらUpdate処理、なかったらCreate
	// そのユーザーがSubmitした問題のanswerを探す
	answer := &models.Answer{}
	if err := db.DB.Where("user_id = ? AND challenge_id = ?", user.ID, ChallengeID).First(&answer).Error; err != nil {
		// なかった場合のCreate
		answer.ChallengeID = ChallengeID
		answer.UserID = user.ID
		answer.Score = acc
		if err := validator.Validate(answer); err != nil {
			return c.HandleBadRequestError(err.Error())
		}
		if err := db.DB.Create(answer).Error; err != nil {
			return c.HandleBadRequestError(err.Error())
		}
		r := Response{ResponseAccuracy{acc}}
		return c.RenderJSON(r)
	}
	// 前回の更新から1分たっているか確認する(解答の解析を出来なくするため
	if time.Since(answer.UpdatedAt) < time.Minute*1 {
		return c.HandleBadRequestError("前回のSubmitから1分が経過していません。")
	}
	// そのユーザーがSubmitした問題のanswerを更新する
	answerNew := &models.Answer{}
	answerNew.ChallengeID = ChallengeID
	answerNew.UserID = user.ID
	if answer.Score < acc {
		answerNew.Score = acc
	} else {
		answerNew.Score = answer.Score
	}
	if err := validator.Validate(answerNew); err != nil {
		return c.HandleBadRequestError(err.Error())
	}
	if err := db.DB.Model(&answer).Update(&answerNew).Error; err != nil {
		return c.HandleNotFoundError(err.Error())
	}
	// PreAnswer
	panswer := &models.PreAnswer{}
	if err := db.DB.Where("user_id = ? AND challenge_id = ?", user.ID, ChallengeID).First(&panswer).Error; err != nil {
		// なかった場合のCreate
		panswer.ChallengeID = ChallengeID
		panswer.UserID = user.ID
		panswer.Score = pacc
		if err := validator.Validate(panswer); err != nil {
			return c.HandleBadRequestError(err.Error())
		}
		if err := db.DB.Create(panswer).Error; err != nil {
			return c.HandleBadRequestError(err.Error())
		}
		r := Response{ResponseAccuracy{pacc}}
		return c.RenderJSON(r)
	}
	// 前回の更新から1分たっているか確認する(解答の解析を出来なくするため
	if time.Since(panswer.UpdatedAt) < time.Minute*1 {
		return c.HandleBadRequestError("前回のSubmitから1分が経過していません。")
	}
	// そのユーザーがSubmitした問題のanswerを更新する
	panswerNew := &models.PreAnswer{}
	panswerNew.ChallengeID = ChallengeID
	panswerNew.UserID = user.ID
	if panswer.Score < pacc {
		panswerNew.Score = pacc
	} else {
		panswerNew.Score = panswer.Score
	}
	if err := validator.Validate(panswerNew); err != nil {
		return c.HandleBadRequestError(err.Error())
	}
	if err := db.DB.Model(&panswer).Update(&panswerNew).Error; err != nil {
		return c.HandleNotFoundError(err.Error())
	}
	if PreRanking() {
		r := Response{ResponseAccuracy{panswerNew.Score}}
		return c.RenderJSON(r)
	}
	r := Response{ResponseAccuracy{answerNew.Score}}
	return c.RenderJSON(r)
}
