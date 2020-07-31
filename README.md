# Microservicio Cafe
Servicio para extraer información de una base de datos MongoDB escrita en Golang

[![Go Report Card](https://goreportcard.com/badge/github.com/pablosilvab/cafe-backend-golang)](https://goreportcard.com/report/github.com/pablosilvab/cafe-backend-golang)

## Comenzando 🚀

Para clonar el proyecto solo debes ejecutar el siguiente comando.

``` 
go get github.com/pablosilvab/cafe-backend-golang
```

### Pre-requisitos 📋

* Go instalado en tu computador.
* Docker instalado en tu computador.

### Instalación 🔧

Para instalar el proyecto y contribuir código debes:

1. Entrar al directorio del proyecto.

```
cd ~/go/src/github.com/pablosilvab/cafe-backend-golang 
```

2. Una vez en el directorio, descargar dependencias.

```
make go-download
```

Para ejecutar el proyecto de forma local, puedes realizarlo mediante un archivo binario o un contenedor de Docker.

#### Ejecución mediante archivo binario

1. En el directorio del proyecto, constuir archivo binario.

```
make go-build
```

2. Ejecutar archivo binario.

```
make go-run-build
```

#### Ejecución mediante Docker

1. En el directorio del proyecto, construir imagen Docker.

```
make docker-build
```

2. Ejecutar contenedor con la imagen construida. 

```
make docker-shell
```

3. Puedes ejecutar el proyecto ya sea con el archivo binario (tal como la sección anterior) o de la forma tradicional.

```
make go-run
```

## Despliegue 📦

Pendiente

* Secret

```
kubectl create secret generic cafe-backend-secret \
  --from-literal mongo="xxxXXxxx"
```

## Construido con 🛠️

* [Golang](https://golang.org) - Lenguaje de programación concurrente y compilado inspirado en la sintaxis de C.
* [Docker](https://www.docker.com) - Proyecto de código abierto que automatiza el despliegue de aplicaciones dentro de contenedores de software
* [MongoDB](https://www.mongodb.com) - MongoDB es un sistema de base de datos NoSQL, orientado a documentos y de código abierto. 
