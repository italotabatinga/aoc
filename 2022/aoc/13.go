package aoc

import (
	"fmt"
	"strconv"
	"strings"

	c "github.com/italotabatinga/aoc/2022/aoc/collections"
)

type Input13 []c.Pair[PacketData]

type Runner13 struct{}

func (r Runner13) FmtInput(input string) Input13 {
	packetPairStrings := strings.Split(input, "\n\n")
	result := []c.Pair[PacketData]{}
	for _, lines := range packetPairStrings {
		packetStrings := strings.Split(lines, "\n")
		firstPacket, _ := ParsePacketString(packetStrings[0], 0)
		secondPacket, _ := ParsePacketString(packetStrings[1], 0)

		result = append(result, c.Pair[PacketData]{First: firstPacket, Second: secondPacket})
	}

	return Input13(result)
}

func (r Runner13) Run1(input Input13, _ bool) int {
	sum := 0
	for i, pair := range input {

		if rightOrder := IsRightOrder(pair.First, pair.Second); rightOrder == 1 {
			sum += i + 1
		}
	}
	return sum
}

func (r Runner13) Run2(input Input13, _ bool) int {
	packets := []PacketData{
		PacketList{Data: []PacketData{PacketList{Data: []PacketData{PacketInt{Data: 2}}}}},
		PacketList{Data: []PacketData{PacketList{Data: []PacketData{PacketInt{Data: 6}}}}},
	}
	for _, pair := range input {
		packets = append(packets, pair.First, pair.Second)
	}
	Sort(packets, func(a, b PacketData) int { return IsRightOrder(a, b) })

	result := 1
	for i, packet := range packets {
		if packet, ok := packet.(PacketList); ok {
			if len(packet.Data) == 1 {
				packet := packet.Data[0]
				if packet, ok := packet.(PacketList); ok {
					if len(packet.Data) == 1 {
						packet := packet.Data[0]
						if packet, ok := packet.(PacketInt); ok {
							if packet.Data == 2 || packet.Data == 6 { // [[2]] or [[6]]
								result *= (i + 1)
							}
						}
					}
				}
			}
		}
	}
	return result
}

func ParsePacketString(str string, i int) (PacketData, int) {
	list := PacketList{Data: []PacketData{}}
	i++
	for i < len(str) {
		elem := rune(str[i])
		switch elem {
		case '[':
			packetData, newI := ParsePacketString(str, i)
			list.Data = append(list.Data, packetData)
			i = newI + 1
		case ']':
			return list, i
		case ',':
			i++
		default:
			j := i
			for rune(str[j]) != ']' && rune(str[j]) != ',' {
				j++
			}
			data, _ := strconv.Atoi(str[i:j])
			list.Data = append(list.Data, PacketInt{Data: data})
			i = j
		}
	}

	return list, i
}

func IsRightOrder(l PacketData, r PacketData) int {
	if l.DataType() == PACKET_INT && r.DataType() == PACKET_INT {
		lInt, _ := l.(PacketInt)
		rInt, _ := r.(PacketInt)
		if lInt.Data > rInt.Data {
			return -1
		} else if lInt.Data < rInt.Data {
			return 1
		} else {
			return 0
		}
	} else if l.DataType() == PACKET_LIST && r.DataType() == PACKET_LIST {
		i := 0
		lList, _ := l.(PacketList)
		rList, _ := r.(PacketList)
		for i < len(lList.Data) && i < len(rList.Data) {
			if cmp := IsRightOrder(lList.Data[i], rList.Data[i]); cmp == 0 {
				i++
			} else {
				return cmp
			}
		}

		if len(lList.Data) > len(rList.Data) {
			return -1
		} else if len(lList.Data) < len(rList.Data) {
			return 1
		} else {
			return 0
		}
	} else if l.DataType() == PACKET_INT {
		return IsRightOrder(PacketList{Data: []PacketData{l}}, r)
	} else if r.DataType() == PACKET_INT {
		return IsRightOrder(l, PacketList{Data: []PacketData{r}})
	} else {
		panic(fmt.Errorf("not expected data types l %v, r %v", l.DataType(), r.DataType()))
	}
}

type PacketData interface {
	DataType() string
	String() string
}

type PacketInt struct {
	Data int
}

func (p PacketInt) DataType() string { return PACKET_INT }

func (p PacketInt) String() string {
	return fmt.Sprintf("%d", p.Data)
}

type PacketList struct {
	Data []PacketData
}

func (p PacketList) DataType() string { return PACKET_LIST }

func (p PacketList) String() string {
	var sb strings.Builder
	sb.WriteRune('[')
	for i, elem := range p.Data {
		sb.WriteString(fmt.Sprintf("%v", elem))
		if i+1 != len(p.Data) {
			sb.WriteRune(',')
		}
	}
	sb.WriteRune(']')
	return sb.String()
}

const (
	PACKET_INT  = "PACKET_INT"
	PACKET_LIST = "PACKET_LIST"
)
