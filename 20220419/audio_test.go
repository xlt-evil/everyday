package _0220419

import (
	"encoding/binary"
	"fmt"
	"math"
	"os"
	"testing"
)

/**
 * @Author: hxy
 * @Description:
 * @File:  audio
 * @Date: 2022/04/19 16:32
 */

const nsamps = 50 // samples to generate

const (
	Duration   = 2
	SampleRate = 44100
	Frequency  = 440  // Pitch Standard
)

func generate() {
	tau := math.Pi * 2
	nsamps := Duration * SampleRate
	var angle float64 = tau / float64(nsamps)
	file := "out.bin"
	f, _ := os.Create(file)
	for i := 0; i < nsamps; i++ {
		sample := math.Sin(angle * Frequency * float64(i))
		var buf [8]byte
		binary.LittleEndian.PutUint32(buf[:],
			math.Float32bits(float32(sample)))
		bw,_ := f.Write(buf[:])
		fmt.Printf("\rWrote: %v bytes to %s", bw, file)
	}
}

func Test_genrate(t *testing.T) {
	generate()
}