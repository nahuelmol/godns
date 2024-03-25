package utils

import (
    "fmt"
    "encoding/base64"
)

func TakeFlags(buffer []byte){
    rawBinary := base64.StdEncoding.EncodeToString(buffer)
    fmt.Println("raw binary: ", rawBinary)

    var response string
    for i:=0;i<8; i++ {
        bit := (buffer >> i) & 1
        fmt.Printf("Bit %d: %d\n", i, bit)

        switch i {
        case 1: //identifier, allows client to match responses
            if bit == 0 {
                response += "qr"
            } else {
                response += "rs"
            }
        case 2://Operation Code the kind of query (4bits)
        case 3:
        case 4:
        case 5:
        
        case 6:
            if bit == 1 {
                //ath -> authorative
                response += " ath"
            }
        case 7:
            //trn truncated
            if bit == 1 {
                response += " trn"
            }
        case 8:
            //rd -> recursion desired
            if bit == 1 {
                response += " rd"
            }
        case 9:
            //recursion available
            //srq -> supports recursive query
            if bit == 1 {
                response += " srq"
            }
        }

        case 10: //reserved for future use (3bits)
        case 11:
        case 12:

        case 13://indicate response status (4bits)
        case 14:
        case 15:
        case 16:
        default:
            response += " end"
    }

    return response
}
