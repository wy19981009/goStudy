package ctxtest

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func handelrequest(ctx context.Context) {
	go writeredis(ctx)
	go writedatabase(ctx)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("handelrequest done.")
			return
		default:
			fmt.Println("handelrequest running")
			time.Sleep(2 * time.Second)
		}
	}
}

func writeredis(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("writeredis done.")
			return
		default:
			fmt.Println("writeredis running")
			time.Sleep(2 * time.Second)
		}
	}
}

func writedatabase(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("writedatabase done.")
			return
		default:
			fmt.Println("writedatabase running")
			time.Sleep(2 * time.Second)
		}
	}
}
func TestCtx(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go handelrequest(ctx)

	time.Sleep(5 * time.Second)
	fmt.Println("it's time to stop all sub goroutines!")
	cancel()

	//just for test whether sub goroutines exit or not
	time.Sleep(5 * time.Second)
}

const shortDuration = 1 * time.Millisecond

func TestCtx2(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), shortDuration)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("在超时时间之后结束")
	case <-ctx.Done():
		fmt.Println("在超时时间结束")
	}
}
