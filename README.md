
# 이 테스트를 하는 이유

stacktrace 를 보고싶어서
[golang error stack trace](https://www.popit.kr/golang-error-stack-trace%EC%99%80-%EB%A1%9C%EA%B9%85/) 글을보고 따라했는데
 controller => service => repository 로 이동도중 에러가 발생하면 지나왔던 경로의 log 는 모두 기록되는건가?
글을 봐도 errors.Wrap 언제사용하지?

작업중인 코드에 무작정 따라해 적용했다.
`stacktace`를 출력시켰다.
문제는
모든 함수에 `log.Errorf("%+v", err)` 를 찍어줬다...
글쓴이에게 문의하니 1번만 찍어주면 된다고 하셨다.
모르겠다 
큰 프로젝트에 적용하기 힘들기 때문에 테스트 만들어서 해보자

1. `errors.New` 는 왜사용?
2. `errors.New` 를 한 로직에 여러번 사용하면?
3. `errors.Wrap` 은 왜사용?
4. `log.Errorf` 왜 써?

