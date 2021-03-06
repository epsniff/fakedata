package fakedata

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/lucapette/fakedata/pkg/data"
)

// A Generator is a func that generates random data along with its description
type Generator struct {
	Func       func() string
	CustomFunc func(string) (func() string, error)
	Desc       string
	Name       string
}

// Generators is an array of Generator
type Generators []Generator

// IsCustom returns a bool indicating whether the generator has a CustomFunc or
// not
func (g Generator) IsCustom() bool {
	return g.CustomFunc != nil
}

// NewGenerators returns the available generators
func NewGenerators() (gens Generators) {
	f := newFactory()

	for _, gen := range f.generators {
		gens = append(gens, gen)
	}

	sort.Slice(gens, func(i, j int) bool { return strings.Compare(gens[i].Name, gens[j].Name) < 0 })
	return gens
}

// WithConstraints returns only the generators that accept constraints
func (gens Generators) WithConstraints() (newGens Generators) {
	for _, gen := range gens {
		if gen.IsCustom() {
			newGens = append(newGens, gen)
		}
	}
	return newGens
}

// FindByName returns, if present, the generator with the name string
func (gens Generators) FindByName(name string) (gen *Generator) {
	for _, g := range gens {
		if g.Name == name {
			return &g
		}
	}
	return gen
}

func withList(list []string) func() string {
	return func() string {
		return list[rand.Intn(len(list))]
	}
}

func ipv4() string {
	return fmt.Sprintf("%d.%d.%d.%d", 1+rand.Intn(253), rand.Intn(255), rand.Intn(255), 1+rand.Intn(253))
}

func ipv6() string {
	return fmt.Sprintf("2001:cafe:%x:%x:%x:%x:%x:%x", rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(255))
}

func mac() string {
	return fmt.Sprintf("%x:%x:%x:%x:%x:%x", rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(255))
}

func latitude() string {
	return strconv.FormatFloat((rand.Float64()*180)-90, 'f', 6, 64)
}

func longitude() string {
	return strconv.FormatFloat((rand.Float64()*360)-180, 'f', 6, 64)
}

func double() string {
	return strconv.FormatFloat(rand.NormFloat64()*1000, 'f', 4, 64)
}

func domain() string {
	return withList([]string{"test", "example"})() + "." + withList(data.TLDs)()
}

func date(options string) (f func() string, err error) {
	var min, max string

	endDate := time.Now()
	startDate := endDate.AddDate(-1, 0, 0)

	dateRange := strings.Split(options, ",")
	min = dateRange[0]

	if len(dateRange) > 1 {
		max = dateRange[1]
	}

	if len(min) > 0 {
		if len(max) > 0 {
			formattedMax := fmt.Sprintf("%sT00:00:00.000Z", max)

			date, err := time.Parse("2006-01-02T15:04:05.000Z", formattedMax)
			if err != nil {
				return nil, fmt.Errorf("problem parsing max date: %v", err)
			}

			endDate = date
		}

		formattedMin := fmt.Sprintf("%sT00:00:00.000Z", min)

		date, err := time.Parse("2006-01-02T15:04:05.000Z", formattedMin)
		if err != nil {
			return nil, fmt.Errorf("problem parsing mix date: %v", err)
		}

		startDate = date
	}

	if startDate.After(endDate) {
		return nil, fmt.Errorf("%v is after %v", startDate, endDate)
	}

	return func() string {
		return startDate.Add(time.Duration(rand.Intn(int(endDate.Sub(startDate))))).Format("2006-01-02")
	}, err
}

func integer(options string) (func() string, error) {
	min := 0
	max := 1000
	var low, high string
	intRange := strings.Split(options, ",")
	low = intRange[0]

	if len(intRange) > 1 {
		high = intRange[1]
	}

	if len(low) > 0 {
		m, err := strconv.Atoi(low)
		if err != nil {
			return nil, fmt.Errorf("could not convert min: %v", err)
		}

		min = m

		if len(high) > 0 {
			m, err := strconv.Atoi(high)
			if err != nil {
				return nil, fmt.Errorf("could not convert max: %v", err)
			}

			max = m
		}
	}

	if min > max {
		return nil, fmt.Errorf("max(%d) is smaller than min(%d)", max, min)
	}

	return func() string { return strconv.Itoa(min + rand.Intn(max+1-min)) }, nil
}

func file(path string) (func() string, error) {
	if path == "" {
		return nil, fmt.Errorf("no file path given")
	}

	filePath := strings.Trim(path, "'\"")

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("could not read file %s: %v", filePath, err)
	}
	list := strings.Split(strings.Trim(string(content), "\n"), "\n")

	return func() string { return withList(list)() }, nil
}

func enum(options string) (func() string, error) {
	list := []string{"foo", "bar", "baz"}
	if options != "" {
		list = strings.Split(options, ",")
	}
	return func() string { return withList(list)() }, nil
}

type generatorsMap map[string]Generator

func (gM generatorsMap) addGen(g Generator) {
	gM[g.Name] = g
}

type factory struct {
	generators generatorsMap
}

func (f factory) extractFunc(key, options string) (fn func() string, err error) {
	gen, ok := f.generators[key]
	if !ok {
		return nil, fmt.Errorf("unknown generator: %s", key)
	}

	if gen.IsCustom() {
		return gen.CustomFunc(options)
	}

	return gen.Func, nil
}

func newFactory() (f factory) {
	generators := make(generatorsMap)

	generators.addGen(Generator{
		Name: "domain.tld",
		Desc: "valid TLD name from https://data.iana.org/TLD/tlds-alpha-by-domain.txt",
		Func: withList(data.TLDs),
	})

	generators.addGen(Generator{Name: "country", Desc: "Full country name", Func: withList(data.Countries)})

	generators.addGen(Generator{Name: "country.code", Desc: "2-digit country code", Func: withList(data.CountryCodes)})

	generators.addGen(Generator{Name: "state", Desc: "Full US state name", Func: withList(data.States)})

	generators.addGen(Generator{Name: "state.code", Desc: "2-digit US state name", Func: withList(data.StateCodes)})

	generators.addGen(Generator{Name: "timezone", Desc: "tz in the form Area/City", Func: withList(data.Timezones)})

	generators.addGen(Generator{Name: "username", Desc: `username using the pattern \w+`, Func: withList(data.Usernames)})

	generators.addGen(Generator{Name: "name.first", Desc: "capitalized first name", Func: withList(data.Firstnames)})

	generators.addGen(Generator{Name: "name.last", Desc: "capitalized last name", Func: withList(data.Lastnames)})

	generators.addGen(Generator{Name: "color", Desc: "one word color", Func: withList(data.Colors)})

	generators.addGen(Generator{
		Name: "event.action",
		Desc: `clicked|purchased|viewed|watched`,
		Func: withList([]string{"clicked", "purchased", "viewed", "watched"}),
	})

	generators.addGen(Generator{
		Name: "http.method",
		Desc: `DELETE|GET|HEAD|OPTION|PATCH|POST|PUT`,
		Func: withList([]string{"DELETE", "GET", "HEAD", "OPTION", "PATCH", "POST", "PUT"}),
	})

	generators.addGen(Generator{
		Name: "name",
		Desc: `name.first + " " + name.last`,
		Func: func() string {
			return withList(data.Firstnames)() + " " + withList(data.Lastnames)()
		},
	})

	generators.addGen(Generator{
		Name: "email",
		Desc: "email",
		Func: func() string {
			return withList(data.Usernames)() + "@" + domain()
		},
	})

	generators.addGen(Generator{Name: "domain", Desc: "domain", Func: domain})

	generators.addGen(Generator{Name: "ipv4", Desc: "ipv4", Func: ipv4})

	generators.addGen(Generator{Name: "ipv6", Desc: "ipv6", Func: ipv6})

	generators.addGen(Generator{Name: "mac.address", Desc: "mac address", Func: mac})

	generators.addGen(Generator{Name: "latitude", Desc: "latitude", Func: latitude})

	generators.addGen(Generator{Name: "longitude", Desc: "longitude", Func: longitude})

	generators.addGen(Generator{Name: "double", Desc: "double number", Func: double})

	generators.addGen(Generator{
		Name: "noun",
		Desc: "noun from https://github.com/dariusk/corpora/blob/master/data/words/nouns.json",
		Func: withList(data.Nouns),
	})

	generators.addGen(Generator{
		Name: "emoji",
		Desc: "emoji from https://github.com/dariusk/corpora/blob/master/data/words/emojis.json",
		Func: withList(data.Emoji),
	})

	generators.addGen(Generator{Name: "animal", Desc: "animal breed", Func: withList(data.Animals)})

	generators.addGen(Generator{Name: "animal.cat", Desc: "random cat breed", Func: withList(data.Cats)})

	generators.addGen(Generator{Name: "animal.dog", Desc: "dog breed", Func: withList(data.Dogs)})

	generators.addGen(Generator{Name: "adjectives", Desc: "adjective", Func: withList(data.Adjectives)})

	generators.addGen(Generator{Name: "industry", Desc: "industry", Func: withList(data.Industries)})

	generators.addGen(Generator{Name: "occupation", Desc: "occupation", Func: withList(data.Occupations)})

	generators.addGen(Generator{Name: "sentence", Desc: "sentence", Func: withList(data.Sentences)})

	// custom generators
	generators.addGen(Generator{
		Name:       "date",
		Desc:       `random date in the format YYYY-MM-DD. By default, it generates dates in the last year`,
		CustomFunc: date,
	})

	generators.addGen(Generator{
		Name:       "int",
		Desc:       "positive integer between 1 and 1000",
		CustomFunc: integer,
	})

	generators.addGen(Generator{
		Name:       "enum",
		Desc:       `value from an enum. By default, the enum is foo,bar,baz. It accepts a list of comma-separated values`,
		CustomFunc: enum,
	})

	generators.addGen(Generator{
		Name:       "file",
		Desc:       `random value from a file. It accepts a file path. It can be either relative or absolute. The file must contain a value per line`,
		CustomFunc: file,
	})

	return factory{generators: generators}
}
