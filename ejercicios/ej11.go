package ejercicios

// Escribir un método recursivo que reciba 2 enteros
// n y b y devuelva true si n es potencia de b.
// Por ejemplo: esPotencia(8, 2) devuelve true.
func EsPotencia(n, b int) bool {
	
	// Caso base: si n es igual a 1 => es potencia de b
	if n == 1{
		return true
		// Caso base: Si n no es divisible por b o es igual a 0 => no es potencia de b
	} else if n % b != 0 || n == 0{
		return false
	} else{
		// Caso recursivo: Si n es divisible por b, se va dividiendo n por b
		return EsPotencia(n / b, b)
	}
}
