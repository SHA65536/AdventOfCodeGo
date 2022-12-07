package day16

import (
	"fmt"
	"strconv"
)

// https://adventofcode.com/2021/day/16

func VersionSum(packets string) (string, error) {
	var sum uint64
	var bin = HexToBin(packets)
	var packet = ParsePacket(&bin)
	packet.SumVer(&sum)
	return strconv.FormatUint(sum, 10), nil
}

func EvalPackets(packets string) (string, error) {
	var res uint64
	var bin = HexToBin(packets)
	var packet = ParsePacket(&bin)
	res = packet.Eval()
	return strconv.FormatUint(res, 10), nil
}

// Parses a single packet and it's sub packets
func ParsePacket(data *BinStream) Packet {
	var res Packet
	version := data.Read(3) // First 3 bits version
	typeId := data.Read(3)  // Another 3 bits for type
	if typeId == 4 {        // Type 4 is literal
		res = ParseLiteral(data)
	} else { // Other types are operators
		lengthId := data.Read(1) // One bit for lengthId
		if lengthId == 0 {       // id 0 is for Length-Bound packets
			res = ParseLengthPacket(data)
		} else { // id 1 is for Count-Bound packets
			res = ParseCountPacket(data)
		}
	}
	res.Type, res.Version = typeId, version
	return res
}

func ParseLiteral(data *BinStream) Packet {
	var stop bool
	var value uint64
	for !stop {
		value <<= 4
		// First bit tells us if it's the last packet
		if data.Read(1) == 0 {
			stop = true
		}
		value += data.Read(4)
	}
	return Packet{LiteralValue: value}
}

func ParseCountPacket(data *BinStream) Packet {
	var sub []Packet
	// First 11 bits are the number of sub packages
	for count := data.Read(11); count > 0; count-- {
		sub = append(sub, ParsePacket(data))
	}
	return Packet{SubPackets: sub}
}

func ParseLengthPacket(data *BinStream) Packet {
	var sub []Packet
	length := data.Read(15)
	// First 15 bits tell us the length of the sub packages
	cur := len(*data)
	for uint64(cur-len(*data)) < length {
		sub = append(sub, ParsePacket(data))
	}
	return Packet{SubPackets: sub}
}

// Reads hexedecimal to a binary stream
func HexToBin(packets string) BinStream {
	var res BinStream = make([]byte, 0)
	for i := range packets {
		v, _ := strconv.ParseUint(packets[i:i+1], 16, 64)
		str := fmt.Sprintf("%04b", v)
		res = append(res, []byte(str)...)
	}
	var start, end = 0, len(res) - 1
	// Reversing so the first part will be at the end
	// ready for reading
	for end > start {
		res[start], res[end] = res[end], res[start]
		start++
		end--
	}
	return res
}

type Packet struct {
	Version      uint64
	Type         uint64
	LiteralValue uint64
	SubPackets   []Packet
}

// Sums up the versions of the packets
func (p *Packet) SumVer(res *uint64) {
	*res += p.Version
	for i := range p.SubPackets {
		p.SubPackets[i].SumVer(res)
	}
}

// Applies operations for the different packet types
func (p *Packet) Eval() uint64 {
	var res uint64
	switch p.Type {
	case 0: // Sum operation
		for i := range p.SubPackets {
			res += p.SubPackets[i].Eval()
		}
	case 1: // Product operation
		res = 1
		for i := range p.SubPackets {
			res *= p.SubPackets[i].Eval()
		}
	case 2: // Min operation
		res = ^uint64(0)
		for i := range p.SubPackets {
			val := p.SubPackets[i].Eval()
			if val < res {
				res = val
			}
		}
	case 3: // Max operation
		for i := range p.SubPackets {
			val := p.SubPackets[i].Eval()
			if val > res {
				res = val
			}
		}
	case 4: // Literal packet
		res = p.LiteralValue
	case 5: // Greater than operation
		if p.SubPackets[0].Eval() > p.SubPackets[1].Eval() {
			res = 1
		}
	case 6: // Less than operation
		if p.SubPackets[0].Eval() < p.SubPackets[1].Eval() {
			res = 1
		}
	case 7: // Equal operation
		if p.SubPackets[0].Eval() == p.SubPackets[1].Eval() {
			res = 1
		}
	}
	return res
}

type BinStream []byte

// Reading l bits from the end
func (s *BinStream) Read(l int) uint64 {
	value := []byte{}
	for i := 0; i < l; i++ {
		value = append(value, (*s)[len(*s)-1])
		*s = (*s)[:len(*s)-1]
	}
	ret, _ := strconv.ParseUint(string(value), 2, 64)
	return ret
}
