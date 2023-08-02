// tiene que colocar en la terminal: go run QuickSort/Go/indexQuickSort.go QuickSort/Go/quickSortGolang.go
package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func nameAlgorithmh() {
	fmt.Println("*************")
	fmt.Println("QUICK SORT GO")
	fmt.Println("*************")
}

// readArrayFromFile reads an integer array from the provided file path.
func readArrayFromFile(filePath string) ([]int, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	data := string(content)
	strNums := strings.Split(data, ",")
	var arr []int

	for _, strNum := range strNums {
		num, err := strconv.Atoi(strings.TrimSpace(strNum))
		if err != nil {
			return nil, err
		}
		arr = append(arr, num)
	}

	return arr, nil
}

// processAndSortArray lee y ordena el arreglo del archivo, y muestra los resultados.
func processAndSortArray(filePath string) {
	unsortedArray, err := readArrayFromFile(filePath)
	if err != nil {
			fmt.Println("Error leyendo el archivo:", err)
			return
	}

	// Número de repeticiones para hallar su promedio de tiempo de ejecución
	numRepetitions := 10
	var totalExecutionTime time.Duration
	var sortedArray []int

	// Guardar los tiempos individuales en un slice en milisegundos
	var executionTimes []float64

	for i := 1; i <= numRepetitions; i++ {
			// fmt.Println("se repite", i, "veces")
			startTimeSort := time.Now()
			sortedArray = quickSort(unsortedArray)
			endTimeSort := time.Now()

			executionTime := endTimeSort.Sub(startTimeSort).Seconds() * 1000 // Convertir a milisegundos
			executionTimes = append(executionTimes, executionTime)

			totalExecutionTime += endTimeSort.Sub(startTimeSort)
	}

	averageExecutionTime := totalExecutionTime / time.Duration(numRepetitions)

	// Calcular la desviación estándar en milisegundos
	var sum float64
	for _, time := range executionTimes {
			sum += time
	}
	mean := sum / float64(len(executionTimes))

	var variance float64
	for _, time := range executionTimes {
			variance += (time - mean) * (time - mean)
	}
	variance = variance / float64(len(executionTimes))

	stdDeviation := math.Sqrt(variance)

	fmt.Println("--------------------")
	fmt.Println(filePath)
	fmt.Println("--------------------")

	fmt.Println("Arreglo no ordenado:", unsortedArray)
	fmt.Println("")
	fmt.Println("Arreglo ordenado:", sortedArray)
	fmt.Println("")
	fmt.Println("Tiempo promedio de ejecución:", averageExecutionTime.Milliseconds(), "milisegundos")
	fmt.Println("Desviación estándar:", stdDeviation, "milisegundos")
	fmt.Println("********************************************************************************")
}

func main() {
	nameAlgorithmh()
	
	// Obtener el directorio actual.
	currentDirectory, err := os.Getwd()
	if err != nil {
			fmt.Println("Error obteniendo el directorio actual:", err)
			return
	}

  // Obtener la ruta absoluta de la carpeta ASSETS
  assetsPath := filepath.Join(currentDirectory, "ASSETS")

	// Definir las rutas de los archivos.
	filePaths := []string{
		filepath.Join(assetsPath, "unsortedArrays", "unsortedArray100.txt"),
		// filepath.Join(assetsPath, "unsortedArrays", "unsortedArray1000.txt"),
		// filepath.Join(assetsPath, "unsortedArrays", "unsortedArray2000.txt"),
		// filepath.Join(assetsPath, "unsortedArrays", "unsortedArray3000.txt"),
		// filepath.Join(assetsPath, "unsortedArrays", "unsortedArray4000.txt"),
		// filepath.Join(assetsPath, "unsortedArrays", "unsortedArray5000.txt"),
		// filepath.Join(assetsPath, "unsortedArrays", "unsortedArray6000.txt"),
		// filepath.Join(assetsPath, "unsortedArrays", "unsortedArray7000.txt"),
		// filepath.Join(assetsPath, "unsortedArrays", "unsortedArray8000.txt"),
		// filepath.Join(assetsPath, "unsortedArrays", "unsortedArray9000.txt"),
		// filepath.Join(assetsPath, "unsortedArrays", "unsortedArray10000.txt"),
		// filepath.Join(assetsPath, "unsortedArrays", "unsortedArray20000.txt"),
		// filepath.Join(assetsPath, "unsortedArrays", "unsortedArray30000.txt"),
		// filepath.Join(assetsPath, "unsortedArrays", "unsortedArray40000.txt"),
		// filepath.Join(assetsPath, "unsortedArrays", "unsortedArray50000.txt"),
	}

	// Procesar y ordenar cada archivo
	for _, file := range filePaths {
		processAndSortArray(file)
}
}