package main

import (
	"context"
	"crud/app/database"
	"crud/app/rpc"
	"crud/app/utils"
	"fmt"
	"net"
	"os"

	pb "crud/app/proto"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	logger := utils.InitLogger()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
	if err != nil {
		logger.Fatal("error while listening tcp", zap.Error(err))
	}

	s := grpc.NewServer()
	pgInstance, err := database.NewPG(context.Background(), os.Getenv("POSTGRES_DSN"), logger)
	if err != nil {
		logger.Fatal("error while connecting to DB", zap.Error(err))
	}
	pb.RegisterCrudServer(s, &rpc.Server{PG: pgInstance, Log: logger})

	if err := insertDefaultUser(pgInstance); err != nil {
		logger.Fatal("error while creating default user", zap.Error(err))
	}

	if err := s.Serve(lis); err != nil {
		logger.Fatal("error while serving server", zap.Error(err))
	}
}
