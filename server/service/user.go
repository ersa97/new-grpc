package service

import (
	"context"
	"errors"

	"github.com/ersa97/new-grpc/server/data"
	"github.com/ersa97/new-grpc/server/utils"
	"github.com/segmentio/ksuid"
)

type UserService struct {
}

type User struct {
	Id       *string
	Name     string
	Email    string
	Password string
}

var USERS = []User{}

func (s *UserService) AddUser(ctx context.Context, req *data.AddUserRequest) (*data.AddUserResponse, error) {
	//validate the user that use the add action
	userid, err := utils.Verify(req.GetAccessToken())

	//catch if there is a mistake in verifying the token
	if err != nil {
		return nil, errors.New("token is invalid")
	}

	//move it to local variable
	var user *data.User
	for _, v := range USERS {
		if v.Id == userid {
			user = &data.User{
				Id:       v.Id,
				Name:     v.Name,
				Email:    v.Email,
				Password: v.Password,
			}
			break
		}
	}
	/*if the user is not found from verifying the JWT token
	then the user is not authorized*/
	if user == nil {
		return nil, errors.New("user is unauthorized")
	}

	//create random id and encrypted password to be stored
	id := ksuid.New().String()
	pass, _ := utils.Encryption(req.User.Password)

	//move all of the data that will be added to a new variable
	newUser := User{
		Id:       &id,
		Name:     req.User.Name,
		Email:    req.User.Email,
		Password: pass,
	}

	//add the data into the array of struct
	USERS = append(USERS, newUser)

	return &data.AddUserResponse{
		Message: "Add User Successful",
		User: &data.User{
			Id:       newUser.Id,
			Name:     newUser.Name,
			Email:    newUser.Email,
			Password: newUser.Password,
		},
	}, nil
}

func (s *UserService) Login(ctx context.Context, req *data.LoginRequest) (*data.LoginResponse, error) {
	var user *data.User
	/*comparing the email and the encrypted password,
	if both of them true then insert it to the local variable*/
	for _, v := range USERS {
		if v.Email == req.Email && utils.Compare(v.Password, req.Password) {
			user = &data.User{
				Id:       v.Id,
				Name:     v.Name,
				Email:    v.Email,
				Password: v.Password,
			}
			break
		}
	}
	//if user not found then maybe they input them wrong
	if user == nil {
		return nil, errors.New("username or password is incorrect")
	}
	//creating token for the user to use inside the app
	token, err := utils.CreateToken(*user.Id)
	if err != nil {
		return nil, errors.New("cannot create token for user")
	}
	return &data.LoginResponse{
		AccessToken: token.AccessToken,
	}, nil
}

func (s *UserService) GetUsers(ctx context.Context, req *data.GetUsersRequest) (*data.GetUsersResponse, error) {
	//create a local array of struct to store all of the data
	var users []*data.User
	for _, v := range USERS {
		users = append(users, &data.User{
			Id:       v.Id,
			Name:     v.Name,
			Email:    v.Email,
			Password: v.Password,
		})
	}
	return &data.GetUsersResponse{
		Message: "Get All Users",
		User:    users,
	}, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *data.UpdateUserRequest) (*data.UpdateUserResponse, error) {
	//check who is doing the update action
	userid, err := utils.Verify(req.GetAccessToken())
	if err != nil {
		return nil, errors.New("token is invalid")
	}

	//move it to local variable
	var user *data.User

	for _, v := range USERS {
		if v.Id == userid {
			user = &data.User{
				Id:       v.Id,
				Name:     v.Name,
				Email:    v.Email,
				Password: v.Password,
			}
			break
		}
	}
	/*check if the user is authorized,
	if user doesn't exist then the token verify that they are not authorized*/
	if user == nil {
		return nil, errors.New("user unauthorized")
	}

	//encrypt the password with RSA before stored
	pass, _ := utils.Encryption(req.User.Password)

	//storing updated data in local variable
	newUser := &User{
		Id:       req.User.Id,
		Name:     req.User.Name,
		Email:    req.User.Email,
		Password: pass,
	}

	//change the existing user by id in the array of struct
	for i, v := range USERS {
		if v.Id == newUser.Id {
			USERS = append(USERS[:i-1], *newUser)
		}
	}

	return &data.UpdateUserResponse{
		Message: "Modify User Successful",
		User: &data.User{
			Id:       newUser.Id,
			Name:     newUser.Name,
			Email:    newUser.Email,
			Password: newUser.Password,
		},
	}, nil

}

func (s *UserService) DeleteUser(ctx context.Context, req *data.DeleteUserRequest) (*data.DeleteUserResponse, error) {
	// get the user who is doing the delete action
	userid, err := utils.Verify(req.GetAccessToken())

	if err != nil {
		return nil, errors.New("token is invalid")
	}

	var user *data.User
	//move it to a local variable
	for _, v := range USERS {
		if userid == v.Id {
			user = &data.User{
				Id:       v.Id,
				Name:     v.Name,
				Email:    v.Email,
				Password: v.Password,
			}
			break
		}
	}
	/*check if the user is authorized,
	if user doesn't exist then the token verify that they are not authorized*/
	if user == nil {
		return nil, errors.New("user is unauthorized to delete another user")
	}

	//delete the user by id from the array of struct
	for i, v := range USERS {
		if v.Id == req.User.Id {
			USERS = append(USERS[:i], USERS[i+1:]...)
			break
		}
	}

	var users []*data.User

	//get all existing user
	for _, v := range USERS {
		users = append(users, &data.User{
			Id:       v.Id,
			Name:     v.Name,
			Email:    v.Email,
			Password: v.Password,
		})
	}

	return &data.DeleteUserResponse{
		Message: "Delete Successful",
		User:    users,
	}, nil

}
