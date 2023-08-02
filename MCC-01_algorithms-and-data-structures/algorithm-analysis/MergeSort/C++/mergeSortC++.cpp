#include <iostream>
#include <vector>

// Función auxiliar para combinar dos sublistas en orden
std::vector<int> merge(const std::vector<int>& left_arr, const std::vector<int>& right_arr) {
    std::vector<int> result;
    size_t left_index = 0, right_index = 0;

    // Comparar elementos de las dos sublistas y combinar en orden,
    // repite el proceso de combinación hasta que solo quede una lista ordenada
    while (left_index < left_arr.size() && right_index < right_arr.size()) {
        if (left_arr[left_index] < right_arr[right_index]) {
            result.push_back(left_arr[left_index]);
            left_index++;
        } else {
            result.push_back(right_arr[right_index]);
            right_index++;
        }
    }

    // Agregar los elementos restantes de las sublistas a 'result'
    while (left_index < left_arr.size()) {
        result.push_back(left_arr[left_index]);
        left_index++;
    }

    while (right_index < right_arr.size()) {
        result.push_back(right_arr[right_index]);
        right_index++;
    }

    return result;
}

// Función Merge Sort
std::vector<int> mergeSort(std::vector<int>& arr) {
    // Validamos si el arreglo tiene más de un elemento
    // Caso base: Si la lista tiene un elemento o está vacía, está ordenada.
    if (arr.size() <= 1) {
        // Si solo tiene un elemento, retornamos el arreglo porque ya estaría ordenado
        return arr;
    }

    // Encontrar el punto medio de la lista
    size_t middle = arr.size() / 2;

    // Dividir la lista en dos partes
    std::vector<int> left(arr.begin(), arr.begin() + middle);
    std::vector<int> right(arr.begin() + middle, arr.end());

    // Aplicar Merge Sort en las dos mitades y combinar los resultados
    return merge(mergeSort(left), mergeSort(right));
}

// *******************************************************************
// Ejemplo de uso
// int main() {
//     std::vector<int> unsorted_list = {38, 27, 43, 3, 9, 82, 10};
//     std::vector<int> sorted_list = mergeSort(unsorted_list);

//     // Imprimir el resultado ordenado
//     for (int num : sorted_list) {
//         std::cout << num << " ";
//     }
//     std::cout << std::endl;

//     return 0;
// }