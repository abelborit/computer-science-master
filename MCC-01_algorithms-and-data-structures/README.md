# Algoritmos y Estructuras de Datos

---

## 1. Análisis de algoritmos (algorithm-analysis)

#### 1.1 Competencias:

- Implementar, analizar y comparar algoritmos de ordenamiento.

#### 1.2 Material:

- Formato del informe:
  - LaTex
- Lenguajes de programación:
  - Pyhton
  - Golang (Go)
  - C++
- Método y/o forma para realizar comparaciones:
  - Notación Big-O
- Gráficos para la comparación:
  - Matplotlib

#### 1.3 Asignación:

- Desarrollar una comparación entre Golang, Python y C++. Dicha comparación debe realizarse en términos de tiempo de procesamiento. Implementar cuatro (04) programas en los tres (03) lenguages y luego medir el tiempo de procesamiento.

- Los programas a implementar pueden ser: Algoritmo de burbuja para ordenar listas, Algoritmo Quick sort, Algoritmo Merge sort, solución al problema de n-reinas, etc. (deben ser los mismos algoritmos en los tres lenguajes)

- Para medir el tiempo de procesamiento considerar:
  - Realizar el experimento en una misma computadora.
  - Ejecutar el código mínimo 5 veces y luego obtener el promedio de tiempo de procesamiento y la desviación estandar.
  - Medir el tiempo para diferentes tamaños de entrada, por ejemplo si implementa algún algoritmo de ordenamiento, realizar los experimentos para listas con tamaño de: 100, 1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000, 9000, 10000, 20000, 30000, 40000 y 50000 elementos.
  - Una vez se obtenga el tiempo de procesamiento para diferentes tamaños de entrada y para los tres lenguajes generar gráficos comparativos usando Matplotlib.

---

## 2. Análisis de grafos (graphs-analysis)

#### 2.1 Competencias:

- Implementar, analizar y comparar algoritmos eficientes para la solución de problemas computacionales.
- Implementa estructuras de datos adecuadas.

#### 2.2 Material:

- Formato del informe:
  - LaTex
- Lenguajes de programación:
  - Golang (Go)
  - JavaScript

#### 2.3 Asignación:

- Implementar una solución al problema del agente viajero. Puede utilizar heurísticas, estas tienen un costo aceptable; sin embargo, no llegan a la solución óptima. Algunas heurísticas que puede utilizar son:
  - Algoritmos genéticos.
  - Busqueda de cuervos.
  - Busquedas locales.
  - Enjambre de partículas.

---

## 3. KD-Tree & KNN Classifier

#### 3.1 Competencias:

- Implementar, analizar el algoritmo KNN Classifier
- Implementa estructuras de datos adecuadas.

#### 3.2 Material:

- Formato del informe:
  - LaTex
- Formato de las diapositivas:
  - LaTex
- Lenguajes de programación:
  - Python

#### 3.3 Asignación:

La estructura KD-Tree es una estructura multidimensional de k dimensiones. Esta permite implementar busquedas por similitud como K Nearest Neighbor o Closest point. Adicionalmente, se puede usar esta estructura como un clasificador. Se debe implementar este clasificador en el tema de su preferencia. A continuación detallamos el algoritmo:

```
------------------------------------------------------------
Algorithm 1: KNN Classifier
------------------------------------------------------------
  Input: X: training data; y: object to be classified.
  Output: Classification for y.
  Extract features of each sample;
  Build KD-Tree;
  Select KNN of y in X;
  Class(y) ← max of classes ( k closest objects );
------------------------------------------------------------
```

Usted es libre de escoger el descriptor. Este descriptor es un método que toma como entrada una muestra de la base de datos y retorna un vector de carácterísticas, luego este vector representa un punto en el KD-Tree.
