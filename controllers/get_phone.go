package controllers

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

type GetPhoneReq struct {
	Code string `json:"code"`
}

type GetPhoneResp struct {
	ErrMsg string `json:"err_msg"`
	Phone  string `json:"phone"`
}

type PhoneInfo struct {
	PhoneNumber string `json:"phoneNumber"`
}
type PhoneResp struct {
	ErrCode   int    `json:"errcode"`
	ErrMsg	  string `json:"errmsg"`
	PhoneInfo PhoneInfo `json:"phone_info" json:"phone_info"`
}

func GetPhone(c *gin.Context) {
	var req GetPhoneReq
	var resp GetPhoneResp
	if err := c.ShouldBindJSON(&req); err != nil {
		logrus.Error(err)
		resp.ErrMsg = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	logrus.Debugf("GetPhone:%v",req)

	jsonStr := fmt.Sprintf(`{ "code": "%s" }`,req.Code)
	url:= "https://api.weixin.qq.com/wxa/business/getuserphonenumber"
	r, err := http.NewRequest("POST", url, bytes.NewBuffer( []byte(jsonStr)))
	r.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	re, err := client.Do(r)
	if err != nil {
		fmt.Printf("err:%s",err.Error())
	}
	defer re.Body.Close()
	if  re.StatusCode != 200 {
		err = fmt.Errorf("get phone err, status:%d", re.StatusCode)
		logrus.Error(err)
		resp.ErrMsg = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	rb, _ := ioutil.ReadAll(re.Body)
	fmt.Println(string(rb))
	pr := PhoneResp{}
	err = json.Unmarshal(rb,&pr)
	if err != nil {
		logrus.Error(err)
		resp.ErrMsg = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp.ErrMsg = "success"
	resp.Phone = pr.PhoneInfo.PhoneNumber
	c.JSON(http.StatusOK, resp)
	logrus.Info("success",resp)
	return

}
