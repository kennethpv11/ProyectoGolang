# Pasos para crear tu protobuf para cliente - servidor grpc

1. Se debe instalar de https://github.com/protocolbuffers/protobuf/releases
la release con la versión compatible de tu sistema operativo (tendrá sufijo win para windows )ejemplo:
protoc-24.4-win64.zip

2. Se debe añadir a las variables de entorno, en el path la ruta del paquete
ejemplo:
C:\Program Files\protoc-24.4-win64\bin

3. Se debe añadir en el archivo .proto la siguiente linea en el inicio:
option go_package = "CRUDSimulador/";
donde CRUDSimulador/ es la ruta del archivo
esto es debido a que protobuf necesita saber el tipo de paquete para poder crear el servidor

4. Se ejecuta el siguiente comando para crear los metodos del servidor para la comunicación
```CMD
protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative  ./methods.proto
```

Donde methods es donde va la ruta del proto

5. instalar 
```go 
go get google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 
```
```go 
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```
```go
go get google.golang.org/grpc
```

El descriptor proto te puede decir mas opciones de configuración
(https://github.com/protocolbuffers/protobuf/blob/main/src/google/protobuf/descriptor.proto)


# Analisis estatico
 
 - se debe instalar la herramienta para ejecutar los linters go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.55.2

 - Posteriormente se debe crear el archivo .golangci.yml para las configuraciones (sino usa la por defecto)

 - golangci-lint run. Con este comando se realiza el analisis estatico para validar los errores

 - https://go.dev/wiki/CodeReviewComments este es un enlace a como deben ser los comentarios

 - go install mvdan.cc/gofumpt@latest (gofumt es un erro tipico de indexado debido a que esta mal organizado el codigo) por ello se instala la herramienta y se ejecuta el comando

 Se usa el comango gofumpt -d <path + filename>.go