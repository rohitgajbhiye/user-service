package service

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
	"user-service/model"

	"github.com/google/uuid"
)

type UserService interface {
	FetchAllUsers(context.Context) (model.Users, error)
	FetchUserById(context.Context, string) (model.User, error)
	RegisterUser(context.Context, model.User) (string, error)
	isExists(context.Context, string) (bool, error)
}

type userService struct {
}

func NewUserService() UserService {
	return &userService{}
}

func (serv *userService) FetchAllUsers(ctx context.Context) (model.Users, error) {
	file, err := os.Open("data/users.json")
	var users model.Users
	if err != nil {
		return users, err
	}
	defer file.Close()
	bytesValue, err := ioutil.ReadAll(file)
	if err != nil {
		return users, err
	}
	json.Unmarshal(bytesValue, &users)
	return users, nil
}

func (serv *userService) FetchUserById(ctx context.Context, id string) (model.User, error) {
	users, err := serv.FetchAllUsers(ctx)
	if err != nil {
		return model.User{}, err
	}
	return users[id], nil
}

func (serv *userService) RegisterUser(ctx context.Context, user model.User) (string, error) {
	file, err := os.Open("data/users.json")
	if err != nil {
		return "", err
	}
	defer file.Close()
	bytesValue, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	var users model.Users
	json.Unmarshal(bytesValue, &users)
	id := uuid.New().String()
	user.Id = id
	users[id] = user
	changed, writeErr := json.Marshal(users)
	if writeErr != nil {
		return "", writeErr
	}
	ioErr := ioutil.WriteFile("data/users.json", changed, os.ModePerm)
	if ioErr != nil {
		return "", ioErr
	}
	return id, nil
}

func (serv *userService) isExists(ctx context.Context, id string) (bool, error) {
	_, err := serv.FetchUserById(ctx, id)
	if err != nil {
		return false, nil
	}
	return true, nil
}
