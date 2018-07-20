package common

import (
	"regexp"
	"strings"
	"github.com/XiuCai/XiuCaiAPI/model"
	"encoding/json"
	"math/rand"
	"time"
)

func CheckPhone(phone string) bool  {

	exp := "\\d{11}"
	ok, err := regexp.MatchString(exp, phone)
	if err != nil {
		return false
	}

	if !ok {
		return false
	}

	if strings.HasPrefix(phone, "11") || strings.HasPrefix(phone, "12") {
		return false
	} else {
		return true
	}
}

func Response(result int, message string, data interface{}) string {
	response := model.ResponseModel {
		Result: result,
		Message: message,
		Data: data,
	}

	jsonStr, _ := json.Marshal(response)

	return string(jsonStr)
}

func GenerateCode() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	return r.Intn(8999) + 1000
}

func BuildImgUrl(urlPrefix, url string) string {
	if len(url) == 0 {
		return ""
	} else if len(url) > 4 && url[0:4] == "http" {
		return  url
	} else {
		return urlPrefix + url
	}
}

func CurrentTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func CurrentDate() string {
	return time.Now().Format("2006-01-02")
}
