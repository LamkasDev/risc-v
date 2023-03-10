package emu_run

import (
	"time"

	"github.com/LamkasDev/awoo-emu/cmd/awoomu/emu"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/internal"
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/rom"
)

func Load(path string) {
	/* program, _ := SelectProgram() */
	emulator := emu.SetupEmulator()
	rom.LoadROMFromPath(&emulator.Internal.ROM, path)
	Run(&emulator)
}

func Run(emulator *emu.AwooEmulator) {
	go func() {
		cycles := emulator.Config.CPU.Speed / 1000
		for emulator.Internal.Executing {
			for i := uint32(0); i < cycles; i++ {
				internal.TickInternal(&emulator.Internal)
				for _, id := range emulator.TickDrivers {
					emulator.Drivers[id] = emulator.Drivers[id].Tick(&emulator.Internal, emulator.Drivers[id])
				}
				emulator.Internal.Executing = emulator.Internal.CPU.Counter < emulator.Internal.ROM.Length
			}
			time.Sleep(time.Millisecond)
		}
	}()
	for emulator.Internal.Running {
		// TODO: this will need a proper lock system, if a driver has both tick and tick long
		for _, id := range emulator.TickLongDrivers {
			emulator.Drivers[id] = emulator.Drivers[id].TickLong(&emulator.Internal, emulator.Drivers[id])
		}
		time.Sleep(time.Millisecond)
	}
	for _, driver := range emulator.Drivers {
		if driver.Clean != nil {
			_, err := driver.Clean(&emulator.Internal, driver)
			if err != nil {
				panic(err)
			}
		}
	}
}
