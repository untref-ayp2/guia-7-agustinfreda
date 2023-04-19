package ejercicios

import "fmt"

// Escribir un método que calcule recursivamente cuántos
// elementos hay en una pila y no altere el contenido de
// la misma. La pila sólo tiene los métodos Push, Pop y isEmpty.
func CantidadDeElementos(pila Stack) int {
	fmt.Println("1:",pila)
	
	// Caso base: si la pila esta vacia retorna 0
	if pila.IsEmpty() {
		return 0
	}
	// Se saca un elemento de la pila y se guarda en una variable
	pila.Pop()
	temp := CantidadDeElementos(pila) + 1
	pila.Push(temp)
	fmt.Println("2:",pila)
	return temp
}
