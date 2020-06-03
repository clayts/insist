package insist

import (
	"log"
	"os"
)

func Is(candidate interface{}, possibilities ...interface{}) {
	for _, p := range possibilities {
		if p == candidate {
			return
		}
	}
	log.Panicln(candidate)
}

func IsNot(candidate interface{}, possibilities ...interface{}) {
	for _, p := range possibilities {
		if p == candidate {
			log.Panicln(candidate)
		}
	}
}

func IsNil(candidate interface{}) {
	if candidate != nil {
		log.Panicln(candidate)
	}
}

func IsNotNil(candidate interface{}) {
	if candidate == nil {
		log.Panicln(candidate)
	}
}

func On(val interface{}, err error) interface{} {
	IsNil(err)
	return val
}

func OnString(val string, err error) string {
	IsNil(err)
	return val
}

func OnInt(val int, err error) int {
	IsNil(err)
	return val
}

func OnFloat32(val float32, err error) float32 {
	IsNil(err)
	return val
}

func OnFloat64(val float64, err error) float64 {
	IsNil(err)
	return val
}

func OnByteSlice(val []byte, err error) []byte {
	IsNil(err)
	return val
}

func OnFilePtr(val *os.File, err error) *os.File {
	IsNil(err)
	return val
}

func OnEnv(key string) string {
	v := os.Getenv(key)
	if len(v) == 0 {
		log.Panicln(key)
	}
	return v
}
