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
        fmt.Printf("err resolving address")
        return
    }

    udpconn, err := net.ListenUDP("udp", udpaddr)
    if err != nil {
        fmt.Printf("error in listener ", err)
    }
    defer udpconn.Close()

    log.Println("UDP is listening", udpaddr)

    buffer := make([]byte, 1024)
    for {
        n, addr, err := udpconn.ReadFromUDP(buffer)
        if err != nil {
            fmt.Println("error reading UDP incomming messagas", err)
        }
        fmt.Printf("\nrequest from: %s\n", addr)
        nodes.AnalizeQuery(buffer[:n])
    }
}
