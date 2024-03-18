package api

import (
	"context"
	"github.com/Njunwa1/fupi.tz/auth/internal/application/core/domain"
	"github.com/Njunwa1/fupi.tz/auth/internal/ports"
	"github.com/Njunwa1/fupitz-proto/golang/user"
	"golang.org/x/crypto/bcrypt"
)

type Application struct {
	db    ports.DBPort
	maker ports.PasetoPort
}

func NewApplication(db ports.DBPort, pasetoMaker ports.PasetoPort) *Application {
	return &Application{db: db, maker: pasetoMaker}
}

func (a *Application) Register(ctx context.Context, request *user.RegisterRequest) (*user.RegisterResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return &user.RegisterResponse{}, err
	}
	newUser := domain.NewUser(
		request.Names,
		request.Company,
		request.Email,
		request.PhoneNumber,
		string(hashedPassword),
		domain.Role{Name: "user"},
	)
	err = a.db.SaveUser(ctx, *newUser)
	if err != nil {
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
	userAccount, err := a.db.GetUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}
	err = userAccount.VerifyPassword(userAccount.Password, password)
	if err != nil {
		return "", err
	}
	token, err := a.maker.CreateToken(userAccount.ID.Hex(), userAccount.Role.Name, 24)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (a *Application) VerifyToken(token string) (*domain.Payload, error) {
	return a.maker.VerifyToken(token)
}
