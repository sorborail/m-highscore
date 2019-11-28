package grpc

import (
	"context"
	"fmt"
	zlog "github.com/rs/zerolog/log"
	highscorepb "github.com/sorborail/m-apis/highscorepb/v1"
	"google.golang.org/grpc"
	"math"
	"net"
)

type Server struct {
	address string
	srv *grpc.Server
	lis net.Listener
}

var HighScore = math.MaxFloat64

func (*Server) SetHighScore(ctx context.Context, req *highscorepb.SetHighScoreRequest) (*highscorepb.SetHighScoreResponse, error) {
	zlog.Info().Msg("SetHighScore request is called.")
	HighScore = req.GetHighScore()
	return &highscorepb.SetHighScoreResponse{Status: true}, nil
}

func (*Server) GetHighScore(ctx context.Context, req *highscorepb.GetHighScoreRequest) (*highscorepb.GetHighScoreResponse, error) {
	zlog.Info().Msg("getHighScore request is called.")
	return &highscorepb.GetHighScoreResponse{HighScore: HighScore}, nil
}

func (s *Server) DoServe() error {
	var err error
	s.lis, err = net.Listen("tcp", s.address)
	if err != nil {
		zlog.Error().Msg("Failed to listen service")
		return fmt.Errorf("failed to listen server port %w", err)
	}
	s.srv = grpc.NewServer()
	highscorepb.RegisterGameServer(s.srv, s)
	zlog.Info().Str("address", s.address).Msg("gRPC Highscore microservice is started.")
	if err = s.srv.Serve(s.lis); err != nil {
		zlog.Error().Msg("Failed to serve server for highscore microservice")
		return fmt.Errorf("failed to serve server for highscore microservice %w", err)
	}
	return nil
}

func NewServer(addr string) *Server {
	return &Server{
		address: addr,
	}
}

func (s *Server) StopServer() {
	zlog.Info().Msg("Stopping the Highscore service Server...")
	s.srv.Stop()
	zlog.Info().Msg("Closing the Listener...")
	_ = s.lis.Close()
	zlog.Info().Msg("End of program.")
}