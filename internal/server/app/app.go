package app

import (
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"

	config "github.com/dimk00z/grpc-filetransfer/config/server"
	"github.com/dimk00z/grpc-filetransfer/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)
	l.Debug("App started")
	ctx, cancel := context.WithCancel(context.Background())

	g := grpc.NewServer()
	reflection.Register(g)
	uploadServer := &GRPCServer{
		l:   l,
		cfg: cfg,
	}
	uploadpb.RegisterFileServiceServer(g, uploadServer)
	// log.Println(server)
	listen, err := net.Listen("tcp", cfg.GRPC.Port)
	if err != nil {
		l.Fatal(err)
	}
	interrupt := make(chan os.Signal, 1)
	shutdownSignals := []os.Signal{
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGINT,
		syscall.SIGQUIT,
	}
	signal.Notify(interrupt, shutdownSignals...)
	go func(g *grpc.Server) {
		l.Info("setGRPC - gRPC server started on " + cfg.GRPC.Port)
		if err := g.Serve(listen); err != nil {
			l.Fatal(err)
		}
	}(g)
	select {
	case killSignal := <-interrupt:
		l.Debug("Got ", killSignal)
		cancel()
	case <-ctx.Done():
	}
}
