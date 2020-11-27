package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	api "./api"
)

const centralAddr = "localhost:8001" // Nodo Central IP aquí
const myPort = "8008"

const (
	cnum = iota // iota genera valores en secuencia y se reinicia en cada bloque const
	opa
	opb
)

type requestNode struct {
	Code int
	Port string
	Data string
}

func main() {
	go server()
	time.Sleep(time.Millisecond * 100)
	fmt.Print("Registration Central: ")
	msg := requestNode{1, myPort, ""}
	send(centralAddr, msg)
	for {

	}
}
func server() {
	if ln, err := net.Listen("tcp", "localhost:"+myPort); err != nil {
		log.Panicln("Can't start listener on", myPort)
	} else {
		defer ln.Close()
		fmt.Println("Listeing on ", myPort)
		for {
			if conn, err := ln.Accept(); err != nil {
				log.Println("Can't accept", conn.RemoteAddr())
			} else {
				go handle(conn)
			}
		}
	}
}
func handle(conn net.Conn) {
	defer conn.Close()
	dec := json.NewDecoder(conn)
	var msg requestNode
	if err := dec.Decode(&msg); err != nil {
		log.Println("Can't decode from", conn.RemoteAddr())
	} else {
		fmt.Println(msg)

		historicoEjes := [][]float32{
			[]float32{80, 71500},
			[]float32{20, 17300},
			[]float32{38, 46327},
			[]float32{70, 124743},
			[]float32{60, 37111},
			[]float32{40, 36566},
			[]float32{55, 69813},
			[]float32{64, 114846},
			[]float32{50, 39706},
		}
		historicoValores := []int{
			7500,
			17200,
			14990,
			8450,
			10950,
			11990,
			9500,
			9950,
			13500,
		}

		k := 3
		datos := 2
		numValores := 17201
		knn := api.Definir(k, datos, numValores)
		err := knn.Aprender(historicoEjes, historicoValores)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		par := []float32{60, 50000}
		respuesta := strconv.Itoa(knn.Predecir(par))
		fmt.Println(respuesta)

		msg := requestNode{4, myPort, respuesta}
		send(centralAddr, msg)
	}
}
func send(remoteAddr string, msg requestNode) {
	if conn, err := net.Dial("tcp", remoteAddr); err != nil {
		log.Println("Can't dial ", remoteAddr)
	} else {
		defer conn.Close()
		fmt.Println("Sending to ", remoteAddr)
		enc := json.NewEncoder(conn)
		enc.Encode(msg)
	}
}
