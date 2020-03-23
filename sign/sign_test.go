package sign

import (
	"fmt"
	"log"
	"testing"
)



func TestGetSignString(t *testing.T) {
	params := make(map[string]string)
	params["recommendId"] = "73"
	params["lastId"] = ""
	si := SignTool.GetSignString(params)
	fmt.Println(si)
	if si != "lastIdrecommendId73048a9c4943398714b356a696503d2d36" {
		log.Fatal("sign string error")
	}
}



func TestGetSign(t *testing.T) {
	params := make(map[string]string)
	//params["spuId"] = "80566"
	//params["productSourceName"] = ""
	//params["propertyValueId"] = ""
	params["os"] = "bdapp"
	params["buildNumber"] = "2.0.0"
	fmt.Println(SignTool.GetSign(SignTool.GetSignString(params)))
	//fmt.Println(ss(1732584193,4023233417, 2562383102, 271733878, 1818850658, 7, 3614090360))
	//fmt.Println(aaa(4023233417, 2562383102, 271733878))
	//fmt.Println(jsAnd(4023233417, 2562383102))
}
