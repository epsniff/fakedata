package main

import (
	"bytes"
	"fmt"
	"github.com/lucapette/fakedata/pkg/fakedata"
	flag "github.com/spf13/pflag"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

var version = "master"

var generatorsFlag = flag.BoolP("generators", "g", false, "lists available generators")
var limitFlag = flag.IntP("limit", "l", 10, "limits rows up to n")
var formatFlag = flag.StringP("format", "f", "", "generators rows in f format. Available formats: csv|tab|sql")
var versionFlag = flag.BoolP("version", "v", false, "shows version information")
var tableFlag = flag.StringP("table", "t", "TABLE", "table name of the sql format")
var templateFlag = flag.StringP("template", "", "", "Use template as input")

func getFormatter(format string) (f fakedata.Formatter) {
	switch format {
	case "csv":
		f = fakedata.NewSeparatorFormatter(",")
	case "tab":
		f = fakedata.NewSeparatorFormatter("\t")
	case "sql":
		f = fakedata.NewSQLFormatter(*tableFlag)
	default:
		f = fakedata.NewSeparatorFormatter(" ")
	}
	return f
}

func generatorsHelp() string {
	generators := fakedata.Generators()
	var buffer bytes.Buffer

	var max int

	for _, gen := range generators {
		if len(gen.Name) > max {
			max = len(gen.Name)
		}
	}

	pattern := fmt.Sprintf("%%-%ds%%s\n", max+2) //+2 makes the output more readable
	for _, gen := range generators {
		fmt.Fprintf(&buffer, pattern, gen.Name, gen.Desc)
	}

	return buffer.String()
}

func isPipe() bool {
	stat, _ := os.Stdin.Stat()
	// Check if template data is piped to fakedata
	return (stat.Mode() & os.ModeCharDevice) == 0
}

func main() {
	if *versionFlag {
		fmt.Println(version)
		os.Exit(0)
	}

	if *generatorsFlag {
		fmt.Print(generatorsHelp())
		os.Exit(0)
	}

	rand.Seed(time.Now().UnixNano())

	if *templateFlag != "" {
		if err := fakedata.ParseTemplate(*templateFlag, *limitFlag); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		os.Exit(0)
	}

	if isPipe() {
		t, err := ioutil.ReadAll(os.Stdin)

		if err != nil {
			fmt.Printf("Unable to read input: %s", err)
			os.Exit(1)
		}

		if err = fakedata.ParseTemplateFromPipe(string(t), *limitFlag); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	if len(flag.Args()) == 0 {
		flag.Usage()
		os.Exit(0)
	}

	columns, err := fakedata.NewColumns(flag.Args())
	if err != nil {
		fmt.Printf("%v\n\nSee fakedata --generators for a list of available generators.\n", err)
		os.Exit(0)
	}

	formatter := getFormatter(*formatFlag)

	for i := 0; i < *limitFlag; i++ {
		fmt.Print(fakedata.GenerateRow(columns, formatter))
	}
}

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: fakedata [option ...] field...\n\n")
		flag.PrintDefaults()
	}
	flag.Parse()
}
