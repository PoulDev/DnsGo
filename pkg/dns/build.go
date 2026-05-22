package dns

import (
	"strings"
	"math/rand"
)

func NewQuery(domain string, qtype uint16, qclass uint16) Packet {
	var domainStrings []String

	for _, part := range strings.Split(domain, ".") {
		domainStrings = append(domainStrings, String{Length: uint8(len(part)), Value: part})
	}

	return Packet{
		TransactionId: uint16(rand.Intn(0xFFFF)),
		Flags:         0x0100,
		NQuestions:    1,
		NAnswers:      0,
		NAuthorities:  0,
		NAdditionals:  0,
		Queries: []Query{
			{
				Domain: domainStrings,
				Type:  qtype,
				Class: qclass,
			},
		},
	}
}
