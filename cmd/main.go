package main

import (
	"fmt"
	"duapp/sign"
)

func main()  {
	url := "https://app.poizon.com/api/v1/h5/index/fire/flow/product/detail?productId=36191&productSourceName=undefined&sign="
	params := make(map[string]string)
	params["productId"] = "36191"
	params["productSourceName"] = "undefined"
	si := sign.GetSignString(params)
	fmt.Println(si)
	sig := sign.GetSign(si)
	url += sig
	fmt.Println(url)
	//url := "https://m.poizon.com/mapi/product/recommendDetail?recommendId=73&lastId=&sign="
	//params := make(map[string]string)
	//params["recommendId"] = "73"
	//params["lastId"] = ""
	//si := sign.GetSignString(params)
	//sig := sign.GetSign(si)
	//url += sig
	//fmt.Println(url)
}

