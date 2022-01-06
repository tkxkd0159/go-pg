# Command
처음에는 프로젝트의 `go.mod`의 모듈이름을 기준으로 임포트된 모듈들을 찾고 없으면 `$GOPATH`, `$GOROOT`에서 찾아본다.  
패키지 이름과 그 패키지를 담는 폴더이름은 달라도 됨.
```bash
go build
go run .
```
