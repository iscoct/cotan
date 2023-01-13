# Inyección de dependencias en Go

En Go la inyección de dependencias se puede lograr mediante el uso de objetos globales dentro de un paquete específico.

La idea es que se creen objetos globales en un paquete específico y luego se importen en los demás paquetes para ser utilizados.
Al tener estos objetos globales en un solo lugar, se facilita su configuración y se evita tener que pasarlos como argumentos a través de todo el código.

Por ejemplo, si queremos utilizar una base de datos en varios puntos de nuestro código, podemos crear un objeto global en un paquete "db" que contenga la configuración de la base de datos y luego importarlo en los demás paquetes para utilizarlo.

```go
package db

import "github.com/jinzhu/gorm"

var DB *gorm.DB

func init() {
    var err error
    DB, err = gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")

    if err != nil {
        panic(err)
    }
}
```

En este ejemplo, se define un objeto global llamado DB en el paquete "db" que contiene una conexión a una base de datos MySQL.
La función init se encarga de inicializar la conexión.

Luego, en cualquier otro paquete de nuestra aplicación, podemos importar este paquete y utilizar el objeto DB para hacer operaciones en la base de datos.

```go
package main

import (
    "log"
    "github.com/ejemplo/db"
)

func main() {
    defer db.DB.Close()

    var user User
    db.DB.First(&user)

    log.Println(user)
}
```

En este caso, el objeto global DB es una dependencia y está siendo inyectado en la función main.

Esto hace que si creamos una interfaz que defina los métodos que queremos podemos utilizar el objeto sin necesidad de saber cuál es el tipo realmente que estamos usando.
