package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strings"
)

const localAddr = "localhost:8001" // su propia IP aquí

const (
	cnum = iota // iota genera valores en secuencia y se reinicia en cada bloque const
	opa
	opb
)

/*
1: Registro de Bloque
2: Eliminación de Bloque
3: Respuesta de Bloque
*/

type tmsg struct {
	Code int
	Port string
	Data string
}
type receiveData struct {
	Code int
	Data string
}

var addrs = []string{"127.0.0.1:8005", "127.0.0.1:8006"}

var chInfo chan map[string]int

func main() {
	chInfo = make(chan map[string]int)

	go func() { chInfo <- map[string]int{} }()

	go server()

	for {
		//fmt.Print("Your option: ")
		/*fmt.Scanf("%d\n", &op)
		msg := tmsg{3, "", "PRUEBA"}
		for _, addr := range addrs {
			send(addr, msg)
		}*/
	}
}
func server() {
	if ln, err := net.Listen("tcp", localAddr); err != nil {
		log.Panicln("Can't start listener on ", localAddr)
	} else {
		defer ln.Close()
		fmt.Println("Listeing on ", localAddr)
		for {
			if conn, err := ln.Accept(); err != nil {
				log.Println("Can't accept ", conn.RemoteAddr())
			} else {
				go handle(conn)
			}
		}
	}
}
func handle(conn net.Conn) {
	defer conn.Close()
	dec := json.NewDecoder(conn)

	var msg tmsg
	//fmt.Println("Dec", dec)
	if err := dec.Decode(&msg); err != nil {
		log.Println("Can't decode from ", conn.RemoteAddr())

		fmt.Println("Code: Recibe Orden2")
		msg := receiveData{1, "PRUEBA"}
		for _, addr := range addrs {
			sendData(addr, msg)
		}

	} else {
		fmt.Println("LLega de mensaje:")
		remote := conn.RemoteAddr().String()
		s := strings.Split(remote, ":")
		youAddress := s[0] + ":" + msg.Port
		fmt.Println("Receive:", youAddress)
		switch msg.Code {
		case 1:
			fmt.Println("Code: Registro")
			registration(conn, youAddress, msg)
			//fmt.Println(list)
			//fmt.Println(chaddrs)
			//go func() { chaddrs <- list }()
		case 2:
			fmt.Println("Code: Eliminación")
		case 3:
			fmt.Println("Code: Recibe Orden")
			msg := receiveData{1, "PRUEBA"}
			for _, addr := range addrs {
				sendData(addr, msg)
			}
		case 4:
			fmt.Println("Code: Concensus")
			fmt.Println(msg)
			concensus(conn, youAddress, msg)
		}
	}
}
func registration(conn net.Conn, remote string, msg tmsg) {
	fmt.Println("Registration1", addrs)
	addrs := append(addrs, remote)
	fmt.Println("bbbb", addrs)
	fmt.Println("Registration2")
}
func concensus(conn net.Conn, remote string, msg tmsg) {
	info := <-chInfo
	info[remote] = 1
	if len(info) == len(addrs) {
		ca, cb := 0, 0
		for _, op := range info {
			if op == opa {
				ca++
			} else {
				cb++
			}
		}
		if ca > cb {
			fmt.Println("GO A!")
		} else {
			fmt.Println("GO B!")
		}
		info = map[string]int{}
	}
	go func() { chInfo <- info }()
}
func sendRegister(remoteAddr string, msg tmsg) {
	if conn, err := net.Dial("tcp", remoteAddr); err != nil {
		log.Println("Can't dial", remoteAddr)
	} else {
		defer conn.Close()
		fmt.Println("Sending to", remoteAddr)
		enc := json.NewEncoder(conn)
		enc.Encode(msg)
	}
}
func sendData(remoteAddr string, msg receiveData) {
	if conn, err := net.Dial("tcp", remoteAddr); err != nil {
		log.Println("Can't dial", remoteAddr)
	} else {
		defer conn.Close()
		fmt.Println("Sending to", remoteAddr)
		enc := json.NewEncoder(conn)
		enc.Encode(msg)
	}
}
