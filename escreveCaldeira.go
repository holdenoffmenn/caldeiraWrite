package main

import(
	"fmt"
	"math/rand"
	"time"

	modbus "github.com/goburrow/modbus"
)

func Escreve(){

	conn := modbus.NewTCPClientHandler("192.168.1.5:502")
	conn.Timeout = 2 * time.Minute
	conn.SlaveId = 1

	err := conn.Connect()

	if err != nil {
		fmt.Println("connect failed, ", err)
		return
	}
	client := modbus.NewClient(conn)

	defer conn.Close()

	//rand.Seed(time.Now().UnixNano())

	for {
		y0 := rand.Intn(2) * 65280 //Queimador caldeira
		y1 := rand.Intn(2) * 65280 //ValvulaAgua
		y2 := rand.Intn(2) * 65280 //ValvulaPressao
		y3 := rand.Intn(2) * 65280 //BombaDAgua

		d1 := rand.Intn(2000-2500+1) + 2000 //NivelAgua
		d2 := rand.Intn(1600-1250+1) + 1250 //TemperaturaCaldeira (1250~1600)
		d3 := rand.Intn(40-11+1) + 11       //PressaoINterna
		d4 := rand.Intn(1600-200+1) + 200   //EnergiaGErada

		client.WriteSingleCoil(40960, uint16(y0))
		client.WriteSingleCoil(40961, uint16(y1))
		client.WriteSingleCoil(40962, uint16(y2))
		client.WriteSingleCoil(40963, uint16(y3))

		client.WriteSingleRegister(0, uint16(d1))
		client.WriteSingleRegister(1, uint16(d2))
		client.WriteSingleRegister(2, uint16(d3))
		client.WriteSingleRegister(3, uint16(d4))

		fmt.Printf("BIT - Y0.0[%d] - Y0.1[%d] - Y0.2[%d] - Y0.3[%d]\n",
			y0, y1, y2, y3)

		fmt.Printf("WORD - D0[%d] - D1[%d] - D2[%d] - D3[%d]\n",
			d1, d2, d3, d4)
		fmt.Printf("\n======>  <======\n")

		time.Sleep(5 * time.Second)

	}


}