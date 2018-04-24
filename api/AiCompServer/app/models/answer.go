package models

type PreAnswer struct {
	BaseModel
	ChallengeID uint64
	UserID      uint64
	Score       float64
}

type Answer struct {
	BaseModel
	ChallengeID uint64
	UserID      uint64
	Score       float64
}
