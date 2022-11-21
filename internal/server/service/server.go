package service

import (
	"bytes"
	"io"

	config "github.com/dimk00z/grpc-filetransfer/config/server"
	"github.com/dimk00z/grpc-filetransfer/pkg/logger"
	uploadpb "github.com/dimk00z/grpc-filetransfer/pkg/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GRPCServer struct {
	uploadpb.UnimplementedFileServiceServer
	l   *logger.Logger
	cfg *config.Config
}

func (g *GRPCServer) Upload(stream uploadpb.FileService_UploadServer) error {
	req, err := stream.Recv()
	if err != nil {
		g.l.Debug(err)
		return status.Errorf(codes.Unknown, "cannot receive image info")
	}
	fileName := req.GetInfo().GetFileName()
	fileData := bytes.Buffer{}
	fileSize := 0

	for {
		err := g.contextError(stream.Context())
		if err != nil {
			return err
		}

		g.l.Debug("waiting to receive more data")

		req, err := stream.Recv()
		if err == io.EOF {
			g.l.Debug("no more data")
			break
		}
		if err != nil {
			return g.logError(status.Errorf(codes.Unknown, "cannot receive chunk data: %v", err))
		}
		chunk := req.GetChunkData()
		size := len(chunk)

		g.l.Debug("received a chunk with size: %d", size)

		fileSize += size

		if _, err = fileData.Write(chunk); err != nil {
			return g.logError(status.Errorf(codes.Internal, "cannot write chunk data: %v", err))
		}
	}

	if err := g.FileSave(fileName, fileData); err != nil {
		return g.logError(status.Errorf(codes.Internal, "cannot save image to the store: %v", err))
	}

	res := &uploadpb.FileUploadResponse{
		FileName: fileName,
		Size:     uint32(fileSize),
	}

	err = stream.SendAndClose(res)
	if err != nil {
		return g.logError(status.Errorf(codes.Unknown, "cannot send response: %v", err))
	}
	g.l.Debug("saved file: %s, size: %d", fileName, fileSize)
	return nil
}
