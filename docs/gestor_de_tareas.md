# Gestor de Tareas

A continuación comento algunos gestores de tareas que se consideran para el proyecto y algunos pros y contras para decidirnos por uno de ellos:

- [Gulp](https://gulpjs.com/docs/en/getting-started/quick-start): Gestor de tareas que te permite [usar tu conocimiento de Javascript](https://gulpjs.com/docs/en/getting-started/javascript-and-gulpfiles) para escribir tareas automatizables que se repiten en tu proyecto y que terminará ejecutando con **gulp <nombre-tarea>**
     - Pros:
          - Si tienes un proyecto muy grande, [Gulp te permite dividir las tareas en múltiples ficheros](https://gulpjs.com/docs/en/getting-started/javascript-and-gulpfiles#splitting-a-gulpfile).
          - Al utilizar Javascript nos permite utilizar la potencia y flexibilidad del lenguaje ([Callbacks, Promesas, Observables](https://gulpjs.com/docs/en/getting-started/creating-tasks)], ...).
          - Tiene [tareas públicas y privadas](https://gulpjs.com/docs/en/getting-started/creating-tasks) que junto con la posibilidad de dividir tareas entre ficheros, nos permite reutilizar tareas entre proyectos.
    - Contras:
          - Debes saber utilizar Javascript, lo cual puede ser un poco más complejo dado la flexibilidad de la herramienta en comparación con otras.
          - Dado que las tareas las debes crear tú, el hacer debug puede no resultar ser sencillo.
          - No suele tener lógicas complejas integradas en la herramienta como ejecutar la compilación sólo si se han modificado algunos de los ficheros.
 - [npm scripts](https://docs.npmjs.com/cli/v8/using-npm/scripts): Te permite ejecutar scripts con npm haciendo **npm run <nombre-script>** u [otros comandos como **npm prepare <nombre-script>**, **npm install**, etc.](https://docs.npmjs.com/cli/v8/using-npm/scripts#life-cycle-scripts) en función del propósito de la tarea. [Aquí](https://docs.npmjs.com/cli/v8/using-npm/scripts#life-cycle-scripts) se pueden ver algunos ejemplos de cómo se escriben las tareas.
      - Pros:
          - Las tareas siguen un propósito concreto cada una de ellas.
          - Se suelen dividir las tareas propósitos muy concretos de las mismas.
          - Tiene tareas para cada parte del ciclo de las tareas.
     - Contras:
          - Normalmente está pensado para ejecutar tareas cortas dado que se escribe en JSON. De hecho, muchas veces se suele usar con [CMake](https://cmake.org/) si se quieren tareas más complejas.
          - No suele tener lógicas complejas integradas en la herramienta como ejecutar la compilación sólo si se han modificado algunos de los ficheros.
 - [CMake](https://cmake.org/): Herramienta de uso genérico, normalmente usada para gestionar código fuente creada en el año 2000.
      - Pros:
           - Permite gestionar dependencias entre tareas.
           - Permite gestionar código fuente de manera eficiente dado que se le puede pedir a la herrramienta que sólo ejecute ciertas tareas si el código fuente se cambió desde la última vez que se ejecutó de manera sencilla.
     - Contras:
          - A diferencia de Gulp con Javascript o npm con npm, nos tendremos que aprender la sintaxis propia de CMake.

Como en mi caso el lenguaje es compilado, considero que me **vendrá bien utilizar CMake** para hacer build del código sólo cuando se haya modificado la última vez el código binario.
Además, CMake no necesita de Javascript u otras dependencias, el número de software a instalar, será menor.

Destacar que hay más herramientas que gestionan tareas pero se considera suficientemente compleja para este proyecto CMake.
