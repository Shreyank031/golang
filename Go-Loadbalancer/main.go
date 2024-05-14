package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

type simpleServer struct {
	addr  string
	proxy httputil.ReverseProxy
}

func newSimpleServer(address string) *simpleServer {

	serverUrl, err := url.Parse(address)
	handleErr(err)

	return &simpleServer{
		addr:  address,
		proxy: *httputil.NewSingleHostReverseProxy(serverUrl),
	}
}

type Loadbalancer struct {
	port       string
	roundRobin int
	servers    []Server
}

type Server interface {
	Address() string
	IsAlive() bool
	Serve(rw http.ResponseWriter, r *http.Request)
}

func NewLoadBalancer(port string, servers []Server) *Loadbalancer {
	return &Loadbalancer{
		port:       port,
		roundRobin: 0,
		servers:    servers,
	}
}

func handleErr(err error) {
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
}

func (s *simpleServer) Address() string {
	return s.addr
}

func (s *simpleServer) IsAlive() bool {
	return true
}

func (s *simpleServer) Serve(rw http.ResponseWriter, req *http.Request) {

	s.proxy.ServeHTTP(rw, req)

}

func (lb *Loadbalancer) getNextAvailableServer() Server {
	server := lb.servers[lb.roundRobin%len(lb.servers)]
	for !server.IsAlive() {
		lb.roundRobin++
		server = lb.servers[lb.roundRobin%len(lb.servers)]
	}
	lb.roundRobin++
	return server

}

func (lb *Loadbalancer) serveProxy(rw http.ResponseWriter, r *http.Request) {
	targetServer := lb.getNextAvailableServer()
	fmt.Printf("\nForwarding request to address: %q\n", targetServer.Address())
	targetServer.Serve(rw, r)
}

func main() {
	servers := []Server{
		newSimpleServer("https://duckduckgo.com"),
		newSimpleServer("https://google.com"),
		newSimpleServer("https://in.pinterest.com"),
		newSimpleServer("https://search.brave.com"),
	}
	lb := NewLoadBalancer("8001", servers)

	handleRedirect := func(rw http.ResponseWriter, r *http.Request) {
		lb.serveProxy(rw, r)
	}

	http.HandleFunc("/", handleRedirect)

	fmt.Printf("serving request at 'localhost:%s'\n", lb.port)
	log.Fatal(http.ListenAndServe(":"+lb.port, nil))
}
