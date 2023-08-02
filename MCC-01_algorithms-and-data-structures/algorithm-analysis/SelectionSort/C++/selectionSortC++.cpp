#include <iostream>
#include <vector>

using namespace std;

// Función de ordenamiento por selección (selection sort)
vector<int> selectionSort(vector<int> arr)
{
  // Obtener la longitud del arreglo
  int n = arr.size();

  // Iterar a través del arreglo
  for (int i = 0; i < n - 1; i++)
  {
    // Encontrar el índice del elemento mínimo en la sublista no ordenada
    int min_index = i; // Suponemos que el índice del elemento mínimo es el índice actual i

    // Bucle interno: Buscar el mínimo en la sublista no ordenada (desde i + 1 hasta el final del arreglo)
    for (int j = i + 1; j < n; j++)
    {
      if (arr[j] < arr[min_index])
      {
        // Si encontramos un elemento menor al elemento en min_index
        min_index = j; // Actualizamos min_index al índice del nuevo elemento mínimo
      }
    }

    // Intercambiar el elemento mínimo con el elemento en la posición actual (i)
    if (min_index != i)
    {
      // Solo intercambiamos si el elemento mínimo no está en su posición actual
      swap(arr[i], arr[min_index]);
    }
  }

  // Devolver el arreglo ordenado
  return arr;
}

// *******************************************************************
// Ejemplo de uso
// int main()
// {
//   vector<int> unsorted_list = {64, 34, 25, 12, 22, 11, 90};
//   vector<int> sorted_list = selectionSort(unsorted_list);

//   // Imprimir el arreglo ordenado
//   for (int num : sorted_list)
//   {
//     cout << num << " ";
//   }
//   cout << endl;

//   return 0;
// }
