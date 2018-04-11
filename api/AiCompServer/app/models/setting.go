package models

import (
	"time"
)

type Setting struct {
	BaseModel
	StartAt         time.Time `json:"start_at"`
	EndAt           time.Time `json:"end_at"`
	ChallangePublic bool      `json:"challenge_public_is"`
	Submit          bool      `json:"submit_is"`
	PreRanking      bool      `json:"preranking_is"`
}
