package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"net/http"

	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

type Dados struct {
	MemoryUsage      string    `json:"memoryUsage"`
	DiskUsage        string    `json:"diskUsage"`
	ConnectedDevices string `json:"connectedDevices"`
}

func GetInfoRasp() {
	for {
		cpuUsage, err := cpu.Percent(0, false)
		if err != nil {
			fmt.Println("Erro ao obter o uso da CPU:", err)
			return
		}

		diskUsage, err := disk.Usage("/")
		if err != nil {
			fmt.Println("Erro ao obter o uso do disco:", err)
			return
		}

		memUsage, err := mem.VirtualMemory()
		if err != nil {
			fmt.Println("Erro ao obter o uso da memória:", err)
			return
		}

		fmt.Printf("Uso da CPU: %.2f%%  Uso do Disco: %.2f%%  Uso da Memória: %.2f%%",
			cpuUsage[0], diskUsage.UsedPercent, memUsage.UsedPercent)

		mem := strconv.Itoa(int(memUsage.UsedPercent))
		disk := strconv.Itoa(int(diskUsage.UsedPercent))
		cpu := strconv.Itoa(int(cpuUsage[0]))

		dados := Dados{
			MemoryUsage:      mem,
			DiskUsage:        disk,
			ConnectedDevices: cpu,
		}
		enviarRequisicaoPost(dados)

		time.Sleep(3 * time.Second) // Atraso de 1 segundo antes da próxima atualização
	}
}

func enviarRequisicaoPost(dados Dados) {
	// Crie uma instância da estrutura Dados com os valores desejados

	// Converta a estrutura em formato JSON
	payload, err := json.Marshal(dados)
	if err != nil {
		panic(err)
	}

	// Crie uma requisição POST com o payload de dados JSON
	req, err := http.NewRequest("POST", "http://localhost:8585/api/v1/resources", bytes.NewBuffer(payload))
	if err != nil {
		panic(err)
	}

	// Defina o cabeçalho Content-Type para application/json
	req.Header.Set("Content-Type", "application/json")

	// Crie um cliente HTTP e envie a requisição
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// Processar a resposta, se necessário
	// ...
}
