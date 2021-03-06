package httpwait

import (
	"flag"
	"fmt"
	"time"

	"github.com/tampopos/httpwait/src/client"
	"github.com/tampopos/httpwait/src/stopwatch"
)

// UseCase は構造体です
type useCase struct {
	stopwatch stopwatch.Stopwatch
	client    client.Client
}

// UseCase はインターフェイスです
type UseCase interface {
	Wait(args *WaitArgs) error
	CreateArgs() (*WaitArgs, error)
}

// CreateUseCase はインスタンスを生成します
func CreateUseCase(stopwatch stopwatch.Stopwatch, client client.Client) UseCase {
	return &useCase{stopwatch: stopwatch, client: client}
}

// WaitArgs はWaitの引数です
type WaitArgs struct {
	client.Request
	Interval   int
	Result     string
	StatusCode int
}

// CreateArgs はコマンドライン引数を取得します
func (useCase *useCase) CreateArgs() (*WaitArgs, error) {
	var args WaitArgs
	flag.StringVar(&args.URL, "url", "", "HTTP Request URL")
	flag.StringVar(&args.Method, "method", "GET", "HTTP Request Method")
	flag.Float64Var(&args.Timeout, "timeout", 60, "Timeout Seconds")
	flag.StringVar(&args.Result, "result", "", "HTTP request Result")
	flag.IntVar(&args.StatusCode, "statusCode", -1, "HTTP request StatusCode")
	flag.IntVar(&args.Interval, "interval", 5, "Request Interval Seconds")
	flag.Parse()

	if args.URL == "" {
		return nil, fmt.Errorf("url is required")
	}
	if args.Result == "" && args.StatusCode == -1 {
		return nil, fmt.Errorf("Either result or statusCode is required")
	}
	return &args, nil
}

// Wait は Timeoutになるまでリクエストします
func (useCase *useCase) Wait(args *WaitArgs) error {
	useCase.stopwatch.Start()
	ch := make(chan bool)
	errCh := make(chan string)

	go useCase.polling(args, ch, errCh)

	return useCase.receive(args, ch, errCh)
}
func (useCase *useCase) receive(args *WaitArgs, ch chan bool, errCh chan string) error {
	timeout := time.Duration(args.Timeout) * time.Second
	select {
	case _ = <-ch:
		fmt.Printf("Success!\n")
		return nil
	case msg := <-errCh:
		return fmt.Errorf(msg)
	case <-time.After(timeout):
		return fmt.Errorf("Timeout")
	}
}
func (useCase *useCase) polling(args *WaitArgs, ch chan bool, errCh chan string) {
	for {
		var elapsed = useCase.stopwatch.GetElapsedSeconds()
		fmt.Printf("elapsed %v sec.\n", elapsed)
		fmt.Printf("check start.\n")
		go useCase.check(args, ch, errCh)
		interval := time.Duration(args.Interval) * time.Second
		fmt.Printf("Wait for %v sec.\n", interval)
		time.Sleep(interval)
	}
}
func (useCase *useCase) check(args *WaitArgs, ch chan bool, errCh chan string) {
	if args.StatusCode != -1 {
		statusCode, err := useCase.client.GetStatusCode(&args.Request)
		if err != nil {
			errCh <- err.Error()
			return
		}

		if statusCode == args.StatusCode {
			ch <- true
			return
		}
		fmt.Printf("Failed: status code is not %v\n", args.StatusCode)
	} else {
		body, err := useCase.client.GetBody(&args.Request)
		if err != nil {
			errCh <- err.Error()
			return
		}
		if body == args.Result {
			ch <- true
			return
		}
		fmt.Printf("Failed: result is not %s\n", args.Result)
	}
}
