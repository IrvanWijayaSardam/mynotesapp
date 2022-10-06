package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/IrvanWijayaSardam/mynotesapp/dto"
	"github.com/IrvanWijayaSardam/mynotesapp/entity"
	"github.com/IrvanWijayaSardam/mynotesapp/helper"
	"github.com/IrvanWijayaSardam/mynotesapp/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type NoteController interface {
	All(context *gin.Context)
	FindById(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type noteController struct {
	noteService service.NoteService
	jwtService  service.JWTService
}

func NewNotesController(noteserv service.NoteService, jwtServ service.JWTService) NoteController {
	return &noteController{
		noteService: noteserv,
		jwtService:  jwtServ,
	}
}

func (c *noteController) All(context *gin.Context) {
	var notes []entity.Notes = c.noteService.All()
	res := helper.BuildResponse(true, "OK!", notes)
	context.JSON(http.StatusOK, res)
}

func (c *noteController) FindById(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("user_id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No Parameter ID was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	var notes entity.Notes = c.noteService.FindByID(id)
	if (notes == entity.Notes{}) {
		res := helper.BuildErrorResponse("Data Not Found", "No Data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", notes)
		context.JSON(http.StatusOK, res)
	}
}

func (c *noteController) Insert(context *gin.Context) {
	var noteCreateDTO dto.NotesCreateDTO
	errDTO := context.ShouldBind(&noteCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		userID := c.getUserIDByToken(authHeader)
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			noteCreateDTO.UserID = convertedUserID
		}
		result := c.noteService.Insert(noteCreateDTO)
		response := helper.BuildResponse(true, "OK!", result)
		context.JSON(http.StatusCreated, response)
	}
}

func (c *noteController) Update(context *gin.Context) {
	var noteUpdateDTO dto.NotesUpdateDTO
	errDTO := context.ShouldBind(&noteUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["userid"])
	if c.noteService.IsAllowedToEdit(userID, noteUpdateDTO.ID) {
		id, errID := strconv.ParseUint(userID, 10, 64)
		if errID == nil {
			noteUpdateDTO.UserID = id
		}
		result := c.noteService.Update(noteUpdateDTO)
		response := helper.BuildResponse(true, "OK!", result)
		context.JSON(http.StatusOK, response)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *noteController) Delete(context *gin.Context) {
	var notes entity.Notes
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed tou get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	notes.ID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["userid"])
	if c.noteService.IsAllowedToEdit(userID, notes.ID) {
		c.noteService.Delete(notes)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *noteController) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["userid"])
	return id
}
