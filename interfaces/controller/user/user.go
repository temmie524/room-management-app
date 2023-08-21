package user

import (
	"room_app_back/domain/model"
	"room_app_back/pkg"
	"room_app_back/usecase/user"
	"strconv"

	"github.com/labstack/echo"
)

type UserController struct {
	uu user.UserUsecase
}

func NewUserController(uu user.UserUsecase) *UserController {
	return &UserController{
		uu: uu,
	}
}

func (uc *UserController) Index(c echo.Context) error {
	users, err := uc.uu.Users()
	if err != nil {
		return c.JSON(500, pkg.NewError(err))
	}

	return c.JSON(200, users)
}

func (uc *UserController) Show(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := uc.uu.UserById(id)
	if err != nil {
		return c.JSON(500, pkg.NewError(err))
	}
	return c.JSON(200, user)
}

func (uc *UserController) Create(c echo.Context) error {
	u := model.User{}
	c.Bind(&u)
	user, err := uc.uu.Add(u)
	if err != nil {
		return c.JSON(500, pkg.NewError(err))
	}
	return c.JSON(201, user)
}

func (uc *UserController) Save(c echo.Context) error {
	u := model.User{}
	c.Bind(&u)
	user, err := uc.uu.Update(u)
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
	err := uc.uu.DeleteById(user)
	if err != nil {
		return c.JSON(500, pkg.NewError(err))
	}
	return c.JSON(201, "deleted")
}
