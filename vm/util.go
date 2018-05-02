package vm

import (
	"strconv"
	"strings"
	"fmt"
)

var commentPrefixes = []string{"#", "//", ";"}

type Nibble byte

func ConvByteToBitStringArray(byte byte) string {
	var strByte = strconv.FormatInt(int64(byte), 2)
	var padding = 4 - len(strByte)
	for i := 0; i < padding; i++ {
		strByte = "0" + strByte
	}
	return strByte
}

func ConvBitStringArrayToByte(bitArray string) byte {
	intByte, _ := strconv.ParseInt(bitArray, 2, 0)
	return byte(intByte)
}

//func RotateLCircular(a *[]bool, i int) {
//	x, b := (*a)[:i], (*a)[i:]
//	*a = append(b, x...)
//}

func RotateRCircular(a *string, i int) {
	for i > len(*a) {
		i = i - len(*a)
	}

	x, b := (*a)[:(len(*a) - i)], (*a)[(len(*a))-i:]
	*a = b + x
}

func SplitByteToNibbles(b byte) []Nibble {
	return []Nibble{Nibble((b >> 4) & 0x0F), Nibble(b & 0x0F)}
}

func CombineNibblesToByte(n1 Nibble, n2 Nibble) byte {
	return byte(n1<<4 | n2)
}

// 1 byte = 1 nibble
func BytesToNibbles(byteArr []byte) (nibbleArr []Nibble) {
	for _, b := range byteArr {
		nibbleArr = append(nibbleArr, Nibble(b))
	}
	return
}

// 1 byte = 2 nibbles
func ByteArrayToNibbleArray(bytes []byte) (nibbles []Nibble) {
	for _, byte := range bytes {
		nibbles = append(nibbles, SplitByteToNibbles(byte)...)
	}
	return nibbles
}

func IsIgnoredLine(item string) bool {
	if strings.TrimSpace(item) == "" {
		return true
	}
	for _, comment := range commentPrefixes {
		if strings.HasPrefix(item, comment) {
			return true
		}
	}
	return false
}

func RemoveTrailingComment(str string) string {
	for _, comment := range commentPrefixes {
		var commentIndex = strings.Index(str, comment)
		if commentIndex < 0 {
			continue
		}
		str = str[:strings.Index(str, comment)]
	}
	return str
}

func PrettyPrintNibbles(nibbles []Nibble) (output string) {

	for _, nibble := range nibbles {
		output += fmt.Sprintf("%X", nibble)
	}
	return output
}