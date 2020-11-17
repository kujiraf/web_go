package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"strconv"
	"web_go/grpc/calculator/calcpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Sum Client")
	nums, err := getArgs()
	if err != nil {
		log.Fatalf("An error occured. %v", err)
	}

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := calcpb.NewCalcServiceClient(conn)
	// doSum(c, nums)
	doFindPrimeComposition(c, nums)
}

func getArgs() ([]int64, error) {
	flag.Parse()
	var nums []int64
	for _, v := range flag.Args() {
		num, err := strconv.Atoi(v)
		if err != nil {
			return nil, errors.New("please input only number")
		}
		nums = append(nums, int64(num))
	}
	return nums, nil
}

func doSum(c calcpb.CalcServiceClient, nums []int64) {
	fmt.Println("[doSum] Starting to do a Unary RPC...")
	req := &calcpb.Nums{
		Nums: nums,
	}
	res, err := c.CalcSum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling CalcSum RPC: %v", err)
	}
	log.Printf("Response from CalcSum: %v", res.Result)
}

func doFindPrimeComposition(c calcpb.CalcServiceClient, nums []int64) {
	fmt.Println("[doFindPrimeComposition] Starting to do a Unary RPC...")
	req := &calcpb.Nums{
		Nums: nums,
	}
	res, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling PrimeNumberDecomposition RPC: %v", err)
	}
	for {
		msg, err := res.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while receiving PrimeNumberDecomposition RPC: %v", err)
		}
		log.Printf("Response from PrimeNumberDecomposition: %v", msg.GetResult())
	}
}
