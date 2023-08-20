package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
	pb "web3/api/service/dex-member"
	"web3/app/service/dex-member/internal/dao"
)

var TokenInvalid = errors.New(403, "AUTH_FAIL", "Token is invalid")

type JwtTokenClaims struct {
	UserID int64 `json:"user_id"`
	jwt.RegisteredClaims
}

const _sKey = "mze7yty.ZHZ1wfm_wyk"

func (s *Service) AuthSignIn(ctx context.Context, req *pb.AuthSignInRequest) (*pb.AuthSignInReply, error) {
	email := req.GetEmail()
	user, err := dao.GlobalDao.MySQLGetUserByEmail(ctx, email)
	if err != nil {
		return nil, errors.New(500, "SYSTEM_ERR", err.Error())
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, &JwtTokenClaims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "https://leek.dex",
			Subject:   "Dex",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(365 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        uuid.NewString(),
		},
	}).SignedString([]byte(_sKey))
	resp := &pb.AuthSignInReply{
		JwtToken: token,
	}
	return resp, nil
}

func (s *Service) AuthSignUp(ctx context.Context, req *pb.AuthSignUpRequest) (*pb.AuthSignUpReply, error) {
	return nil, errors.New(500, "SYSTEM_ERR", "not support this opt")
}
func (s *Service) AuthVerify(ctx context.Context, req *pb.AuthVerifyRequest) (*pb.AuthVerifyReply, error) {
	resp := &pb.AuthVerifyReply{
		User: &pb.UserItem{},
	}
	token, err := _parseJwtToken(req.GetJwtToken())
	if err != nil {
		return nil, TokenInvalid
	}
	user, err := dao.GlobalDao.MySQLGetUser(ctx, token.UserID)
	if err != nil {
		return nil, TokenInvalid
	}
	cTime := user.CreatedAt
	uTime := user.UpdatedAt
	resp.User.Id = user.ID
	resp.User.NickName = user.NickName
	resp.User.AvatarUrl = user.AvatarURL
	resp.User.Email = user.Email
	resp.User.CreatedAt = timestamppb.New(*cTime)
	resp.User.UpdatedAt = timestamppb.New(*uTime)

	return resp, nil
}

func _parseJwtToken(token string) (*JwtTokenClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &JwtTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(_sKey), nil
	})
	if err != nil {
		return nil, err
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*JwtTokenClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
