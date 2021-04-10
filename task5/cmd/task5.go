package cmd

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Input struct {
	This struct {
		Is map[string][]string `yaml:"is"`
	} `yaml:"this"`
}

func getInput(filename string, username string) Input {
	var c Input
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	//m := make(map[interface{}]interface{})
	// err = yaml.Unmarshal([]byte(yamlFile), &m)
	// if err != nil {
	// 	log.Fatalf("error: %v", err)
	// }
	//m["this"]["is"][username] = m["this"]["is"]["my"]
	//fmt.Printf("--- m:\n%v\n\n", m)
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	c.This.Is[username] = c.This.Is["my"]
	delete(c.This.Is, "my")
	return c
}

// func main() {
// 	// c := getInput()
// 	// o := getOutput(c)
// 	// output, err := yaml.Marshal(o)
// 	// if err != nil {
// 	// 	log.Printf("yamlFile.Get err   #%v ", err)
// 	// }
// 	// fmt.Printf("%+v\n", string(output))
// 	cmd.Execute()
// }
