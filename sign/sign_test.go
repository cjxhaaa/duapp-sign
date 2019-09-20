package sign

import (
	"fmt"
	"log"
	"testing"
)

func TestUtf8StringToBytes(t *testing.T) {
	fmt.Println(Utf8StringToBytes("cjxh"))
}

func TestBinStringToBytes(t *testing.T) {
	fmt.Println(^0x0f)
	fmt.Println(BinStringToBytes("cjxh"))
}

func TestAll(t *testing.T) {
	//fmt.Println(uint32(4294967295))
	var (
		e = 1732584193
		n = -271733879
		tt = -1732584194
		r = 271733878
		o = 1953718636
		i = 7
		a = -680876936
	)
	if ff(e, n, tt, r, o, i, a) != -556753618 {
		log.Fatal("ff error")
	}

	if gg(e, n, tt, r, o, i, a) != -978951642 {
		log.Fatal("gg error")
	}
	if hh(e, n, tt, r, o, i, a) != 603168406 {
		log.Fatal("hh error")
	}
	if ii(e, n, tt, r, o, i, a) != 1025366430 {
		log.Fatal("ii error")
	}



}

func TestGetSignString(t *testing.T) {
	params := make(map[string]string)
	params["recommendId"] = "73"
	params["lastId"] = ""
	si := GetSignString(params)
	if si != "lastIdrecommendId73048a9c4943398714b356a696503d2d36" {
		log.Fatal("sign string error")
	}
}



func TestGetSign(t *testing.T) {
	//fmt.Println(GetSign("lastIdrecommendId73048a9c4943398714b356a696503d2d36"))
	ee := Utf8StringToBytes("lastIdrecommendId73048a9c4943398714b356a696503d2d36")
	s := BytesToWords(ee)
	c := 8 * len(ee)
	u := 1732584193
	l := -271733879
	d := -1732584194
	f := 271733878

	for p := 0; p < len(s); p ++ {
		s[p] = 16711935 & (jsLeft(int(s[p]), 8) | jsRight(int(s[p]), 24)) | 4278255360 & (jsLeft(int(s[p]), 24) | jsRight(int(s[p]), 8))
	}
	s[jsRight(c,5)] |= jsLeft(128, c % 32)
	s = append(s, 0)
	s = append(s, c)
	for p := 0; p < len(s); p++ {
		_A := u
		y := l
		b := d
		_C := f

		u = ff(u, l, d, f, s[p + 0], 7, -680876936)
		if u != -556753618 {
			log.Fatal("u error s[p + 0]")
		}
		f = ff(f, u, l, d, s[p + 1], 12, -389564586)
		if f != 1383034883 {
			log.Fatal("f error s[p + 1]")
		}
		d = ff(d, f, u, l, s[p + 2], 17, 606105819)
		if d != 66831381 {
			log.Fatal("d error s[p + 2]")
		}

		//[-271733879, 66831381, 1383034883, -556753618, 1231318629]

		l = ff(l, d, f, u, s[p + 3], 22, -1044525330)
		if l != -2047730971 {
			log.Fatal("l error s[p + 3]")
		}
		u = ff(u, l, d, f, s[p + 4], 7, -176418897)
		if u != -980140271 {
			log.Fatal("u error s[p + 4]")
		}
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
		l = ff(l, d, f, u, 0, 22, 1236535329)
		u = gg(u, l, d, f, s[p + 1], 5, -165796510)
		f = gg(f, u, l, d, s[p + 6], 9, -1069501632)
		d = gg(d, f, u, l, s[p + 11], 14, 643717713)
		l = gg(l, d, f, u, s[p + 0], 20, -373897302)
		u = gg(u, l, d, f, s[p + 5], 5, -701558691)
		f = gg(f, u, l, d, s[p + 10], 9, 38016083)
		d = gg(d, f, u, l, 0, 14, -660478335)
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
		d = hh(d, f, u, l, 0, 16, 530742520)
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
		f = ii(f, u, l, d, 0, 10, -30611744)
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
	if u != 3459472981 || l != 2897701410 || d != 1828735935 || f != 1670330112 {
		log.Fatal("u, l, d, f error")
	}

	kk := Endian([]int{u, l, d, f})
	fmt.Println(kk)
	if ee ,ok := kk.([]int); ok {
		r := WordsToBytes(ee)
		fmt.Println(BytesToHex(r))
	}

}

func TestGetSign1(t *testing.T) {
	//[-271733879, 66831381, 1383034883, -556753618, 1231318629]
	fmt.Println(aa(-648330745, 22, 66831381))
	var (
		s = -648330745
		i = 22
		n = 66831381
	)
	fmt.Println((jsLeft(s, i) | jsPositiveRight(s, 32-i))+ n)
	fmt.Println("<<",jsLeft(s, i))

	fmt.Println(ff(-271733879, 66831381, 1383034883, -556753618,1231318629, 22, -1044525330))
}
