package controller

import (
	"Ecommerce-User/dto"
	"Ecommerce-User/helper"
	"Ecommerce-User/service"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserController interface {
	Register(context *gin.Context)
	Login(context *gin.Context)
	Update(context *gin.Context)
	Profile(context *gin.Context)
}

type userController struct {
	userService service.UserService
	jwtService  service.JWTService
}

func NewUserController(userService service.UserService, jwtService service.JWTService) UserController {
	return &userController{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (controller *userController) Register(context *gin.Context) {
	var registerUserDTO dto.RegisterUserDTO
	err := context.ShouldBind(&registerUserDTO)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if !controller.userService.IsDuplicateEmail(registerUserDTO.Email) {
		response := helper.BuildErrorResponse("Failed to process request", "Email already exists", helper.EmptyObj{})
		context.JSON(http.StatusConflict, response)
	} else {
		createdUser := controller.userService.Register(registerUserDTO)
		token := controller.jwtService.GenerateToken(strconv.FormatUint(createdUser.ID, 10))
		createdUser.Token = token
		response := helper.BuildResponse(true, "User created successfully", createdUser)
		context.JSON(http.StatusCreated, response)
	}
}

func (controller *userController) Login(context *gin.Context) {
	var loginUserDTO dto.LoginUserDTO
	err := context.ShouldBind(&loginUserDTO)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	user := controller.userService.Login(loginUserDTO)
	if user.ID == 0 {
		response := helper.BuildErrorResponse("Email or password is incorrect", "Failed to process request", helper.EmptyObj{})
		context.JSON(http.StatusUnauthorized, response)
		return
	}
	token := controller.jwtService.GenerateToken(strconv.FormatUint(user.ID, 10))
	user.Token = token
	response := helper.BuildResponse(true, "User logged in successfully", user)
	context.JSON(http.StatusOK, response)
}

func (controller *userController) Update(context *gin.Context) {
	var updateUserDTO dto.UpdateUserDTO
	errDTO := context.ShouldBind(&updateUserDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authHeader := context.GetHeader("Authorization")
	token, errToken := controller.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		response := helper.BuildErrorResponse("Failed to process request", errToken.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	updateUserDTO.Id = id
	updatedUser := controller.userService.Update(updateUserDTO)
	res := helper.BuildResponse(true, "User updated successfully", updatedUser)
	context.JSON(http.StatusOK, res)
}

func (controller *userController) Profile(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	token, errToken := controller.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		response := helper.BuildErrorResponse("Failed to process request", errToken.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	id, errClaims := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	if errClaims != nil {
		response := helper.BuildErrorResponse("Failed to process request", errClaims.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	idString := strconv.FormatUint(id, 10)
	user := controller.userService.Profile(idString)
	res := helper.BuildResponse(true, "User profile retrieved successfully", user)
	context.JSON(http.StatusOK, res)

}
