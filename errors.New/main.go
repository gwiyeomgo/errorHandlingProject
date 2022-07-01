
package main
import (
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
	return errors.New("에러 발생")
	//알게된 사실
	// New returns an error with the supplied message.
	// New also records the stack trace at the point it was called.
}
