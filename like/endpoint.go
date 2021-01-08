package like

import (
	"context"

	"github.com/bhuvnesh13396/PlayMySong/common/kit"
)

type GetEndpoint kit.Endpoint
type AddEndpoint kit.Endpoint
type DeleteEndpoint kit.Endpoint
type UpdateCountEndpoint kit.Endpoint

type Endpoint struct {
	GetEndpoint
	AddEndpoint
	DeleteEndpoint
	UpdateCountEndpoint
}

type getRequest struct {
	ActivityID string `schema:"activity_id"`
}

type getResponse struct {
	Count int `json:"count"`
}

func MakeGetEndpoint(s Service) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getRequest)
		count, err := s.Get(ctx, req.ActivityID)
		return getResponse{Count: count}, err
	}
}

func (e GetEndpoint) Get(ctx context.Context, activityID string) (count int, err error) {
	request := getRequest{
		ActivityID: activityID,
	}
	response, err := e(ctx, request)
	resp := response.(getResponse)
	return resp.Count, err
}

type addRequest struct {
	Type       string `json:"type"`
	ActivityID string `json:"activity_id"`
	UserID     string `json:"user_id"`
}

type addResponse struct {
}

func MakeAddEndpoint(s Service) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addRequest)
		err := s.Add(ctx, req.Type, req.ActivityID, req.UserID)
		return addResponse{}, err
	}
}

func (e AddEndpoint) Add(ctx context.Context, activityType string, activityID string, userID string) (err error) {
	request := addRequest{
		Type:       activityType,
		ActivityID: activityID,
		UserID:     userID,
	}
	_, err = e(ctx, request)
	return err
}

type deleteRequest struct {
	ActivityID string `json:"activity_id"`
	UserID     string `json:"user_id"`
}

type deleteResponse struct {
}

func MakeDeleteEndpoint(s Service) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteRequest)
		err := s.Delete(ctx, req.ActivityID, req.UserID)
		return deleteResponse{}, err
	}
}

func (e DeleteEndpoint) Delete(ctx context.Context, activityID string, userID string) (err error) {
	request := deleteRequest{
		ActivityID: activityID,
		UserID:     userID,
	}
	_, err = e(ctx, request)
	return err
}

type updateCountRequest struct {
	ActivityID string `json:"activity_id"`
	Count      int    `json:"count"`
}

type updateCountResponse struct {
}

func MakeUpdateCountEndpoint(s Service) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateCountRequest)
		err := s.UpdateCount(ctx, req.ActivityID, req.Count)
		return updateCountResponse{}, err
	}
}

func (e UpdateCountEndpoint) UpdateCount(ctx context.Context, activityID string, count int) (err error) {
	request := updateCountRequest{
		ActivityID: activityID,
		Count:      count,
	}

	_, err = e(ctx, request)
	return err
}
