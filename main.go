package main

import (
    "time"

    "github.com/claudemuller/chip-8-emu-go/hardware"
)

const FPS = 60
const FPS_INTERVAL = 1000 / FPS

func main() {
    renderer := hardware.NewRenderer(10)
    cpu := hardware.NewCPU(renderer)
    err := cpu.LoadSpritesIntoMemory()
    if err != nil {
        panic(err)
    }

    tn := time.Now()
    then := tn.Unix()

    step(cpu, then)
}

func step(cpu *hardware.CPU, then int64) {
    tn := time.Now()
    now := tn.Unix()
    elapsed := now - then
    if elapsed > FPS_INTERVAL {
        cpu.Cycle()
    }
}
