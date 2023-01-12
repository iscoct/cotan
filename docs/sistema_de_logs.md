# Sistema de logs

A continuación indico los criterios que tendré en cuenta para escoger la librería de logs:

- Que podamos esperar un soporte continuado en el tiempo.
- Que nos permita especificar distintos tipos de niveles de impresión: logs, debug, warning y fail o error.
- Que nos permita imprimir o guardar los logs en diferentes lugares. Por ejemplo:
    - Imprimir por consola.
    - Guardar los logs en ficheros.
    - Enviar dichos ficheros a un servidor remoto.
    - ...

Algunas opciones son:

- [log](https://pkg.go.dev/log): Librería estándar de Go que proporciona una interfaz simple para la impresión de registros.
- [glog](https://pkg.go.dev/github.com/golang/glog): Librería de logs desarrollada por Google que proporciona características avanzadas, como la rotación de registros y el seguimiento automático de la ubicación de la llamada.

## Librería log

No tiene un gran conjunto de configuraciones, pero permite configurar lo siguiente:

- Nivel de registro: La librería log permite establecer el nivel de registro (por ejemplo, información, advertencia, error) para controlar qué mensajes de registro se imprimen.
- Salida de registro: La librería log imprime los registros a la salida estándar por defecto, pero también permite establecer un archivo de salida específico o una conexión de red para enviar los registros a través de ella.
- Formato de salida: La librería log utiliza un formato de salida predeterminado que incluye la fecha y hora, el nivel de registro y el mensaje de registro, pero no tiene la facilidad de cambiar el formato de salida
- Prefixo de registro: La librería log permite agregar un prefijo personalizado a los mensajes de registro, lo que puede ser útil para identificar fácilmente de qué parte de la aplicación provienen los registros.

## Librería glog

Permite todo lo anterior más la característica interesante de rotación automática de registros.

Esto significa que glog hace automáticamente que los registros antiguos se mueven a archivos de registros secundarios para evitar que el archivo principal crezca demasiado.

## Decisión

Dado los criterios anterios se escoge log porque:
- Podemos esperar soporte ya que pertenece al estándar del lenguaje.
- Puede tener diferentes niveles de verbosidad si creamos diferentes objetos logger y le añadimos un prefijo que indique el tipo.
- Permite configurar diferentes lugares donde poner los registros.
- Si le ponemos el prefijo de la máquina en la que esté, podemos implementar un sistema de logs distribuido imprimiendo ese formato en un sistema remoto (lugar donde se agrupen los logs).

## Comentarios

Al tener distintos niveles de impresión (log, info, warning y error) podemos especificar para distintas máquina o entornos cuántos logs queremos recibir.
Por ejemplo, si está todo testeado y se garantiza que el software está bastante bien, podemos imprimir warning y error sólo.
Mientras que si hemos encontrado un bug, a lo mejor nos interesa imprimir más información, hasta log o info.

La decisión deberá tomarla el equipo de desarrollo que use este modelo de varios niveles.
