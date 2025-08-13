package gapi

import (
	db "github.com/AtoyanMikhai/bank_api/db/sqlc"
	"github.com/AtoyanMikhai/bank_api/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user db.User) *pb.User {
	return &pb.User{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordChangedAt: timestamppb.New(user.PasswordChagedAt),
		CreatedAt:         timestamppb.New(user.CreatedAt),
	}
}
