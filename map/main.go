package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

var speakers = map[string]bool{
	"Eduardo Alves":   true,
	"Bastian Winkler": true,
	"Xabi":            true,
	"The Italian Trio. Filippo Valsorda, Roberto Clapis, Giovanni Bajo": true,
	"Tiago Mendes":          true,
	"Denys Nahurnyi":        true,
	"Syamala Umamaheswaran": true,
	"Takuya Ueda":           true,
	"Wojtek Siudzinski":     true,
	"Haiko Schol":           true,
	"Egon Elbre":            true,
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	// first speaker is fixed
	announce("Joonatan Samuel")
	s.Scan()
	// second speaker is fixed
	announce("Carlos Guzman")
	// the rest... random baby!
	for name := range speakers {
		s.Scan() // they pressed enter
		announce(name)
	}
}

func announce(name string) {
	fmt.Println("Next up:", name)
	exec.Command("say", "Next up, please welcome the wonderful "+name).Run()
}
