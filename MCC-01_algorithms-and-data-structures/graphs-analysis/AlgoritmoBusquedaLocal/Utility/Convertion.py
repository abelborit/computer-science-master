import os

input_path  = os.path.join('graphs-analysis', 'Assets', 'generateRandomGraph', 'randomGraph640.txt')
output_path = os.path.join('graphs-analysis', 'Results', 'Dictionaries.txt')

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

# Escribir el diccionario resultante en el archivo de salida
with open(output_path, 'a') as archivo_salida:
    archivo_salida.write(str(grafo_diccionario) +'\n\n')

print("Archivo procesado y escrito exitosamente.")
