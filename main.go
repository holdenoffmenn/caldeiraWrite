package main

import (
	"fmt"
	"math/rand"
	"time"

	modbus "github.com/goburrow/modbus"
)

func main() {
	conn := modbus.NewTCPClientHandler("172.16.16.61:502")

	err := conn.Connect()
	if err != nil {
		fmt.Println("connect failed, ", err)
		return
	}
	client := modbus.NewClient(conn)

	defer conn.Close()

	//rand.Seed(time.Now().UnixNano())

	for {
		y0 := rand.Intn(2) * 65280
		y1 := rand.Intn(2) * 65280

		d1000 := rand.Intn(161) + 20
		d2000 := rand.Intn(601) + 300

		client.WriteSingleCoil(40960, uint16(y0))
		client.WriteSingleCoil(20484, uint16(y1))

		client.WriteSingleRegister(1000, uint16(d1000))
		client.WriteSingleRegister(2000, uint16(d2000))

		fmt.Printf("Y0[%d] - Y1[%d] - D1000[%d] - D2000[%d]\n", y0, y1, d1000, d2000)

		time.Sleep(5 * time.Second)

	}

}
