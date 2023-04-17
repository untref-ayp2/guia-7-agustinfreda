package ejercicios

// Escriba un método recursivo que tome una cadena y
// devuelva como resultado la cadena invertida.
func Invertir(cadena string) string { // O(n)
	if len(cadena) == 0{
		return cadena
	}
	return Invertir(cadena[1:]) + string(cadena[0])
}
