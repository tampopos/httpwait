package httpwait

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	var args, err = getArgs()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
	err = wait(args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Success!\n")
}

// Args はコマンドライン引数です
type Args struct {
	Timeout    float64
	Method     string
	URL        string
	Result     string
	StatusCode int
}

func getArgs() (*Args, error) {
	var args Args
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

func wait(args *Args) error {
	var start = time.Now()
	for {
		res, err := request(args)
		if err != nil {
			return err
		}
		if res {
			return nil
		}
		var elapsed = getElapsedSeconds(start)
		fmt.Printf("elapsed %v sec.\n", elapsed)
		if args.Timeout <= elapsed {
			return fmt.Errorf("Timeout")
		}
		fmt.Printf("Wait for 5sec.\n")
		time.Sleep(5 * time.Second)
	}
}

func request(args *Args) (bool, error) {
	var req, err = http.NewRequest(args.Method, args.URL, nil)
	if err != nil {
		return false, fmt.Errorf("Failed: create request : %s", err)
	}
	fmt.Printf("Requesting [%s] %s\n", req.Method, req.URL)

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Failed: request: %s\n", err)
		return false, nil
	}
	defer resp.Body.Close()

	fmt.Printf("StatusCode: %d\n", resp.StatusCode)
	if args.StatusCode != -1 {
		if resp.StatusCode == args.StatusCode {
			return true, nil
		}
		fmt.Printf("Failed: status code is not %v\n", args.StatusCode)
		return false, nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Failed: read body: %s\n", err)
		return false, nil
	}

	strBody := string(body)
	fmt.Printf("Result: %v\n", strBody)
	if strBody != args.Result {
		fmt.Printf("Failed: result is not %s\n", args.Result)
		return false, nil
	}

	return true, nil
}

func getElapsedSeconds(start time.Time) float64 {
	return (time.Now().Sub(start)).Seconds()
}
