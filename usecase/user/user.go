package user

import (
	"os"
	"room_app_back/domain/model"
	"room_app_back/domain/repository/user"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	Add(u *model.User) (*model.User, error)
	Update(u *model.User) (*model.User, error)
	DeleteById(id int) error
	Users() (*model.Users, error)
	UserById(id int) (*model.User, error)
	SignUp(u *model.User) (*model.User, error)
	Login(u *model.User) (string, error)
}

type UserUsecase struct {
	ur user.UserRepository
}

func NewUserUsecase(ur user.UserRepository) *UserUsecase {
	return &UserUsecase{ur}
}

func (uu *UserUsecase) Add(u *model.User) (*model.User, error) {
	return uu.ur.Store(u)
}

func (uu *UserUsecase) Update(u *model.User) (*model.User, error) {
	return uu.ur.Update(u)
}

func (uu *UserUsecase) DeleteById(id int) error {
	return uu.ur.DeleteById(id)
}

func (uu *UserUsecase) Users() (*model.Users, error) {
	return uu.ur.FindAll()
}

func (uu *UserUsecase) UserById(id int) (*model.User, error) {
	return uu.ur.FindById(id)
}

func (uu *UserUsecase) SignUp(u *model.User) (*model.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		return nil, err
	}
	newUser := model.User{
		Email:    u.Email,
		Password: string(hash),
	}

	resUser, err := uu.ur.Store(&newUser)
	if err != nil {
		return &model.User{}, err
	}
	return resUser, nil

}

func (uu *UserUsecase) Login(u *model.User) (string, error) {
	storedUser := model.User{}
	if err := uu.ur.FindByEmail(u.Email); err != nil {
		return "", nil
	}
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(u.Password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", nil
	}
	return tokenString, nil
}
