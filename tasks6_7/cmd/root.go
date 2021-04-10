package cmd

import (
	"fmt"
	"log"
	"regexp"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var (
	// Used for flags.
	rootCmd = &cobra.Command{
		Use:   "task5",
		Short: "task5 application",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			validateUsername(args[1])
			validateYamlFile(args[0])
			C := getInput(args[0], args[1])
			output, err := yaml.Marshal(C)
			if err != nil {
				log.Printf("yamlFile.Get err   #%v ", err)
			}
			//C.This.Is = map[string][]string{args[1]: C.This.Is["username"]}
			//fmt.Println(C)
			fmt.Printf("%+v\n", string(output))
		},
	}
)

func validateUsername(username string) {
	r, _ := regexp.Compile("^[a-zA-z][a-zA-z_0-9]{0,10}")
	if !r.MatchString(username) {
		log.Fatalf("Provide a valid username")
	}
	if len(username) > 10 || len(username) < 4 {
		log.Fatalf("Provide at least 4 characters, or your username is too long")
	}
}

func validateYamlFile(filename string) {
	r, _ := regexp.Compile("^.*(yaml)$")
	if !r.MatchString(filename) {
		log.Fatalf("Provide a yaml file")
	}
}

func Execute() error {
	return rootCmd.Execute()
}
