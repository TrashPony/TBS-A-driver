package TransportData

import (
	"github.com/tarm/serial"
	"strconv"
	"time"
)



func SelectPort() (scalePort *Port) {
	println("Поиск портов")
	portClass := []string{"/dev/ttyS", "/dev/ttyACM", "/dev/ttyUSB"}

	for {
		for _, nameClass := range portClass {
			for i := 0; i < 10; i++ {

				portName := nameClass + strconv.Itoa(i)

				if scalePort == nil {
					scalePort = FindScale(portName)
				}

				if scalePort != nil {
					println("Весы подключены.")
					return
				}
			}
		}
	}
}

func FindScale(portName string) (port *Port) {

	weightConfig := &serial.Config{Name: portName,
		Baud: 4800,
		Parity: 'E',
		ReadTimeout: time.Millisecond * 200}

	port = &Port{Name:portName, Config:weightConfig}
	connect := port.Connect()
	if connect == nil {
		return nil
	}

	connect.Flush()

	_, err := connect.Write([]byte{0x48})
	if err != nil {
		connect.Close()
		return nil
	}

	buf := make([]byte, 2)
	n, err := connect.Read(buf)

	if err != nil {
		connect.Close()
		return nil
	} else {
		if n == 2 && (buf[0] == 128 || buf[0] == 192) {
			println("Весы подключены к порту " + portName)
			return port
		} else {
			return nil
		}
	}
}
