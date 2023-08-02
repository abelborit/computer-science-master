# Obtenemos el arreglo de elementos
def bubbleSort(arr):
    # Se declara una variable n que almacenará la longitud del arreglo para evitar calcularla repetidamente
    n = len(arr)
    # Se declara una variable swapped para controlar si se realizaron cambios en el arreglo durante un pase completo del algoritmo. Si en algún pase no se realizaron intercambios, significa que el arreglo está ordenado y el algoritmo puede detenerse
    swapped = True

    # Se inicia un bucle while, que asegura que el algoritmo se ejecute al menos una vez y luego se repita mientras haya intercambios (es decir, swapped sea True)
    while swapped:
        # Se establece swapped como False al inicio de cada iteración para verificar si ocurren intercambios en el bucle for
        swapped = False

        # Se recorre la lista usando un bucle for desde el índice 0 hasta el penúltimo elemento (n - 1)
        for i in range(n - 1):
            # En cada iteración del bucle for, se compara el elemento actual arr[i] con el siguiente elemento arr[i + 1]
            if arr[i] > arr[i + 1]:
                # Si el elemento actual es mayor que el siguiente elemento, se realiza un intercambio de posiciones usando la asignación múltiple de Python

                # Intercambio de elementos si están en el orden incorrecto
                arr[i], arr[i + 1] = arr[i + 1], arr[i]

                # Se actualiza la variable swapped a True para indicar que se realizó un intercambio en esta iteración
                swapped = True

    # Si no ocurre ningún intercambio durante una pasada completa del bucle for, swapped seguirá siendo False, y el bucle while se detendrá, ya que la lista estará completamente ordenada
    
    # Finalmente, cuando el bucle while termine, se devolverá la lista ordenada
    return arr


# *******************************************************************
# Ejemplo de uso
# unsorted_list = [64, 34, 25, 12, 22, 11, 90]
# sorted_list = bubbleSort(unsorted_list)
# print(sorted_list)  # Output: [11, 12, 22, 25, 34, 64, 90]
