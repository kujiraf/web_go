package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
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

	c := calcpb.NewSumServiceClient(conn)
	doSum(c, nums)
}

func getArgs() ([]int, error) {
	flag.Parse()
	args := flag.Args()
	if len(args) < 2 {
		return nil, errors.New("please input more than 2 numbers args")
	}
	var nums []int
	for _, v := range args {
		num, err := strconv.Atoi(v)
		if err != nil {
			return nil, errors.New("please input only number")
		}
		nums = append(nums, num)
	}
	return nums, nil
}

func doSum(c calcpb.SumServiceClient, nums []int) {
	fmt.Println("[doSum] Starting to do a Unary RPC...")
	req := &calcpb.SumRequest{
		Nums: &calcpb.Nums{
			Num_1: int64(nums[0]),
			Num_2: int64(nums[1]),
		},
	}
	res, err := c.CalcSum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling CalcSum RPC: %v", err)
	}
	log.Printf("Response from CalcSum: %v", res.Result)
}
