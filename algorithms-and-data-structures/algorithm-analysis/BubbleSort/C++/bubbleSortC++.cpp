#include <iostream>
#include <vector>

using namespace std;

// Implementación de la función de ordenamiento burbuja
vector<int> bubbleSort(vector<int>& arr) {
    int n = arr.size();
    bool swapped = true;

    while (swapped) {
        swapped = false;

        for (int i = 0; i < n - 1; i++) {
            if (arr[i] > arr[i + 1]) {
                // Intercambio de elementos si están en el orden incorrecto
                swap(arr[i], arr[i + 1]);

                swapped = true;
            }
        }
    }

    return arr;
}

// *******************************************************************
// Ejemplo de uso
// int main() {
//     vector<int> unsorted_list = {64, 34, 25, 12, 22, 11, 90};
//     bubbleSort(unsorted_list);

//     // Imprimir la lista ordenada
//     for (int num : unsorted_list) {
//         cout << num << " ";
//     }
//     cout << endl;

//     return 0;
// }