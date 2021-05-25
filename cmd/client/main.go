package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/jirawan-chuapradit/grpc-golang-pcbook/pb"
	"github.com/jirawan-chuapradit/grpc-golang-pcbook/sample"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	serverAddress := flag.String("address", "", "the server address")
	flag.Parse()
	log.Printf("dial server %s", *serverAddress)

	conn, err := grpc.Dial(*serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}
	defer conn.Close()

	laptopClient := pb.NewLaptopServiceClient(conn)
	// set time out
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	CreateNewLaptopRequestWithID(laptopClient, ctx)
	CreateNewLaptopRequestWithoutID(laptopClient, ctx)
	

}

func CreateNewLaptopRequestWithID(laptopClient pb.LaptopServiceClient, ctx context.Context){
	log.Println("======> CreateNewLaptopRequestWithID")
	laptop := sample.NewLaptop()
	
	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	res, err := laptopClient.CreateLaptop(ctx, req)
	if err != nil {
		st, ok := status.FromError(err) // convest err to object
		if !ok && st.Code() == codes.AlreadyExists {
			// not a big deal
			log.Println("laptop already exists")
		} else {
			log.Fatal("cannot create laptop: ", err)
		}
		return

	}

	log.Printf("created laptop with id: %s", res.Id)
}

func CreateNewLaptopRequestWithoutID(laptopClient pb.LaptopServiceClient, ctx context.Context){
	log.Println("======> CreateNewLaptopRequestWithoutID")

	laptop := sample.NewLaptop()
	laptop.Id = "9b8e6280-ef47-42a8-858e-8d94fcd4c03a"

	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	res, err := laptopClient.CreateLaptop(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			// not a big deal
			log.Println("laptop already exists")
		}else {
			log.Fatal("cannot create laptop: ", err)
		}
		return
	}
	log.Printf("created laptop with id: %s", res.Id)
}