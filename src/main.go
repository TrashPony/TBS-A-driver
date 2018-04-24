package main

import (
	"./TransportData"
	"./ParseData"
	"./InputData"
	"strconv"
	"time"
)

var scalePort *TransportData.Port

func main() {
	Controller()
}

func Controller() {

	for {

		if scalePort == nil {

			scalePort = TransportData.SelectPort()

		} else {

			scaleResponse := TransportData.SendScaleCommand(scalePort)

			if scaleResponse == nil {
				println("Весы отвалились")
				scalePort = nil
			}

			if scalePort != nil {

				weightBox := ParseData.ParseScaleData(scaleResponse)

				correctWeight := int(weightBox * 100) * 10 // TODO исправить протокол для 60 кг весов
				
				checkData := ParseData.CheckData(correctWeight)

				if checkData {

					InputData.ToClipBoard(strconv.Itoa(correctWeight))
					InputData.ToClipBoard("_ESC_Save")
					time.Sleep(time.Second * 3)
				}
			}
		}
	}
}