package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Input struct {
	This struct {
		Is []string `yaml:"is"`
	} `yaml:"this"`
}

type Output struct {
	This struct {
		Is struct {
			My []string `yaml:"my"`
		} `yaml:"is"`
	} `yaml:"this"`
}

func getInput() Input {
	var c Input
	yamlFile, err := ioutil.ReadFile("task4.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}

func getOutput(c Input) Output {
	var o Output = Output{}
	o.This.Is.My = c.This.Is
	o.This.Is.My = []string{c.This.Is[1], c.This.Is[2]}
	return o
}

func main() {
	c := getInput()
	o := getOutput(c)
	output, err := yaml.Marshal(o)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	fmt.Printf("%+v\n", string(output))
}
