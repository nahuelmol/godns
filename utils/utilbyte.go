package utils

func TakeFlags(buffer []byte) string {

    var response string
    i := 1
    for _, bit := range buffer {

        switch i {
        case 1: //identifier, allows client to match responses
            if bit == 1 {
                response += " id"
            }
        case 2://Operation Code the kind of query (4bits)
            if bit == 1 {
                response += " q1"
            }
        case 3:
            if bit == 1 {
                response += " q2"
            }
        case 4:
            if bit == 1 {
                response += " q3"
            }
        case 5:
            if bit == 1 {
                response += " q4"
            }
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

        case 10: //reserved for future use (3bits)
            if bit == 1 {
                response += " rsrv1"
            }
        case 11:
            if bit == 1 {
                response += " rsrv2"
            }
        case 12:
            if bit == 1 {
                response += " rsrv3"
            }

        case 13://indicate response status (4bits)
            if bit == 1 {
                response += " rstatus1"
            }
        case 14:
            if bit == 1 {
                response += " rstatus2"
            }
        case 15:
            if bit == 1 {
                response += " rstatus3"
            }
        case 16:
            if bit == 1 {
                response += " rstatus4"
            }
        }
        i+=1
    }

    return response
}
