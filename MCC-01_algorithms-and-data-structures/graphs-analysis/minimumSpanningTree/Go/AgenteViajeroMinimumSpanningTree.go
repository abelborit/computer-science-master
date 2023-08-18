package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func nameAlgorithm() {
	fmt.Println("**************************************")
	fmt.Println("Agente Viajero - Minimum Spanning Tree")
	fmt.Println("**************************************")
}

// readArrayFromFileNodes leerá un archivo .txt referente a los nodos y devuelve una lista de nodos como strings
func readArrayFromFileNodes(filePath string) ([]string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	contentStr := string(content)
	contentStr = strings.TrimSpace(contentStr)
	nodes := strings.Split(contentStr, ",")

	return nodes, nil
}

// readArrayFromFileGraphs leerá un archivo .txt referente a las matrices y devuelve una matriz de enteros
func readArrayFromFileGraphs(filePath string) ([][]int, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")
	var result [][]int

	for _, line := range lines {
		if line == "" {
			continue
		}

		numStrings := strings.Split(line, ",")
		var numArray []int

		for _, numStr := range numStrings {
			num, err := strconv.Atoi(strings.TrimSpace(numStr))
			if err != nil {
				return nil, err
			}
			numArray = append(numArray, num)
		}

		result = append(result, numArray)
	}

	return result, nil
}

// Calcula la distancia total recorrida en una ruta. Itera a través de la ruta y suma las distancias entre nodos consecutivos
func calcularDistancia(ruta []int, distancias [][]int) int {
	distanciaTotal := 0
	for i := 0; i < len(ruta)-1; i++ {
		distanciaTotal += distancias[ruta[i]][ruta[i+1]]
	}
	distanciaTotal += distancias[ruta[len(ruta)-1]][ruta[0]]
	return distanciaTotal
}

// Implementar el algoritmo del Agente Viajero basado en el Minimum Spanning Tree
func agenteViajeroMinimumSpanningTree(distancias [][]int, nombresCiudades []string) map[string]interface{} {
	startTime := time.Now()

	// Obtiene el número de ciudades
	numCiudades := len(distancias)

	// Crea una lista con los índices de las ciudades que deben ser visitadas
	ciudadesPorVisitar := make([]int, numCiudades)
	for i := 0; i < numCiudades; i++ {
		ciudadesPorVisitar[i] = i
	}

	// Crea una matriz para almacenar el árbol de expansión mínima
	arbolExpMinimo := make([][]int, numCiudades)
	for i := 0; i < numCiudades; i++ {
		arbolExpMinimo[i] = make([]int, numCiudades)
	}

	// Inicializa la ruta y el nodo actual
	ruta := []int{}
	nodoActual := 0
	indexNodoActual := -1

	// Encuentra el índice del nodo actual en la lista de ciudades por visitar
	for i, v := range ciudadesPorVisitar {
		if v == nodoActual {
			indexNodoActual = i
			break
		}
	}

	// Elimina el nodo actual de la lista de ciudades por visitar
	ciudadesPorVisitar = append(ciudadesPorVisitar[:indexNodoActual], ciudadesPorVisitar[indexNodoActual+1:]...)

	// Mientras hayan ciudades por visitar
	for len(ciudadesPorVisitar) > 0 {
		// Encuentra la ciudad más cercana a la actual
		ciudadMasCercana := ciudadesPorVisitar[0]
		distanciaMinima := distancias[nodoActual][ciudadMasCercana]

		for _, ciudad := range ciudadesPorVisitar {
			distanciaActual := distancias[nodoActual][ciudad]
			if distanciaActual < distanciaMinima {
				ciudadMasCercana = ciudad
				distanciaMinima = distanciaActual
			}
		}

		// Agrega la distancia al árbol de expansión mínima
		arbolExpMinimo[nodoActual][ciudadMasCercana] = distanciaMinima
		arbolExpMinimo[ciudadMasCercana][nodoActual] = distanciaMinima

		// Agrega el nodo actual a la ruta y actualiza el nodo actual
		ruta = append(ruta, nodoActual)
		nodoActual = ciudadMasCercana

		// Encuentra el índice del nodo actual en la lista de ciudades por visitar
		index := -1
		for i, v := range ciudadesPorVisitar {
			if v == nodoActual {
				index = i
				break
			}
		}

		// Elimina el nodo actual de la lista de ciudades por visitar
		ciudadesPorVisitar = append(ciudadesPorVisitar[:index], ciudadesPorVisitar[index+1:]...)
	}

	// Agrega el último nodo a la ruta y cierra el ciclo
	ruta = append(ruta, nodoActual)
	ruta = append(ruta, ruta[0])

	// Calcula la distancia total de la ruta
	distanciaTotal := calcularDistancia(ruta, distancias)

	// Convierte los índices de la ruta en nombres de ciudades
	rutaNombres := make([]string, len(ruta))
	for i, v := range ruta {
		rutaNombres[i] = nombresCiudades[v]
	}

	// Calcula el tiempo de ejecución.
	// elapsedTime := float64(time.Since(startTime).Milliseconds())
	elapsedTime := time.Since(startTime)

	// Crea un mapa con los resultados
	resultado := map[string]interface{}{
		"ruta":      rutaNombres,
		"distancia": distanciaTotal,
		"tiempo_ms": elapsedTime,
	}

	return resultado
}

func main() {
	// Obtener el directorio actual
	currentDirectory, err := os.Getwd()
	if err != nil {
		fmt.Println("Error obteniendo el directorio actual:", err)
		return
	}

	// Obtener la ruta absoluta de la carpeta Assets
	assetsPath := filepath.Join(currentDirectory, "MCC-01_algorithms-and-data-structures", "graphs-analysis", "minimumSpanningTree", "Assets")

	// Definir la cantidad de nodos que va a procesar
	quantityNodes := "1"

	// Definir las rutas de los archivos para los nodos
	filePathsNodes := filepath.Join(assetsPath, "generateNodesGraph", "nodesGraph"+quantityNodes+".txt")
	resultNodes, err := readArrayFromFileNodes(filePathsNodes)
	if err != nil {
		log.Fatal(err)
	}
	resultNodesToFunction := resultNodes

	// Definir las rutas de los archivos para las matrices
	filePathsGraphs := filepath.Join(assetsPath, "generateRandomGraph", "randomGraph"+quantityNodes+".txt")
	resultGraphs, err := readArrayFromFileGraphs(filePathsGraphs)
	if err != nil {
		log.Fatal(err)
	}
	resultGraphsToFunction := resultGraphs

	// fmt.Println("resultNodesToFunction:", resultNodesToFunction)
	// fmt.Println("resultGraphsToFunction:", resultGraphsToFunction)

	resultado := agenteViajeroMinimumSpanningTree(resultGraphsToFunction, resultNodesToFunction)

	fmt.Println("Mejor ruta:", resultado["ruta"])
	fmt.Println("Distancia:", resultado["distancia"])
	// fmt.Println("Tiempo de ejecución:", resultado["tiempo_ms"], "ms")
	fmt.Println("Tiempo de ejecución:", resultado["tiempo_ms"])
}
