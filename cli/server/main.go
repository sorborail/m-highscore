package main

import (
	"flag"
	zlog "github.com/rs/zerolog/log"
	"github.com/sorborail/m-highscore/internal/server/grpc"
	"os"
	"os/signal"
)

func main() {
	zlog.Info().Msg("Begin starting highscore service server...")
	var addrPtr = flag.String("address", ":50051", "address to connect highscore service")
	flag.Parse()
	zlog.Debug().Msgf("Value of addrPtr - %s", *addrPtr)
	srv := grpc.NewServer(*addrPtr)
	go func() {
		if err := srv.DoServe(); err != nil {
			zlog.Fatal().Msgf("Error while starting server for highscore service - %v", err)
		}
	}()

	// Wait for Control C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	//Block until a signal is received
	<- ch
	srv.StopServer()
}
