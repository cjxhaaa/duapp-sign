package main

import (
	"duapp/sign"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/url"
)

func main()  {
	engine := gin.Default()
	engine.GET("/getsign",GetSign)
	engine.GET("/get_detail_sign",GetDetailSign)
	engine.Run(":9999")
}


func GetSign(c *gin.Context) {
	var tmp = struct{
		Url  string  `form:"url" binding:"required"`
	}{}
	if err := c.ShouldBindQuery(&tmp); err != nil {
		log.Fatal(err)
	}

	uu, err := url.Parse(tmp.Url)
	if err != nil {
		panic(err)
	}

	uq := uu.Query()
	fmt.Println(uq)
	params := make(map[string]string)
	for k, v := range uu.Query() {
		params[k] = v[0]
	}
	si := sign.SignTool.GetSignString(params)
	sig := sign.SignTool.GetSign(si)
	uq.Add("sign", sig)
	uu.RawQuery = uq.Encode()

	c.String(http.StatusOK, uu.String())
}



func GetDetailSign(c *gin.Context) {
	var pa struct {
		Params  string  `form:"params" binding:"required"`
	}
	if err := c.ShouldBindQuery(&pa); err != nil {
		log.Fatal(err)
	}

	params := make(map[string]string)
	qp, _ := url.QueryUnescape(pa.Params)
	if err := json.Unmarshal([]byte(qp), &params); err != nil {
		panic(err)
	}

	si := sign.SignTool.GetSignString(params)
	sig := sign.SignTool.GetSign(si)

	c.String(http.StatusOK, sig)
}

