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
    log.Println("UDP is listening", udpaddr)
    buffer := make([]byte, 1024)
    for {
        n, addr, err := udpconn.ReadFromUDP(buffer)
        if err != nil {
            fmt.Println("error reading UDP buffer", err)
            fmt.Println("buffer size: %d", n)
        }
        fmt.Printf("\nrequest from: %s", addr)
        nodes.AnalizeQuery(buffer[:n])
    }
}
