package main //Computadora B

import (
    "fmt"
    "net"
    "bufio"
    "os"
)

func main() {
    conn, err := net.Dial("tcp4", "IP") //Cambia IP por la IP de computadora A
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("[!]Conectado")
	go back(conn) // llama a la función back en una goroutine para manejar la comunicación con el cliente
	us(conn) // llama a la función return en una goroutine para manejar la comunicación con el cliente
}

func back(conn net.Conn) {
    scanner := bufio.NewScanner(conn)
    for scanner.Scan() {
        fmt.Println("You:", scanner.Text())
		if scanner.Text() == "" {
			fmt.Fprintln(conn, "Empty message")
		}
    }
}

func us(conn net.Conn) {
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        fmt.Println("Them:", scanner.Text())   // 1. muestra en tu pantalla
        fmt.Fprintln(conn, "Them:", scanner.Text())    // 2. manda por la red
    }
}