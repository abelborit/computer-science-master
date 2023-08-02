# Obtenemos el arreglo de elementos
def selectionSort(arr):
    # Obtener la longitud del arreglo
    n = len(arr)  

    # Iterar a través del arreglo
    for i in range(n - 1):
        # Encontrar el índice del elemento mínimo en la sublista no ordenada
        min_index = i  # Suponemos que el índice del elemento mínimo es el índice actual i

        # Bucle interno: Buscar el mínimo en la sublista no ordenada (desde i + 1 hasta el final del arreglo)
        for j in range(i + 1, n):
            if arr[j] < arr[min_index]:
                # Si encontramos un elemento menor al elemento en min_index
                min_index = j  # Actualizamos min_index al índice del nuevo elemento mínimo

        # Intercambiar el elemento mínimo con el elemento en la posición actual (i)
        if min_index != i:
            # Solo intercambiamos si el elemento mínimo no está en su posición actual
            arr[i], arr[min_index] = arr[min_index], arr[i]

    # Devolver el arreglo ordenado
    return arr  

# *******************************************************************
# Ejemplo de uso
# unsorted_list = [64, 34, 25, 12, 22, 11, 90]
# sorted_list = selectionSort(unsorted_list)
# print(sorted_list)  # Output: [11, 12, 22, 25, 34, 64, 90]
