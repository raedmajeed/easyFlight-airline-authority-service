package pkg

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/raedmajeed/admin-servcie/config"
	"github.com/raedmajeed/admin-servcie/pkg/api/bookingHandlers"
	"github.com/raedmajeed/admin-servcie/pkg/api/handlers"
	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
	"github.com/raedmajeed/admin-servcie/pkg/service/interfaces"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	E   *gin.Engine
	cfg *config.ConfigParams
}

func NewServer(cfg *config.ConfigParams, handler *handlers.AdminAirlineHandler,
	svc interfaces.AdminAirlineService, bHandler *bookingHandlers.BookingHandler) {
	//newContext, cancel := context.WithTimeout(context.Background(), time.Second*1500)
	//defer cancel()

	signalChan := make(chan os.Signal, 2)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	//go config.KafkaReaders(newContext, svc, signalChan)

	go NewBookingGrpcServer(cfg, bHandler)
	go NewGrpcServer(cfg, handler)
	<-signalChan
}

func NewGrpcServer(cfg *config.ConfigParams, handler *handlers.AdminAirlineHandler) {
	log.Println("connecting to gRPC server")
	addr := fmt.Sprintf(":%s", cfg.ADMINPORT)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Println("error Connecting to gRPC server")
		return
	}
	grp := grpc.NewServer()
	pb.RegisterAdminAirlineServer(grp, handler)
	if err != nil {
		log.Println("error connecting to gRPC server")
		return
	}

	log.Printf("listening on gRPC server %v", cfg.ADMINPORT)
	err = grp.Serve(lis)
	if err != nil {
		log.Println("error connecting to gRPC server")
		return
	}
	return
}

func NewBookingGrpcServer(cfg *config.ConfigParams, handler *bookingHandlers.BookingHandler) {
	log.Println("connecting to Booking gRPC server")
	addr := fmt.Sprintf(":%s", cfg.ADMINBOOKINGPORT)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Println("error listening to booking service")
		return
	}
	grpcServer := grpc.NewServer()
	pb.RegisterAdminServiceServer(grpcServer, handler)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Println("error connecting to booking  grpc server")
		return
	}
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
