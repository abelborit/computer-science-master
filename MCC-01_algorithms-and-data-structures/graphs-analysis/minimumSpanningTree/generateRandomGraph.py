# Importamos el módulo random, que nos permite generar números aleatorios
import random
import os

# Definir las variables
num_nodes = 10 # Número de nodos en el grafo
min_weight = 1 # Valor mínimo para los pesos de las aristas
max_weight = 600 # Valor máximo para los pesos de las aristas
fileNameGraph = f"randomGraph{num_nodes}" # Nombre del archivo que vamos a generar para los grafos
fileNameNodes = f"nodesGraph{num_nodes}" # Nombre del archivo que vamos a generar para los nodos
currentDirectory = os.path.dirname(os.path.abspath(__file__)) # Nombre del directorio actual
createInDirectoryGraph = os.path.join(currentDirectory, f"./Assets/generateRandomGraph/{fileNameGraph}.txt") # Nombre del directorio donde quiero que se creen los archivos para los grafos
createInDirectoryNode = os.path.join(currentDirectory, f"./Assets/generateNodesGraph/{fileNameNodes}.txt") # Nombre del directorio donde quiero que se creen los archivos para los nodos

# Creamos la matriz de adyacencia como una lista de listas. Inicialmente, todos los valores se establecen en cero
matriz_adyacencia = [[0 for _ in range(num_nodes)] for _ in range(num_nodes)]

# Llenar la matriz con pesos aleatorios
# Utilizamos dos bucles anidados para recorrer todas las combinaciones posibles de nodos. Si i es diferente de j, esto asegura que no estamos creando un lazo (arista que conecta un nodo consigo mismo). Generamos un peso aleatorio entre min_weight y max_weight para las aristas
for i in range(num_nodes):
    for j in range(num_nodes):
        if i != j:  # Evitar lazos
            peso = random.randint(min_weight, max_weight)
            matriz_adyacencia[i][j] = peso

# Imprimir los números de nodos como un arreglo
# Crea una cadena de caracteres que representa el arreglo. Hemos utilizado join() y map() para convertir los números a cadenas y unirlos con comas. Luego, agregamos los corchetes de apertura y cierre manualmente
# array_nodes = '[' + ', '.join(map(str, range(1, num_nodes + 1))) + ']'
array_nodes = ', '.join(map(str, range(1, num_nodes + 1)))

# Abrir creando un nuevo archivo .txt y poder escribir dentro de este
generateNodesGraph = open(createInDirectoryNode,"w")
generateNodesGraph.write(array_nodes)
# Cerramos la interacción con el archivo para que el contenido sea finalmente escrito en el archivo y además que se pueda liberar el archivo para ser usado luego
generateNodesGraph.close()

# Abrir creando un nuevo archivo .txt y poder escribir dentro de este
generateRandomGraph = open(createInDirectoryGraph,"w")
# Imprimir la matriz de adyacencia en forma de arreglos en el archivo .txt
# Iteramos a través de cada fila en la matriz de adyacencia. Convertimos los elementos de la fila en cadenas y los unimos utilizando ', '.join(). Agregamos corchetes al principio y al final de la cadena generada para formar un arreglo. Agregamos una coma al final de cada arreglo para cumplir con tu requerimiento de terminar con una coma. Imprimimos la cadena resultante que representa una fila de la matriz de adyacencia en forma de arreglo
for fila in matriz_adyacencia:
    fila_str = ', '.join(map(str, fila))
    generateRandomGraph.write(fila_str + "\n")

# Leemos el archivo antes creado para quitarle al último elemento los caracteres de separación (coma y espacio)
generateRandomGraph = open(createInDirectoryGraph,"r")
elementsFile = generateRandomGraph.read()

# Con rstrip eliminaremos los caracteres del lado derecho de la cadena que se le asigna
elementsWithOutLastSpace = elementsFile.rstrip(elementsFile[-1]) # elimina el espacio o último salto de linea

# Volvemos a escribir en el archivo los elementos leídos pero ahora estarán sin los caracteres de separación
generateRandomGraph = open(createInDirectoryGraph,"w") 
generateRandomGraph.write(elementsWithOutLastSpace)

# Cerramos la interacción con el archivo para que el contenido sea finalmente escrito en el archivo y además que se pueda liberar el archivo para ser usado luego
generateRandomGraph.close()



