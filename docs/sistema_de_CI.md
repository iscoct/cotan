# Sistema de Integración continua

## Criterios de selección

Como hemos indicado en [este issue](https://github.com/iscoct/cotan/issues/43), lo que vamos a tener en cuenta para la selección de la plataforma de CI será:

- Que sea gratis o tenga una capa gratuita: Lo cual es importante para proyectos públicos pequeños.
- Que tenga integración con GitHub y use Github Check: No sólo porque el profesor va a asegurarse de que se utilizan por la plataforma, sino, porque como es recomendable que sólo se hagan cambios una vez garantizado que se han pasado los tests, no se debería permitir que se haga ninguno hasta que las pipelines que ejecutan los tests den el visto bueno de que el proyecto está testeado y de que se puede mergear sin problemas.
- Que podamos guardar la plantilla en Github para poder modificarla desde ahí o desde la misma plataforma de CI.
- Que tenga lo necesario que nos permita testear Go y Docker (que ejecutemos docker con docker run).

## Versiones de Go testeadas

En este [página web](https://endoflife.date/go) se indica las versiones soportadas de Go y las que no.

Como se puede ver ahí, las únicas versiones estables y soportadas de Go ahora mismo son Go 1.19 y Go 1.18, así que éstas serán las que se vayan a testear.

## Sistemas estudiados

### Gitlab CI

Se ha probado a integrar Gitlab CI con Github porque Gitlab CI tiene una capa gratuita para los proyectos públicos, permite guardar la plantilla (.gitlab-ci.yml) y tiene todo lo necesario para testear con Go y Docker.

Lo que sucede es que para ejecutar Gitlab CI desde un código Github, hay que hacer un mirroring, véase [esta documentación](https://docs.gitlab.com/ee/ci/ci_cd_for_external_repos/github_integration.html) y esta funcionalidad es sólo para Gitlab Premium.

### Github Actions

Github actions sí tiene capa gratuita, se integra con el propio Github Check, permite guardar las plantillas (que estarán dentro de la carpeta .github/workflows) y tiene todo lo necesario para testear Go y Docker.

Por tanto, Github Actions es una de las opciones utilizadas como sistema de integración continua.

Para Github actions se buildea la imagen y se pushea cada vez que hay cambios en go.mod, go.sum o Dockerfile.

También testea desde ubuntu el proyecto instalando las dependencias y ejecutando Go en las versiones 1.19 y 1.18.

### Semaphore CI

Semaphore CI se ha configurado e incluido en el repositorio porque con los requisitos de permitirnos testear Go y Docker.

De hecho, semaphore CI hará el build de la imagen de docker nuestra y ejecutará el comando que ejecuta: task docker
También ejecutará Go en la versión 1.19.

El problema y por lo que no es realmente válido es que aunque salga abajo en los checks al lado de Merge Request, no aparece arriba en los Checks.
Se ha configurado con Github Apps pero no se ha conseguido que aparezca arriba.

### Drone.io

Estudiada también, se ha visto que no tiene integración con Github Checks por lo que descarta su selección.

### Azure Pipelines

Cumple con todos los requisitos, de hecho está configurado en el proyecto y nos testea Go en la vesión 1.19.3.

Dicho sea de paso que aunque ya no tiene capa gratuita directamente, sino que hay que pedirle permiso a Azure Pipelines, cosa que he hecho, explicándole la razón de la petición y el propósito del proyecto.
