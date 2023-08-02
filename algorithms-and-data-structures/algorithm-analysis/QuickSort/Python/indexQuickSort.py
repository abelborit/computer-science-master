import os
import time
import statistics
from quickSortPython import quickSort

print("*****************")
print("QUICK SORT PYTHON")
print("*****************")

def readArrayFromFile(filePath):
    try:
        # Este código utilizará la función open para abrir el archivo y luego leerá su contenido
        with open(filePath, 'r') as file:
            data = file.read()
            # Los elementos se separan por comas en el archivo de texto, por lo que se utilizan split(",") para separar los elementos y int(num.strip()) para convertir cada elemento en un número entero. 
            arr = [int(num.strip()) for num in data.split(",")]
            return arr
    except Exception as error:
        # Si ocurre algún error al leer el archivo, se mostrará un mensaje de error y se devolverá una lista vacía
        print("Error reading file:", error)
        return []

def processAndSortArray(filePath):
    unsortedArray = readArrayFromFile(filePath)

    numRepetitions = 10
    totalExecutionTime = 0
    sortedArray = []

    # Almacenar los tiempos individuales de ejecución en una lista en milisegundos
    executionTimes = []

    for i in range(1, numRepetitions + 1):
        # print("se repite", i, "veces")

        # usar time.perf_counter(), que devuelve en segundos, en lugar de time.time() para medir el tiempo transcurrido, proporciona una mayor precisión en la medición del tiempo y puede ser más adecuado para tareas de medición de rendimiento
        startTimeSort = time.perf_counter() * 1000 # Convertir a milisegundos
        sortedArray = quickSort(unsortedArray)
        endTimeSort = time.perf_counter() * 1000 # Convertir a milisegundos

        executionTime = (endTimeSort - startTimeSort) 
        executionTimes.append(executionTime)

        totalExecutionTime += (endTimeSort - startTimeSort)

    averageExecutionTime = totalExecutionTime / numRepetitions

    # Calcular la desviación estándar en milisegundos
    stdDeviation = statistics.stdev(executionTimes)

    print("--------------------")
    print(file_path)
    print("--------------------")

    print("Arreglo no ordenado:", unsortedArray)
    print("")
    print("Arreglo ordenado:", sortedArray)
    print("")
    print("Tiempo promedio de ejecución:", averageExecutionTime, "milisegundos")
    print("Desviación estándar:", stdDeviation, "milisegundos")
    print("********************************************************************************")

# Directorio donde se encuentra el script de Python
currentDirectory = os.path.dirname(os.path.abspath(__file__))

# Constante para hacer un formateo de strings o concatener de una forma especial
fileName = "../../Assets/unsortedArrays/";

# Ruta relativa al archivo "unsortedArray100.txt" (os.path.join construirá la ruta correcta basada en la ubicación actual del script de Python y la ruta proporcionada para el archivo)
unsorted_array_file_paths = [
  # os.path.join(currentDirectory, f"{fileName}unsortedArray100.txt"),
  # os.path.join(currentDirectory, f"{fileName}unsortedArray1000.txt"),
  # os.path.join(currentDirectory, f"{fileName}unsortedArray2000.txt"),
  # os.path.join(currentDirectory, f"{fileName}unsortedArray3000.txt"),
  # os.path.join(currentDirectory, f"{fileName}unsortedArray4000.txt"),
  # os.path.join(currentDirectory, f"{fileName}unsortedArray5000.txt"),
  # os.path.join(currentDirectory, f"{fileName}unsortedArray6000.txt"),
  # os.path.join(currentDirectory, f"{fileName}unsortedArray7000.txt"),
  # os.path.join(currentDirectory, f"{fileName}unsortedArray8000.txt"),
  # os.path.join(currentDirectory, f"{fileName}unsortedArray9000.txt"),
  # os.path.join(currentDirectory, f"{fileName}unsortedArray10000.txt"),
  # os.path.join(currentDirectory, f"{fileName}unsortedArray20000.txt"),
  # os.path.join(currentDirectory, f"{fileName}unsortedArray30000.txt"),
  # os.path.join(currentDirectory, f"{fileName}unsortedArray40000.txt"),
  os.path.join(currentDirectory, f"{fileName}unsortedArray50000.txt"),
]

# Procesar y ordenar cada archivo
for file_path in unsorted_array_file_paths:
    processAndSortArray(file_path)
