# Obtenemos el arreglo de elementos
def mergeSort(arr):
    # Validamos si el arreglo tiene más de un elemento
    # Caso base: Si la lista tiene un elemento o está vacía, está ordenada.
    if len(arr) <= 1:
        # Si solo tiene un elemento retornanos el arreglo porque ya estaría ordenado
        return arr

    # Encontrar el punto medio de la lista
    middle = len(arr) // 2

    # Dividir la lista en dos partes
    left = arr[:middle]
    right = arr[middle:]

    # Función auxiliar para combinar dos sublistas en orden
    def merge(left_arr, right_arr):
        result = []
        left_index, right_index = 0, 0

        # Comparar elementos de las dos sublistas y combinar en orden, repite el proceso de combinación hasta que solo quede una lista ordenada
        while left_index < len(left_arr) and right_index < len(right_arr):
            if left_arr[left_index] < right_arr[right_index]:
                result.append(left_arr[left_index])
                left_index += 1
            else:
                result.append(right_arr[right_index])
                right_index += 1

        result.extend(left_arr[left_index:])
        result.extend(right_arr[right_index:])
        return result

    # Aplicar Merge Sort en las dos mitades y combinar los resultados
    return merge(mergeSort(left), mergeSort(right))

# *******************************************************************
# Ejemplo de uso
# unsorted_list = [38, 27, 43, 3, 9, 82, 10]
# sorted_list = mergeSort(unsorted_list)
# print(sorted_list)  # Output: [3, 9, 10, 27, 38, 43, 82]

