package service

import (
	"Ecommerce-User/dto"
	"Ecommerce-User/model"
	"Ecommerce-User/repository"
	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type UserService interface {
	FindByEmail(email string) model.User
	IsDuplicateEmail(email string) bool
	Register(user dto.RegisterUserDTO) model.User
	Login(user dto.LoginUserDTO) model.User
	Update(user dto.UpdateUserDTO) model.User
	Profile(userID string) model.User
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (service *userService) FindByEmail(email string) model.User {
	return service.userRepository.FindByEmail(email)
}

func (service *userService) IsDuplicateEmail(email string) bool {
	res := service.userRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

func (service *userService) Register(user dto.RegisterUserDTO) model.User {
	userDTOtoModel := model.User{}
	err := smapping.FillStruct(&userDTOtoModel, smapping.MapFields(&user))
	if err != nil {
		return model.User{}
	}
	res := service.userRepository.InsertUser(userDTOtoModel)
	return res
}

func (service *userService) Login(user dto.LoginUserDTO) model.User {
	res := service.userRepository.VerifyCredential(user.Email)
	if v, ok := res.(model.User); ok {
		comparePassword := comparePassword(v.Password, []byte(user.Password))
		if v.Email == user.Email && comparePassword {
			return v
		}
		return model.User{
			ID: 0,
		}
	}
	return model.User{}
}

func (service *userService) Update(user dto.UpdateUserDTO) model.User {
	userDTOtoModel := model.User{}
	err := smapping.FillStruct(&userDTOtoModel, smapping.MapFields(&user))
	if err != nil {
		return model.User{}
	}
	userDTOtoModel.ID = user.Id
	res := service.userRepository.UpdateUser(userDTOtoModel)
	log.Println(userDTOtoModel)
	return res
}

func (service *userService) Profile(userID string) model.User {
	return service.userRepository.ProfileUser(userID)
}

func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		return false
	}
	return true
}
