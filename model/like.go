package model

type Like struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	ActivityID string `json:"activity_id"`
	UserID     string `json:"user_id"`
}

type LikeRepo interface {
	Add(like Like) (err error)
	Delete(activityID string, userID string) (err error)
	Get(activityID string) (count int, err error)
	UpdateCount(ActivityID string, count int) (err error)
}
