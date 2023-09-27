package reservation

import (
	"backend/config"
	"backend/domain/model"
	"backend/domain/repository/reservation"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

type IReservationUsecase interface {
	Add(input *AddInput, id uint) (*AddOutput, error)
	Update(input *AddInput) (*AddOutput, error)
	DeleteById(id int) error
	Reservations() (*AddOutputs, error)
	ReservationById(id int) (*AddOutput, error)
	JwtToUserId(tokenString string) (uint, error)
}

type ReservationUsecase struct {
	rr  reservation.ReservationRepository
	cnf *config.AppConfig
}

func NewReservationUsecase(rr reservation.ReservationRepository, cnf *config.AppConfig) IReservationUsecase {
	return &ReservationUsecase{
		rr:  rr,
		cnf: cnf}
}

func (ru *ReservationUsecase) Add(input *AddInput, userId uint) (*AddOutput, error) {
	output, err := ru.rr.Store(&model.Reservation{
		ID:        input.ID,
		Purpose:   input.Purpose,
		StartTime: input.StartTime,
		EndTime:   input.EndTime,
		UserID:    userId,
		RoomID:    input.RoomID,
		CreatedAt: input.CreatedAt,
		UpdatedAt: input.UpdatedAt,
		User:      input.User,
		Room:      input.Room,
	})
	if err != nil {
		return nil, err
	}
	return &AddOutput{output}, nil
}

func (ru *ReservationUsecase) Update(input *AddInput) (*AddOutput, error) {
	output, err := ru.rr.Update(&model.Reservation{
		ID:        input.ID,
		Purpose:   input.Purpose,
		StartTime: input.StartTime,
		EndTime:   input.EndTime,
		UserID:    input.UserID,
		RoomID:    input.RoomID,
		CreatedAt: input.CreatedAt,
		UpdatedAt: input.UpdatedAt,
		User:      input.User,
		Room:      input.Room,
	})
	if err != nil {
		return nil, err
	}
	return &AddOutput{output}, nil
}

func (ru *ReservationUsecase) DeleteById(id int) error {
	return ru.rr.DeleteById(id)
}

func (ru *ReservationUsecase) Reservations() (*AddOutputs, error) {
	users, err := ru.rr.FindAll()
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

func (ru *ReservationUsecase) ReservationById(id int) (*AddOutput, error) {
	output, err := ru.rr.FindById(id)
	if err != nil {
		return nil, err
	}
	return &AddOutput{output}, nil
}

func (ru *ReservationUsecase) JwtToUserId(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(ru.cnf.SecretKey), nil
	})
	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("invalid token claims")
	}

	userID, ok := claims["user_id"].(float64)
	if !ok {
		return 0, fmt.Errorf("user_id not found in token or not a number")
	}

	return uint(userID), nil
}
