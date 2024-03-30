package nodes

type ReferralQuery struct {
    id uint16
    query uint16
    
    domain string
    tld string
}

func CreateReferral(ID uint16, DOM, TLD string) *ReferralQuery {
    var query uint16 = 1
    domain := DOM
    tld := TLD
    id := ID

    return &ReferralQuery {
        id,
        query,
        domain,
        tld,
    }
}
