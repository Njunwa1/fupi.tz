package api

import (
	"context"
	"github.com/Njunwa1/fupi.tz/auth/internal/application/core/domain"
	"github.com/Njunwa1/fupi.tz/auth/internal/ports"
	"log/slog"
)

type Application struct {
	db    ports.DBPort
	maker ports.PasetoPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{db: db}
}

func (a *Application) Register(ctx context.Context, user domain.User) (domain.User, error) {
	hashedPassword, err := user.HashPassword(user.Password)
	if err != nil {
		slog.Error("Error while hashing password: ", err)
		return domain.User{}, err
	}
	newUser := domain.NewUser(
		user.Names,
		user.Company,
		hashedPassword,
		user.Email,
		user.PhoneNumber,
		user.Role,
	)
	err = a.db.SaveUser(ctx, *newUser)
	if err != nil {
		slog.Error("Error while saving user: ", err)
		return domain.User{}, err
	}
	return *newUser, nil
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
	token, err := a.maker.CreateToken(user.ID, user.Role.Name, 24)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (a *Application) VerifyToken(token string) (*domain.Payload, error) {
	return a.maker.VerifyToken(token)
}
