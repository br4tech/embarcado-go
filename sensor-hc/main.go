package main

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/hcsr04"
)

const (
	triggerPin = machine.D3
	echoPin    = machine.D2
)

const (
	ledVerde    = machine.D4
	ledAmarelo  = machine.D5
	ledVermelho = machine.D6
)

func main() {

	sensor := hcsr04.New(triggerPin, echoPin)

	ledVerde.Configure(machine.PinConfig{Mode: machine.PinOutput})
	ledAmarelo.Configure(machine.PinConfig{Mode: machine.PinOutput})
	ledVermelho.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		distanciaMM := sensor.ReadDistance()
		println("Distância (mm):", distanciaMM)
		ledDistancia(distanciaMM)

		time.Sleep(1000 * time.Millisecond) // tempo de espera aumentado
	}
}

func ledDistancia(distanciaMM int32) { // usando uint8
	ledVerde.Low()
	ledAmarelo.Low()
	ledVermelho.Low()

	switch {
	case distanciaMM < 200: // 10 cm em milímetros
		ledVermelho.High()
	case distanciaMM <= 200: // 20 cm em milímetros
		ledAmarelo.High()
	case distanciaMM <= 300: // 30 cm em milímetros
		ledVerde.High()

	}
}
