package api

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"
	"shProxy/serializer"
	"strings"

	"github.com/gin-gonic/gin"
)

//MyPing 测试
func MyPing(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Code: 200,
		Msg:  "hello",
	})
}

// TestProxy 代理测试
func TestProxy(c *gin.Context) {
	//localhost:3000/shProxyApi/v1/test?test=1
	test := c.Query("test")
	log.Println("参数测试", test)

	url := "https://fanyi-api.baidu.com/api/trans/vip/translate"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println(err)
	}
	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		log.Println(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	log.Println("响应:", string(body))
	c.Data(200, "application/json", body)
}

// GetProxy 代理Get方式 传递一个url
func GetProxy(c *gin.Context) {
	//读取参数
	base64Url := c.Query("url")
	if strings.TrimSpace(base64Url) == "" {
		c.JSON(http.StatusBadRequest, serializer.Response{
			Code: 400,
			Msg:  "url参数不能为空",
		})
		return
	}
	//解码
	urlBytes, decodeErr := base64.URLEncoding.DecodeString(base64Url)
	if decodeErr != nil {
		log.Println("参数解码失败：", decodeErr)
		c.JSON(http.StatusBadRequest, ErrorResponse(decodeErr))
		return
	}
	url := string(urlBytes)
	log.Println("接收到的url", base64Url, "\n解码后", url)
	//建立请求
	request, reqErr := http.NewRequest(http.MethodGet, url, nil)
	if reqErr != nil {
		log.Println("建立请求失败：", reqErr)
		c.JSON(http.StatusBadRequest, ErrorResponse(reqErr))
		return
	}

	//发送请求
	client := &http.Client{}
	response, repErr := client.Do(request)
	if repErr != nil {
		log.Println("发送请求失败：", repErr)
		c.JSON(http.StatusExpectationFailed, ErrorResponse(repErr))
		return
	}
	defer response.Body.Close()

	//读取响应数据
	data, redErr := ioutil.ReadAll(response.Body)
	if redErr != nil {
		log.Println("数据读取失败：", redErr)
		c.JSON(http.StatusExpectationFailed, ErrorResponse(redErr))
		return
	}

	//返回数据
	c.Data(http.StatusOK, "application/json", data)
	return
}
