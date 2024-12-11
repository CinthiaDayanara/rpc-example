package main

import (
	"fmt"
	"net"
	"net/rpc"
)

// Definir la estructura que implementa el servicio RPC
type Greeter struct{}

// Método que se va a exponer por RPC
func (g *Greeter) SayHello(name string, reply *string) error {
	*reply = fmt.Sprintf("Hola, %s!", name)
	return nil
}

func main() {
	// Crear una instancia de Greeter
	greeter := new(Greeter)

	// Registrar el servicio RPC
	err := rpc.Register(greeter)
	if err != nil {
		fmt.Println("Error al registrar el servicio:", err)
		return
	}

	// Escuchar en el puerto 1234
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("Error al escuchar en el puerto 1234:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Servidor RPC escuchando en puerto 1234")

	// Aceptar las conexiones entrantes
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error en la conexión:", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
