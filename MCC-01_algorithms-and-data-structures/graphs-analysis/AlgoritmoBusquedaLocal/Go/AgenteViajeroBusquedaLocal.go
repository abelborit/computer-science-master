package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Graph map[int]map[int]int

func guardarEnArchivo(ciclo []int, costo int, tiempo float64, cantidadNodos int) error {
	output := "Results\\Results_Go.txt"
	file, err := os.OpenFile(output, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	resultado := map[string]interface{}{
		"CicloHamiltoniano": ciclo,
		"Costo":             costo,
		"TiempoEjecucion":   tiempo,
		"CantidadNodos":     cantidadNodos,
	}

	resultadoJSON, err := json.Marshal(resultado)
	if err != nil {
		return err
	}

	_, err = file.Write(append(resultadoJSON, []byte("\n")...))
	if err != nil {
		return err
	}

	return nil
}

func pesoCiclo(ciclo []int, grafo Graph) int {
	peso := 0
	for i := 0; i < len(ciclo)-1; i++ {
		peso += grafo[ciclo[i]][ciclo[i+1]]
	}
	peso += grafo[ciclo[len(ciclo)-1]][ciclo[0]]
	return peso
}

func busquedaLocal(grafo Graph, maxIteraciones int) ([]int, int, float64) {
	cicloActual := generarCicloInicial(grafo)
	costoActual := pesoCiclo(cicloActual, grafo)

	startTime := time.Now()
	for i := 0; i < maxIteraciones; i++ {
		idx1 := rand.Intn(len(cicloActual))
		idx2 := rand.Intn(len(cicloActual))
		nuevoCiclo := append([]int{}, cicloActual...)
		nuevoCiclo[idx1], nuevoCiclo[idx2] = nuevoCiclo[idx2], nuevoCiclo[idx1]
		nuevoCosto := pesoCiclo(nuevoCiclo, grafo)

		if nuevoCosto < costoActual {
			cicloActual = nuevoCiclo
			costoActual = nuevoCosto
		}
	}
	endTime := time.Now()
	tiempoProcesamiento := endTime.Sub(startTime).Seconds()
	return cicloActual, costoActual, tiempoProcesamiento
}

func generarCicloInicial(grafo Graph) []int {
	vertices := make([]int, 0, len(grafo))
	for vertice := range grafo {
		vertices = append(vertices, vertice)
	}

	rand.Seed(time.Now().UnixNano())
	verticeInicial := vertices[rand.Intn(len(vertices))]

	ciclo := []int{verticeInicial}
	nodosNoVisitados := make(map[int]bool)
	for _, vertice := range vertices {
		if vertice != verticeInicial {
			nodosNoVisitados[vertice] = true
		}
	}

	for len(nodosNoVisitados) > 0 {
		for vecino := range grafo[ciclo[len(ciclo)-1]] {
			if nodosNoVisitados[vecino] {
				ciclo = append(ciclo, vecino)
				delete(nodosNoVisitados, vecino)
				break
			}
		}
	}
	return ciclo
}

func main() {
	input := "Assets\\generateNodesGraphBL\\nodeGraph5120.json"
	fileData, err := os.ReadFile(input)
	if err != nil {
		fmt.Println("Error al leer el archivo JSON:", err)
		return
	}

	var grafo Graph
	err = json.Unmarshal(fileData, &grafo)
	if err != nil {
		fmt.Println("Error al decodificar el archivo JSON:", err)
		return
	}

	ciclo, costo, tiempo := busquedaLocal(grafo, 10000)
	fmt.Printf("Ciclo hamiltoniano encontrado: %v -> %v con peso %v\n", ciclo, ciclo[0], costo)
	fmt.Printf("Tiempo de procesamiento: %.6f segundos\n", tiempo)

	cantidadNodos := len(grafo)
	err = guardarEnArchivo(ciclo, costo, tiempo, cantidadNodos)
	if err != nil {
		fmt.Println("Error al guardar en el archivo:", err)
	} else {
		fmt.Println("Información guardada exitosamente en ResultsGo.json")
	}
}
