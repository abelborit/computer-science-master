package main

// Obtenemos el arreglo de elementos
func mergeSort(arr []int) []int {
	// Validamos si el arreglo tiene más de un elemento
	// Caso base: Si la lista tiene un elemento o está vacía, está ordenada.
	if len(arr) <= 1 {
		// Si solo tiene un elemento retornanos el arreglo porque ya estaría ordenado
		return arr
	}

	// Encontrar el punto medio de la lista
	middle := len(arr) / 2

	// Dividir la lista en dos partes
	left := arr[:middle]
	right := arr[middle:]

	// Función auxiliar para combinar dos sublistas en orden
	merge := func(leftArr, rightArr []int) []int {
		result := make([]int, 0)
		leftIndex, rightIndex := 0, 0

		// Comparar elementos de las dos sublistas y combinar en orden, repite el proceso de combinación hasta que solo quede una lista ordenada
		for leftIndex < len(leftArr) && rightIndex < len(rightArr) {
			if leftArr[leftIndex] < rightArr[rightIndex] {
				result = append(result, leftArr[leftIndex])
				leftIndex++
			} else {
				result = append(result, rightArr[rightIndex])
				rightIndex++
			}
		}

		result = append(result, leftArr[leftIndex:]...)
		result = append(result, rightArr[rightIndex:]...)
		return result
	}

	// Aplicar Merge Sort en las dos mitades y combinar los resultados
	return merge(mergeSort(left), mergeSort(right))
}

// *******************************************************************
// func main() {
// 	// Ejemplo de uso
// 	unsortedSlice := []int{38, 27, 43, 3, 9, 82, 10}
// 	sortedSlice := mergeSort(unsortedSlice)
// 	fmt.Println(sortedSlice) // Output: [3 9 10 27 38 43 82]
// }