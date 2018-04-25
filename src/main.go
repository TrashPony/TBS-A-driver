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

				correctWeight := int(weightBox)
				
				checkData := ParseData.CheckData(correctWeight)
				time.Sleep(time.Second * 3)
				if checkData {
/*
80
_ESC_Save
8
_ESC_Save
_ESC_Save
_ESC_Save
90
_ESC_Save
90
_ESC_Save
450
_ESC_Save
440
_ESC_Save
90
_ESC_Save
90
_ESC_Save
100
_ESC_Save
130
_ESC_Save
480
_ESC_Save

*/
					InputData.ToClipBoard(strconv.Itoa(correctWeight))
					InputData.ToClipBoard("_ESC_Save")

				}
			}
		}
	}
}