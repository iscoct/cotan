# Gestor de Tareas

A continuación comento algunos gestores de tareas que se consideran para el proyecto y algunos pros y contras para decidirnos por uno de ellos:

 - [Make](https://makefiletutorial.com/):
     - Herramienta de uso genérico, normalmente usada para ayudar a decidir que partes de un programa necesita ser recompilado.
     - Usada sobre todo en proyectos cuyo lenguaje de programación es C, C++, Scala o Fortran, aunque también se puede usar para simplificar tareas que requieran de diversos comandos o scripts.
     - Pros:
          - Permite gestionar dependencias entre tareas.
          - Permite gestionar código fuente de manera eficiente dado que se le puede pedir a la herrramienta que sólo ejecute ciertas tareas si el código fuente se cambió desde la última vez que se ejecutó de manera sencilla.
     - Contras:
          - Hay que aprenderse la sintaxis para escribir el Makefile (fichero donde indicamos las tareas creadas, además de cómo y cuándo ejecutarlas).
          - Como el objetivo es compilar aquello que sea necesario, lenguajes interpretados como Python, Javascript, etc. no necesitan tanto de Make.
- [Task](https://taskfile.dev/):
     - La funcionalidad es muy parecida a Make siendo su objetivo ser más simple y fácil de usar que Make.
     - Pros:
          - Dado que está escrito en Go, Task es sólo un fichero binario con ninguna dependencia.
          - La sintaxis para escribir el Taskfile.yml es YAML.
          - Es compatible en múltiples plataformas y fácil de instalar.
          - Permite ejecutar tareas sólo cuando algunos ficheros cambian como Make y permite declarar dependencias entre tareas de manera sencilla.
     - Cons:
          - Debes de saber YAML.
- [Realize](https://github.com/oxequa/realize): Aunque es una herramienta interesante, sobre todo el Live Reload (te permite recompilar proyectos cuando se modifican), está más pensada para controlar múltiples proyectos al mismo tiempo, lo cual para este proyecto en concreto, no será necesario.

Por tanto, en este proyecto se usará **Task** dado que ya se está usando Go y porque así podemos hacer build del código sólo cuando se hayan modificado los ficheros del proyecto.