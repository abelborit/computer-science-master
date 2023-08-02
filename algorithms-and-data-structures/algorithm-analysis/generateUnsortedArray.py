import random

# Definir las variables
nameFile = "unsortedArray100"
quantityElements = 10
separator = ", "
minRangeNumber = 0
maxRangeNumber = 50

# Abrir creando un nuevo archivo .txt y poder escribir dentro de este
generateArrayUnsorted = open(f"Assets/unsortedArrays/{nameFile}.txt","w") 

# Según la cantidad de elementos que queremos tener agregará un nuevo elemento
for i in range(quantityElements):
    generateArrayUnsorted.write(f"{random.randint(minRangeNumber, maxRangeNumber)}{separator}")

# Leemos el archivo antes creado para quitarle al último elemento los caracteres de separación (coma y espacio)
generateArrayUnsorted = open(f"Assets/unsortedArrays/{nameFile}.txt","r")
elementsFile = generateArrayUnsorted.read()

# Con rstrip eliminaremos los caracteres del lado derecho de la cadena que se le asigna
elementsWithOutLastSpace = elementsFile.rstrip(elementsFile[-1]) # primero elimina el espacio
elementsWithOutLastComma = elementsWithOutLastSpace.rstrip(elementsWithOutLastSpace[-1]) # ahora elimina la coma

# Volvemos a escribir en el archivo los elementos leídos pero ahora estarán sin los caracteres de separación
generateArrayUnsorted = open(f"Assets/unsortedArrays/{nameFile}.txt","w") 
generateArrayUnsorted.write(elementsWithOutLastComma)

# Cerramos la interacción con el archivo para que el contenido sea finalmente escrito en el archivo y además que se pueda liberar el archivo para ser usado luego
generateArrayUnsorted.close()