package service

import (
	"context"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"

	uploadpb "github.com/dimk00z/grpc-filetransfer/pkg/proto"
	"google.golang.org/grpc"
)

type ClientService struct {
	addr     string
	filePath string
	client   uploadpb.FileServiceClient
}

func New(addr string, filePath string) *ClientService {
	return &ClientService{
		addr:     addr,
		filePath: filePath,
	}
}

func (s *ClientService) SendFile() error {
	log.Println(s.addr, s.filePath)
	conn, err := grpc.Dial(s.addr, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()
	s.client = uploadpb.NewFileServiceClient(conn)
	interrupt := make(chan os.Signal, 1)
	shutdownSignals := []os.Signal{
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGINT,
		syscall.SIGQUIT,
	}
	signal.Notify(interrupt, shutdownSignals...)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func(s *ClientService) {
		if err = s.upload(ctx, cancel); err != nil {
			log.Fatal(err)
			cancel()
		}
	}(s)

	select {
	case killSignal := <-interrupt:
		log.Println("Got ", killSignal)
		cancel()
	case <-ctx.Done():
	}
	return nil
}

func (s *ClientService) upload(ctx context.Context, cancel context.CancelFunc) error {
	stream, err := s.client.Upload(ctx)
	if err != nil {
		return err
	}
	file, err := os.Open(s.filePath)
	if err != nil {
		return err
	}
	buf := make([]byte, 64*1024)
	for {
		num, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if err := stream.Send(&uploadpb.FileUploadRequest{FileName: s.filePath, Chunk: buf[:num]}); err != nil {
			return err
		}
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}
	log.Printf("Sent - %v bytes - %s\n", res.GetSize(), res.GetFileName())
	cancel()
	return nil
}
