package main //Computadora A

import (
	"fmt"
	"net" // importa el net para lidar com conexiones de rede
	"bufio" // importa el paquete de  bufio para leer datos de la conexión de forma eficiente
	"os"
)


func main() {
    listener, err := net.Listen("tcp4", ":6969") // listener es igual a net.Listen("tcp4", ":6969"), que crea un listener TCP4 en el puerto 6969
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Esperando conexión...")

    conn, err := listener.Accept() // conn es igual a listener.Accept(), que espera a que un cliente se conecte y devuelve una conexión
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("[!]Conexion establecida con MacBook / Areia", conn.RemoteAddr())
	fmt.Fprintln(conn, "Chat start here")

	go back(conn) // llama a la función back en una goroutine para manejar la comunicación con el cliente
	us(conn) // llama a la función return en una goroutine para manejar la comunicación con el cliente
}

func back(conn net.Conn) {
    scanner := bufio.NewScanner(conn)
    for scanner.Scan() {
        fmt.Println("Them:", scanner.Text())
		if scanner.Text() == "" {
			fmt.Fprintln(conn, "Empty message")
		}
    }
}

func us(conn net.Conn) {
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        fmt.Println("You:", scanner.Text())   // 1. muestra en tu pantalla
        fmt.Fprintln(conn, "You:", scanner.Text())    // 2. manda por la red
    }
}
