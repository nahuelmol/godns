package nodes

import (
    "fmt"
    "encoding/binary"

    "dnsservice/node/utils"
)

type Query struct {
    MAChead     uint16
    IPv         uint16
    UDPheader   uint16

    checksum    uint16
    
    //dns things
    trans_id    uint16
    flags       string

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

    var message_n = len(message)
    fmt.Printf("message lenght:%d\n", message_n)

    available_classes := [...]string{"INET"}
    fmt.Printf("%s\n",available_classes[0])

    MAChead := binary.BigEndian.Uint16(message[:14])

    ipHeader := 14
    ipVersion := "unknown IPv"
    ipLength := 20

    if message[ipHeader]&0xF0 == 0x40 {
        ipVersion = "IPv4"
        ipLength = 20
    } else if message[ipHeader]&0xF0 == 0x60 {
        ipVersion = "IPv6"
        ipLength = 40
    }

    fmt.Printf("the ip version is: %s\n", ipVersion)


    ipEnd := 14 + ipLength
    IPv    := binary.BigEndian.Uint16(message[14:ipEnd])

    UDPstt := ipEnd
    UDPend := UDPstt + 8

    UDPheader := binary.BigEndian.Uint16(message[UDPstt:UDPend])

    DNSstt := UDPend
    chksum := DNSstt + 10

    checksum := binary.BigEndian.Uint16(message[DNSstt:chksum])

    trans_id   := binary.BigEndian.Uint16(message[10:12])
    flags      := utils.TakeFlags(message[12:14])

    questions:= binary.BigEndian.Uint16(message[14:16])
    ansRR   := binary.BigEndian.Uint16(message[16:18])
    authRR  := binary.BigEndian.Uint16(message[18:20])
    addRR   := binary.BigEndian.Uint16(message[20:22])

    domain      := string(message[22:24])
    queryType   := string(message[24:26])
    queryClass  := string(message[26:28])
    //here I must convert byte data into string or legible data in general

    return &Query {

        MAChead,
        IPv,
        UDPheader, 

        checksum,
            
        //dns things
        trans_id,
        flags,

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
    fmt.Printf("IPv4:%d\n ", query.IPv)
    fmt.Printf("Mac Head:%d\n ", query.MAChead)
    fmt.Printf("UDP head:%d\n ", query.UDPheader)

    fmt.Printf("questions:%d\n ", query.questions)
    fmt.Printf("answer:%d\n ", query.ansRR)
    fmt.Printf("auth:%d\n ", query.authRR)
    fmt.Printf("add:%d\n ", query.addRR)

    fmt.Printf("flags:%s\n", query.flags)
    fmt.Printf("domain:%s\n", query.domain)
    fmt.Printf("query type:%s\n", query.queryType)
    fmt.Printf("query class:%s\n", query.queryClass)

    //check if I have an IP asociated to the domain,    
    //counter returns a referral to other authorative server

    //SendReferral(domain)
}

func SendReferral(domain string) string {
    myTLD := GetTLD(domain)
    fmt.Printf("sending the referral...")

    referral := CreateReferral(14, domain, myTLD)
    fmt.Println(referral)
    //I must research about anamtomy of a reerral query
    return "referral"
}

func GetTLD(domain string) string {
    var TLDs = []string{"com", "net", "edu", "org"}
    var TLD string
    var aux string

    for i:= len(domain) -1; i >= 0; i-- {
        if(domain[i] == '.'){
            break;
        }
        aux = string(domain[i])
        aux += TLD
        TLD = aux
    }

    for _, tld := range TLDs{
        if tld == TLD {
            return TLD
        }
    }
    return TLD
}


