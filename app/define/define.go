package define

// ReturnDto 结构体
type ReturnDto struct {
	Success bool        `json:"success"`
	Message string      `json:"messgae"`
	Data    interface{} `json:"data"`
}

func ReturnDefault(msg string) *ReturnDto {
	return &ReturnDto{
		Success: true,
		Message: msg,
		Data:    nil,
	}
}

func ReturnOk(o interface{}) *ReturnDto {
	return &ReturnDto{
		Success: true,
		Message: "ok",
		Data:    o,
	}
}

func ReturnFail(msg string) *ReturnDto {
	return &ReturnDto{
		Success: false,
		Message: msg,
	}
}
