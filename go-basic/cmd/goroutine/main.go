package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2) // {{} 8589934592 0}

	ctx, cancel := context.WithCancel(context.Background())
	subctx := context.WithValue(ctx, "subkey", "Conveyed value")
	newctx, cancel2 := context.WithCancel(context.Background())

	i := 0
	go func(ctx context.Context) {
		defer wg.Done()

		go func(ctx context.Context) {
			defer wg.Done()
		subloop:
			for {
				i++

				select {
				case <-subctx.Done():
					i -= 10000000000
					log.Println(subctx.Value("subkey"))
					log.Println("sub-subgoroutine is done")
					break subloop
				case <-newctx.Done():
					log.Println("Newctx is triggered")
					break subloop

				default:

				}
			}
		}(subctx)

		select {
		case <-ctx.Done():
			log.Println("sub goroutine is done")
		default:
			log.Println("there is no done signal from context")
		}
		// Done 실행 이후 wg == {{} 4294967296 0}
	}(ctx)

	//time.Sleep(1 * time.Second)
	//cancel2()                    -> newctx.Done() is triggered
	time.Sleep(3 * time.Second)
	cancel() // -> ctx.Done(), substx.Done() are triggered
	fmt.Println("reach here after context is cancelled")

	// wg로 묶어서 고루틴을 기다리는 형태로 만들지 않으면
	// cancel 이후 ctx.Done()에 해당하는 조건문 실행하다가 메인 함수 종료되서 같이 종료될 수 있음 (실행 보장 x)
	wg.Wait()
	log.Println(i)

}
