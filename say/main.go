package main

import (
	"context"
	"flag"

	"google.golang.org/grpc"

	"fmt"

	"github.com/Sirupsen/logrus"
	math "github.com/grpc-tutorial/app/api/math"
	say "github.com/grpc-tutorial/app/api/say"
)

func main() {
	mathBackend := flag.String("mb", "localhost:8080", "address of Math backend")
	sayBackend := flag.String("sb", "localhost:9090", "address of Say backend")
	flag.Parse()

	connToMath, err := grpc.Dial(*mathBackend, grpc.WithInsecure())
	if err != nil {
		logrus.Fatalf("could not connect to %s: %v", *mathBackend, err)
	}
	defer connToMath.Close()

	client := math.NewAdditionClient(connToMath)
	val := &math.BinaryValue{A: 100, B: 200}
	// res, err := client.Hello(context.Background(), text)
	res, err := client.Add(context.Background(), val)
	if err != nil {
		logrus.Fatalf("could not add %s: %v", val, err)
	}
	fmt.Println(res)

	un := &math.UnaryValue{A: 100}
	res, err = client.Incr(context.Background(), un)
	if err != nil {
		logrus.Fatalf("could not incr %s: %v", un, err)
	}
	fmt.Println(res)

	////////////////////////////////////

	connToSay, err := grpc.Dial(*sayBackend, grpc.WithInsecure())
	if err != nil {
		logrus.Fatalf("could not connect to %s: %v", *sayBackend, err)
	}
	defer connToSay.Close()

	sayClient := say.NewSayHelloClient(connToSay)
	text := &say.Text{Text: "how are you"}
	sayRes, err := sayClient.Hello(context.Background(), text)
	if err != nil {
		logrus.Fatalf("could not add %s: %v", val, err)
	}
	fmt.Println(sayRes)

}
