package main

import (
	"io/ioutil"
	"log"

	rta "demo/conf"
	"gopkg.in/yaml.v2"
)

func main() {
	c := rta.Conf{}
	conf, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(conf, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	rta.GetYaml(&c)
	rta.Hello()
}
