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
	line := strings.Fields(s)
	// Parsing numerico preventivo (fino a 5 argomenti, default 0 se non presenti o non validi)
	nums := make([]int, 5)
	for i := 1; i < len(line) && i <= 5; i++ {
		nums[i-1], _ = strconv.Atoi(line[i])
	}
	if len(line) > 0 {
		char := []rune(line[0])[0]
		switch char {
		case 's':
			fmt.Println(p.Stato(nums[0], nums[1]))
		case 'S':
			p.Stampa()
		case 'a':
			p.Automa(nums[0], nums[1], line[3])
		case 'o':
			p.Ostacolo(nums[0], nums[1], nums[2], nums[3])
		case 'r':
			p.Richiamo(nums[0], nums[1], line[3])
		case 'p':
			fmt.Println(p.Posizioni(line[1]))
		case 'e':
			esiste := p.EsistePercorso(nums[0], nums[1], line[3])
			if esiste {
				fmt.Println("SI")
			} else {
				fmt.Println("NO")
			}
		case 'f':
			os.Exit(0)
		case 't':
			fmt.Println(0) // Placeholder per test
		default:
		}
	}

}

type piano = utils.Piano
