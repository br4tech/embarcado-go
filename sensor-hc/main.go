package main

import (
	"machine"
	"time"

	"github.com/br4tech/embarcado-go/sensor/hcs04"
)

const (
	triggerPin = machine.D3
	echoPin    = machine.D2
)

const (
	ledVerde    = machine.D10
	ledAmarelo  = machine.D9
	ledVermelho = machine.D8
)

func main() {

	sensor := hcs04.New(triggerPin, echoPin)

	ledVerde.Configure(machine.PinConfig{Mode: machine.PinOutput})
	ledAmarelo.Configure(machine.PinConfig{Mode: machine.PinOutput})
	ledVermelho.Configure(machine.PinConfig{Mode: machine.PinOutput})

	println("Lendo dados do sensor...")
	for {
		distanciaCM := sensor.Read(29)
		println("Valor do trigger:", triggerPin)
		println("Valor do echo:", echoPin)
		ledDistancia(distanciaCM)

		time.Sleep(1000 * time.Millisecond) // tempo de espera aumentado
	}
}

func ledDistancia(distanciaCM uint16) { // usando uint8
	ledVerde.Low()
	ledAmarelo.Low()
	ledVermelho.Low()

	println("Distância (cm):", distanciaCM)

	// Acende o LED correspondente à distância
	if distanciaCM <= 30 && distanciaCM >= 20 {
		ledVerde.High()
	} else if distanciaCM <= 20 && distanciaCM >= 10 {
		ledAmarelo.High()
	} else if distanciaCM < 10 {
		ledVermelho.High()
	}
}
