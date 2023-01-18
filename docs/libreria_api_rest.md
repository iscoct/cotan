# Librería API REST

En este documento se detallarán los criterios para escoger un framework u otro en Go y finalmente el framework escogido.

## Frameworks disponibles

Los frameworks más populares para crear una API REST en Go son:

- Gin: Framework minimalista que se enfoca en la velocidad y el rendimiento.
- Echo: Framework de alto rendimiento que admite una gran cantidad de características y configuraciones avanzadas.
- Revel: Framework completo que proporciona todas las herramientas y funcionalidades necesarias para construir una aplicación web, incluyendo una interfaz de programación de aplicaciones (API) RESTful.
- Gorilla: Paquete de go que proporciona un conjunto de herramientas para escribir aplicaciones web más fácilmente, incluyendo una interfaz mux y un sistema de rutas.
- Beego: Framework web para Go que proporciona una gran cantidad de funciones para desarrollar aplicaciones web, incluyendo una interfaz de programación de aplicaciones (API) RESTful.

## Criterios para seleccionar una librería u otra

- Rendimiento: El rendimiento es un factor importante a considerar al elegir un framework. Se recomienda elegir un framework que sea rápido y eficiente en el uso de recursos.
- Facilidad de uso: Un framework fácil de usar y configurar puede ahorrar una gran cantidad de tiempo y esfuerzo. Es importante elegir un framework con una curva de aprendizaje baja y una documentación clara y fácil de seguir.
- Características y funcionalidades: Considera las características y funcionalidades que necesitas en tu aplicación. Algunos frameworks proporcionan un conjunto completo de herramientas, mientras que otros son más minimalistas.
- Comunidad y soporte: Una comunidad activa y un buen soporte pueden ser importantes para resolver problemas y obtener ayuda rápidamente.
- Escalabilidad: Si tu aplicación tiene un gran volumen de tráfico o preves un crecimiento significativo en el futuro, es importante elegir un framework escalable.

## Comparativa según los criterios seleccionados

### Rendimiento

En términos de rendimiento, Gin se destaca por ser un framework minimalista que se enfoca en la velocidad y el rendimiento.
Es rápido y eficiente en el uso de recursos, lo que lo hace una excelente opción para proyectos que requieren alta velocidad y rendimiento.
Echo también es un framework de alto rendimiento, pero se enfoca más en las características y configuraciones avanzadas.

### Facilidad de uso

En cuanto a la facilidad de uso, Gin es fácil de usar y configurar, con una curva de aprendizaje baja y una documentación clara y fácil de seguir.
Por otro lado, Echo y Revel pueden ser un poco más difíciles de usar y configurar debido a la gran cantidad de funciones y configuraciones avanzadas que ofrecen.
Gorilla es también fácil de usar y configurar, pero carece de algunas características avanzadas.

### Funcionalidades disponibles

En cuanto a las características y funcionalidades, Echo y Revel ofrecen un conjunto completo de herramientas, mientras que Gin, Gorilla y Beego son más minimalistas. Sin embargo, si necesitas una gran cantidad de funciones, Echo y Revel podrían ser una mejor opción.

### Comunidad y soporte

En cuanto a la comunidad y soporte, todos los frameworks mencionados tienen una comunidad activa y buen soporte, por lo que esto no debería ser un factor determinante a la hora de elegir un framework.

### Escalabilidad

En cuanto a escalabilidad, Echo y Revel son escalables, mientras que Gin y Gorilla son más adecuados para proyectos de menor escala.

## Conclusión

Teniendo en cuenta estos criterios, en mi caso, Gin es la opción que escogeré para crear mi API REST debido a su rendimiento, así como su facilidad de uso y configuración.
Sin embargo, si se requieran características y funcionalidades avanzadas y escalabilidad, Echo y Revel podrían ser una mejor opción.
