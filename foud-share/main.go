package main

import (
	"flag"
	//	"fmt"
	"io/ioutil"
	"log"
	"sort"

	"gopkg.in/yaml.v2"
)

var (
	config = flag.String("c", "foud.yaml", "foud config pathname")
)

type Foud struct {
	Name   string
	Amount float32
}

type FoudArgs struct {
	Name       string  `yaml:"name"`
	TodayQuote float32 `yaml:"todayQuote"`
	WeekQuote  float32 `yaml:"weekQuote"`
	MonthQuote float32 `yaml:"monthQuote"`
	SumQuote   float32
}

type SortBy func(p, q *FoudArgs) bool

type FoudPoolArgs struct {
	Name        string    `yaml:"name"`
	SalaryBase  float32   `yaml:"salaryBase"`
	FoudPercent float32   `yaml:"foudPercent"`
	Scores      []float32 `yaml:"scores"`
	QuoteType   int       `yaml:"quoteType"` // 0使用weekQuote 1使用monthQuote
	by          func(p, q *FoudArgs) bool
	Fouds       []*FoudArgs `yaml:"fouds"`
}

func (sw FoudPoolArgs) Len() int {
	return len(sw.Fouds)
}

func (sw FoudPoolArgs) Swap(i, j int) {
	sw.Fouds[i], sw.Fouds[j] = sw.Fouds[j], sw.Fouds[i]
}

func (sw FoudPoolArgs) Less(i, j int) bool {
	return sw.by(sw.Fouds[i], sw.Fouds[j])
}

func SortSumQuote(fouds []*FoudArgs, by SortBy) {
	sort.Sort(FoudPoolArgs{Fouds: fouds, by: by})
}

type FoudConfig struct {
	Todo  string          `yaml:"todo"`
	Pools []*FoudPoolArgs `yaml:"pools"`
}

var conf FoudConfig

func main() {
	flag.Parse()
	file, err := ioutil.ReadFile(*config)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(file, &conf)
	if err != nil {
		panic(err)
	}
	//log.Printf("conf: %+v", conf)

	log.Printf("foud todo: %s", conf.Todo)
	for _, fa := range conf.Pools {
		FoudPool(fa)
	}
}

func FoudPool(pool *FoudPoolArgs) {
	if len(pool.Scores) > len(pool.Fouds) {
		log.Fatalf("number Score > Fouds")
		return
	}

	var quoteSum float32 = 0.0
	for _, q := range pool.Scores {
		quoteSum += q
	}
	amount := pool.SalaryBase * pool.FoudPercent
	eachAmount := amount / quoteSum
	log.Printf("基金池:%s 总金额:%f 每份金额:%f 总份数:%f scores:%v", pool.Name, amount, eachAmount, quoteSum, pool.Scores)

	for _, foud := range pool.Fouds {
		if pool.QuoteType == 0 {
			foud.SumQuote = foud.TodayQuote + foud.WeekQuote
		} else {
			foud.SumQuote = foud.TodayQuote + foud.MonthQuote
		}
		//log.Printf("foud: %v", foud)
	}
	tempFouds := pool.Fouds
	SortSumQuote(tempFouds, func(p, q *FoudArgs) bool {
		return q.SumQuote > p.SumQuote // SumQuote递减
	})
	fouds := make([]Foud, 0, len(pool.Scores))
	for i := 0; i < len(pool.Scores); i++ {
		fouds = append(fouds, Foud{Name: tempFouds[i].Name, Amount: pool.Scores[i] * eachAmount})
	}
	log.Printf("%s基金池配额如下：", pool.Name)
	for _, foud := range fouds {
		log.Printf("%s:\t%f", foud.Name, foud.Amount)
	}
}
