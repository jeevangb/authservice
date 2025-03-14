package services

import (
	"context"
	"errors"

	"github.com/jeevangb/authservice/internal/grpc/proto"
	"github.com/jeevangb/authservice/internal/models"
	"github.com/jeevangb/authservice/internal/utils"
	"gorm.io/gorm"
)

func (s *GrpcImpl) SignUp(ctx context.Context, req *proto.SignUpRequest) (*proto.UserResponse, error) {
	exisUser, err := s.usrsrc.FindUserByEmail(req.Email)
	if err == nil && exisUser.Email != "" {
		return &proto.UserResponse{
			Message: "User already exist",
		}, nil
	}
	hashPass, err := utils.HashPassword(req.Password)
	if err != nil {
		return &proto.UserResponse{
			Message: "internal server error",
		}, err
	}
	user := &models.User{
		Email:    req.Email,
		Password: string(hashPass),
		Name:     req.Name,
	}
	if err = s.usrsrc.Createuser(*user); err != nil {
		return &proto.UserResponse{
			Message: "internal system error",
		}, nil
	}

	return &proto.UserResponse{
		Message: "SignUp successful",
	}, nil
}

func (s *GrpcImpl) Login(ctx context.Context, req *proto.LoginRequest) (*proto.UserResponse, error) {
	// var user models.User
	user, err := s.usrsrc.FindUserByEmail(req.Email)
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, errors.New("user not exists")
	}
	if !utils.CheckPasswordHash(user.Password, req.Password) {
		return &proto.UserResponse{
			Message: "Inavlid password , Please try again",
		}, nil
	}
	return &proto.UserResponse{
		Message: "Login Successfull",
	}, nil
}
