package user

import (
	"net/http"
	"os"
	"room_app_back/domain/model"
	"room_app_back/pkg"
	"room_app_back/usecase/user"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
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

func (uc *UserController) SignUp(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(500, pkg.NewError(err))
	}
	userRes, err := uc.uu.SignUp(user)
	if err != nil {
		return c.JSON(500, pkg.NewError(err))
	}
	return c.JSON(201, userRes)
}

func (uc *UserController) LogIn(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(500, pkg.NewError(err))
	}
	tokenString, err := uc.uu.Login(user)
	if err != nil {
		return c.JSON(500, pkg.NewError(err))
	}
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true //postman確認時コメントアウト
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)

}

func (uc *UserController) LogOut(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Expires = time.Now()
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true //postman確認時コメントアウト
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)

}

func (uc *UserController) CsrfToken(c echo.Context) error {
	token := c.Get("csrf").(string)
	return c.JSON(http.StatusOK, echo.Map{
		"csrf_token": token,
	})
}
