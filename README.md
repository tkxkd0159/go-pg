Implement the regular backend server written by Go

# Setting
When your code imports packages contained in other modules, you manage those dependencies through your code's own module.  
Create `go.mod` files:
  
```bash
go mod init <github_src_path> # github.com/tkxkd0159/go-chain
go list -m all                # current module's dependencies
go list -m -versions <pkg>    # Show available versions of that package
go mod tidy                   # Remove unused dependencies & Update dependencies
go get [-d] <pkg>[@v1.0.0]    # Update dependency or Get a specific version of dependency
                              # d 플래그 붙이면 $GOPATH/src에 소스파일만 다운로드, 안붙이면 빌드해서 $GOPATH/bin에 바이너리까지 생성
			      # u 플래그 붙이면 해당 패키지 및 종속성 업데이트, v 플래그 붙이면 로그 출력
			      
go test                       # Test module (suffix : <filename>_test.go)
go run <target go file>       # Execute main.go
go build [mod_name]           # Build binary file from main.go => go build, go build . for cwd
go install [mod_name]         # go build + $GOPATH/bin 
                              # and caches all non-main packages which are imported to $GOPATH/pkg
go doc <pkg>                  # Show the document of pkg
```

```go
package main

import (
	"fmt"

	"rsc.io/quote"
)


func main() {
    fmt.Println(quote.Go())
}
```
