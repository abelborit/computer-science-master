package main

// import "fmt"

// Obtenemos el arreglo de elementos
func selectionSort(arr []int) []int {
	// Obtener la longitud del arreglo
	n := len(arr) 

	// Iterar a través del arreglo
	for i := 0; i < n-1; i++ {
		// Encontrar el índice del elemento mínimo en la sublista no ordenada
		minIndex := i // Suponemos que el índice del elemento mínimo es el índice actual i

		// Bucle interno: Buscar el mínimo en la sublista no ordenada (desde i + 1 hasta el final del arreglo)
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				// Si encontramos un elemento menor al elemento en minIndex
				minIndex = j // Actualizamos minIndex al índice del nuevo elemento mínimo
			}
		}

		// Intercambiar el elemento mínimo con el elemento en la posición actual (i)
		if minIndex != i {
			// Solo intercambiamos si el elemento mínimo no está en su posición actual
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
	}

	return arr // Devolver el arreglo ordenado
}


// *******************************************************************
// func main() {
// 	// Ejemplo de uso
//     unsortedArray := []int{64, 34, 25, 12, 22, 11, 90}
//     selectionSort(unsortedArray)
//     fmt.Println(unsortedArray) // Output: [11 12 22 25 34 64 90]
// }











