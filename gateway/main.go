package main

import (
	"context"
	"log"
	"net/http"

	"google.golang.org/grpc"

	"fmt"

	"strconv"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	math "github.com/grpc-tutorial/app/api/math"
)

func handlerAdd(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	a, _ := strconv.Atoi(vars["a"])
	b, _ := strconv.Atoi(vars["b"])
	mathBackend := "backend:8080"
	connToMath, err := grpc.Dial(mathBackend, grpc.WithInsecure())
	if err != nil {
		logrus.Fatalf("could not connect to %s: %v", mathBackend, err)
	}
	defer connToMath.Close()

	client := math.NewAdditionClient(connToMath)
	val := &math.BinaryValue{A: int32(a), B: int32(b)}
	res, err := client.Add(context.Background(), val)
	if err != nil {
		logrus.Fatalf("could not add %s: %v", val, err)
	}
	w.Write([]byte(strconv.Itoa(int(res.R))))
	w.WriteHeader(http.StatusOK)
	return
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	r.HandleFunc("/add/{a:[0-9]+}/{b:[0-9]+}", handlerAdd).Methods("GET")
	fmt.Println("listening on 1111")
	err := http.ListenAndServe(":1111", r)
	if err != nil {
		log.Fatal("could not start the server")
	}
}