package like

import (
	"context"
	"errors"

	"github.com/bhuvnesh13396/PlayMySong/common/id"
	"github.com/bhuvnesh13396/PlayMySong/model"
)

var (
	errInvalidArgument = errors.New("Invalid Arguments.")
)

type Service interface {
	Add(ctx context.Context, activityType string, activityID string, userID string) (err error)
	Delete(ctx context.Context, activityID string, userID string) (err error)
	Get(ctx context.Context, activityID string) (count int, err error)
	UpdateCount(ctx context.Context, activityID string, count int) (err error)
}

type service struct {
	likeRepo model.LikeRepo
}

func NewService(likeRepo model.LikeRepo) Service {
	return &service{
		likeRepo: likeRepo,
	}
}

func (s *service) Get(ctx context.Context, activityID string) (count int, err error) {
	return s.likeRepo.Get(activityID)
}

func (s *service) Add(ctx context.Context, activityType string, activityID string, userID string) (err error) {

	if len(activityType) < 1 || len(activityID) < 1 || len(userID) < 1 {
		err = errInvalidArgument
		return
	}

	like := model.Like{
		ID:         id.New(),
		Type:       activityType,
		ActivityID: activityID,
		UserID:     userID,
	}

	// Add the entry in the likes table
	err = s.likeRepo.Add(like)
	if err != nil {
		return
	}

	// Increment the count of likes in the aggregation collection
	// for the given activity
	currentLikes, err := s.Get(ctx, activityID)
	if err != nil {
		return err
	}

	return s.UpdateCount(ctx, activityID, currentLikes+1)
}

func (s *service) Delete(ctx context.Context, activityID string, userID string) (err error) {

	if len(activityID) < 1 || len(userID) < 1 {
		err = errInvalidArgument
		return
	}

	err = s.likeRepo.Delete(activityID, userID)
	if err != nil {
		return
	}

	// Decrement the count of likes in the aggregation collection
	// for the given activity
	currentLikes, err := s.Get(ctx, activityID)
	if err != nil {
		return err
	}

	return s.UpdateCount(ctx, activityID, currentLikes-1)
}

func (s *service) UpdateCount(ctx context.Context, activityID string, count int) (err error) {
	if len(activityID) < 1 {
		err = errInvalidArgument
		return
	}

	return s.likeRepo.UpdateCount(activityID, count)
}
