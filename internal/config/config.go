package config

import (
	"flag"
	"os"
)

const (
	input  = "input/example.json"
	schema = "input/schema.json"
	output = "output"
)

type Config struct {
	Input   string
	Output  string
	Schema  string
	Include []string
}

type stringsFlags []string

func (s *stringsFlags) String() string {
	return "my string representation"
}

func (s *stringsFlags) Set(value string) error {
	*s = append(*s, value)
	return nil
}

var includeFlags stringsFlags

func New() (Config, error) {
	var i, o, s string

	flag.StringVar(&i, "input", "", "Path to intermediate representation")
	flag.StringVar(&o, "output", "", "Directory for generated code files")
	flag.StringVar(&s, "schema", "", "Path or URL to intermediate representation JSON schema")
	flag.Var(&includeFlags, "include", "Specify which data sources, provider and resources to include on the basis of name")
	flag.CommandLine.Parse(os.Args[2:])

	config := Config{
		Input:  input,
		Schema: schema,
		Output: output,
	}

	if i != "" {
		config.Input = i
	}

	if s != "" {
		config.Schema = s
	}

	if o != "" {
		config.Output = o
	}

	return config, nil
}
