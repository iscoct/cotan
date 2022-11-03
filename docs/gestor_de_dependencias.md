# Gestor de Dependencias

Como ya se decidió en este issue: #18, el lenguaje de programación a usar será [Go](https://go.dev/).

Anteriormente, Go usaba varios gestores de dependencias que no formaban parte del standard del lenguaje, como [dep](https://golang.github.io/dep/docs/introduction.html) o [glide](https://github.com/Masterminds/glide), ([aquí](https://go.libhunt.com/go-modules-alternatives) se pueden ver los mejor valorados por la comunidad).
No obstante, como bien ponen en las descripciones de dichos gestores, desde la introducción de [los módulos de Go](https://go.dev/ref/mod), estos gestores de dependencias quedaron deprecados a favor de Go modules.

Por tanto, el gestor de dependencias que usaremos será el nativo de Go: [Go modules](https://go.dev/ref/mod).
