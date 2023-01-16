# Sistema de Configuración

Las herramientas de configuración de aplicaciones permiten almacenar y acceder a diferentes configuraciones de una aplicación de manera organizada y segura.
Estas configuraciones pueden incluir valores como la dirección de servidor, credenciales de base de datos, opciones de aplicación, etc.

Al utilizar una biblioteca de configuración, se asegura que las configuraciones de la aplicación están separadas del código y se pueden modificar fácilmente sin necesidad de recompilar la aplicación
Además, estas librerías permiten cargar configuraciones desde diferentes fuentes, como archivos, variables de entorno, bases de datos, etc. lo que permite tener diferentes configuraciones para diferentes entornos.

Para nuestra aplicación se han estudiado las diferentes bibliotecas de Go.

## Godotenv

godotenv es una librería simple y fácil de usar que se enfoca en la carga de variables de entorno desde archivos .env.
Además esta librería permite especificar rutas de archivo personalizadas y cargar múltiples archivos .env.
Sin embargo, no soporta otras fuentes de configuración además de los archivos .env.

## Viper
viper es una librería más potente y flexible, que permite la carga de variables de entorno desde diferentes fuentes, como archivos .env, archivos de configuración JSON, YAML o TOML, variables de entorno del sistema, etc.
Además, viene con funciones adicionales para acceder a las variables de configuración, como la posibilidad de acceder a las variables de configuración utilizando notación de punto, lo que facilita su lectura y acceso.
Sin embargo, es un poco más compleja de configurar y utilizar que godotenv.

## etcd/clientv3

etcd es un sistema de almacenamiento de clave-valor distribuido que se utiliza para mantener la configuración y el estado de un sistema distribuido.

Para utilizar etcd desde Go, se puede utilizar la biblioteca oficial de clientes de etcd llamada "etcd/clientv3". Esta biblioteca proporciona una interfaz fácil de usar para interactuar con un cluster de etcd utilizando el protocolo de comunicación gRPC.

La biblioteca clientv3 también proporciona una variedad de métodos para realizar operaciones avanzadas, como transacciones, seguimiento de cambios y más.

# Decisión

En nuestro caso, vamos a optar por la solución más simple, utilizar la librería godotenv, aunque vamos a abstraer la lectura de configuraciones a un paquete específico, para en caso de que necesitáramos otra forma de acceder a nuestra configuración, la aplicación no cambie.
