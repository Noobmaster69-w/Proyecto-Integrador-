package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// servidorChat encapsula la configuración del servidor
type ServidorChat struct {
	puerto string
}

// iniciar arranca el servidor y escucha nuevas conexiones
func (s *ServidorChat) Iniciar() {
	listener, err := net.Listen("tcp", s.puerto)
	if err != nil {
		log.Fatalf("Error al iniciar el servidor en el puerto %s: %v\n", s.puerto, err)
	}
	// cerramos el puerto cuando el servidor se detenga
	defer listener.Close()

	fmt.Printf("Servidor de Chat Empresarial iniciado. Escuchando en el puerto %s...\n", s.puerto)

	// bucle infinito para aceptar multiples conexiones
	for {
		conexion, err := listener.Accept()
		if err != nil {
			log.Printf("Error al aceptar conexión: %v\n", err)
			continue
		}

		// goroutine: maneja al nuevo cliente de forma concurrente (en segundo plano)
		go s.manejarCliente(conexion)
	}
}

// la funcion manejarCliente procesa los mensajes entrantes de una conexión específica
func (s *ServidorChat) manejarCliente(conn net.Conn) {
	defer conn.Close()

	clienteID := conn.RemoteAddr().String()
	fmt.Printf("Nuevo cliente conectado desde: %s\n", clienteID)

	// enviar un mensaje de bienvenida al cliente
	conn.Write([]byte("Bienvenido al servidor del Chat Empresarial.\nEscribe tu mensaje: "))

	// leer continuamente lo que el cliente escribe
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		mensaje := scanner.Text()

		// mostrar en la consola del servidor lo que llegó
		fmt.Printf("[%s]: %s\n", clienteID, mensaje)

		// enviar una confirmacion de vuelta al cliente
		respuesta := fmt.Sprintf("Servidor recibió: %s\n", mensaje)
		conn.Write([]byte(respuesta))
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error leyendo del cliente %s: %v\n", clienteID, err)
	}

	fmt.Printf("Cliente desconectado: %s\n", clienteID)
}

func main() {
	// instanciar el servidor en el puerto 8080
	servidor := ServidorChat{
		puerto: ":8080",
	}

	// ejecutar
	servidor.Iniciar()
}
