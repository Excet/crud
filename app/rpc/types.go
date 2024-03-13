package rpc

import (
	"crud/app/models"
	pb "crud/app/proto"

	"go.uber.org/zap"
)

type Server struct {
	pb.UnimplementedCrudServer
	PG  RPCIface
	Log *zap.Logger
}

type RPCIface interface {
	GetNodeParents(node models.GraphNode) ([]models.GraphNode, error)
	GetNodePrice(node models.GraphNode) (error, price int32)
	SetNodePrice(node models.GraphNode) error
}
