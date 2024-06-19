To find samplerate and duration of M4a type audio format pls follow below steps

Install   go get github.com/haritechversant/M4aDecoder/m4adecoder

package main

import (
	"fmt"

	"github.com/haritechversant/M4aDecoder/m4adecoder"
)

func main() {

	sampleRate, duration, err := m4adecoder.GetAudioMetadata("sample.m4a")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sampleRate)
	fmt.Println(duration)

}
