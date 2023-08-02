// para correr los programas g++ indexQuickSort.cpp quickSortC++.cpp -o programa
#include <iostream>
#include <fstream>
#include <vector>
#include <string>
#include <chrono>
#include <numeric>
#include <cmath>
#include "selectionSortC++.cpp"

using namespace std;

// Declaración de funciones
vector<int> readArrayFromFile(const string &filePath);
void processAndSortArray(const string &filePath);

int main()
{
  cout << "******************" << endl;
  cout << "SELECTION SORT C++" << endl;
  cout << "******************" << endl;

  // Ruta relativa al archivo "unsortedArray100.txt"
  vector<string> unsortedArrayFilePaths = {
      "../../Assets/unsortedArrays/unsortedArray100.txt",
      // "../../Assets/unsortedArrays/unsortedArray1000.txt",
      // "../../Assets/unsortedArrays/unsortedArray2000.txt",
      // "../../Assets/unsortedArrays/unsortedArray3000.txt",
      // "../../Assets/unsortedArrays/unsortedArray4000.txt",
      // "../../Assets/unsortedArrays/unsortedArray5000.txt",
      // "../../Assets/unsortedArrays/unsortedArray6000.txt",
      // "../../Assets/unsortedArrays/unsortedArray7000.txt",
      // "../../Assets/unsortedArrays/unsortedArray8000.txt",
      // "../../Assets/unsortedArrays/unsortedArray9000.txt",
      // "../../Assets/unsortedArrays/unsortedArray10000.txt",
      // "../../Assets/unsortedArrays/unsortedArray20000.txt",
      // "../../Assets/unsortedArrays/unsortedArray30000.txt",
      // "../../Assets/unsortedArrays/unsortedArray40000.txt",
      // "../../Assets/unsortedArrays/unsortedArray50000.txt",
  };

  // Procesar y ordenar cada archivo
  for (const auto &filePath : unsortedArrayFilePaths)
  {
    processAndSortArray(filePath);
  }

  return 0;
}

vector<int> readArrayFromFile(const string &filePath)
{
  vector<int> arr;
  // Abrir y leer el archivo especificado por filePath. Una clase para la lectura de archivos
  ifstream file(filePath);

  if (file.is_open())
  {
    string line;

    // La función getline lee una línea completa del archivo hasta que encuentra el delimitador ',' (coma). Cada vez que se lee una línea, el contenido se convierte a un entero utilizando stoi(line) y luego se agrega al final del vector arr utilizando arr.push_back(stoi(line))
    while (getline(file, line, ','))
    {
      arr.push_back(stoi(line));
    }

    // Una vez que se han leído todas las líneas del archivo y se han agregado al vector arr, se cierra el archivo usando file.close()
    file.close();
  }

  // Finalmente, la función devuelve el vector arr que contiene los números leídos desde el archivo
  return arr;
}

void processAndSortArray(const string &filePath)
{
  vector<int> unsortedArray = readArrayFromFile(filePath);

  int numRepetitions = 10;
  double totalExecutionTime = 0;
  vector<int> sortedArray;

  // Almacenar los tiempos individuales de ejecución en una lista en milisegundos
  vector<double> executionTimes;

  for (int i = 1; i <= numRepetitions; i++)
  {
    int n = unsortedArray.size();

    // Iniciar el temporizador
    auto startTimeSort = chrono::high_resolution_clock::now();

    // Ordenar el arreglo usando selectionSort desde selectionSort.cpp
    sortedArray = selectionSort(unsortedArray);

    // Detener el temporizador
    auto endTimeSort = chrono::high_resolution_clock::now();

    // Calcular el tiempo transcurrido en milisegundos
    double executionTime = chrono::duration<double, milli>(endTimeSort - startTimeSort).count();
    executionTimes.push_back(executionTime);
    totalExecutionTime += executionTime;
  }

  double averageExecutionTime = totalExecutionTime / numRepetitions;

  // Calcular la desviación estándar en milisegundos
  double sumOfSquares = 0.0;
  // iterar a través de una colección de elementos
  for (double time : executionTimes)
  {
    sumOfSquares += pow(time - averageExecutionTime, 2);
  }
  double stdDeviation = sqrt(sumOfSquares / numRepetitions);

  cout << "--------------------" << endl;
  cout << filePath << endl;
  cout << "--------------------" << endl;

  cout << "Arreglo no ordenado:";
  for (int num : unsortedArray)
  {
    cout << " " << num;
  }
  cout << endl
       << endl;

  cout << "Arreglo ordenado:";
  for (int num : sortedArray)
  {
    cout << " " << num;
  }
  cout << endl
       << endl;

  cout << "Tiempo promedio de ejecución: " << averageExecutionTime << " milisegundos" << endl;
  cout << "Desviación estándar: " << stdDeviation << " milisegundos" << endl;
  cout << "********************************************************************************" << endl;
}
