# httpwait

指定した URL にリクエストして結果が帰ってくるまでタスクをスリープさせるためのツール

## Required

go(v1.11 以降)

## Command

#### Run

```sh
go run src/main.go
```

#### Build

```sh
go build -o dist/httpwait
```

#### Test

```sh
go test ./src/httpwait/tests/ -test.v
```

#### Release

```sh
git tag 1.0.0
git push origin --tags
```

## Usage

- **method** `string`  
  HTTP Request Method (default "GET")
- **result** `string`  
  HTTP request Result
- **statusCode** `int`  
  HTTP request StatusCode (default -1)
- **timeout** `float`  
  Timeout Seconds (default 60)
- **url** `string`  
  HTTP Request URL
- **interval** `int`  
  Request Interval Seconds
