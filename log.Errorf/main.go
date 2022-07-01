
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
		//fmt.Printf("%+v", err) // 2.err 에 message 와 stack 을 console 에 출력
		//로그 패키지 (http://golang.site/go/article/114-Logging)
		log.Errorf("%+v", err)
	}
}
type Controller struct {}
func (Controller) Get() error {
	err := Service{}.GetData()
	if err != nil {
		return err
	}
	return nil
}

type Service struct {}
func (Service) GetData() error  {
	var s = `{"a:a"}`
	var result = map[string]interface{}{}
	err := json.Unmarshal([]byte(s), &result)
	if err != nil {
		return errors.Wrap(err, "오류 메세지~")
	}

	return nil
}
