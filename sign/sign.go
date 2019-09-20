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
	s := e + (n & t | ^n & r) + (o >> 0) + a
	return aa(s, i, n)
}

func gg(e, n, t, r, o, i, a int) int {
	s := e + (n & r | t & ^r) + (o >> 0) + a
	return aa(s, i, n)
}

func hh(e, n, t, r, o, i, a int) int {
	s := e + (n ^ t ^ r) + (o >> 0) + a
	return aa(s, i, n)
}

func ii(e, n, t, r, o, i, a int) int {
	s := e + (t ^ (n | ^r)) + (o >> 0) + a
	return aa(s, i, n)
}

func GetSign(e string) string  {
	ee := Utf8StringToBytes(e)
	s := BytesToWords(ee)
	c := 8 * len(ee)
	u := 1732584193
	l := -271733879
	d := -1732584194
	f := 271733878

	for p := 0; p < len(s); p ++ {
		s[p] = 16711935 & (jsLeft(s[p], 8) | jsPositiveRight(s[p], 24)) | 4278255360 & (jsLeft(s[p], 24) | jsPositiveRight(s[p], 8))
	}

	s[jsPositiveRight(c,5)] |= jsLeft(128, c % 32)
	s[14 + jsLeft(jsPositiveRight(c + 64, 9), 4)] = c
	for p := 0; p < len(s); p++ {
		_A := u
		y := l
		b := d
		_C := f

		u = ff(u, l, d, f, s[p + 0], 7, -680876936)
		f = ff(f, u, l, d, s[p + 1], 12, -389564586)
		d = ff(d, f, u, l, s[p + 2], 17, 606105819)
		l = ff(l, d, f, u, s[p + 3], 22, -1044525330)
		u = ff(u, l, d, f, s[p + 4], 7, -176418897)
		f = ff(f, u, l, d, s[p + 5], 12, 1200080426)
		d = ff(d, f, u, l, s[p + 6], 17, -1473231341)
		l = ff(l, d, f, u, s[p + 7], 22, -45705983)
		u = ff(u, l, d, f, s[p + 8], 7, 1770035416)
		f = ff(f, u, l, d, s[p + 9], 12, -1958414417)
		d = ff(d, f, u, l, s[p + 10], 17, -42063)
		l = ff(l, d, f, u, s[p + 11], 22, -1990404162)
		u = ff(u, l, d, f, s[p + 12], 7, 1804603682)
		f = ff(f, u, l, d, s[p + 13], 12, -40341101)
		d = ff(d, f, u, l, s[p + 14], 17, -1502002290)
		l = ff(l, d, f, u, s[p + 15], 22, 1236535329)
		u = gg(u, l, d, f, s[p + 1], 5, -165796510)
		f = gg(f, u, l, d, s[p + 6], 9, -1069501632)
		d = gg(d, f, u, l, s[p + 11], 14, 643717713)
		l = gg(l, d, f, u, s[p + 0], 20, -373897302)
		u = gg(u, l, d, f, s[p + 5], 5, -701558691)
		f = gg(f, u, l, d, s[p + 10], 9, 38016083)
		d = gg(d, f, u, l, s[p + 15], 14, -660478335)
		l = gg(l, d, f, u, s[p + 4], 20, -405537848)
		u = gg(u, l, d, f, s[p + 9], 5, 568446438)
		f = gg(f, u, l, d, s[p + 14], 9, -1019803690)
		d = gg(d, f, u, l, s[p + 3], 14, -187363961)
		l = gg(l, d, f, u, s[p + 8], 20, 1163531501)
		u = gg(u, l, d, f, s[p + 13], 5, -1444681467)
		f = gg(f, u, l, d, s[p + 2], 9, -51403784)
		d = gg(d, f, u, l, s[p + 7], 14, 1735328473)
		l = gg(l, d, f, u, s[p + 12], 20, -1926607734)
		u = hh(u, l, d, f, s[p + 5], 4, -378558)
		f = hh(f, u, l, d, s[p + 8], 11, -2022574463)
		d = hh(d, f, u, l, s[p + 11], 16, 1839030562)
		l = hh(l, d, f, u, s[p + 14], 23, -35309556)
		u = hh(u, l, d, f, s[p + 1], 4, -1530992060)
		f = hh(f, u, l, d, s[p + 4], 11, 1272893353)
		d = hh(d, f, u, l, s[p + 7], 16, -155497632)
		l = hh(l, d, f, u, s[p + 10], 23, -1094730640)
		u = hh(u, l, d, f, s[p + 13], 4, 681279174)
		f = hh(f, u, l, d, s[p + 0], 11, -358537222)
		d = hh(d, f, u, l, s[p + 3], 16, -722521979)
		l = hh(l, d, f, u, s[p + 6], 23, 76029189)
		u = hh(u, l, d, f, s[p + 9], 4, -640364487)
		f = hh(f, u, l, d, s[p + 12], 11, -421815835)
		d = hh(d, f, u, l, s[p + 15], 16, 530742520)
		l = hh(l, d, f, u, s[p + 2], 23, -995338651)
		u = ii(u, l, d, f, s[p + 0], 6, -198630844)
		f = ii(f, u, l, d, s[p + 7], 10, 1126891415)
		d = ii(d, f, u, l, s[p + 14], 15, -1416354905)
		l = ii(l, d, f, u, s[p + 5], 21, -57434055)
		u = ii(u, l, d, f, s[p + 12], 6, 1700485571)
		f = ii(f, u, l, d, s[p + 3], 10, -1894986606)
		d = ii(d, f, u, l, s[p + 10], 15, -1051523)
		l = ii(l, d, f, u, s[p + 1], 21, -2054922799)
		u = ii(u, l, d, f, s[p + 8], 6, 1873313359)
		f = ii(f, u, l, d, s[p + 15], 10, -30611744)
		d = ii(d, f, u, l, s[p + 6], 15, -1560198380)
		l = ii(l, d, f, u, s[p + 13], 21, 1309151649)
		u = ii(u, l, d, f, s[p + 4], 6, -145523070)
		f = ii(f, u, l, d, s[p + 11], 10, -1120210379)
		d = ii(d, f, u, l, s[p + 2], 15, 718787259)
		l = ii(l, d, f, u, s[p + 9], 21, -343485551)
		u = jsPositiveRight(u + _A, 0)
		l = jsPositiveRight(l + y, 0)
		d = jsPositiveRight(d + b, 0)
		f = jsPositiveRight(f + _C,0)
		p += 16
		break
	}

	if ed, ok := Endian([]int{u, l, d, f}).([]int); ok {
		r := WordsToBytes(ed)
		return BytesToHex(r)
	} else {
		panic("nibi")
	}
	panic("err return")
}


