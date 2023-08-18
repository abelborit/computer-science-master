import json

# Lee el archivo de texto
with open('Assets\generateNodesGraphBL\\nodesGraph5120p.txt', 'r') as file:
    contenido = file.read()

# Convierte el contenido a un diccionario
diccionario = eval(contenido)

# Convierte el diccionario a formato JSON
json_data = json.dumps(diccionario, indent=4)

# Escribe el JSON en un archivo de texto
with open('Assets\generateNodesGraphBL\\nodeGraph5120.json', 'w') as json_file:
    json_file.write(json_data)

print("Archivo JSON creado exitosamente.")


