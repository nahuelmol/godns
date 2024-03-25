package nodes

import (
    "fmt"
    "strconv"
)

type Record struct {
    hashmap map[string]int
    id int
}

func newRec(ip int, domain string) *Record {
    hm := make(map[string]int)
    hm[domain] = ip
    return &Record {
        hashmap:hm,
        id:1,
    }
}

func StartRecording (ip string, domain string) {

    ipint, err := strconv.Atoi(ip)
    if err != nil {
        fmt.Printf("there's an error %s \nconverting %s \n", err,  ip)
    }
    record:= newRec(ipint, domain)
    fmt.Printf("%s\n",record.id)

    //fmt.Printf("%s {}\n", record.hashmap['www.google.com'])
}
