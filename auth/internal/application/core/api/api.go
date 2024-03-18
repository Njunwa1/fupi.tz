package api

import (
	"context"
	"github.com/Njunwa1/fupi.tz/auth/internal/application/core/domain"
	"github.com/Njunwa1/fupi.tz/auth/internal/ports"
	"github.com/Njunwa1/fupitz-proto/golang/user"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
)

type Application struct {
	db    ports.DBPort
	maker ports.PasetoPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{db: db}
}

func (a *Application) Register(ctx context.Context, request *user.RegisterRequest) (*user.RegisterResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		slog.Error("Error while hashing password: ", err)
		return &user.RegisterResponse{}, err
	}
	newUser := domain.NewUser(
		request.Names,
		request.Company,
		string(hashedPassword),
		request.Email,
		request.PhoneNumber,
		domain.Role{Name: "user"},
	)
	err = a.db.SaveUser(ctx, *newUser)
	if err != nil {
		slog.Error("Error while saving user: ", err)
		return &user.RegisterResponse{}, err
	}
	return &user.RegisterResponse{
		Names:       newUser.Names,
		Company:     newUser.Company,
		Email:       newUser.Email,
		PhoneNumber: newUser.PhoneNumber,
	}, nil
}

func (a *Application) Login(ctx context.Context, email, password string) (string, error) {
	user, err := a.db.GetUserByEmail(ctx, email)
	if err != nil {
		slog.Error("Error while getting user by email: ", err)
		return "", err
	}
	err = user.VerifyPassword(user.Password, password)
	if err != nil {
		return "", err
	}
	token, err := a.maker.CreateToken(user.ID.Hex(), user.Role.Name, 24)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (a *Application) VerifyToken(token string) (*domain.Payload, error) {
	return a.maker.VerifyToken(token)
}
