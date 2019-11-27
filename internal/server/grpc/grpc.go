package grpc

import (
	"context"
	"github.com/rs/zerolog/log"
	highscorepb "github.com/sorborail/m-apis/highscorepb/v1"
	"google.golang.org/grpc"
	"math"
)

type Grpc struct {
	address string
	srv *grpc.Server
}

var HighScore = math.MaxFloat64

func (*Grpc) SetHighScore(ctx context.Context, req *highscorepb.SetHighScoreRequest) (*highscorepb.SetHighScoreResponse, error) {
	log.Info().Msg("SetHighScore request is called.")
	HighScore = req.GetHighScore()
	return &highscorepb.SetHighScoreResponse{Status: true}, nil
}

func (*Grpc) GetHighScore(ctx context.Context, req *highscorepb.GetHighScoreRequest) (*highscorepb.GetHighScoreResponse, error) {
	log.Info().Msg("getHighScore request is called.")
	return &highscorepb.GetHighScoreResponse{HighScore: HighScore}, nil
}