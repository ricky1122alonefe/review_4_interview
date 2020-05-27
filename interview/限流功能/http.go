package main

import (

	"net"
	"net/http"

	"golang.org/x/net/netutil"
)

func main(){
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		fmt.Fatalf("Listen: %v", err)
	}
	defer l.Close()
	l = netutil.LimitListener(l, 5)

	http.Serve(l, http.HandlerFunc())


}
