package service

import (
	"context"
	om "delinkcious/pkg/object_model"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

type followRequest struct {
	Followed string `json:"followed"`
	Follower string `json:"follower"`
}

type followResponse struct {
	Err string `json:"err"`
}

type unfollowRequest struct {
	Followed string `json:"followed"`
	Follower string `json:"follower"`
}

type unfollowResponse struct {
	Err string `json:"err"`
}

type getByUsernameRequest struct {
	Username string `json:"username"`
}

type getFollowersResponse struct {
	Followers map[string]bool `json:"followers"`
	Err       string          `json:"err"`
}

type getFollowingResponse struct {
	Following map[string]bool `json:"following"`
	Err       string          `json:"err"`
}

func decodeFollowRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request followRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, err
	}
	return request, nil
}

func decodeUnfollowRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request unfollowRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, err
	}
	return request, nil
}

func decodeGetFollowingRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request getByUsernameRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, err
	}
	return request, nil
}

func decodeGetFollowersRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request getByUsernameRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func makeFollowEndpoint(svc om.SocialGraphManager) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(followRequest)
		err := svc.Follow(req.Followed, req.Follower)
		res := followResponse{}
		if err != nil {
			res.Err = err.Error()
		}
		return res, nil
	}
}

func makeUnfollowEndpoint(svc om.SocialGraphManager) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(unfollowRequest)
		err := svc.Unfollow(req.Followed, req.Follower)
		res := unfollowResponse{}
		if err != nil {
			res.Err = err.Error()
		}
		return res, nil
	}
}

func makeGetFollowingEndpoint(svc om.SocialGraphManager) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(getByUsernameRequest)
		followingMap, err := svc.GetFollowing(req.Username)
		res := getFollowingResponse{Following: followingMap}
		if err != nil {
			res.Err = err.Error()
		}
		return res, nil
	}
}

func makeGetFollowersEndpoint(svc om.SocialGraphManager) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(getByUsernameRequest)
		followersMap, err := svc.GetFollowers(req.Username)
		res := getFollowersResponse{Followers: followersMap}
		if err != nil {
			res.Err = err.Error()
		}
		return res, nil
	}
}