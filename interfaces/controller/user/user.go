package user

import (
	"net/http"
	"room_app_back/config"
	"room_app_back/usecase/user"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	uu  user.IUserUsecase
	cnf *config.AppConfig
}

func NewUserController(uu user.IUserUsecase, cnf *config.AppConfig) *UserController {
	return &UserController{
		uu:  uu,
		cnf: cnf,
	}
}

func (uc *UserController) Index(c echo.Context) error {
	users, err := uc.uu.Users()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, users)
}

func (uc *UserController) Show(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	user, err := uc.uu.UserById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, user)
}

func (uc *UserController) Create(c echo.Context) error {
	u := user.AddInput{}
	if err := c.Bind(&u); err != nil {
		return err
	}
	output, err := uc.uu.Add(&u)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, output)
}

func (uc *UserController) Save(c echo.Context) error {
	input := user.AddInput{}
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	output, err := uc.uu.Update(&input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, output)
}

func (uc *UserController) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := uc.uu.DeleteById(id); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.NoContent(http.StatusNoContent)
}

func (uc *UserController) SignUp(c echo.Context) error {
	input := user.AddInput{}
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	output, err := uc.uu.SignUp(&input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, output)
}

func (uc *UserController) LogIn(c echo.Context) error {
	input := user.AddInput{}
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	tokenString, err := uc.uu.Login(&input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	cookie := http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  time.Now().Add(24 * time.Hour),
		Path:     "/",
		Domain:   uc.cnf.ApiDomain,
		Secure:   true, //TODO:postman確認時コメントアウト
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	}
	c.SetCookie(&cookie)
	return c.NoContent(http.StatusOK)

}

func (uc *UserController) LogOut(c echo.Context) error {
	cookie := http.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now(),
		Path:     "/",
		Domain:   uc.cnf.ApiDomain,
		Secure:   true, //TODO:postman確認時コメントアウト
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	}
	c.SetCookie(&cookie)
	return c.NoContent(http.StatusOK)

}

func (uc *UserController) CsrfToken(c echo.Context) error {
	token, ok := c.Get("csrf").(string)
	if !ok {
		return c.JSON(http.StatusNotFound, echo.Map{
			"error": "Resource not found",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"csrf_token": token,
	})
}
