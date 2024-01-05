# GRPC Use

- Definiendo archivo .proto:

1. Definir la sintaxis (es decir la version del proto): syntax = "proto3";
2. Definir el go_package , cada lenguaje tiene su paquete y distinciones, este se necesita para golang
3. Definir mensajes de entrada y salida al metodo rpc
4. Definir los servicios a utilizar (servicio = stub)
5. Definir los rpc (remote procedure control) para montar el servidor

- Después de tener el archivo .proto:

1. Debemos instalar ya sea en la maquina o mediante alguna extensión para el lenguaje un compilador de 
protoc el cual se encargará de convertir el archivo .proto a metodos y mensajes del lenguaje de programación

2. En golang se utiliza el comando
` protoc --go_out=./   --go-grpc_out=./ crudSimulador.proto `

Esto generará los archivos grpc.pb.go en la ruta = --go-grpc_out=./ y pb.go = --go_out=./ los cuales son
la definición en codigo del mensaje y el servicio para el cliente y servidor

3. Posteriormente instalaremos la dependencia de google para poder utilizar los metodos generados
`go get "google.golang.org/grpc/"`

4. Debemos crear un servidor para poder instanciar un cliente

## Ejemplo de cliente:
En la libreria de google.golang.org/grpc/ podemos encontrar documentación sobre
anti patrones, cliente y servidor, este es un ejemplo de servidor donde siempre
se utilizará como punto de entrada grpc.Dial
[Ejemplo de cliente en grpc](https://github.com/grpc/grpc-go/blob/master/examples/helloworld/greeter_client/main.go)

## Ejemplo de servidor:
En la libreria de google.golang.org/grpc/ podemos encontrar documentación sobre el servidor, el helloword siempre será la forma mas sencilla de hacer algo, a continuación se encuentra el [ejemplo del servidor en grpc](https://github.com/grpc/grpc-go/blob/master/examples/helloworld/greeter_server/main.go)

