package controllers

import (
	"encoding/json"
	"net/http"
	"taskmanager/common"
	"taskmanager/data"
	"taskmanager/models"
)

const dbCollectionUser = "users"

// Register handler for HTTP POST - /users/register
// add a new user doc
func Register(w http.ResponseWriter, r *http.Request) {
	var dataResource UserResource
	// decode incoming User json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"invalid user data",
			500,
		)
		return
	}
	user := &dataResource.Data
	ctx := NewContext()
	defer ctx.Close()
	// get users collection
	c := ctx.DbCollection(dbCollectionUser)
	repo := &data.UserRepository{C: c}
	// insert user document
	user, err = repo.CreateUser(user)
	if err != nil {
		return
	}
	// clean-up the hashpassword from response
	user.HashPassword = nil
	if j, err := json.Marshal(UserResource{Data: *user}); err != nil {
		common.DisplayAppError(
			w,
			err,
			"an unexpected error has occurred",
			500,
		)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, err := w.Write(j)
		if err != nil {
			return
		}
	}
}

// Login handler for HTTP POST - /users/login
// authenticate with username and password
// TODO: If there are no such email registered, route to register page.
func Login(w http.ResponseWriter, r *http.Request) {
	var dataResource LoginResource
	var token string
	// decode the incoming json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"invalid login data",
			500,
		)
		return
	}
	loginModel := dataResource.Data
	loginUser := models.User{
		Email:    loginModel.Email,
		Password: loginModel.Password,
	}
	ctx := NewContext()
	defer ctx.Close()
	c := ctx.DbCollection(dbCollectionUser)
	repo := &data.UserRepository{C: c}

	// auth the login user
	if user, err := repo.Login(&loginUser); err != nil {
		common.DisplayAppError(
			w,
			err,
			"invalid login credentials",
			401,
		)
		return
	} else {
		token, err = common.GenerateJWT(user.Email, "member")
		if err != nil {
			common.DisplayAppError(
				w,
				err,
				"error while generating access token",
				500,
			)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		user.HashPassword = nil
		authUser := AuthUserModel{
			User:  *user,
			Token: token,
		}
		j, err := json.Marshal(AuthUserResource{Data: authUser})
		if err != nil {
			common.DisplayAppError(
				w,
				err,
				"an unexpected error has occurred",
				500,
			)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(j)
		if err != nil {
			return
		}
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var dataResource UserResource

	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"invalid user data",
			500,
		)
		return
	}
	userModel := dataResource.Data
	ctx := NewContext()
	defer ctx.Close()
	c := ctx.DbCollection(dbCollectionUser)
	repo := &data.UserRepository{C: c}

	if err := repo.Delete(&userModel); err != nil {
		common.DisplayAppError(
			w,
			err,
			"invalid user credentials",
			401,
		)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
