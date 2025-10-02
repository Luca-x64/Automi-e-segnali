// GHIRIMOLDI LUCA 31974A

package utils


// import "container/list"

type Piano struct {
	automi   *Albero
	ostacoli *Ostacoli
}
type Ostacoli = []Ostacolo

type Automa struct {
	nome string
	pos  Coordinata
}

type Ostacolo struct {
	p0, p1 Coordinata
}

// Albero è un tipo che rappresenta un albero binario
type Albero = struct {
	radice *Nodo
}

type Nodo struct {
	sx  *Nodo
	dx  *Nodo
	val *Automa
}

// Inserisce l'automa nell'albero binario
func (p *Piano) inserisci(automa *Automa) {
	node := p.automi.radice
	for _, char := range automa.nome {

		if char == '0' {
			if node.sx == nil {
				node.sx = &Nodo{sx: nil, dx: nil, val: nil}
			}
			node = node.sx
		} else if char == '1' {
			if node.dx == nil {
				node.dx = &Nodo{sx: nil, dx: nil, val: nil}
			}
			node = node.dx
		}

	}
	node.val = automa
}

// cerca un nodo, dato il prefisso
func (p *Piano) cerca(prefisso string) (node *Nodo) {
	node = p.automi.radice
	if node == nil {
		return nil
	}
	for _, char := range prefisso {
		if char == '0' {
			if node.sx == nil {
				return nil
			}
			node = node.sx
		} else if char == '1' {
			if node.dx == nil {
				return nil
			}
			node = node.dx
		}
	}
	return node
}

// ritorna true se la coorodinate di posizione {x,y} è libera (non è contenuta in ostacoli)
func (piano *Piano) coordinateLibere(x, y int) bool {
	point := Coordinata{x, y}
	for _, v := range *piano.ostacoli {
		if point.isInCoord(v.p0, v.p1) {
			return false
		}
	}
	return true
}

// implementa BFS ottimizzato per trovare il percorso libero, restituisce la distanza e -1 se non esiste un percorso libero
func (piano *Piano) bfs(start, end Coordinata) int {
	if start == end {
		return 0
	}

	type NodoPercorso struct {
		posizione Coordinata
		distanza  int
	}

	direzioni := []Coordinata{}
	if end.y > start.y { direzioni = append(direzioni, Coordinata{0, 1}) // Su
    } else { direzioni = append(direzioni, Coordinata{0, -1}) }// Giù 
	if end.x > start.x { direzioni = append(direzioni, Coordinata{1, 0}) // Destra
    } else  { direzioni = append(direzioni, Coordinata{-1, 0}) } // Sinistra 
    
	visitati := make(map[Coordinata]bool)
	coda := &Coda[*NodoPercorso]{}

	coda.push(&NodoPercorso{start, 0})
	
	for !coda.isEmpty() {
		nodo := coda.pop()

		for _, direzione := range direzioni {
			
			nuovaPosizione := Coordinata{nodo.posizione.x + direzione.x, nodo.posizione.y + direzione.y}

			if nuovaPosizione == end {
				return nodo.distanza + 1
			}
			if !visitati[nuovaPosizione] && nuovaPosizione.Distanza(end) < nodo.posizione.Distanza(end) && piano.coordinateLibere(nuovaPosizione.x, nuovaPosizione.y) {
				visitati[nuovaPosizione] = true
				coda.push(&NodoPercorso{nuovaPosizione, nodo.distanza + 1})
			}
			
		}
	}

	return -1
}
