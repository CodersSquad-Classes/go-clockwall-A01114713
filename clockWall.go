package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

type reloj struct {
	lugar, localhost string
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Introduzca al menos una timezone y su respectivo localhost")
		os.Exit(0)
	}

	var relojes []reloj
	for _, parametro := range os.Args[1:] {
		info := strings.Split(parametro, "=")
		if len(info) == 1 {
			fmt.Println("Example of expected argument: Tokyo=localhost:3030")
			os.Exit(0)
		}
		relojes = append(relojes, reloj{info[0], info[1]})
	}
	for _, r := range relojes {
		conn, err := net.Dial("tcp", r.localhost)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		go mustCopy(os.Stdout, conn)
	}
	for {
		time.Sleep(1 * time.Second)
	}
}
