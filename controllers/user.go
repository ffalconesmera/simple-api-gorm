package controllers

import (
	"net/http"
	"strconv"

	"simple-api-gorm/config"
	"simple-api-gorm/models"
	"simple-api-gorm/responses"
	"simple-api-gorm/services"

	"github.com/gin-gonic/gin"
)

type User struct {
	responses.Response
}

func (a *User) Index(c *gin.Context) {
	var params *models.User
	c.ShouldBindJSON(&params)
	result, list := new(services.Users).All(params)
	a.Json(c, result, list)
}

func (a *User) Store(c *gin.Context) {
	var params *models.User
	c.ShouldBindJSON(&params)

	validateUser := params.ValidateInsert()

	if validateUser != nil {
		a.JsonFail(c, 504, validateUser.Error())
		return
	}

	result, data := new(services.Users).Create(params)
	a.Json(c, result, data)
}

func (a *User) Update(c *gin.Context) {
	var params *models.User
	c.ShouldBindJSON(&params)

	validateUser := params.ValidateBase()

	if validateUser != nil {
		a.JsonFail(c, 504, validateUser.Error())
		return
	}

	result, data := new(services.Users).Update(params)
	a.Json(c, result, data)
}

func (a *User) Destroy(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err == nil {
		result, data := new(services.Users).Delete(id)
		a.Json(c, result, data)
		return
	}
	a.JsonFail(c, http.StatusBadRequest, "Not have params")
}

func (a *User) Show(c *gin.Context) {
	var account models.User
	result := config.Db().First(&account, c.Params.ByName("id"))
	a.Json(c, result, account)
}
