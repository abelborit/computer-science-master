# Obtenemos el arreglo de elementos
def quickSort(arr):
    # Validamos si el arreglo tiene más de un elemento
    if len(arr) <= 1:
        # Si solo tiene un elemento retornanos el arreglo porque ya estaría ordenado
        return arr

    # Definimos la variable pivote que en este caso sera el primer elemento (este nos servirá para poder hacer las comparaciones)
    pivot = arr[0]

    # Definimos un arreglo para elementos hacia la izquierda del pivote (los elementos menores)
    left = []
    # Definimos un arreglo para elementos hacia la derecha del pivote (los elementos iguales o mayores)
    right = []

    # Iteramos el arreglo comparando los elementos con el pivote
    for element in arr[1:]:
        if element < pivot:
            # Si el elemento a comparar es menor que el pivote, lo añadimos al arrego de elementos hacia la izquierda
            left.append(element)  # Elementos más pequeños que el pivote
        else:
            # Si el elemento a comparar es mayor que el pivote, lo añadimos al arrego de elementos hacia la derecha
            right.append(element)  # Elementos más grandes que o iguales al pivote
    
    # Aplicamos recursivamente el Quick Sort a los sub-arrays izquierdo y derecho
    sortedLeft = quickSort(left)
    sortedRight = quickSort(right)

    # Concatenamos los sub-arrays ordenados con el pivote en el medio y retornamos el array completo ordenado
    return sortedLeft + [pivot] + sortedRight

# *******************************************************************
# Ejemplo de uso
# unsortedList = [7, 2, 1, 2, 10, 6, 8, 5, 3, 4]
# sortedList = quickSort(unsortedList)
# print(sortedList)  # Output: [1, 2, 2, 3, 4, 5, 6, 7, 8, 10]
