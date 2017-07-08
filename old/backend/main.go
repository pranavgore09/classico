package main

import (
	"flag"
	"fmt"
	"net"

	"google.golang.org/grpc"

	"github.com/Sirupsen/logrus"
	math "github.com/grpc-tutorial/app/api/math"
	say "github.com/grpc-tutorial/app/api/say"
	context "golang.org/x/net/context"
)

func startMathserver(port *int, s *grpc.Server) {
	logrus.Infof("Math server : listening to port %d", *port)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		logrus.Fatalf("could not listen to port %d: %v", *port, err)
	}
	math.RegisterAdditionServer(s, Adder{})
	err = s.Serve(lis)
	if err != nil {
		logrus.Fatalf("could not serve : %v", err)
	}
}

func main() {
	portForMath := flag.Int("pm", 8080, "port for Math server to listen to")
	portForSay := flag.Int("ps", 9090, "port for Say server to listen to")
	flag.Parse()
	s := grpc.NewServer()

	go startMathserver(portForMath, s)

	logrus.Infof("Say server : listening to port %d", *portForSay)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *portForSay))
	if err != nil {
		logrus.Fatalf("could not listen to port %d: %v", *portForSay, err)
	}
	say.RegisterSayHelloServer(s, server{})

	err = s.Serve(lis)
	if err != nil {
		logrus.Fatalf("could not serve : %v", err)
	}
}

type server struct{}

func (s server) Hello(c context.Context, t *say.Text) (*say.Text, error) {
	// return nil, fmt.Errorf("not implemented")
	return &say.Text{Text: fmt.Sprintf("awesomeness of %s", t.Text)}, nil
}

type Adder struct{}

func (a Adder) Add(c context.Context, v *math.BinaryValue) (*math.Value, error) {
	return &math.Value{R: v.A + v.B}, nil
}

func (a Adder) Incr(c context.Context, v *math.UnaryValue) (*math.Value, error) {
	return &math.Value{R: v.A + 1}, nil
}

// fmt.Println("hello world")
