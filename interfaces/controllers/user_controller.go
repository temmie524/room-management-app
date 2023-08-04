package controllers

import (
	"room_app_back/domain/model"
	"room_app_back/interfaces/database"
	"room_app_back/usecases"

	"strconv"

	"github.com/labstack/echo"
)

type UserController struct {
	Intetacter usecases.UserInteractor
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		Intetacter: usecases.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *UserController) Index(c echo.Context) (err error) {
	users, err := controller.Intetacter.Users()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, users)
	return
}

func (controller *UserController) Show(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := controller.Intetacter.UserById(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, user)
	return
}

func (controller *UserController) Create(c echo.Context) (err error) {
	u := model.User{}
	c.Bind(&u)
	user, err := controller.Intetacter.Add(u)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, user)
	return
}

func (controller *UserController) Save(c echo.Context) (err error) {
	u := model.User{}
	c.Bind(&u)
	user, err := controller.Intetacter.Update(u)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, user)
	return
}

func (controller *UserController) Delete(c echo.Context) (err error) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	user := model.User{
		ID: uint(id),
	}
	err = controller.Intetacter.DeleteById(user)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, "deleted")
	return
}
