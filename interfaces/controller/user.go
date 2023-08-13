package controller

import (
	"room_app_back/domain/model"
	"room_app_back/pkg"
	"room_app_back/usecases/interactors"
	"strconv"

	"github.com/labstack/echo"
)

type UserController struct {
	ui interactors.UserInteractor
}

func NewUserController(userInteractor interactors.UserInteractor) *UserController {
	return &UserController{
		ui: userInteractor,
	}
}

func (uc *UserController) Index(c echo.Context) error {
	users, err := uc.ui.Users()
	if err != nil {
		return c.JSON(500, pkg.NewError(err))
	}

	return c.JSON(200, users)
}

func (uc *UserController) Show(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := uc.ui.UserById(id)
	if err != nil {
		return c.JSON(500, pkg.NewError(err))
	}
	return c.JSON(200, user)
}

func (uc *UserController) Create(c echo.Context) error {
	u := model.User{}
	c.Bind(&u)
	user, err := uc.ui.Add(u)
	if err != nil {
		return c.JSON(500, pkg.NewError(err))
	}
	return c.JSON(201, user)
}

func (uc *UserController) Save(c echo.Context) error {
	u := model.User{}
	c.Bind(&u)
	user, err := uc.ui.Update(u)
	if err != nil {
		return c.JSON(500, pkg.NewError(err))
	}
	return c.JSON(201, user)
}

func (uc *UserController) Delete(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	user := model.User{
		ID: uint(id),
	}
	err := uc.ui.DeleteById(user)
	if err != nil {
		return c.JSON(500, pkg.NewError(err))
	}
	return c.JSON(201, "deleted")
}
