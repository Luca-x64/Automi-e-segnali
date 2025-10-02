// GHIRIMOLDI LUCA 31974A

package utils

import (
	"fmt"
	"strings"
)

// crea un nuovo piano
func Crea() *Piano {
	p := &Piano{&Albero{&Nodo{}}, &Ostacoli{}}
	return p
}

/*
restituisce un carattere che indica cosa si trova nella posizione (x, y): A se un automa, O se un ostacolo,
E se la posizione è vuota.
*/
func (p *Piano) Stato(x, y int) string {
	it := p.automi.radice.Iteratore()
	for it.HasNext() {
		automa := it.Next()
		if automa.pos.x == x && automa.pos.y == y {
			return "A"
		}
	}
	if !p.coordinateLibere(x, y) {
		return "O"
	}

	return "E"
}

/*
Stampa l’elenco degli automi seguito dall’elenco degli ostacoli, secondo quanto indicato nelle Speci-
fiche di implementazione.
*/
func (piano *Piano) Stampa() {
	sb := strings.Builder{}

	it := piano.automi.radice.Iteratore()

	sb.WriteString("(\n")
	for it.HasNext() {
		automa := it.Next()
		sb.WriteString(fmt.Sprintf("%s: %d,%d\n", automa.nome, automa.pos.x, automa.pos.y))
	}
	sb.WriteString(")\n[\n")
	for _, v := range *piano.ostacoli {
		sb.WriteString(fmt.Sprintf("(%d,%d)(%d,%d)\n", v.p0.x, v.p0.y, v.p1.x, v.p1.y))
	}
	sb.WriteRune(']')
	fmt.Println(sb.String())

}

/*
Se il punto (x, y) è contenuto in qualche ostacolo, allora non esegue alcuna operazione. Altrimenti,
se non esiste alcun automa di nome η lo crea e lo pone in (x, y). Se η esiste già, lo riposiziona nel
punto (x, y).
*/
func (piano *Piano) Automa(x, y int, n string) {
	point := Coordinata{x, y}
	if piano.coordinateLibere(x, y) {
		nodo := piano.cerca(n)
		if nodo == nil {
			piano.inserisci(&Automa{n, point})
		} else {
			nodo.val = &Automa{n, point}
		}

	}
}

/*Se i punti nel rettangolo R(x0 , y0 , x1 , y1 ) non contengono alcun automa, inserisce nel piano l’ostacolo
rappresentato da R(x0 , y0 , x1 , y1 ), altrimenti non compie alcuna operazione. Si può assumere che
x0 < x1 e y0 < y1 .
*/

func (piano *Piano) Ostacolo(x0, y0, x1, y1 int) {

	p0 := Coordinata{x0, y0}
	p1 := Coordinata{x1, y1}
	trovato := false
	it := piano.automi.radice.Iteratore()
	for it.HasNext() {
		automa := it.Next()
		if automa.pos.isInCoord(p0, p1) {
			trovato = true
			break
		}
	}

	if !trovato {
		*piano.ostacoli = append(*piano.ostacoli, Ostacolo{p0, p1})
	}
}

/*
Viene emesso il segnale di richiamo α dal punto (x, y) (ovviamente,
se il punto (x, y) appartiene a qualche ostacolo, esso non `e raggiungibile da alcun automa).
*/
func (piano *Piano) Richiamo(x int, y int, s string) { // (x,y) è il punto di richiamo, s è il prefisso degli automi che possono raggiungere il punto di richiamo

	if !piano.coordinateLibere(x, y) { // se il punto (x, y) appartiene a qualche ostacolo, esso non `e raggiungibile da alcun automa).
		return
	}

	coordRichiamo := Coordinata{x, y}
	nodo := piano.cerca(s)
	var automiRaggiungibili []*Automa
	it := nodo.Iteratore()
	for it.HasNext() { // itera su tutti gli automi con prefisso s
		automa := it.Next()
		if piano.EsistePercorso(coordRichiamo.x, coordRichiamo.y, automa.nome) {
			automiRaggiungibili = append(automiRaggiungibili, automa)
		}
	}

	if len(automiRaggiungibili) == 0 {
		return
	}

	distanzaMinima := automiRaggiungibili[0].pos.Distanza(coordRichiamo)

	// trova la distanza minima
	
	for _, automa := range automiRaggiungibili[1:] {
		distanza := automa.pos.Distanza(coordRichiamo)
		if distanza < distanzaMinima {
			distanzaMinima = distanza
		}
	}
	
	

	for _, automa := range automiRaggiungibili {
		if automa.pos.Distanza(coordRichiamo) == distanzaMinima {
		automa.pos = coordRichiamo
	}
	}
}

/*
restituisce le posizioni di tutti gli automi η tali che α è un prefisso di η, secondo il formato definito
nelle note della sezione Specifiche di implementazione.
*/
func (piano *Piano) Posizioni(n string) string {
	sb := strings.Builder{}
	nodo := piano.cerca(n)
	it := nodo.Iteratore()
	sb.WriteString("(\n")
	for it.HasNext() {
		automa := it.Next()
		sb.WriteString(fmt.Sprintf("%s: %d,%d\n", automa.nome, automa.pos.x, automa.pos.y))
	}
	sb.WriteString(")")
	return sb.String()
}

/*
Restituisce true se esiste almeno un percorso libero da P(η) a (x, y) di lunghezza D(P(η),(x, y)), false
in caso contrario (ovviamente, false se η non esiste o se (x, y) `e un punto all’interno di un
ostacolo).
*/
func (piano *Piano) EsistePercorso(x int, y int, n string) bool {
	if !piano.coordinateLibere(x, y) { // contenuto in ostacoli
		return false
	}

	nodo := piano.cerca(n)
	if nodo == nil || nodo.val == nil {
		return false // non esiste automa di nome n
	}
	if len(*piano.ostacoli) == 0 {
		return true // non ci sono ostacoli
	}

	dest := Coordinata{x, y}
	if nodo.val.pos == dest { return false}
	pos := nodo.val.pos
	if pos.x == x { // Se i due punti hanno la stessa x
		min := min(pos.y,y)
		max := max(pos.y, y)
		for i:= min; i <= max ; i += 2 {
			if !piano.coordinateLibere(x, i) {
				return false
			}
		}
		distanzaPercorsa := max-min
		return pos.Distanza(dest) == distanzaPercorsa
	}

	if pos.y == y { // Se i due punti hanno la stessa y
		min := min(pos.x, x)
		max := max(pos.x, x)
		for i := min; i <= max; i += 2 {
			if !piano.coordinateLibere(i, y) {
				return false
			}
		}
		return pos.Distanza(dest) == max-min 
	}


	percorsoCalcolato := piano.bfs(pos, dest)
	return percorsoCalcolato != -1
}
