Learn-Go
=========================
[![Build Status](https://img.shields.io/badge/build-passing-brightgreen.svg)](https://github.com/MrDongdongLin/Learn-Go#learn-go)
[![Packagist](https://img.shields.io/badge/paskagist-v1.1.0-blue.svg)](https://github.com/MrDongdongLin/Learn-Go/releases)
[![Powerby](https://img.shields.io/badge/powerby-DongdongLin-orange.svg)](https://github.com/MrDongdongLin)

This repository records some skills about Go programming.

## [yaml io](https://github.com/MrDongdongLin/Learn-Go/tree/master/yaml%20io)
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
## CHANGELOG
[CHANGELOG](https://github.com/MrDongdongLin/Learn-Go/releases)