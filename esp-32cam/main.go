package main

import (
	"machine"
	"time"
)

func main() {
	// Configurar o Wi-Fi
	machine.InitADC()
	net := machine.WLAN
	net.Configure(machine.WLANConfig{SSID: "Joao", Password: "2578466517"})

	// Tentativa de conexão
	println("Conectando à rede Wi-Fi...")
	for {
		if net.Connected() {
			println("Conectado!")
			ip, _, _, _ := net.GetIP()
			println("Endereço IP:", ip.String())
			break
		}
		time.Sleep(time.Second)
	}

	// Seu código aqui pode usar a conexão Wi-Fi agora.
	// Por exemplo, você poderia implementar um cliente TCP simples
	// ou buscar uma implementação mínima de um cliente HTTP.
	for {
		time.Sleep(time.Second * 10)
	}
}
