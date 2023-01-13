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

## Mas documentación

- [Contexto](./contexto.md): Se describe el contexto del problema, información sobre el sector de la gastronomía y beneficios de nuestro sistema.
- [Personas](./personas.md)
- [Planificación](./planificacion.md)
- [Ejemplo Situaciones Actuales vs Ideales](./ejem_situaciones_actual_vs_ideal.md): Para ejemplificar los beneficios del sistema, se compara la organización actual de ciertos bares y restaurantes con la organización que este sistema propone. Se exponen los beneficios de ejecutar la organización que este sistema propone.
- [Test runner y librería de aserciones](./testing.md)
- [Selección y razonamiento de la imagen base del Dockerfile escogida](./seleccion_imagen.md)
- [Selección y razonamiento de la libreriá de logs escogida](./sistema_de_logs.md)
- [Explicación de la estrategia de inyección de dependencias escogida en Go](./inyeccion_de_dependencias.md)
