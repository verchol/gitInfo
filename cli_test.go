package main

import (
	"fmt"
	"io/ioutil"
	"testing"

	yaml "gopkg.in/yaml.v2"
)

type OnlyTriggers struct {
	Triggers []struct {
		Name string `yaml:"name"`
		Spec struct {
			TriggerType string   `yaml:"triggerType"`
			Provider    string   `yaml:"provider"`
			Repo        string   `yaml:"repo"`
			Events      []string `yaml:"events"`
			BranchRegex string   `yaml:"branchRegex"`
		} `yaml:"spec"`
	} `yaml:"Triggers"`
}

func TestYaml(t *testing.T) {
	values := "./triggers.yaml"
	valuesData, err := ioutil.ReadFile(values)
	if err != nil {
		fmt.Print(err)
	}
	v := OnlyTriggers{}
	err = yaml.Unmarshal([]byte(valuesData), &v)
	if err != nil {
		panic(err)
	}
	fmt.Printf("valuesFile %v\n", values)
	fmt.Printf("values: %s\n", (string)(valuesData))
	fmt.Printf("valuesObj: %v\n", v)

	fmt.Printf("trigger name is  %s\n", v.Triggers[0].Name)

}
