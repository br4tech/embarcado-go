package hcs04

import (
	"machine"
	"time"
)

// Define as unidades de medida de distância
const (
	CM  = 29 // Velocidade do som em cm/us
	INC = 74 // Velocidade do som em polegadas/us
)

const TIMEOUT = 23324 // max sensing distance (4m)

// Device holds the pins
type Ultrasonic struct {
	trigger machine.Pin
	echo    machine.Pin
	timeout time.Duration
}

// New returns a new ultrasonic driver given 2 pins
func New(trigger, echo machine.Pin) *Ultrasonic {
	return &Ultrasonic{
		trigger: trigger,
		echo:    echo,
		timeout: 20000 * time.Microsecond,
	}
}

// Read lê a distância do sensor na unidade especificada (CM ou INC).
func (u *Ultrasonic) Read(unit int) uint16 {
	// Envia um pulso trigger
	u.trigger.Configure(machine.PinConfig{Mode: machine.PinOutput})
	u.trigger.High()
	time.Sleep(10 * time.Microsecond)
	u.trigger.Low()

	// Aguarda o pulso echo
	u.echo.Configure(machine.PinConfig{Mode: machine.PinInput})
	startTime := time.Now()
	for u.echo.Get() == false {
		if time.Since(startTime) > u.timeout {
			return 0
		}
	}

	// Mede a duração do pulso echo
	startTime = time.Now()
	for u.echo.Get() == true {
		if time.Since(startTime) > u.timeout {
			return 0
		}
	}
	duration := time.Since(startTime)

	// Calcula a distância
	var distance uint16
	switch unit {
	case CM:
		distance = uint16(duration.Microseconds() / 2 / CM)
	case INC:
		distance = uint16(duration.Microseconds() / 2 / INC)
	default:
		return 0
	}

	return distance
}

// SetTimeout define o tempo limite para a leitura do sensor.
func (u *Ultrasonic) SetTimeout(timeout time.Duration) {
	u.timeout = timeout
}

// SetMaxDistance define a distância máxima que o sensor pode ler.
func (u *Ultrasonic) SetMaxDistance(dist uint16) {
	u.timeout = time.Duration(dist) * CM * 2 * time.Microsecond
}
