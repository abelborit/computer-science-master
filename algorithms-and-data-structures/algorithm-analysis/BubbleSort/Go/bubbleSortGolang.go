package main

// import "fmt"

// Obtenemos el arreglo de elementos
func bubbleSort(arr []int) []int {
	// Se declara una variable n que almacenará la longitud del arreglo para evitar calcularla repetidamente
	n := len(arr)
	// Se declara una variable swapped para controlar si se realizaron cambios en el arreglo durante un pase completo del algoritmo. Si en algún pase no se realizaron intercambios, significa que el arreglo está ordenado y el algoritmo puede detenerse
	swapped := true

	// Se inicia un bucle for que continuará hasta que no se realicen más intercambios (es decir, swapped sea false)
	for swapped {
		// Dentro del bucle for, se establece swapped como false al inicio de cada iteración para verificar si ocurren intercambios en el bucle anidado
			swapped = false

			// Se recorre la slice usando otro bucle for desde el índice 0 hasta el penúltimo elemento (n - 1)
			for i := 0; i < n-1; i++ {
				// En cada iteración del bucle anidado, se compara el elemento actual arr[i] con el siguiente elemento arr[i + 1]
					if arr[i] > arr[i+1] {
						// Si el elemento actual es mayor que el siguiente elemento, se realiza un intercambio de posiciones utilizando la asignación múltiple de Go

							// Intercambio de elementos si están en el orden incorrecto
							arr[i], arr[i+1] = arr[i+1], arr[i]

							// Se actualiza la variable swapped a true para indicar que se realizó un intercambio en esta iteración
							swapped = true
					}
				}
				// Si no ocurre ningún intercambio durante una pasada completa del bucle anidado, swapped seguirá siendo false, y el bucle principal se detendrá, ya que la slice estará completamente ordenada
	}

	// Finalmente, cuando el bucle for termine, se devolverá el arreglo ordenado
	return arr
}

// *******************************************************************
// func main() {
// 	// Ejemplo de uso
//     unsortedArray := []int{64, 34, 25, 12, 22, 11, 90}
//     bubbleSort(unsortedArray)
//     fmt.Println(unsortedArray) // Output: [11 12 22 25 34 64 90]
// }



