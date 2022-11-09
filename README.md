# Comandos Importantes

## Hacer build

```go
go build
```

## Ejecutar tras haber hecho build

```bash
./<nombre-ejecutable> # By default: ./cotan
```

## Hacer Build y Ejecutar

```go
go run .
```

## Comprobación de que el programa compila

```bash
task check
```

## Testear

Sin coverage:

```bash
task test
```

```bash
task coverage
```

## Docker

Si se prefiere, se puede trabajar directamente con [docker](https://www.docker.com/).

Para ello, se ha creado un [Dockerfile](./Dockerfile) cuya ejecución hace los tests de todo el código del proyecto.

También está automatizado con task, así que para poder ejecutar el testeo desde docker, sólo tienes que tener instalado [task](https://taskfile.dev/) y docker.

Si quieres lanzarlo sin Task, puedes ejecutar el siguiente comando:

```bash
docker build --tag iscoct/cotan:latest . && docker run iscoct/cotan
```

Si tienes instalado Task:

```bash
task docker
```

O dado que se puede encontrar el contenedor [aquí](https://hub.docker.com/repository/docker/iscoct/cotan), se puede hacer directamente:

```bash
docker run -u 1001 -t -v `pwd`:/app/test iscoct/cotan
```

## Más documentación

- [Resto de documentación](./docs/README.md)
- [Instalar Go](https://go.dev/doc/install)
- [Instalar Task - Task Runner](https://taskfile.dev/installation/)
