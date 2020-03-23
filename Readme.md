### 2020.3.23毒app的sign规则改了个版


## 写成了接口的形式

```$xslt
编译
cd cmd

go build main.go

运行

./main

流程很简单，拿着要请求的参数去请求一下加密api即可

如想得到https://m.poizon.com/router/product/ProductDetail?spuId=10242商品数据

打开浏览器请求http://localhost:9999/get_detail_sign?params={"spuId":"10242","productSourceName":"","propertyValueId":""}
直接返回sign参数文本信息9ca955484240e758f7f5b876ec72fd74

post：https://app.poizon.com/api/v1/h5/index/fire/flow/product/detail
data: {"spuId":"10242","productSourceName":"","propertyValueId":"","sign":"9ca955484240e758f7f5b876ec72fd74"}

发起请求即可获得官方商品数据


```