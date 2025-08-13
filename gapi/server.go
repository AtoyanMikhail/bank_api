package gapi

import (
	"fmt"

	db "github.com/AtoyanMikhai/bank_api/db/sqlc"
	"github.com/AtoyanMikhai/bank_api/pb"
	"github.com/AtoyanMikhai/bank_api/token"
	"github.com/AtoyanMikhai/bank_api/util"
	"github.com/AtoyanMikhai/bank_api/worker"
)

// Server serves gRPC requests for banking service
type Server struct {
	pb.UnimplementedBasicBankServer
	config          util.Config
	store           db.Store
	tokenMaker      token.Maker
	taskDistributor worker.TaskDistributor
}

// NewServer creates a new HTTP server and setup routing
func NewServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:          config,
		store:           store,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
	}

	return server, nil
}
