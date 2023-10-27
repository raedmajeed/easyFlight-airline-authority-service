package pkg

import (
	"fmt"
	"log"
	"net"

	"github.com/gin-gonic/gin"
	"github.com/raedmajeed/admin-servcie/config"
	"github.com/raedmajeed/admin-servcie/pkg/api/handlers"
	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
	"google.golang.org/grpc"
)

type Server struct {
	E   *gin.Engine
	cfg *config.ConfigParams
}

func NewServer(cfg *config.ConfigParams, handler *handlers.AdminAirlineHandler) (*Server, error) {
	err := NewGrpcServer(cfg, handler)
	if err != nil {
		log.Println("error connecting to gRPC server")
		return nil, err
	}
	r := gin.Default()
	return &Server{
		E:   r,
		cfg: cfg,
	}, nil
}

func NewGrpcServer(cfg *config.ConfigParams, handler *handlers.AdminAirlineHandler) error {
	log.Println("connecting to gRPC server")
	addr := fmt.Sprintf(":%s", cfg.ADMINPORT)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Println("error Connecting to gRPC server")
		return err
	}
	grp := grpc.NewServer()
	pb.RegisterAdminAirlineServer(grp, handler)
	if err != nil {
		log.Println("error connecting to gRPC server")
		return err
	}

	log.Printf("listening on gRPC server %v", cfg.ADMINPORT)
	err = grp.Serve(lis)
	if err != nil {
		log.Println("error connecting to gRPC server")
		return err
	}
	return nil
}

func (s *Server) ServerStart() error {
	err := s.E.Run(":" + s.cfg.PORT)
	if err != nil {
		log.Println("error starting server")
		return err
	}
	log.Println("Server started")
	return nil
}
