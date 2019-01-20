# httpwait

指定したサービスのが立ち上がるまでポーリングする

## Required

go

## Command

#### Init

```sh
go mod init
```

#### Run

```sh
go run src/main.go
```

#### Build

```sh
go build -o dist/httpwait
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
