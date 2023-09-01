package user

import (
	"backend/config"
	"backend/domain/model"
	"backend/domain/repository/user"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	Add(input *AddInput) (*AddOutput, error)
	Update(input *AddInput) (*AddOutput, error)
	DeleteById(id int) error
	Users() (*AddOutputs, error)
	UserById(id int) (*AddOutput, error)
	SignUp(input *AddInput) (*AddOutput, error)
	Login(input *AddInput) (string, error)
}

type UserUsecase struct {
	ur  user.UserRepository
	cnf *config.AppConfig
}

func NewUserUsecase(ur user.UserRepository, cnf *config.AppConfig) IUserUsecase {
	return &UserUsecase{
		ur:  ur,
		cnf: cnf,
	}
}

func (uu *UserUsecase) Add(input *AddInput) (*AddOutput, error) {
	output, err := uu.ur.Store(&model.User{
		ID:        input.ID,
		Email:     input.Email,
		LastName:  input.LastName,
		FirstName: input.FirstName,
		Password:  input.Password,
		Age:       input.Age,
		Role:      input.Role,
		IdNumber:  input.IdNumber,
		CreatedAt: input.CreatedAt,
		UpdatedAt: input.UpdatedAt,
	})
	if err != nil {
		return nil, err
	}
	return &AddOutput{output}, err
}

func (uu *UserUsecase) Update(input *AddInput) (*AddOutput, error) {
	output, err := uu.ur.Update(&model.User{
		ID:        input.ID,
		Email:     input.Email,
		LastName:  input.LastName,
		FirstName: input.FirstName,
		Password:  input.Password,
		Age:       input.Age,
		Role:      input.Role,
		IdNumber:  input.IdNumber,
		CreatedAt: input.CreatedAt,
		UpdatedAt: input.UpdatedAt,
	})
	if err != nil {
		return nil, err
	}
	return &AddOutput{output}, err
}

func (uu *UserUsecase) DeleteById(id int) error {
	return uu.ur.DeleteById(id)
}

func (uu *UserUsecase) Users() (*AddOutputs, error) {
	users, err := uu.ur.FindAll()
	if err != nil {
		return nil, err
	}
	var outputs AddOutputs
	for _, _u := range *users {
		u := _u
		outputs = append(outputs, AddOutput{&u})
	}
	return &outputs, nil
}

func (uu *UserUsecase) UserById(id int) (*AddOutput, error) {
	user, err := uu.ur.FindById(id)
	if err != nil {
		return nil, err
	}

	return &AddOutput{User: user}, err
}

func (uu *UserUsecase) SignUp(input *AddInput) (*AddOutput, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
	if err != nil {
		return nil, err
	}
	newUser := model.User{
		Email:    input.Email,
		Password: string(hash),
	}

	resUser, err := uu.ur.Store(&newUser)
	if err != nil {
		return nil, err
	}
	return &AddOutput{User: resUser}, nil

}

func (uu *UserUsecase) Login(input *AddInput) (string, error) {
	storedUser, err := uu.ur.FindByEmail(input.Email)
	if err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(input.Password)); err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})
	tokenString, err := token.SignedString([]byte(uu.cnf.SecretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
