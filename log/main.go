package main

import (
	"encoding/json"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors" // 1. errors 패키지 대체
)

func main() {
	controller := Controller{}
	err := controller.Get()
	if err != nil {
		log.Errorf("%+v", err)
		//{"time":"2022-07-04T14:40:00.1899807+09:00","level":"ERROR","prefix":"-","file":"main.go","line":"13","message":"invalid character '}' after object key\n오류 메세지~\nmain.Service.GetData\n\tC:/gwiyeom/errorHandlingProject/log/main.go:31\nmain.Controller.Get\n\tC:/gwiyeom/errorHandlingProject/log/main.go:18\nmain.main\n\tC:/gwiyeom/errorHandlingProject/log/main.go:11\nruntime.main\n\tC:/Program Files/Go/src/runtime/proc.go:225\nruntime.goexit\n\tC:/Program Files/Go/src/runtime/asm_amd64.s:1371"}
		//출력됨
	}
}

type Controller struct{}

func (Controller) Get() error {
	err := Service{}.GetData()
	if err != nil {
		return err
	}
	return nil
}

type Service struct{}

func (Service) GetData() error {
	var s = `{"a:a"}`
	var result = map[string]interface{}{}
	err := json.Unmarshal([]byte(s), &result)
	if err != nil {
		return errors.Wrap(err, "오류 메세지~")
	}

	return nil
}
