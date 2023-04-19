package ejercicios

import "fmt"

// Escribir un método recursivo que dado un número
// entero decimal devuelva su equivalente en
// binario en forma de string
func DecimalABinario(decimal int) string {
	fmt.Println("num:", decimal, "resto:",decimal%2)
	if decimal == 0 {
		return "0"
	}
	if decimal == 1 {
		return "1"
	}
	resultado := fmt.Sprintf("%s%d", DecimalABinario(decimal/2), decimal%2)
	return resultado
}
