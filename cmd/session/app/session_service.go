package app

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"mock-grpc/pkg/proto"
	"time"
)

type SessionService struct {
	SecretKey string
}

func (s *SessionService) CreateSession(ctx context.Context, request *proto.CreateSessionRequest) (*proto.CreateSessionResponse, error) {
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	token.Claims = jwt.MapClaims{
		"username": request.Username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	}
	signedToken, err := token.SignedString([]byte(s.SecretKey))
	if err != nil {
		return nil, err
	}
	return &proto.CreateSessionResponse{Token: signedToken}, nil
}

func (s *SessionService) VerifySession(ctx context.Context, request *proto.VerifySessionRequest) (*proto.VerifySessionResponse, error) {
	_, err := jwt.Parse(request.Token, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.SecretKey), nil
	})
	return &proto.VerifySessionResponse{IsValid: err == nil}, nil
}

func (s *SessionService) GetSessionRoles(ctx context.Context, request *proto.GetUserRolesRequest) (*proto.GetUserRolesResponse, error) {
	panic("implement me")
}
