package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"
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
	calcpb.RegisterCalcServiceServer(s, &server{})

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

func (*server) PrimeNumberDecomposition(req *calcpb.Nums, stream calcpb.CalcService_PrimeNumberDecompositionServer) error {
	nums := req.GetNums()
	for _, v := range nums {
		num := int(v)
		var results []string
		send(strconv.Itoa(num)+" is composed by ...", stream)
		if num == 1 {
			results = append(results, "1")
		} else {
			results = calcPrimeNumberDecomposition(num, stream)
		}
		send(fmt.Sprintf("%d is composed by %s", num, results), stream)
	}
	return nil
}

func calcPrimeNumberDecomposition(num int, stream calcpb.CalcService_PrimeNumberDecompositionServer) []string {
	var target int
	if num < 0 {
		target = -num
	} else {
		target = num
	}
	var result []string
	divisor := 2
	for target > 1 {
		if target%divisor == 0 {
			// fmt.Printf("target:%d, i:%d\n", target, i)
			strnum := strconv.Itoa(divisor)
			result = append(result, strnum)
			target /= divisor
			send(strnum, stream)
		} else {
			divisor++
		}
		time.Sleep(200 * time.Millisecond)
	}
	return result
}

func send(msg string, stream calcpb.CalcService_PrimeNumberDecompositionServer) {
	res := &calcpb.PrimeNumberDecompositionResponse{
		Result: msg,
	}
	stream.Send(res)
}
