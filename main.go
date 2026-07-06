package main

import (
    "fmt"
    "log"
    "net"
    "dnsservice/node/nodes"
)

func main(){
    udpaddr, err := net.ResolveUDPAddr("udp", ":8080")
    if err != nil {
        fmt.Printf("err resolving udp:8080 address")
        return
    }
    udpconn, err := net.ListenUDP("udp", udpaddr)
    if err != nil {
        fmt.Printf("error listeninig %s", err)
    }
    defer udpconn.Close()
    log.Println("listening on ", udpaddr)
    buffer := make([]byte, 1024)

    for {
        n, addr, err := udpconn.ReadFromUDP(buffer)
        if err != nil {
            fmt.Println("error reading UDP buffer", err)
			continue
        }

		packet := make([]byte, n)
		copy(packet, buffer[:n])

		go func(data []byte, client *net.UDPAddr) {
			fmt.Printf("\nrequest from: %s", client)
			query := nodes.AnalizeQuery(data)
			udpconn.WriteToUDP(query.Response, addr)
		}(packet, addr)
    }
}
