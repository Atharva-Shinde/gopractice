package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	doTCP   bool
	doUDP   bool
	doHTTP  bool
	doClose bool
	port    int
)

func init() {
	flag.BoolVar(&doTCP, "tcp", false, "Serve raw over TCP.")
	flag.BoolVar(&doUDP, "udp", false, "Serve raw over UDP.")
	flag.BoolVar(&doHTTP, "http", true, "Serve HTTP.")
	flag.BoolVar(&doClose, "close", false, "Close connection per each HTTP request.")
	flag.IntVar(&port, "port", 9376, "Port number.")
}

func main() {
	flag.Parse()
	if doHTTP && (doTCP || doUDP) {
		log.Fatalf("Can't server TCP/UDP mode and HTTP mode at the same time")
	}

	hostname, err := os.Hostname()
	if err != nil {
		log.Fatalf("Error from os.Hostname(): %s", err)
	}

	if doTCP {
		listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err != nil {
			log.Fatalf("Error from net.Listen(): %s", err)
		}
		go func() {
			for {
				conn, err := listener.Accept()
				if err != nil {
					log.Fatalf("Error from Accept(): %s", err)
				}
				log.Printf("TCP request from %s", conn.RemoteAddr().String())
				conn.Write([]byte(hostname))
				conn.Close()
			}
		}()
	}
	if doUDP {
		addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", port))
		if err != nil {
			log.Fatalf("Error from net.ResolveUDPAddr(): %s", err)
		}
		sock, err := net.ListenUDP("udp", addr)
		if err != nil {
			log.Fatalf("Error from ListenUDP(): %s", err)
		}
		go func() {
			var buffer [16]byte
			for {
				_, cliAddr, err := sock.ReadFrom(buffer[0:])
				if err != nil {
					log.Fatalf("Error from ReadFrom(): %s", err)
				}
				log.Printf("UDP request from %s", cliAddr.String())
				sock.WriteTo([]byte(hostname), cliAddr)
			}
		}()
	}
	if doHTTP {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			log.Printf("HTTP request from %s", r.RemoteAddr)

			if doClose {
				// Add this header to force to close the connection after serving the request.
				w.Header().Add("Connection", "close")
			}

			fmt.Fprintf(w, "%s", hostname)
		})
		go func() {
			// Run in a closure so http.ListenAndServe doesn't block
			http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
			log.Printf("ListenAndServe() returned")
		}()
	}
	log.Printf("Serving on port %dn", port)
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM)
	sig := <-signals
	log.Printf("Shutting down after receiving signal: %s", sig)
	log.Printf("Awaiting pod deletion")
	time.Sleep(60 * time.Second)
}