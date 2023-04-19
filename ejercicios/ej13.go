package ejercicios

// Escriba un método recursivo para determinar si
// un elemento está en un arreglo utilizando búsqueda binaria.
func BusquedaBinaria(arreglo []int, elemento int) bool {
	return busquedaBinaria(arreglo, 0, len(arreglo)-1, elemento)
}

func busquedaBinaria(array []int, inicio int, fin int, x int) bool{
	if inicio > fin{
		return false;
	}
	medio:=(inicio+fin)/2
	if array[medio] > x{
		return busquedaBinaria(array, inicio, medio-1, x)
	} else if array[medio] < x{
		return busquedaBinaria(array, medio+1, fin, x)
	}
	return true
}