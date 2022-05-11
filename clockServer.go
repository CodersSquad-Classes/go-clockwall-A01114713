// Clock Server is a concurrent TCP server that periodically writes the time.
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

func handleConn(c net.Conn, timeZone string) {
	defer c.Close()

	t := time.Now()
	loc, err := time.LoadLocation(timeZone)
	if err == nil {
		t = t.In(loc)
	}

	for {
		t = time.Now()
		_, err = io.WriteString(c, timeZone+": "+t.Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	//FLAG
	port := flag.String("port", "", "port number for listener")
	flag.Parse()
	portStr := *port

	if portStr == "" {
		fmt.Println("Missing -port flag")
		os.Exit(0)
	}
	portInt, err := strconv.Atoi(portStr)
	if err != nil {
		fmt.Println("Solo se aceptan numeros")
		os.Exit(0)
	} else if portInt < 0 {
		fmt.Println("Numero de puerto erroneo")
		os.Exit(0)
	} else if portInt <= 1024 {
		fmt.Println("Puertos menores o iguales a 1024 son puertos privilegiados")
		os.Exit(0)
	}
	localhost := "localhost:" + portStr

	//Environment variable
	tz, thereIsVariable := os.LookupEnv("TZ")
	if !thereIsVariable {
		fmt.Println("Missing variable TZ")
		os.Exit(0)
	}

	listener, err := net.Listen("tcp", localhost)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn, tz) // handle connections concurrently
	}

}
