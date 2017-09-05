# Learn-Go
[![Build Status](https://img.shields.io/badge/build-passing-brightgreen.svg)](https://github.com/MrDongdongLin/Learn-Go#learn-go)
[![Packagist](https://img.shields.io/badge/packgist-v1.3.1-orange.svg)](https://github.com/MrDongdongLin/Learn-Go/releases)
[![Powerby](https://img.shields.io/badge/powerby-DongdongLin-blue.svg)](https://github.com/MrDongdongLin)

This repository records some skills about Go programming.

## [yaml](https://github.com/MrDongdongLin/Learn-Go/tree/master/yaml)
This project describes how to read a yaml file and print it to screen.
- Package need to import:
  - `io/ioutil`
  - `gopkg.in/yaml.v2`
- Read yaml file:
  ``` go
  conf, err := ioutil.ReadFile("conf.yaml")
  ```
- Convert yaml data to struct:
  ```go
  err = yaml.Unmarshal(conf, &c)
  ```
where `c` is a struct object that defined according to the format of yaml data by us. Suppose yaml data is
```
person:
  name: 'Flower'
  age:  '25'
```
then our struct must define as
```go
type Conf struct {
  Person struct {
    Name string `yaml:"name"`
    Age  string `yaml:"age"`
  }
}
```

## [httprouter](https://github.com/MrDongdongLin/Learn-Go/tree/master/httprouter)
HttpRouter is a lightweight high performance HTTP request router (also called multiplexer or just mux for short) for Go.

### Installation
```go
go get github.com/julienschmidt/httprouter
```
Create a new httprouter object with `New` function, and we can create request methods like `GET`, `POST` and so on. Let's see an example:
```go
router := httprouter.New()
router.GET("/get/:requestid", controller.GetTaskProcess)
router.POST("/post/:requestid", controller.PostTaskProcess)
...
```
then writing methods `GetTaskProcess` and `PostTaskProcess` with a simple response respectively:
```go
func GetTaskProcess(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  fmt.Fprintf(w, "hello, %s!\n", ps.ByName("requestid"))
}

func PostTaskProcess(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  body, err := ioutil.ReadAll(r.Body)
  if err != nil {
    fmt.Fprint(w, "hello, something wrong happened for sending message to %s!\n", ps.ByName("requestid"))
  } else {
    fmt.Fprint(w, "Your enter is %s\n", []byte(body))
  }
}
```
At last, we use `curl` to test this project:
```
curl -XGET localhost:8002/get/world
```
then we get
```
hello, world
```

## [gorm](https://github.com/MrDongdongLin/Learn-Go/tree/master/gorm)
The fantastic ORM library for Golang, aims to be developer friendly.

### Installation
```go
go get -u github.com/jinzhu/gorm
```

Ditail operations can be found [here](http://jinzhu.me/gorm/).

## [redis](https://github.com/MrDongdongLin/Learn-Go/tree/master/redis)
Here, taking the [gosexy redis](https://github.com/gosexy/redis) for example, to introduce how to use redis in golang.

### Installation
- redis
```Shell Session
$ wget http://download.redis.io/releases/redis-2.8.3.tar.gz
$ tar xzf redis-2.8.3.tar.gz
$ cd redis-2.8.3
$ make
$ redis-server
```
- redis client
```go
go get menteslibres.net/gosexy/redis
```
then import redis with
```go
import "menteslibres.net/gosexy/redis"
```
- usage
```go
var client *redis.Client

client = redis.New()

err = client.Connect(host, port)

if err != nil {
  log.Fatalf("Connect failed: %s\n", err.Error())
  return
}

log.Println("Connected to redis-server.")

log.Printf("Sending PING...\n")
s, err = client.Ping()

if err != nil {
  log.Fatalf("Could not ping: %s\n", err.Error())
  return
}

log.Printf("Received %s!\n", s)

client.Quit()
```

### A RESTful api using httprouter and redis

## CHANGELOG
[CHANGELOG](https://github.com/MrDongdongLin/Learn-Go/releases)

## Contact me
![](https://github.com/MrDongdongLin/src/blob/master/weixin/lindd.jpg)
