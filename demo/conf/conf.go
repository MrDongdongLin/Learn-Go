package test

import (
	"fmt"
)

// Conf store local conf
type Conf struct {
	Person struct {
		Name string `yaml:"name"`
		Age  string `yaml:"age"`
	}
}

func GetYaml(conf *Conf) {
	fmt.Printf("%s is %s years old.\n", conf.Person.Name, conf.Person.Age)
}
