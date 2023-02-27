package controllers

import "taskmanager/models"

type (
	// UserResource for POST - /user/register
	UserResource struct {
		Data models.User `json:"data"`
	}
	// LoginResource for POST - /user/login
	LoginResource struct {
		Data LoginModel `json:"data"`
	}
	// AuthUserResource response for authorized user POST - /user/login
	AuthUserResource struct {
		Data AuthUserModel `json:"data"`
	}
	// LoginModel model for auth
	LoginModel struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	// AuthUserModel model for auth user with access token
	AuthUserModel struct {
		User  models.User `json:"user"`
		Token string      `json:"token"`
	}
	// TaskResource for POST/PUT - /tasks
	// for GET - /tasks/id
	TaskResource struct {
		Data models.Task `json:"data"`
	}
	// TasksResource for GET - /tasks
	TasksResource struct {
		Data []models.Task `json:"data"`
	}
)
