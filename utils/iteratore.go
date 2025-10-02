// GHIRIMOLDI LUCA 31974A

package utils

type Iteratore struct {
	coda *Coda[*Nodo]
}

func (n *Nodo) Iteratore() *Iteratore {
	it := &Iteratore{&Coda[*Nodo]{}}
	if n != nil {
		it.aggiungi(n)
	}
	return it
}

func (it *Iteratore) aggiungi(nodoRadice *Nodo) {
	if nodoRadice != nil {
		if nodoRadice.val != nil {
			it.coda.push(nodoRadice)
		}
		if nodoRadice.dx != nil {
			it.aggiungi(nodoRadice.dx)
		}
		if nodoRadice.sx != nil {
			it.aggiungi(nodoRadice.sx)
		}
	}
}

func (it *Iteratore) HasNext() bool {
	return !it.coda.isEmpty()
}

func (it *Iteratore) Next() *Automa {
	if !it.HasNext() {
		return nil
	}

	current := it.coda.pop()

	return current.val
}
