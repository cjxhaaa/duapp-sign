package main

import (
	"duapp/sign"
	"fmt"
)

func main() {
	url := "https://m.poizon.com/mapi/product/recommendDetail?recommendId=137&sign="
	params := make(map[string]string)
	params["recommendId"] = "137"
	si := sign.GetSignString(params)
	sig := sign.GetSign(si)
	url += sig
	fmt.Println(url)
}
