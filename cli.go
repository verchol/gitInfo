package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"text/template"

	yaml "gopkg.in/yaml.v2"
)

type ProjectInfo struct {
	Project struct {
		Name string `yaml:"name"`
	} `yaml:"Project"`
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
	Steps struct {
		Location string `yaml:"location"`
	} `yaml:"Steps"`
}

func processTemplateFile() {

	type Values struct {
		ReplicaCount int `yaml:"replicaCount"`
		Image        struct {
			Repository string `yaml:"repository"`
			Tag        string `yaml:"tag"`
			PullPolicy string `yaml:"pullPolicy"`
		} `yaml:"image"`
		Service struct {
			Name         string `yaml:"name"`
			Type         string `yaml:"type"`
			ExternalPort int    `yaml:"externalPort"`
			InternalPort int    `yaml:"internalPort"`
		} `yaml:"service"`
		Ingress struct {
			Enabled     bool        `yaml:"enabled"`
			Hosts       []string    `yaml:"hosts"`
			Annotations interface{} `yaml:"annotations"`
			TLS         interface{} `yaml:"tls"`
		} `yaml:"ingress"`
		Resources struct {
		} `yaml:"resources"`
	}

	info := new(Values)
	file := path.Join("./", "/test/demochart/templates/deployment.yaml")
	tmpl := path.Join("./", "/test/demochart/templates/_helpers.tpl")
	f := func(string, interface{}) string { return "not implemented" }
	t := template.New("mytemplate")
	funcMap := template.FuncMap{
		"toToml":     f,
		"toYaml":     f,
		"fromYaml":   f,
		"toJson":     f,
		"fromJson":   f,
		"indent":     f,
		"default":    f,
		"trunc":      f,
		"trimSuffix": f,

		// This is a placeholder for the "include" function, which is
		// late-bound to a template. By declaring it here, we preserve the
		// integrity of the linter.
		"include":  func(string, interface{}) string { return "not implemented" },
		"required": func(string, interface{}) interface{} { return "not implemented" },
		"tpl":      func(string, interface{}) interface{} { return "not implemented" },
	}
	t.Funcs(funcMap)
	templates, err := t.ParseFiles(file, tmpl)

	if err != nil {
		panic(err)
	}
	templates.ExecuteTemplate(os.Stdout, "t1", info)
}

func setupFlags() {
	flag.String("valuesFile", "values.yaml", "file with input parameters")
}
func setupProject() (values *string, name *string) {

	flagSet := flag.NewFlagSet("project", flag.ExitOnError)
	flagSet.String("valuesFile", "values.yaml", "file with input parameters")
	flagSet.String("name", "", "name of the project")

	return
}

type Values map[string]interface{}

func main() {

	flagSet := flag.NewFlagSet("project", flag.ContinueOnError)
	values := flagSet.String("values", "values.yaml", "file with input parameters")
	projectName := flagSet.String("name", "noname", "name of the project")
	pipsFolder := flagSet.String("pipsFolder", "./", "folder to load pipfiles")
	flagSet.Parse(os.Args[2:])
	pips, err := ioutil.ReadDir(*pipsFolder)
	valuesData, err := ioutil.ReadFile(*values)
	if err != nil {
		fmt.Print(err)
	}
	v := ProjectInfo{}
	err = yaml.Unmarshal([]byte(valuesData), &v)
	if err != nil {
		panic(err)
	}
	fmt.Printf("valuesFile %v\n", *values)
	fmt.Printf("values: %s\n", (string)(valuesData))
	fmt.Printf("valuesObj, triggers spec %+v\n", v)
	fmt.Printf("%s, %s\n", *values, *projectName)
	t := template.New("cfproject")
	_ = t
	var files []string
	for _, file := range pips {
		files = append(files, path.Join(*pipsFolder, file.Name()))
	}
	fmt.Printf("files to process %v\n", files)
	templates, err := t.ParseFiles(files...)
	fmt.Printf("%v", templates)
	if err != nil {
		panic(err)
	}

	err = templates.ExecuteTemplate(os.Stdout, "deploypip.yaml", v)
	if err != nil {
		panic(err)
	}
}
