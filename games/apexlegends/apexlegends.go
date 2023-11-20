package apexlegends

import (
	"encoding/json"
	"math/rand"
	"os"
)

const (
	BaseURI = "https://api.mozambiquehe.re"
)

var (
	ApexLegendsConfig = &Apex{}
)

type Apex struct {
	Legends []*Legend
}

func (apex *Apex) IsEmpty() bool {
	return len(apex.Legends) <= 0
}

func (apex *Apex) init() error {
	b, err := os.ReadFile("./games/apexlegends/apexlegends.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, &apex.Legends)
	if err != nil {
		panic(err)
	}
	return nil
}

func (apex *Apex) RandomLegend() *Legend {
	if apex.IsEmpty() {
		apex.init()
	}
	num := rand.Intn(len(apex.Legends))
	return apex.Legends[num]
}

func (apex *Apex) RandomLegendByClass(Class string) *Legend {
	if apex.IsEmpty() {
		apex.init()
	}
	var legends []*Legend
	for _, legend := range apex.Legends {
		if legend.Is(Class) {
			legends = append(legends, legend)
		}
	}
	num := rand.Intn(len(legends))
	return legends[num]
}

type Legend struct {
	Name  string `json:"Name"`
	Class string `json:"Class"`
}

func (legend *Legend) Is(Class string) bool {
	return legend.Class == Class
}
