package dns

type String struct {
	Length uint8
	Value  string
}

type Query struct {
	Domain []String
	Type   uint16
	Class  uint16
}

type Packet struct {
	TransactionId uint16
	Flags         uint16
	NQuestions    uint16
	NAnswers      uint16
	NAuthorities  uint16
	NAdditionals  uint16

	Queries []Query
}

func appendUint16(buf []byte, val uint16) []byte {
	// dns uses big endian
	return append(buf, byte(val>>8), byte(val))
}

func (q *Query) Dump() []byte {
	var buf []byte

	for _, domain := range q.Domain {
		buf = append(buf, domain.Length)
		buf = append(buf, []byte(domain.Value)...)
	}
	buf = append(buf, 0x00)

	buf = appendUint16(buf, q.Type)
	buf = appendUint16(buf, q.Class)

	return buf
}

func (q *Packet) Dump() []byte {
	var buf []byte

	buf = appendUint16(buf, q.TransactionId)
	buf = appendUint16(buf, q.Flags)
	buf = appendUint16(buf, q.NQuestions)
	buf = appendUint16(buf, q.NAnswers)
	buf = appendUint16(buf, q.NAuthorities)
	buf = appendUint16(buf, q.NAdditionals)

	for _, query := range q.Queries {
		buf = append(buf, query.Dump()...)
	}

	return buf
}
