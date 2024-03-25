package nodes

import (
    "fmt"
    "encoding/binary"
)

type Query struct {
    MAChead     uint16
    IPv4        uint16
    UDPsport    uint16
    UDPdport    uint16

    checksum    uint16
    
    //dns things
    trans_id    uint16
    parameters  uint16

    questions   uint16
    ansRR       uint16
    authRR      uint16
    addRR       uint16

    //the query
    domain      string
    queryType   string
    queryClass  string
}

func newQuery(message []byte) *Query {
    var MAChead uint16
    var IPv4 uint16
    var UDPsport uint16
    var UDPdport uint16

    var checksum uint16

    var trans_id    uint16
    var parameters  uint16

    var questions   uint16
    var ansRR       uint16
    var authRR      uint16
    var addRR       uint16

    //the query
    var domain      string
    var queryType   string
    var queryClass  string

    available_classes := [...]string{"INET"}
    fmt.Printf("%s\n",available_classes[0])

    MAChead = binary.BigEndian.Uint16(message[:2])
    IPv4    = binary.BigEndian.Uint16(message[2:4])

    UDPsport = binary.BigEndian.Uint16(message[4:6])
    UDPdport = binary.BigEndian.Uint16(message[6:8])

    checksum = binary.BigEndian.Uint16(message[:2])

    trans_id   = binary.BigEndian.Uint16(message[:2])
    parameters = binary.BigEndian.Uint16(message[:2])
 

    questions= binary.BigEndian.Uint16(message[:2])
    ansRR   = binary.BigEndian.Uint16(message[:2])
    authRR  = binary.BigEndian.Uint16(message[:2])
    addRR   = binary.BigEndian.Uint16(message[:2])

    domain      = string(message[:2])
    queryType   = string(message[:2])
    queryClass  = string(message[:2])
    //here I must convert byte data into string or legible data in general

    return &Query {

        MAChead,
        IPv4,
        UDPsport, 
        UDPdport,

        checksum,
            
        //dns things
        trans_id,
        parameters,

        questions,
        ansRR,
        authRR,
        addRR,

        //the query
        domain,
        queryType,
        queryClass,
    }
}


func AnalizeQuery(message []byte) {

        
    query := newQuery(message);
    fmt.Printf("IPv4:%d\n ", query.IPv4)
    fmt.Printf("Mac Head:%d\n ", query.MAChead)
    fmt.Printf("UDP source:%d\n ", query.UDPsport)
    fmt.Printf("UDP destination:%d\n ", query.UDPdport)

    //check if I have an IP asociated to the domain, counter return the 
    //referal

    returnReferal()
}

func returnReferal(){
    fmt.Printf("returning referal")
}


