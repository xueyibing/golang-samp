package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// JsonResult 返回结构
type userInfoResult struct {
	Code     int         `json:"code"`
	ErrorMsg string      `json:"errorMsg,omitempty"`
	Data     interface{} `json:"data"`
}


// CounterHandler 计数器接口
func GetUserinfoHandler(w http.ResponseWriter, r *http.Request) {
	res := &userInfoResult{}
	fmt.Println("GetUserinfoHandler")
	if r.Method == http.MethodPost {

		decoder := json.NewDecoder(r.Body)
		body := make(map[string]interface{})
		if err := decoder.Decode(&body); err != nil {
			fmt.Println("decode error")
			return
		}
		defer r.Body.Close()
		code, ok := body["code"]
		if !ok {
			fmt.Println("缺少 code 参数")
			res.ErrorMsg = "缺少 code 参数"
			return
		}
		fmt.Printf("code:%s",code)

		jsonStr := fmt.Sprintf(`{ "code": "%s" }`,code)


		url:= "https://api.weixin.qq.com/wxa/business/getuserphonenumber"
		req, err := http.NewRequest("POST", url, bytes.NewBuffer( []byte(jsonStr)))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("err:%s",err.Error())
		}
		defer resp.Body.Close()


		statuscode := resp.StatusCode
		hea := resp.Header
		rb, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(rb))
		fmt.Println(statuscode)
		fmt.Println(hea)





	} else {
		res.Code = -1
		res.ErrorMsg = fmt.Sprintf("请求方法 %s 不支持", r.Method)
	}

	msg, err := json.Marshal(res)
	if err != nil {
		fmt.Fprint(w, "内部错误")
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(msg)
}
