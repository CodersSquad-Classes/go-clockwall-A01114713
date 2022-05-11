package main

import (
	"bytes"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func (r *reloj) mustCopy(dst io.Writer, src io.Reader) {
	//srcString := reader2string(src)
	//output := r.lugar + ": " + srcString

	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
	//fmt.Fprintf(dst, "%s\n", output)
}

func reader2string(src io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(src)
	return buf.String()
}

type reloj struct {
	lugar, localhost string
}

func main() {
	var relojes []reloj
	for _, parametro := range os.Args[1:] {
		info := strings.Split(parametro, "=")
		relojes = append(relojes, reloj{info[0], info[1]})
	}
	for _, r := range relojes {
		conn, err := net.Dial("tcp", r.localhost)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		go r.mustCopy(os.Stdout, conn)
	}
	for {
		time.Sleep(1 * time.Second)
	}
}
