# Creador de archivo .csv
# Las siglas CSV vienen del inglés "Comma Separated Values" y significan valores separados por comas. Dicho esto, un archivo CSV es cualquier archivo de texto en el cual los caracteres están separados por comas, haciendo una especie de tabla en filas y columnas. 
# Las columnas quedan definidas por cada punto y coma (;), mientras que cada fila se define mediante una línea adicional en el texto. De esta manera, se pueden crear archivos CSV con gran facilidad. Es por esto que los archivos .csv están asociados directamente a la creación de tablas de contenido.
# Los archivos CSV sirven para manejar una gran cantidad de datos en formato tabla, sin que ello conlleve sobrecoste computacional alguno. Es tremendamente sencillo generar una tabla a partir de un documento de texto, con tan solo delimitar cada celda requerida con un punto y coma (en el caso de Europa) o con una coma (en países de habla inglesa).
import csv
import os
import random

# Definir las variables
quantityData = 600
fileName = f"dataBaseElements{quantityData}.csv"
currentDirectory = os.path.dirname(os.path.abspath(__file__)) # Nombre del directorio actual
createInDirectory = os.path.join(currentDirectory, f"./Assets/{fileName}")

def generateData(quantityDataParam):
    data = []

    for _ in range(quantityDataParam):
        edad = random.randint(18, 60) # rango de edad entre 18 y 60 años
        credito = random.randint(50000, 600000) # rango de crédito entre 50 000 y 600 000 soles
        cumplio = random.randint(0, 1) # cumplió es 1 / no cumplió es 0

        data.append([edad, credito, cumplio])
    return data

def saveCSVFile(data, routeDirectory):
    with open(routeDirectory, mode='w', newline='') as CSV_File:
        writer = csv.writer(CSV_File)
        writer.writerow(['Edad', 'Credito', 'Cumplio'])
        writer.writerows(data)

dataGenerated = generateData(quantityData)
saveCSVFile(dataGenerated, createInDirectory)

print(f"Se generaron y guardaron {quantityData} datos en el archivo '{fileName}'.")