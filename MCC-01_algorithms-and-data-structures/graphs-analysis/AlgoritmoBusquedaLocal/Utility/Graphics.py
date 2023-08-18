import matplotlib.pyplot as plt

# Valores en el eje X Tamaños de entrada
x = [10, 20, 40, 80, 160, 320, 640, 1280, 2560, 5120]

#Tiempo de procesamiento (Segundos)
y = [0.0042656,0.0071241,0.0138356,0.0365711,0.0842718,0.2938134,0.8536143,2.7220933,8.9074015,25.8581051]

# Crear el gráfico
plt.plot(x, y, 'o-', label='Busqueda Local', color='red', linewidth=2)
plt.xticks(x)
plt.xticks(fontsize='x-small')

# Agregar títulos y etiquetas
plt.xlabel('Cantidad de nodos')
plt.ylabel('Tiempo de Procesamiento (s)')
plt.legend()

# Añadir rejilla para mejorar la legibilidad
plt.grid(True)

# Mostrar el gráfico
plt.tight_layout()
plt.show()