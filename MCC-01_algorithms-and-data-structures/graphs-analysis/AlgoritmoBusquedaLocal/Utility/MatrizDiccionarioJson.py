import os
import json

input_path = os.path.join('Assets', 'generateRandomGraphBL', 'randomGraph5120.txt')
output_path = os.path.join('Assets', 'generateNodesGraphBL', 'nodeGraph5120.json')

# Leer la matriz desde el archivo de entrada
matriz_str = None
with open(input_path, 'r') as archivo_entrada:
    matriz_str = archivo_entrada.read()

# Convertir la matriz en una lista de listas
matriz = eval(matriz_str)

# Convertir la matriz en un diccionario
grafo_diccionario = {}
for i in range(len(matriz)):
    grafo_diccionario[i + 1] = {}
    for j in range(len(matriz[i])):
        if i != j and matriz[i][j] != 0:
            grafo_diccionario[i + 1][j + 1] = matriz[i][j]

# Guardar el diccionario en formato JSON en el archivo de salida
with open(output_path, 'w') as archivo_salida:
    json.dump(grafo_diccionario, archivo_salida, indent=4)

print("Archivo procesado y escrito exitosamente en formato JSON.")
