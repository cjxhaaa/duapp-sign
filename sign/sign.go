package sign

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"unsafe"
)

func GetSignString(params map[string]string) string {
	signString := ""

	keys := make([]string, 0, len(params))
	for k, _ := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		signString += k + params[k]
	}
	signString += "048a9c4943398714b356a696503d2d36"
	return signString
}


func Rotl(e, n int) int {
	return jsLeft(e, n)| jsRight(e, 32 - n)
}

func Endian(e interface{}) interface{} {
	if ee, ok := e.(int); ok {
		return 16711935 & Rotl(ee, 8) | 4278255360 & Rotl(ee, 24)
	}
	if ee, ok := e.([]int); ok {
		for n := 0; n < len(ee); n++ {
			if eee, ok := Endian(ee[n]).(int); ok {
				ee[n] = eee
			} else {
				panic(errors.New("err Endian"))
			}

		}
		return ee
	}
	return nil
}

func BytesToWords(e []byte) []int {
	n := []int{}
	for i := 0; i < len(e); i++ {
		n = append(n, 0)
	}
	var r int = 0
	for t := 0; t < len(e); t++ {
		n[jsPositiveRight(r, 5)] |= (jsLeft(int(e[t]), 24 - r % 32) )

		r += 8
	}
	return n
}

func BytesToHex(e []byte) string {
	n := []string{}

	for t := 0; t < len(e); t++ {
		s1 := fmt.Sprintf("%x", jsRight(int(e[t]), 4))
		n = append(n,  string(s1[len(s1)-1]))
		s2 := fmt.Sprintf("%x", 15 & e[t])
		n = append(n, string(s2[len(s2)-1]))
	}
	return strings.Join(n, "")
}

func WordsToBytes(e []int) []byte {
	n := []uint8{}

	for t := 0; t < 32 * len(e); t+=8 {
		n = append(n, uint8(jsRight(e[jsRight(t, 5)], 24 - t % 32 & 255)))
	}
	return n
}

func Utf8StringToBytes(e string) []byte {
	ee, err := url.PathUnescape(url.PathEscape(e))
	if err != nil {
		panic(err)
	}
	return BinStringToBytes(ee)
}

func BinStringToBytes(e string) []byte {
	n := []byte{}
	for t := 0; t < len(e); t++ {
		n = append(n, 255 & uint8(e[t]))
	}
	return n
}

func BinBytesToString(e []byte) string {
	return string(e)
}

// js << 操作
func jsLeft(a, i int) int {
	r := jsint32(jsBInt32(a) << uint(i))
	if r >= 1 << 31 {
		return r - 1 << 32
	} else {
		return r
	}
}

// js >>> 操作
func jsPositiveRight(a, i int) int {
	r := jsint32(jsBInt32(a) >> uint(i))
	return r
}

// js >> 操作
func jsRight(a, i int) int {
	r := jsint32(jsBInt32(a) >> uint(i))
	if r >= 1 << 31 {
		return r - 1 << 32
	} else {
		return r
	}
}

// 转js int32
func jsint32(s int) int {
	ss := fmt.Sprintf("%08b",s)
	if len(ss) > 32 {
		kk, _ := strconv.ParseUint(ss[len(ss)-32:], 2, 32)
		return int(kk)
	} else {
		return s
	}
}

// (负数问题)转js 无符号int32
func jsBInt32(n int) int {
	var binStr string
	binStr = strconv.FormatUint(uint64(*(*uint64)(unsafe.Pointer(&n))), 2)
	if len(binStr) > 32 {
		binStr = binStr[len(binStr)-32:]
	}
	i, err := strconv.ParseUint(binStr, 2, 32)
	if err != nil {
		log.Fatal(err)
	}
	return int(i)
}

func aa(s, i, n int) int {
	return (jsLeft(s, i) | jsPositiveRight(s, 32-i))+ n

}

func ff(e, n, t, r, o, i, a int) int {
	s := e + (n & t | ^n & r) + jsPositiveRight(o, 0) + a
	//fmt.Println("o:", o, "s: ", s)
	return aa(s, i, n)
}

func gg(e, n, t, r, o, i, a int) int {
	s := e + (n & r | t & ^r) + jsPositiveRight(o, 0) + a
	return aa(s, i, n)
}

func hh(e, n, t, r, o, i, a int) int {
	s := e + (n ^ t ^ r) + jsPositiveRight(o, 0) + a
	return aa(s, i, n)
}

func ii(e, n, t, r, o, i, a int) int {
	s := e + (t ^ (n | ^r)) + jsPositiveRight(o, 0) + a
	return aa(s, i, n)
}

func arrayIndex(s []int, i int) int {
	if i >= len(s) {
		return 0
	} else {
		return s[i]
	}
}

func GetSign(e string) string  {
	ee := Utf8StringToBytes(e)
	s := BytesToWords(ee)
	c := 8 * len(ee)
	l := 1732584193
	d := -271733879
	u := -1732584194
	p := 271733878

	for k := 0; k < len(s); k ++ {
		s[k] = 16711935 & (jsLeft(s[k], 8) | jsPositiveRight(s[k], 24)) | 4278255360 & (jsLeft(s[k], 24) | jsPositiveRight(s[k], 8))
	}

	s[jsPositiveRight(c,5)] |= jsLeft(128, c % 32)

	s[14 + jsLeft(jsPositiveRight(c + 64, 9), 4)] = c

	var (
		A, b, y, x int
	)

	for h := 0; h < countLen(s); h+=16 {
		A = l
		b = d
		y = u
		x = p

		l = ff(l, d, u, p, arrayIndex(s, h + 0), 7, -680876936)
		p = ff(p, l, d, u, arrayIndex(s, h + 1), 12, -389564586)
		u = ff(u, p, l, d, arrayIndex(s, h + 2), 17, 606105819)
		d = ff(d, u, p, l, arrayIndex(s, h + 3), 22, -1044525330)
		l = ff(l, d, u, p, arrayIndex(s, h + 4), 7, -176418897)
		p = ff(p, l, d, u, arrayIndex(s, h + 5), 12, 1200080426)
		u = ff(u, p, l, d, arrayIndex(s, h + 6), 17, -1473231341)
		d = ff(d, u, p, l, arrayIndex(s, h + 7), 22, -45705983)
		l = ff(l, d, u, p, arrayIndex(s, h + 8), 7, 1770035416)
		p = ff(p, l, d, u, arrayIndex(s, h + 9), 12, -1958414417)
		u = ff(u, p, l, d, arrayIndex(s, h + 10), 17, -42063)
		d = ff(d, u, p, l, arrayIndex(s, h + 11), 22, -1990404162)
		l = ff(l, d, u, p, arrayIndex(s, h + 12), 7, 1804603682)
		p = ff(p, l, d, u, arrayIndex(s, h + 13), 12, -40341101)
		u = ff(u, p, l, d, arrayIndex(s, h + 14), 17, -1502002290)
		d = ff(d, u, p, l, arrayIndex(s, h + 15), 22, 1236535329)
		l = gg(l, d, u, p, arrayIndex(s, h + 1), 5, -165796510)
		p = gg(p, l, d, u, arrayIndex(s, h + 6), 9, -1069501632)
		u = gg(u, p, l, d, arrayIndex(s, h + 11), 14, 643717713)
		d = gg(d, u, p, l, arrayIndex(s, h + 0), 20, -373897302)
		l = gg(l, d, u, p, arrayIndex(s, h + 5), 5, -701558691)
		p = gg(p, l, d, u, arrayIndex(s, h + 10), 9, 38016083)
		u = gg(u, p, l, d, arrayIndex(s, h + 15), 14, -660478335)
		d = gg(d, u, p, l, arrayIndex(s, h + 4), 20, -405537848)
		l = gg(l, d, u, p, arrayIndex(s, h + 9), 5, 568446438)
		p = gg(p, l, d, u, arrayIndex(s, h + 14), 9, -1019803690)
		u = gg(u, p, l, d, arrayIndex(s, h + 3), 14, -187363961)
		d = gg(d, u, p, l, arrayIndex(s, h + 8), 20, 1163531501)
		l = gg(l, d, u, p, arrayIndex(s, h + 13), 5, -1444681467)
		p = gg(p, l, d, u, arrayIndex(s, h + 2), 9, -51403784)
		u = gg(u, p, l, d, arrayIndex(s, h + 7), 14, 1735328473)
		d = gg(d, u, p, l, arrayIndex(s, h + 12), 20, -1926607734)
		l = hh(l, d, u, p, arrayIndex(s, h + 5), 4, -378558)
		p = hh(p, l, d, u, arrayIndex(s, h + 8), 11, -2022574463)
		u = hh(u, p, l, d, arrayIndex(s, h + 11), 16, 1839030562)
		d = hh(d, u, p, l, arrayIndex(s, h + 14), 23, -35309556)
		l = hh(l, d, u, p, arrayIndex(s, h + 1), 4, -1530992060)
		p = hh(p, l, d, u, arrayIndex(s, h + 4), 11, 1272893353)
		u = hh(u, p, l, d, arrayIndex(s, h + 7), 16, -155497632)
		d = hh(d, u, p, l, arrayIndex(s, h + 10), 23, -1094730640)
		l = hh(l, d, u, p, arrayIndex(s, h + 13), 4, 681279174)
		p = hh(p, l, d, u, arrayIndex(s, h + 0), 11, -358537222)
		u = hh(u, p, l, d, arrayIndex(s, h + 3), 16, -722521979)
		d = hh(d, u, p, l, arrayIndex(s, h + 6), 23, 76029189)
		l = hh(l, d, u, p, arrayIndex(s, h + 9), 4, -640364487)
		p = hh(p, l, d, u, arrayIndex(s, h + 12), 11, -421815835)
		u = hh(u, p, l, d, arrayIndex(s, h + 15), 16, 530742520)
		d = hh(d, u, p, l, arrayIndex(s, h + 2), 23, -995338651)
		l = ii(l, d, u, p, arrayIndex(s, h + 0), 6, -198630844)
		p = ii(p, l, d, u, arrayIndex(s, h + 7), 10, 1126891415)
		u = ii(u, p, l, d, arrayIndex(s, h + 14), 15, -1416354905)
		d = ii(d, u, p, l, arrayIndex(s, h + 5), 21, -57434055)
		l = ii(l, d, u, p, arrayIndex(s, h + 12), 6, 1700485571)
		p = ii(p, l, d, u, arrayIndex(s, h + 3), 10, -1894986606)
		u = ii(u, p, l, d, arrayIndex(s, h + 10), 15, -1051523)
		d = ii(d, u, p, l, arrayIndex(s, h + 1), 21, -2054922799)
		l = ii(l, d, u, p, arrayIndex(s, h + 8), 6, 1873313359)
		p = ii(p, l, d, u, arrayIndex(s, h + 15), 10, -30611744)
		u = ii(u, p, l, d, arrayIndex(s, h + 6), 15, -1560198380)
		d = ii(d, u, p, l, arrayIndex(s, h + 13), 21, 1309151649)
		l = ii(l, d, u, p, arrayIndex(s, h + 4), 6, -145523070)
		p = ii(p, l, d, u, arrayIndex(s, h + 11), 10, -1120210379)
		u = ii(u, p, l, d, arrayIndex(s, h + 2), 15, 718787259)
		d = ii(d, u, p, l, arrayIndex(s, h + 9), 21, -343485551)
		l = jsPositiveRight(l + A, 0)
		d = jsPositiveRight(d + b, 0)
		u = jsPositiveRight(u + y, 0)
		p = jsPositiveRight(p + x,0)

	}

	if ed, ok := Endian([]int{l, d, u, p}).([]int); ok {
		r := WordsToBytes(ed)
		return BytesToHex(r)
	} else {
		panic("nibi")
	}
	panic("err return")
}


// 按js的方式计算len
func countLen(s []int) int {
	for i := len(s)-1; i >=0; i-- {
		if s[i] != 0 {
			return i
		}
	}
	return 0
}

