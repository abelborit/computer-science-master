#include <iostream>
#include <vector>

using namespace std;

// Función que implementa el algoritmo Quick Sort
vector<int> quickSort(vector<int> arr) {
    // Validamos si el vector tiene más de un elemento
    if (arr.size() <= 1) {
        // Si solo tiene un elemento, retornamos el vector porque ya estaría ordenado
        return arr;
    }

    // Definimos la variable pivote que en este caso será el primer elemento (este nos servirá para poder hacer las comparaciones)
    int pivot = arr[0];

    // Definimos un vector para elementos hacia la izquierda del pivote (los elementos menores)
    vector<int> left;
    // Definimos un vector para elementos hacia la derecha del pivote (los elementos iguales o mayores)
    vector<int> right;

    // Iteramos el vector comparando los elementos con el pivote
    for (size_t i = 1; i < arr.size(); ++i) {
        if (arr[i] < pivot) {
            // Si el elemento a comparar es menor que el pivote, lo añadimos al vector de elementos hacia la izquierda
            left.push_back(arr[i]);  // Elementos más pequeños que el pivote
        } else {
            // Si el elemento a comparar es mayor que el pivote, lo añadimos al vector de elementos hacia la derecha
            right.push_back(arr[i]);  // Elementos más grandes que o iguales al pivote
        }
    }

    // Aplicamos recursivamente el Quick Sort a los sub-vectores izquierdo y derecho
    vector<int> sortedLeft = quickSort(left);
    vector<int> sortedRight = quickSort(right);

    // Concatenamos los sub-vectores ordenados con el pivote en el medio y retornamos el vector completo ordenado
    vector<int> sortedVector;
    // La función reserve se utiliza para reservar espacio en la memoria para el vector sortedVector antes de agregar elementos. Al reservar espacio de antemano, se evitan posibles reasignaciones y copias innecesarias a medida que se agregan elementos, lo que puede mejorar el rendimiento. Se está calculando el tamaño total necesario para sortedVector, que será la suma de los tamaños de los vectores sortedLeft y sortedRight, más 1 para el elemento pivote
    sortedVector.reserve(sortedLeft.size() + sortedRight.size() + 1);
    sortedVector.insert(sortedVector.end(), sortedLeft.begin(), sortedLeft.end());
    sortedVector.push_back(pivot);
    sortedVector.insert(sortedVector.end(), sortedRight.begin(), sortedRight.end());

    return sortedVector;
}

// *******************************************************************
// Ejemplo de uso
// int main() {
//     vector<int> unsortedList = {7, 2, 1, 2, 10, 6, 8, 5, 3, 4};
//     vector<int> sortedList = quickSort(unsortedList);
    
//     // Imprimimos el vector ordenado
//     for (const int& num : sortedList) {
//         cout << num << " ";
//     }
//     cout << endl;

//     return 0;
// }
