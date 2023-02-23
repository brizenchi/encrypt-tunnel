package main

import (
	"encrypt-tunnel/relay"
	"flag"
	"fmt"
	"net"
)

func main() {
	//listenAddr := os.Getenv("listenAddr")
	//remoteAddr := os.Getenv("remoteAddr")
	//role := os.Getenv("role")
	//secret := os.Getenv("secret")
	//tunnel := os.Getenv("tunnel")

	listenAddr := flag.String("listenAddr", "127.0.0.1:2000", "")
	remoteAddr := flag.String("remoteAddr", "127.0.0.1:2001", "")
	role := flag.String("role", "A", "A or B")
	secret := flag.String("secret", "", "")
	tunnel := flag.String("tunnel", "", "")
	flag.Parse()

	if *secret == "" {
		fmt.Println("Please specify a secret.")
		return
	}
	//GlobalKey = []byte(fmt.Sprintf("%x", md5.Sum([]byte(*secret))))
	fmt.Printf("[%s] -> [%s], role = %s, secret = %s, tunnel = %s \n", *listenAddr, *remoteAddr, *role, *secret, *tunnel)

	server, err := net.Listen("tcp", *listenAddr)
	if err != nil {
		fmt.Printf("Listen failed: %v\n", err)
		return
	}

	for {
		client, err := server.Accept()
		if err != nil {
			fmt.Printf("Accept failed: %v", err)
			continue
		}
		go relay.Relay(client, *remoteAddr, *role, *secret, *tunnel)
	}
}
