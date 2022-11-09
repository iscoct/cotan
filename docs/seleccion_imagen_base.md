# Elección imagen

## Consideraciones

### Debe haberse creado no hace mucho tiempo

No es una buena elección coger imágenes muy antiguas cuando tenemos otras más nuevas.

Las razones principales suelen ser:

- Seguridad: Nuevas versiones o parches pueden suplir errores de seguridad encontrados en un momento concreto.
- Rendimiento: Esto no siempre sucede pero el hecho de que sea más nuevo suele hacer que sea más rápido, por ejemplo, en nuestro caso en Go, esto puede ser el compilador.
- Estándard: Nuevas funcionalidades del lenguajes pueden ser útiles para seguir mejores maneras de desarrollar, y tener una versión avanzada nos permitirá sacar provecho de esas nuevas funcionalidades.

Sobre el rendimiento hay que tener cuidado porque imágenes de versiones más nuevas hace que muchas veces el tamaño de sus imágenes sea mayor.

Véase las de [Ubuntu](https://hub.docker.com/_/ubuntu) para corrobarlo.

### Preferiblmente debe tener Go ya instalado

Las imágenes de docker están diseñadas para ofrecernos lo mínimo imprescindible para poder hacer uso del software que queremos con el mínimo número de dependencias posibles.

Por tanto, es muy probable que una imagen oficial del lenguaje ofrezca un menor número de dependencia, y por tanto ocupar menos los contenedores e imágenes que utilicemos.

De todas formas, esto no siempre puede que sea así y deberemos de estudiar cuán bien están conseguidas las imágenes oficiales.

### Debe de ser lo más ligera posible

Cuanta más ligera, más portable será puesto que algunos sistemas pueden tener tamaños de memoria o procesamiento muy limitados.

Además, cuanto más ligera sea, más rápido arrancarán los contenedores que partan de dichas imágenes, lo cual es importante.

### La imagen base debe de ser estable

Aunque nos podría valer en nuestro caso porque estamos testeando, es preferible no escoger imágenes en fase alpha o beta dado que pueden contener errores, cosa que aunque también puede pasarle a una imagen estable, se supone que está en una fase que debería tener menos.

### La imagen tiene una entidad que la soporta

En caso de que se encuente un problema en la imagen, de seguridad o cualquier tipo, estaría bien que hubiera una entidad, organización o empresa, que resuelva el problema y que podamos actualizar la imagen fácilmente.

## Imágenes candidatas

Las imágenes candidatas son las que forman parte de [Golang](https://hub.docker.com/_/golang), [Ubuntu](https://hub.docker.com/_/ubuntu) o [Alpine](https://hub.docker.com/_/alpine).

### Alpine

Es imagen que ocupa sólo 5 MB en tamaño.

Los problemas de escoger Alpine serían:

- Que no es tan sencilla de utilizar como Ubuntu.
- Que deberíamos de instalar y configurar Go nosotros mismos.

### Ubuntu

Los problemas de Ubuntu serían:

- Más grande la imagen de lo necesario.
- Deberíamos de instalar y configurar Go nosotros mismo.

### Golang

En mi caso, para paliar los problemas de Alpine y Ubuntu, utilizaré [Golang 1.19 con Alpine 3.16](https://github.com/docker-library/golang/blob/30403f1c144bf7773508cfbab5de09ecf4dbddf9/1.19/alpine3.16/Dockerfile) dado que tiene lo mejor de ambos mundos:

- Tiene ya instalado Go, hace que nuestro Dockerfile sea más fácil de mantener al tener menos líneas de código.
- La imagen suficientement pequeña como para que nos dé un buen rendimiento.

#### ¿Por qué escoger Go 1.19?

[Aquí] están las notas de la release de Go 1.19.

Algunas de las características más interesantes son:

- El compilador es 20% más rápido.
- Inclusión de tipo atómicos para actualizar objetos sequencialmente.
- Corrección de bugs, aunque esto fue desarrollado más desde [Go 1.18](https://tip.golang.org/doc/go1.18) hacia detrás.

Lógicamente, en cada versión habrán ido mejorando el lenguaje, compilar, ejecutor de test, etc. por lo que es importante escoger una imagen lo más nueva posible del lenguaje.

## Tests realizados

- Tamaño imagen alpine:3.16.2: 5.544 MB
- Tamaño imagen ubuntu:20.04: 72.786162 MB
- Tamaño imagen ubuntu:22.04: 77.792502 MB
- Tamaño imagen golang:1.19.3-bullseye: 992.331472 MB
- Tamaño imagen golang:1.19.3-buster: 934.62343 MB
- Tamaño imagen golang:1.19-alpine: 352.188475 MB
- Tamaño de mi imagen: 590.574365 MB

Es más, sólo con estos comandos:

```shell
FROM alpine:3.16.2
RUN apk add go
RUN apk add build-base
RUN go install github.com/go-task/task/v3/cmd/task@latest
```

La imagen resultante de ese dockerfile es: 627.354749 MB.

Como el tamaño de golang-1.19-alpine es menor que las demás, está en la última versión tanto de alpine y Go, se escoge esta.