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
	for {
		if args.StatusCode != -1 {
			statusCode, err := useCase.client.GetStatusCode(&args.Request)
			if err != nil {
				return err
			}

			fmt.Printf("StatusCode: %d\n", statusCode)
			if statusCode == args.StatusCode {
				return nil
			}
			fmt.Printf("Failed: status code is not %v\n", args.StatusCode)
		} else {
			body, err := useCase.client.GetBody(&args.Request)
			if err != nil {
				return err
			}
			fmt.Printf("Result: %v\n", body)
			if body == args.Result {
				return nil
			}
			fmt.Printf("Failed: result is not %s\n", args.Result)
		}

		var elapsed = useCase.stopwatch.GetElapsedSeconds()
		fmt.Printf("elapsed %v sec.\n", elapsed)
		if args.Timeout <= elapsed {
			return fmt.Errorf("Timeout")
		}
		fmt.Printf("Wait for 5sec.\n")
		time.Sleep(5 * time.Second)
	}
}
