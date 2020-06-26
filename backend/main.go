package main

import (
	"flag"
	"fmt"
	"net"

	"github.com/Sirupsen/logrus"
)

func main() {
	port := flag.Int("p", 8080, "port to listen to")
	flag.Parse()

	net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		logrus.Fatalf("cound not listen to port %d: %v", *port, err)
	}

	s := grpc.NewServer()
}

//cmd := exec.Command("flite", "-t", os.Args[1], "-o", "output.wav")
//cmd.Stdout = os.Stdout
//cmd.Stderr = os.Stderr
//if err := cmd.Run(); err != nil {
//	log.Fatal(err)
//}
