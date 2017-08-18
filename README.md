# Learn-Go
This repository records some skills about Go programming.

## yaml io
This project describes how to read a yaml file and print it to screem.

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

where `c` is a struct object that defined by us, according to the format of yaml data. Suppose the yaml data is
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