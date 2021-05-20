package sample

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/jirawan-chuapradit/grpc-golang-pcbook/pb"
)

// NewKeyboard returns a new sample keyboard
func NewKeyboard() *pb.Keyboard {
	keyboard := &pb.Keyboard{
		Layout:  randomKeyboardLayout(),
		Backlit: randomBool(),
	}
	return keyboard
}

// NewCPU returns a new sample CPU
func NewCPU() *pb.CPU {
	brand := randomCPUBrand()
	name := randomCPUName(brand)

	numberCores := randomInt(2, 8)
	numberThreads := randomInt(numberCores, 12)

	minGhz := randomFloat64(2.0, 3.5)
	maxGhz := randomFloat64(minGhz, 5.0)

	cpu := &pb.CPU{
		Brand:         brand,
		Name:          name,
		NumberCores:   uint32(numberCores),
		NumberThreads: uint32(numberThreads),
		MinGhz:        minGhz,
		MaxGhz:        maxGhz,
	}

	return cpu
}

// NewGPU returns a new sample GPU
func NewGPU() *pb.GPU {
	brand := randomGPUBrand()
	name := randomGPUName(brand)
	minGhz := randomFloat64(1.0, 1.5)
	maxGhz := randomFloat64(minGhz, 2.0)

	memory := &pb.Memory{
		Value: uint64(randomInt(2, 6)),
		Unit:  pb.Memory_GIGABYTE,
	}

	gpu := &pb.GPU{
		Brand:  brand,
		Name:   name,
		MinGhz: minGhz,
		MaxGhz: maxGhz,
		Memory: memory,
	}

	return gpu

}

// NewRam returns a new sample Ram
func NewRam() *pb.Memory {
	ram := &pb.Memory{
		Value: uint64(randomInt(4,64)),
		Unit: pb.Memory_GIGABYTE,
	}
	return ram
}

// NewSSD returns a new sample SSD storage
func NewSSD() *pb.Storage {
	ssd := &pb.Storage{
		Driver: pb.Storage_SSD,
		Memory: &pb.Memory{
			Value: uint64(randomInt(128, 1024)),
			Unit: pb.Memory_GIGABYTE,
		},
	}
	return ssd
}

// NewHDD returns a new sample HDD storage
func NewHDD() *pb.Storage {
	hdd := &pb.Storage{
		Driver: pb.Storage_HDD,
		Memory: &pb.Memory{
			Value: uint64(randomInt(1,6)),
			Unit: pb.Memory_GIGABYTE,
		},
	}
	return hdd
}


func NewScreen() *pb.Screen {
	screen := &pb.Screen {
		SizeInch: randomfloat32(13,17),
		Resolution: randomScreenResolution(),
		Panel: randomScreenPanel(),
		Multitouch: randomBool(),
	}
	return screen
}

func NewLaptop() *pb.Laptop {
	brand := randomLaptopBrand()
	name := randomLaptopName(brand)
	laptop := &pb.Laptop{
		Id: radnomID(),
		Brand: brand,
		Name: name,
		Cpu: NewCPU(),
		Ram: NewRam(),
		Gpus: []*pb.GPU{NewGPU()},
		Storage: []*pb.Storage{NewSSD(),NewHDD()},
		Screen: NewScreen(),
		Keyboard: NewKeyboard(),
		Weight: &pb.Laptop_WeightKg{
			WeightKg: randomFloat64(1.0, 3.0),
		},
		PriceUsd: randomFloat64(1500, 3000),
		ReleaseYear: uint32(randomInt(2015,2019)),
		UpdatedAt: ptypes.TimestampNow(),
	}
	return laptop
}