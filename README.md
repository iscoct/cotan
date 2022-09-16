# Optimizador de Restaurantes

## Definición del Problema

### Definición del Problema General

Tiempo medio que se espera entre que se pide una comanda en un bar o restaurante y el tiempo en que se reciben las bebidas, platos o tapas y postres de dicha comanda.

### Definición del Problema en el que nos vamos a enfocar

Tiempo medio de comunicación del staff (camareros, pinches, chefs, ...) y priorización de pasos a seguir para la ejecución de las comandas (servir bebidas, dar postres, pasos intermedios de platos, ...) por parte del staff de cocina (pinches, chefs, ...).

### Problemas relacionados con los temas anteriores a resolver

1. Lógica:
    - Priorización de los pasos a realizar.
    - Priorización de comandas según las subpartes de éstas (bebidas, entrantes, primeros platos, ...).
    - Asignación de tareas.
    - Concurrencia entre todas las tareas, pasos y comandas.
2. Transmisión y almacenamiento de información:
    - Envío de las comandas.
    - Formato de visualización de las comandas.
    - Registro de la información de cada comanda: Tiempo medio de cada tarea, salida de cada parte de la comanda (bebida, entrante, primer plato, ...).

## Contexto

### Suposiciones

- La persona propietaria o administradora del software describe las bebidas, entrantes, platos, postres.
- La persona propietaria o administradora del software describe los pasos de cada plato.
- La persona propietaria o administradora del software define cuál es su personal y de qué se encarga cada uno.
- Los camareros envían las comandas con el software.
- El staff de cocina indican cuando comienzan y terminan cada paso.
