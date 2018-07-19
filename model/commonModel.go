package model

type (
	ResponseModel struct{
		Result int `json:"result"`
		Message string `json:"message"`
		Data interface{} `json:"data"`
	}
	
	CodeModel struct {
		Code int `json:"code"`
	}
)
