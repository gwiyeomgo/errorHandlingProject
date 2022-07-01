
package main
import (
	"encoding/json"
	"fmt"
"github.com/pkg/errors" // 1. errors 패키지 대체
)
func main() {
	controller := Controller{}
	err := controller.Get()
	if err != nil {
		fmt.Printf("%+v", err) //  2.err 에 message 와 stack 을 console 에 출력
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
	// 아...encoding/json 패키지의 json 모듈로 Unmarshal 시 발생한 오류는 stack 이 없다.
	// stack 이 없는 발생한 오류를 감싸면
	// 직접입력한 오류메세지와 err 내용이 같이 출력된다.
	// 즉 가장마지막에 써줌
	return nil
}
