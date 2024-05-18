package main

import (
	"context"
	"log"
	"net"

	excelreader "github.com/vinodyadavjella/excelreader/src/proto"
	"google.golang.org/grpc"
)

type server struct {
	excelreader.UnimplementedExcelReaderServer
}

func (s *server) UploadExcelFile(ctx context.Context, file *excelreader.ExcelFile) (*excelreader.ProcessResponse, error) {
	jsonData, err := processExcelFile(file.Content)
	if err != nil {
		return nil, err
	}

	return &excelreader.ProcessResponse{
		Message:  "File processed successfully",
		JsonData: jsonData,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	excelreader.RegisterExcelReaderServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
