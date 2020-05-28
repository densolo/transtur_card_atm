package utils

import (
	"regexp"
	"encoding/binary"
)

func ReadInt16(data []byte, pos int)(num int) {
	return int(binary.BigEndian.Uint16(data[pos:pos+2]))
}

func ReadInt8(data []byte, pos int)(num int) {
	return int(data[pos])
}

func ReadString(data []byte, pos int, size int)(txt string) {
	txt = string(data[pos:pos+size])
	txt = regexp.MustCompile(`\s+$`).ReplaceAllString(txt, ``)
	return txt
}

func ReadCodeString(data []byte, pos int, size int)(txt string) {
	txt = string(data[pos:pos+size])
	txt = txt[1:len(txt)-1]
	txt = regexp.MustCompile(`\0`).ReplaceAllString(txt, ` `)
	txt = regexp.MustCompile(`\s+$`).ReplaceAllString(txt, ``)
	return txt
}
