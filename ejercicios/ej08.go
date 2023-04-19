package ejercicios

// Escribir un método recursivo que devuelva el
// cociente y el resto de la división entera mediante
// restas sucesivas
func DivisionEntera(dividendo, divisor int) (cociente, resto int) {
	
	// Si el numerador es 0, devuelve 0
	if dividendo == 0 {
		return 0, 0
	}
	// Si el denominador es 0, devuelve 0 y como resto el numerador.
	if divisor == 0 || dividendo < divisor{  // funciona sacando el dividendo < divisor (?)
		return 0, dividendo
		// Deberia ser error, pero no se como idearlo por ahora.
	}

	coc, res := DivisionEntera(dividendo-divisor, divisor)
	return coc + 1, res 

}
