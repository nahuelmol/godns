package nodes

import (
    "fmt"
    "encoding/binary"
    "errors"
    //"dnsservice/node/utils"
)

type Query struct {
    IP  [20]byte
    MAC [14]byte
    UDP [8]byte

    UDPheader   uint16 // this is 8 bytes

    checksum    uint16
    
    //dns things
    //trans_id    uint16
    //flags       string

    questions   uint16
    ansRR       uint16
    authRR      uint16
    addRR       uint16

    //the query
    domain      string
    queryType   string
    queryClass  string

	Response  	[]byte
}

func newQuery(packet []byte) *Query {

	header := packet[:12];
	id := binary.BigEndian.Uint16(header[0:2])
	//flags := binary.BigEndian.Uint16(header[2:4])
	//qdCount := binary.BigEndian.Uint16(header[4:6])
	//anCount := binary.BigEndian.Uint16(header[6:8])
	//nsCount := binary.BigEndian.Uint16(header[8:10])
	//arCount := binary.BigEndian.Uint16(header[10:12])

	//_qr := (flags >> 15) & 1
	//_opcode := (flags >> 11) & 0xF
	//_aa := (flags >> 10) & 1
	//_tc := (flags >> 9) & 1
	//_rd := (flags >> 8) & 1
	//_ra := (flags >> 7) & 1
	//_ad := (flags >> 5) & 1
	//_cd := (flags >> 4) & 1
	//_rcode := flags & 0xF

	offset := 12
	domain := ""
	for {
		length := int(packet[offset]);
		offset++

		if length == 0 {
			domain = domain[1:]
			break;
		}

		label := packet[offset:offset+length];
		offset+= length;
		domain = domain + "." + string(label);
	}
	qtypecode := binary.BigEndian.Uint16(packet[offset : offset+2])
	offset += 2
	qclasscode := binary.BigEndian.Uint16(packet[offset : offset+2])
	offset += 2

	var qtype string
	var qclass string 

	switch qtypecode {
	case 1:
		qtype = "IPv4"
	case 28:
		qtype = "IPv6"
	case 15:
		qtype = "MX"
	case 5:
		qtype = "CNAME"
	}
	
	switch qclasscode {
	case 1:
		qclass = "IN"
	default:
		fmt.Println("\nNot recognized QClass")
	}

	fmt.Println("\nQType:", qtype)
	fmt.Println("\nQClass:", qclass)
	
	response := make([]byte, 0, 512)
	response = binary.BigEndian.AppendUint16(response, id)
	response = binary.BigEndian.AppendUint16(response, 0x8180)
	response = binary.BigEndian.AppendUint16(response, 1) // QDCOUNT
	response = binary.BigEndian.AppendUint16(response, 1) // ANCOUNT
	response = binary.BigEndian.AppendUint16(response, 0) // NSCOUNT
	response = binary.BigEndian.AppendUint16(response, 0) // ARCOUNT

	response = append(response, packet[12:offset]...)
	response = append(response, 0xC0, 0x0C)
	response = binary.BigEndian.AppendUint16(response, 1)
	response = binary.BigEndian.AppendUint16(response, 1)
	response = binary.BigEndian.AppendUint32(response, 300)
	response = binary.BigEndian.AppendUint16(response, 4)
	response = append(response, 1, 2, 3, 4)

	return &Query {
		Response:response,
	}
}


func AnalizeQuery(packet []byte) *Query {
	query := newQuery(packet);
	return query 
}

func SendReferral(domain string) string {
    tld, err := GetTLD(domain)
    if err != nil {
        fmt.Println("err: %s", err)
    }
    fmt.Printf("sending the referral...")

    referral := CreateReferral(14, domain, tld)
    fmt.Println(referral)
    return "referral"
}

func GetTLD(domain string) (string, error) {
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
    for _, tld := range TLDs {
        if tld == TLD {
            return TLD, nil
        }
    }
    return "", errors.New("the domain is not recognized")
}
