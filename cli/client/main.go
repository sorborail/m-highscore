package main

import (
	"context"
	"flag"
	zlog "github.com/rs/zerolog/log"
	highscorepb "github.com/sorborail/m-apis/highscorepb/v1"
	"google.golang.org/grpc"
	"time"
)

func main() {
	zlog.Info().Msg("Begin starting highscore client...")
	var addrPtr = flag.String("address", "localhost:50051", "address to connect highscore service")
	flag.Parse()
	zlog.Debug().Msgf("Value of addrPtr - %s", *addrPtr)
	opts := grpc.WithInsecure()
	conn, err := grpc.Dial(*addrPtr, opts)
	if err != nil {
		zlog.Fatal().Err(err).Str("address", *addrPtr).Msg("could not connect to the highscore server")
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			zlog.Error().Err(err).Str("address", *addrPtr).Msg("Failed to close connection")
		}
	}()
	cl := highscorepb.NewGameClient(conn)
	if cl == nil {
		zlog.Error().Msg("Cannot connection to the highscore service")
	} else {
		zlog.Info().Msg("Highscore client is started")
		doGetHighscore(cl, addrPtr)
	}
}

func doGetHighscore(cl highscorepb.GameClient, addr *string) {
	zlog.Info().Msg("Begin get highscore request...")
	req := &highscorepb.GetHighScoreRequest{}
	ctx, cancel := context.WithTimeout(context.Background(), 20 * time.Second)
	defer cancel()
	res, err := cl.GetHighScore(ctx, req)
	if err != nil {
		zlog.Fatal().Err(err).Str("address", *addr).Msg("error happened while get highscore response")
	}
	zlog.Info().Interface("highscore", res.GetHighScore()).Msg("Value from highscore service")
}