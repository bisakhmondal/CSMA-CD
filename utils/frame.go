package utils

import (
	"strings"
)
import "CSMA/utils/crc"
type mac=int

const FRAMELENGTH = 72*8
func MakeFrame(src, dest mac, data int, nodeinfo int ) Bytestream{
	frame := ""
	SFD := "10101011"
	preamble := []rune(encode(0, 7*8)) //7 Byte
	for i :=0;i<len(preamble);i+=2{
		preamble[i]='1'
	}
	frame+=string(preamble)
	frame += SFD
	frame += encode(dest, 6*8)
	frame += encode(src, 6*8)
	frame += encode(nodeinfo, 2*8)
	frame += encode(data, 46*8)
	frame = crc.Generate(frame, crc.CRC32)
	return frame
}

func GetData(b Bytestream) int{
	offset := 8+7*8 + 2*6*8 + 2*8
	dataarr := []rune(b)[offset: offset+46*8]

	return decode(dataarr)
}
func GetNodeInfo(b Bytestream) int{
	offset := 8+7*8 + 2*6*8
	dataarr := []rune(b)[offset: offset+ 2*8]

	return decode(dataarr)
}

func extract(d int)[]string{
	if d==0{
		return []string{}
	}
	rem := d%2

	slice := extract(d/2)
	if rem==1{
		return append(slice, "1")
	}
	return append(slice, "0")
}
func decode(arr []rune) int{
	var value int
	for i:=0; i<len(arr); i++{
		value *=2
		if arr[i]=='1'{
			value +=1
		}
	}

	return value
}
func encode(data int, lenT int) Bytestream{
	var encoded []string = extract(data)

	return strings.Repeat("0", lenT - len(encoded)) + strings.Join(encoded,"")
}

