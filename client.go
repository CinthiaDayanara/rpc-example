package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	// Establecer una conexión con el servidor RPC
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Error al conectar con el servidor:", err)
		return
	}
	defer client.Close()

	// Llamar al método SayHello del servidor
	var reply string
	err = client.Call("Greeter.SayHello", "Mundo", &reply)
	if err != nil {
		fmt.Println("Error al llamar al método:", err)
		return
	}

	// Mostrar la respuesta del servidor
	fmt.Println(reply) // Esto debería mostrar: "Hola, Mundo!"
}
