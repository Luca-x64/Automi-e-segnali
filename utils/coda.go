package utils


type Coda[T any] struct {
	elements []T // Slice per memorizzare gli elementi generici
	head    int // Indice dell'inizio della coda
}

// push aggiunge un elemento alla fine della coda
func (c *Coda[T]) push(element T) {
	c.elements = append(c.elements, element)
}

// pop rimuove e restituisce l'elemento in testa alla coda
func (c *Coda[T]) pop() T {
	if c.isEmpty() {
		var zeroValue T
		return zeroValue // Restituisce il valore zero per il tipo generico
	}

	elem := c.elements[c.head]
	c.head++ // Sposta l'indice di inizio in avanti

	// Riallinea la coda per risparmiare memoria quando necessario
	if c.head > len(c.elements)/2 {
		c.elements = c.elements[c.head:]
		c.head = 0
	}
	return elem
}

// isEmpty controlla se la coda è vuota
func (c *Coda[T]) isEmpty() bool {
	return c.head >= len(c.elements)
}
