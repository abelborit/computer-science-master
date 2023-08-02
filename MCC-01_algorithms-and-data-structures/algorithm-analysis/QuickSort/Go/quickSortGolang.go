package main

// import "fmt"

// Obtenemos el arreglo de elementos
func quickSort(arr []int) []int {
	// Validamos si el arreglo tiene más de un elemento
	if len(arr) <= 1 {
		// Si solo tiene un elemento retornanos el arreglo porque ya estaría ordenado
		return arr
	}

	// Definimos la variable pivote que en este caso sera el primer elemento (este nos servirá para poder hacer las comparaciones)
	pivot := arr[0]

	// Definimos un arreglo para elementos hacia la izquierda del pivote (los elementos menores)
	left := make([]int, 0)
	// Definimos un arreglo para elementos hacia la derecha del pivote (los elementos iguales o mayores)
	right := make([]int, 0) 

	// Iteramos el arreglo comparando los elementos con el pivote
	for _, element := range arr[1:] {
		if element < pivot {
			// Si el elemento a comparar es menor que el pivote, lo añadimos al arrego de elementos hacia la izquierda
			left = append(left, element)
		} else {
			// Si el elemento a comparar es mayor que el pivote, lo añadimos al arrego de elementos hacia la derecha
			right = append(right, element)
		}
	}

	// Aplicamos recursivamente QuickSort en las sublistas izquierda y derecha
	return append(append(quickSort(left), pivot), quickSort(right)...)
}

// *******************************************************************
// func main() {
// 	// Ejemplo de uso
// 	unsortedSlice := []int{7, 2, 1, 2, 10, 6, 8, 5, 3, 4}
// 	sortedSlice := quickSort(unsortedSlice)
// 	fmt.Println(sortedSlice) // Output: [1 2 2 3 4 5 6 7 8 10]
// }
