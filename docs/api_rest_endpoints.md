# API REST Endpoints

Este documento indica los endpoints a crear en base a los recursos con sus verbos y sus respuestas en formato Swagger:

```yaml
swagger: '2.0'
info:
  version: '0.1'
  title: API
host: localhost:8000
basePath: /
schemes:
  - http
paths:
    /platos/{nombre}:
        get:
            summary: "Obtener información de un plato específico"
            parameters:
                - name: "nombre"
                in: "path"
                required: true
                type: "string"
            responses:
                200:
                    description: "Información del plato"
                    schema:
                        $ref: '#/definitions/Plato'
                404:
                    description: Plato no encontrado
        put:
            summary: "Actualizar información de un plato específico"
            parameters:
                - name: "nombre"
                  in: "path"
                  required: true
                  type: "string"
                - name: "tipo"
                  in: "body"
                  schema:
                    type: "string"
                    enum: ["entrante", "postre", "primero", "segundo"]
                - name: "pasos"
                  in: "body"
                  schema:
                    type: "array"
                    items:
                        $ref: '#/definitions/Paso'
            responses:
                201:
                    description: "Información del plato actualizada"
                404:
                    description: Plato no encontrado

        delete:
            summary: "Eliminar un plato específico"
            parameters:
                - name: "nombre"
                in: "path"
                required: true
                type: "string"
            responses:
                202:
                    description: "Plato eliminado"
                404:
                    description: Plato no encontrado
    /comandas:
        put:
            summary: "Agregar platos a una comanda"
            parameters:
                - name: "platos"
                  in: "body"
                  required: true
                  schema:
                    type: "array"
                    items:
                        type: "string"
            responses:
                201:
                    description: "Platos agregados a la comanda"
                404:
                    description: "Algún plato no se corresponde con los actuales"
        get:
            summary: "Obtener información de una comanda"
            responses:
                200:
                    description: "Información de la comanda"
                    schema:
                        type: "array"
                        items:
                            $ref: '#/definitions/Plato'

    /paso/fin/{id}:
        post:
            summary: Fin del paso que se especifica en el id
            operationId: finPaso
            responses:
                202:
                    description: Fin del paso
                404:
                    description: Paso no encontrado

    /comandas/ejecutandose:
        get:
            summary: Obtener información sobre platos y pasos
            description: Obtener información sobre los platos y pasos que se están ejecutando
            operationId: getComandasEjecutandose
            responses:
                200:
                    description: Información de las comandas
                    schema:
                        type: array
                        items:
                            type: object
                            properties:
                                $ref: '#definitions/PlatoEjecutandose'

    /instrumento/{nombre}:
        get:
            summary: Obtener información de un instrumento específico
            description: Devuelve la información de un instrumento dado su nombre
            operationId: getInstrumento
            parameters:
                - name: nombre
                  in: path
                  description: Nombre del instrumento a obtener
                  required: true
                  type: string
            responses:
                200:
                    description: Información del instrumento
                    schema:
                        $ref: '#/definitions/Instrumento'
                404:
                    description: Instrumento no encontrado
        put:
            summary: Actualizar información de un instrumento específico
            description: Actualiza la información de un instrumento dado su nombre
            operationId: updateInstrumento
            parameters:
                - name: nombre
                  in: path
                  description: Nombre del instrumento a actualizar
                  required: true
                  type: string
                - name: instrumento
                  in: body
                  description: Información actualizada del instrumento
                  required: true
                  schema:
                    $ref: '#/definitions/Instrumento'
            responses:
                201:
                    description: Instrumento actualizado exitosamente
                404:
                    description: Instrumento no encontrado
        delete:
            summary: Eliminar un instrumento específico
            description: Elimina un instrumento dado su nombre
            operationId: deleteInstrumento
            parameters:
                - name: nombre
                in: path
                description: Nombre del instrumento a eliminar
                required: true
                type: integer
            responses:
                202:
                    description: Instrumento eliminado exitosamente
                404:
                    description: Instrumento no encontrado
definitions:
    Instrumento:
        type: object
        properties:
            nombre:
                type: string
            tiempo_para_estar_libre:
                type: string
    Paso:
        type: "object"
        properties:
            nombre:
                type: "string"
            instrumento:
                type: "string"
            duracion:
                type: "integer"
            dependencias:
                type: "array"
                items:
                    type: "string"
    Plato:
        type: "object"
        properties:
            nombre:
                type: "string"
            tiempo_inicio:
                type: "string"
                format: "date-time"
            paso:
                $ref: '#/definitions/Paso'
    PasoEjecutandose:
        type: "object"
        properties:
            id:
                type: "integer"
            nombre:
                type: "string"
            tiempo_inicio:
                type: "date-time"
            duracion:
                type: "date-time"
    PlatoEjecutandose:
        type: "object"
        properties:
            nombre:
                type: "string"
            tiempo_inicio:
                type: "date-time"
            pasos:
                type: "array"
                items:
                    $ref: '#definitions/PasoEjecutandose'
```
