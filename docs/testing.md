# Task Runner

## Notes

## Segunda revisión

### Test runner

## Tercera revisión - Resumen del proyecto del tío de Go del anterior objetivo

- gocheck: Informes de errores útiles para ayudar a resolver problemas.
- go-testdeep: Permite probar fácilmente las rutas HTTP API, utilizando operadores flexibles y el ayudante tdhttp.

# Task Runner

En Go, el test runner más utilizado es el nativo de Go.

Hay otros como [el Test Runner de GoLand propietario de JetBrains](https://www.jetbrains.com/help/go/test-runner-tab.html) pero para utilizarlo necesitas GoLand.
También hay otros como [Github - exercism/go-test-runner](https://github.com/exercism/go-test-runner) aunque realmente utilizan el runner de Go.

La documentación de este runner se puede ver [aquí](https://github.com/exercism/go-test-runner), tanto su comando en CLI como el cómo busca este comando los tests en Go.

# Testing en Go

A continuación se dicen algunas peculiaridades de cómo se testea nativamente en Go (sin paquetes externos):

- El resultado de los test se hace con operadores de comparación, no con aserciones.
- Como Go no tiene excepciones, deberemos de hacer un diseño para controlar los errores devolviendo valores específicos, ya sea devolviendo un valor y error (Go permite devolver más de un valor para una función) o valores específicos para indicar el error. Haciendo uso de los valores que nos devuelven las funciones, comprobaremos dichos valores para comprobar el correcto funcionamiento de nuestro código.
- Go no tiene excepciones por lo que los errores se harán con comprobaciones de los datos que se devuelvan, es decir, habrá que diseñar y estudiar cómo se van a gestionar los errores en Go devolviendo errores o valores específicos.
- Para implementar mocks se usan muchos las interfaces de Go porque permiten crear tipos con las funciones que se requieren sólo y enviar dicho tipo a las funciones.
- Go nos permite ejecutar [benchmaks](https://pkg.go.dev/testing#hdr-Benchmarks) para medir el rendimiento de nuestro código.
- Go nativamente nos da la funcionalidad de utilizar la técnica [Fuzzing](https://pkg.go.dev/testing#hdr-Fuzzing), una técnica donde se llama a una función con valores random y reporta bugs no controlados por los tests unitarios.
- [Posibilidad de ejecutar subtests y subbenchmarks](https://pkg.go.dev/testing#hdr-Subtests_and_Sub_benchmarks)
- Nos permite la posibilidad de un programa test o benchmark de hacer extra setups o teardown antes o después de que se ejecute. Para ello, se llama a TestMain y este llamará a los tests, de esa manera se hace antes y después las cosas. Si TestMain vuelve, el wrapper pasará el resultado de m.Run a os.Exit. Más documentación se puede ver [aquí](https://pkg.go.dev/testing#hdr-Main)

Ahora, si bien con las funcionalidades que no da Go nativamente, podemos testear todo lo que necesitemos, hay librerías de aserciones y paquetes en Go que nos pueden hacer la vida más fácil.

## Librerías y paquetes en Go

Sea como sea que se haga, una buena librería o paquete debería de realizar de manera sencilla lo siguiente:
- Mocking: Esta técnica nos permite sustituir código, valores de objetos o funciones, para testear únicamente partes de nuestro código y no depender de terceros o ejecutar partes que no queremos. Por ejemplo, si parte del código hace llamadas a API, podemos sustituir la llamada y devolver valores, normalmente los que esperaríamos de dicha API.
- Tener soporte para ejecutar código antes de lanzar tests específicos o conjunto de tests (setup o teardown). Esto es así porque hay veces que necesitamos preparar datos o mocks antes de tests o eliminar datos después de tests, como utilizar cierta base de datos y eliminar los datos que hayamos introducido luego, por lo que es importante que dé soporte a esta funcionalidad.
- Hacer comprobaciones o aserciones para comprobar los valores de los resultados de las funciones.

A continuación veremos ciertos paquetes en Go que no dan funcionalidad para testear de forma sencilla nuestro código.
Destacar que todos los paquetes que se documentan son usados por el test runner nativo de Go.

### Testify

[Testify](https://github.com/stretchr/testify) es un conjunto de paquetes que nos permite realizar fáciles aserciones, mocking, y ejecutar código antes y después de conjuntos de tests (suite).

Los siguientes paquetes de Testify nos permiten realizar todo lo anterior:
- [assert package](https://github.com/stretchr/testify#assert-package): Nos permite hacer uso de aserciones,el cual se puede combinar con [require](https://github.com/stretchr/testify#require-package) para terminar los tests que fallen las aserciones.
- [mock](https://github.com/stretchr/testify#require-package) provee mecanismos para mockear objetos de manera sencilla.
- [suite](https://github.com/stretchr/testify#suite-package) es el paquete que nos permite ejecutar código antes y después de ciertos tests de manera sencilla.

### Gocheck

[Gocheck](https://labix.org/gocheck) es un paquete que nos permite todo lo anterior proporcionado por Testify (con otra nomenclatura), pero con más funcionalidad, por ejemplo:
- Benchmarks para lógice suite (conjunto de tests).
- Gestión de directorios temporales.
- Saltarse funciones de testeo.
- Más información en la ejecución de los tests.

Personalmente, sobre Gocheck, lo que me parece más interesante es la [gestión de directorios temporales](https://pkg.go.dev/gopkg.in/check.v1#C.MkDir).
Aún así, me parece que la sintaxis no es tan sencilla como Testify.

[Aquí](https://pkg.go.dev/gopkg.in/check.v1) la documentación de la API

### Goblin

[Goblin](https://github.com/franela/goblin) es un framework de testeo BDD al estilo [Mocka](https://mochajs.org/) que no require de dependencias adicionales.

Lo más interesante de este paquete es la manera en la que se testea, muy similar a como se hace en Jest, permitiéndonos describir el propósito de cada test que hacemos de manera sencilla con strings.

Lo malo de este paquete en comparación con otros paquetes es que no tiene tanta funcionalidad como suites o mocking.

### Decisión

Tras haber analizado algunos de los paquetes de testing más utilizados por Go, voy a utilizar Testify para mi proyecto, dado que es sencillo y provee la función que necesito.

No obstante, otra persona podría haber utilizado el testing nativo de Go o los otros paquetes mencionados.

## Recursos Adicionales

- [Linode - Golang Unit Testing](https://www.linode.com/docs/guides/golang-unit-testing/)
