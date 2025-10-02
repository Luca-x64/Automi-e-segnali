// GHIRIMOLDI LUCA 31974A

package utils


type Coda[T any] struct {
	elementi []T // Slice per memorizzare gli elementi generici
	testa    int // Indice dell'inizio della coda
}

// push aggiunge un elemento alla fine della coda
func (c *Coda[T]) push(element T) {
	c.elementi = append(c.elementi, element)
}

// pop rimuove e restituisce l'elemento in testa alla coda
func (c *Coda[T]) pop() T {
	if c.isEmpty() {
		var zeroValue T
		return zeroValue // Restituisce il valore zero per il tipo generico
	}

	elem := c.elementi[c.testa]
	c.testa++ // Sposta l'indice di inizio in avanti

	// Riallinea la coda per risparmiare memoria quando necessario
	if c.testa > len(c.elementi)/2 {
		c.elementi = c.elementi[c.testa:]
		c.testa = 0
	}
	return elem
}

// isEmpty controlla se la coda è vuota
func (c *Coda[T]) isEmpty() bool {
	return c.testa >= len(c.elementi)
}
