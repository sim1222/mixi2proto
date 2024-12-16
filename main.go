package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	pb "github.com/sim1222/mixi2proto/proto"
	"google.golang.org/protobuf/proto"
)

// load from env

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// data = append([]byte{0, 0, 0, 0, 8}, data...)

	// fmt.Printf("Encoded message: %x\n", data)

	// // save to file
	// err = os.WriteFile("create_post1.bin", data, 0644)
	// if err != nil {
	// 	log.Fatal("write file error: ", err)
	// }

	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: main.go <message>")
		os.Exit(1)
	}

	create_post(args[1])
}

func create_post(msg string) {
	m := &pb.CreatePost{
		Body: msg,
		Nazoone: &pb.CreatePost_Nazo1{
			Nazo1: 0,
		},
	}

	// Encode the message to binary
	data, err := proto.Marshal(m)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	data = build_grpc_binary(data)

	fmt.Printf("Encoded message: %x\n", data)

	send_grpc_api("https://rpc.mixi.social/com.mixi.mercury.api.MercuryService/CreatePost", data)
}

func build_grpc_binary(data []byte) []byte {
	len := len(data)

	// 00 00 00 00 len

	data = append([]byte{0, 0, 0, 0, byte(len)}, data...)

	// data = append([]byte{len}, data...)
	// data = append([]byte{0, 0, 0, 0, 8}, data...)
	return data
}

func send_grpc_api(url string, data []byte) {

	mercuryAuthorization := os.Getenv("MERCURY_AUTHORIZATION")
	authKey := os.Getenv("AUTH_KEY")

	reader := bytes.NewReader(data)

	req, _ := http.NewRequest("POST", url, reader)
	req.Header.Set("grpc-timeout", "7000m")
	req.Header.Set("content-type", "application/grpc")
	req.Header.Set("te", "trailers")
	req.Header.Set("user-agent", "Mercury-ios/1.2.0")
	req.Header.Set("x-mercury-authorization", mercuryAuthorization)
	req.Header.Set("x-auth-key", authKey)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal("http request error: ", err)
	}

	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Printf("response body: %v\n", resp.Body)
}