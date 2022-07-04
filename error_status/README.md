#  erros.New 를 사용해서 발생한 에러는 internal server error 만 가능한가?

=>  github/pkg/errors 를 통해서 New 에러를 만들 때 code 값도 추가 가능하다고 함
=> but 다만 꼭 필요한 경우가 아니면 권하지 않음  error 를 핸들링하기 까다로워요


# error 를 쓰는 방식을

오픈소스에 가보니 중요 api 의 경우 이렇게 상태를 보냄
이따는 stacktrace 를 남기지않고 error 만 발생 왲?
```

		userClaim, err := jwtAuthentication.ConvertTokenUserClaim(accessToken)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
			}

```