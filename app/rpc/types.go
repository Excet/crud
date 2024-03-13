package rpc

import (
	"context"
	"crud/app/models"
	pb "crud/app/proto"

	"go.uber.org/zap"
)

type Server struct {
	pb.UnimplementedCrudServer
	PG  RPCIface
	Log *zap.Logger
}

// GetNodeParents implements proto.CrudServer.
// Subtle: this method shadows the method (UnimplementedCrudServer).GetNodeParent of Server.UnimplementedCrudServer.
func (s *Server) GetNodeParent(context.Context, *pb.Node) (*pb.Node, error) {
	panic("unimplemented")
}

// GetNodePrice implements proto.CrudServer.
// Subtle: this method shadows the method (UnimplementedCrudServer).GetNodePrice of Server.UnimplementedCrudServer.
func (s *Server) GetNodePrice(context.Context, *pb.Node) (*pb.Price, error) {
	panic("unimplemented")
}

// SetNodePrice implements proto.CrudServer.
// Subtle: this method shadows the method (UnimplementedCrudServer).SetNodePrice of Server.UnimplementedCrudServer.
func (s *Server) SetNodePrice(context.Context, *pb.Node) (*pb.Ok, error) {
	panic("unimplemented")
}

// mustEmbedUnimplementedCrudServer implements proto.CrudServer.
// Subtle: this method shadows the method (UnimplementedCrudServer).mustEmbedUnimplementedCrudServer of Server.UnimplementedCrudServer.
func (s *Server) mustEmbedUnimplementedCrudServer() {
	panic("unimplemented")
}

type RPCIface interface {
	GetNodeParent(node models.GraphNode) (models.GraphNode, error)
	GetNodePrice(node models.GraphNode) (int32, error)
	SetNodePrice(node models.GraphNode) error
}
