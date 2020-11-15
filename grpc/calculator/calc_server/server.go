package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"web_go/grpc/calculator/calcpb"

	"google.golang.org/grpc"
)

type server struct{}

func main() {
	fmt.Println("Calc Server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calcpb.RegisterSumServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (*server) CalcSum(ctx context.Context, req *calcpb.Nums) (*calcpb.SumResponse, error) {
	fmt.Printf("Calc function was invoked with %v\n", req)
	nums := req.GetNums()
	var result int
	for _, v := range nums {
		result += int(v)
	}
	res := &calcpb.SumResponse{
		Result: fmt.Sprintf("ans=%d", result),
	}
	return res, nil
}
