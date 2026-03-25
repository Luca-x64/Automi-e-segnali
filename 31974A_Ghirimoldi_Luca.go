package main

import (
	"bufio"
	"fmt"
	"os"
	"solution/utils"
	"strconv"
	"strings"
)

func main() {
	var piano piano = newPiano()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		data := scanner.Text()
		if data == "c" {
			piano = newPiano()
		} else {
			esegui(piano, data)
		}
	}
}

func newPiano() piano {
	return *utils.Crea()
}

/*
funzione che applica al piano rappresentato da p l’operazione associata dalla stringa s, secondo quanto
specificato nella Tabella 1.
*/
func esegui(p piano, s string) {
	line := strings.Fields(s) //separo gli elementi della linea in base alla presenza di uno o più spazi
	if len(line) > 0 {
		char := []rune(line[0])[0]

		switch char {
		case 's':
			x, _ := strconv.Atoi(line[1])
			y, _ := strconv.Atoi(line[2])

			fmt.Println(p.Stato(x, y))
		case 'S':

			p.Stampa()
		case 'a':
			a, _ := strconv.Atoi(line[1])
			b, _ := strconv.Atoi(line[2])

			p.Automa(a, b, line[3])
		case 'o':
			a, _ := strconv.Atoi(line[1])
			b, _ := strconv.Atoi(line[2])
			c, _ := strconv.Atoi(line[3])
			d, _ := strconv.Atoi(line[4])

			p.Ostacolo(a, b, c, d)

		case 'r':
			a, _ := strconv.Atoi(line[1])
			b, _ := strconv.Atoi(line[2])

			p.Richiamo(a, b, line[3])
		case 'p':
			fmt.Println(p.Posizioni(line[1]))

		case 'e':
			a, _ := strconv.Atoi(line[1])
			b, _ := strconv.Atoi(line[2])

			esiste := p.EsistePercorso(a, b, line[3])
			if esiste {
				fmt.Println("SI")
			} else {
				fmt.Println("NO")
			}
		case 'f':
			os.Exit(0)
		case 't': // per febbraio
			fmt.Println(0)
		default:

		}
	}

}

type piano = utils.Piano
