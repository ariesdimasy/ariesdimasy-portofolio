package controllers

import "ariesdimasy-portofolio/repositories"

type UserController struct {
	UserRepository repositories.UserRepository
}

func NewUserController(userRepository repositories.UserRepository) *UserController {
	return &UserController{UserRepository: userRepository}
}
