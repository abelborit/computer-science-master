# IMPLEMENTACIÓN PARA EL PROBLEMA DEL AGENTE VIAJERO

---

# 1. INTRODUCCIÓN

- ## Introducción al problema del agente viajero

  El problema del agente viajero (Traveling Salesman Problem, TSP) es un problema de optimización combinatoria (se trata de un problema de optimización que involucra una cantidad finita o numerable de soluciones posibles) en el campo de la teoría de la computación y la investigación de operaciones. Se trata de encontrar la ruta más corta posible que permita a un agente visitar un conjunto dado de ciudades y regresar a la ciudad de origen, pasando por cada ciudad exactamente una vez.

- ## Introducción a heurísticas

  Cuando un algoritmo usa una heurística, ya no necesita buscar de manera exhaustiva todas las soluciones posibles, y por tanto puede encontrar soluciones aproximadas mas rápido. Una heuristica es un atajo que sacrifica exactitud por rapidez.

- ## Introducción a metaheurísticas

  Los metaheurísticos son métodos aproximados diseñados para resolver problemas de optimización combinatoria, en los que los heurísticos clásicos no son efectivos.

---

# 2. PROBLEMA DEL AGENTE VIAJERO

- El objetivo es encontrar una forma de hacer su viaje más eficiente (en términos de la distancia total recorrida o del costo total).
- El problema se puede modelar mediante un grafo etiquetado (las aristas tienen distancias o costos asociados a ellas), en el cual buscamos el ciclo hamiltoneano más eficiente. En la Figura 1, se muestra esta representación.
- Para un grafo de n vértices, tenemos (n − 1)! posibles ciclos hamiltonianos.

Para resolver el problema del agente viajero, se utilizaron, como prueba y estructura para resolver el problema, los siguientes datos:

- Nodo Grand Rapids:

  - Desde Grand Rapids a Grand Rapids es 0
  - Desde Grand Rapids a Saginaw es 113
  - Desde Grand Rapids a Detroit es 147
  - Desde Grand Rapids a Toledo es 167
  - Desde Grand Rapids a Kalamazoo es 56

- Nodo Saginaw:

  - Desde Saginaw a Grand Rapids es 113
  - Desde Saginaw a Saginaw es 0
  - Desde Saginaw a Detroit es 98
  - Desde Saginaw a Toledo es 142
  - Desde Saginaw a Kalamazoo es 137

- Nodo Detroit:

  - Desde Detroit a Grand Rapids es 147
  - Desde Detroit a Saginaw es 98
  - Desde Detroit a Detroit es 0
  - Desde Detroit a Toledo es 58
  - Desde Detroit a Kalamazoo es 135

- Nodo Toledo:

  - Desde Toledo a Grand Rapids es 167
  - Desde Toledo a Saginaw es 142
  - Desde Toledo a Detroit es 58
  - Desde Toledo a Toledo es 0
  - Desde Toledo a Kalamazoo es 133

- Nodo Kalamazoo:

  - Desde Kalamazoo a Grand Rapids es 56
  - Desde Kalamazoo a Saginaw es 137
  - Desde Kalamazoo a Detroit es 135
  - Desde Kalamazoo a Toledo es 133
  - Desde Kalamazoo a Kalamazoo es 0

|              | Grand Rapids | Saginaw | Detroit | Toledo | Kalamazoo |
| ------------ | ------------ | ------- | ------- | ------ | --------- |
| Grand Rapids | 0            | 113     | 147     | 167    | 56        |
| Saginaw      | 113          | 0       | 98      | 142    | 137       |
| Detroit      | 147          | 98      | 0       | 58     | 135       |
| Toledo       | 167          | 142     | 58      | 0      | 133       |
| Kalamazoo    | 56           | 137     | 135     | 133    | 0         |

![Problema del Agente Viajero](/MCC-01_algorithms-and-data-structures/graphs-analysis/Assets/ProblemaAgenteViajero.jpg "Problema del Agente Viajero")

---

# 3. ALGORITMOS

- **ALGORITMO DEL ÁRBOL DE EXPANSIÓN MÍNIMA (MINIMUM SPANNING TREE, MST):** Se construye un árbol de expansión mínima a partir de las ciudades y luego se recorre para formar una ruta. El algoritmo encuentra un subconjunto de aristas que forman un árbol con todos los vértices, donde el peso total de todas las aristas en el árbol es el mínimo posible.

  Este algoritmo es ampliamente utilizado en teoría de grafos y optimización combinatoria para encontrar un subconjunto de aristas de un grafo conexo que conecta todos los vértices y tiene un peso total mínimo.

  El Algoritmo del Árbol de Expansión Mínima es mejor para grandes cantidades de datos:

  - **Optimización garantizada:** El MST garantiza una solución óptima para el problema del Agente Viajero si el gráfico (grafo completo con pesos en este caso) cumple con la propiedad del triángulo. Esta propiedad establece que para cualquier conjunto de tres ciudades, el camino más largo entre dos de ellas no es mayor que la suma de los otros dos caminos. Esto asegura que el MST proporcionará una solución cercana a la óptima para TSP.

  - **Enfoque heurístico:** Aunque no es necesariamente la solución más óptima en todos los casos, el MST se comporta bien en la práctica para muchas instancias de TSP, especialmente para casos con un gran número de ciudades.

  - **Escalabilidad:** A medida que el número de ciudades aumenta, la diferencia en la eficiencia entre los algoritmos se hace más evidente. El MST es más escalable en términos de tiempo de ejecución.

  ![Problema del Agente Viajero Aplicando Minimum Spanning Tree](/MCC-01_algorithms-and-data-structures/graphs-analysis/Matplotlib/Graphics/GraphicMST.png "Problema del Agente Viajero Aplicando Minimum Spanning Tree")

---

# 4. TABLAS COMPARATIVAS

- ### COMPARATIVA DEL TIEMPO DE EJECUCIÓN EN GO (milisegundos):

| Cantidad de Nodos | Minimum Spanning Tree | XXXXXXXXXX | XXXXXXXXXX |
| ----------------- | --------------------- | ---------- | ---------- |
| 10                | 0.0000                |            |            |
| 20                | 0.0000                |            |            |
| 40                | 0.0000                |            |            |
| 80                | 0.0000                |            |            |
| 160               | 0.0000                |            |            |
| 320               | 0.5024                |            |            |
| 640               | 0.6837                |            |            |
| 1280              | 4.0896                |            |            |
| 2560              | 12.5886               |            |            |
| 5120              | 33.6501               |            |            |

- ### GRÁFICA DEL TIEMPO PARA EL PROBLEMA DEL AGENTE VIAJERO EN GO (milisegundos):

![Tiempo Promedio para el Problema del Agente Viajero](/MCC-01_algorithms-and-data-structures/graphs-analysis/Matplotlib/Graphics/GraphicMST.png "Tiempo Promedio para el Problema del Agente Viajero")

---

# 5. CONCLUSIONES
