package cli

import (
	"os"
	"github.com/urfave/cli"
	"github.com/go-yaml/yaml"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

func Read() {
	cli.VersionFlag = cli.BoolFlag{
		Name: "version", // we'll use -v for verbose output.
		Usage: "print the version",
	}
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("version=%s\n", c.App.Version)
	}

	app := cli.NewApp()
	app.Name = "yaml-reader"
	app.Usage = "Read YAML formatted files."
	app.Version = "0.0.1"

	/*
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "file, f",
			Usage: "Try to parse `FILE`",
		},
	}
	*/

	fileFlag := cli.StringFlag{
		Name: "file, f",
		Value: "file.yml",
		// TODO how to make this flag mandatory
	}

	app.Commands = []cli.Command{
		{
			Name:    "json",
			Usage:   "Convert to json",
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "pretty, p"},
				fileFlag,
			},
			Action:  func(c *cli.Context) error {
				fileName := c.String("f")
				fmt.Printf("Reading YML file '%s'...\n", fileName)
				// Read it into an obj, map, interface etc.
				bytea := readFile(fileName)
				t := arbitrary(bytea)
				fmt.Println("Converting into JSON...")
				// Marshal it to JSON & print
				jsonRepr, _ := json.Marshal(t)
				fmt.Println(string(jsonRepr))
				return nil
			},

		},
		{
			Name: "read",
			Usage: "Just read the file, return error if it is not properly formatted.",
			Flags: []cli.Flag{fileFlag},
			Action: func(c *cli.Context) error {
				// Read it into an obj, map, interface etc.
				fileName := c.String("f")
				bytea := readFile(fileName)
				mapVal := arbitrary(bytea)
				print_v(mapVal)
				return nil
			},
		},

	}

	app.Run(os.Args)
}

type T struct {
	A string
	B struct {
		  RenamedC int   `yaml:"c"`
		  D        []int `yaml:",flow"`
	  }
}

func readFile(fileName string) []byte {
	// http://stackoverflow.com/questions/1821811/how-to-read-write-from-to-file
	fmt.Printf("Trying to read file '%s' into a byte array...\n", fileName)
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return b
}

func arbitrary(bytea []byte) map[string]interface{} {
	// https://michaelheap.com/golang-encodedecode-arbitrary-json/
	var hede map[string]interface{}
	fmt.Println("Trying to convery yaml byte array into object...")
	err := yaml.Unmarshal(bytea, &hede)
	if err != nil {
		panic(err)
	}
	return hede
}

func print_v(mapVal map[string]interface{}) {
	fmt.Print(mapVal)
}

func readYaml(bytea []byte) T {
	// https://github.com/go-yaml/yaml
	t := T{}
	fmt.Println("Trying to convery yaml byte array into object...")
	err := yaml.Unmarshal(bytea, &t)
	if err != nil {
		panic(err)
	}
	return t
}
