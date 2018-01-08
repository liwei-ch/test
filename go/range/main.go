package main

import (
	"fmt"
)

type printer interface {
	printf()
}

type Myprinter struct {
}

func (p *Myprinter) printf() {
	fmt.Printf("Myprinter printf\n")
}

type Myprinter2 struct {
}

func (p *Myprinter2) printf() {
	fmt.Printf("Myprinter2 printf\n")
}

type Myprinter3 struct {
}

func (p *Myprinter3) printf() {
	fmt.Printf("Myprinter3 printf\n")
}

func main() {

	maps := map[int]printer{}

	maps[1] = &Myprinter{}
	maps[2] = &Myprinter2{}
	maps[3] = &Myprinter3{}

	for i, p := range maps {
		_, _ = i, p
		/*go func(p printer) {
			p.printf()
		}(p)*/
		go func() { p.printf() }()
		go p.printf()
		//go maps[i].printf()
	}
	for {
		select {}
	}
}
