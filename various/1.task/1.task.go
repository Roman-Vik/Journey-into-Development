package main
// Задача 1: Провести ревью кода.

import (
	"math"
	"errors"
)

type parrotType int

const (
  TypeEuropean      parrotType = iota+ 1
  TypeAfrican       
  TypeNorwegianBlue
)

const (
  EuropeanSpeed   = 12.0
  LoadFactor  = 9.0
)

// Parrot has a Speed.
type Parrot interface {
  Speed() (float64, error)
}

type mixedParrot struct {
  Type            parrotType
  NumberOfCoconuts int
  Voltage          float64
  Nailed           bool
}

func CreateParrot(pt parrotType, numberOfCoconuts int, voltage float64, nailed bool) Parrot {
  return mixedParrot{
    Type:             pt,// Тип попугая
    NumberOfCoconuts: numberOfCoconuts, // Число кокосов
    Voltage:          voltage, // Активный
    Nailed:           nailed, // Прикереплен
  }
}

func (parrot mixedParrot) Speed() (float64, error) {
  switch parrot.Type {
  case TypeEuropean:
    return parrot.europeanSpeed(), nil 
  case TypeAfrican:
    return parrot.africanSpeed(), nil
  case TypeNorwegianBlue:
    if parrot.Nailed {
      return 0, nil
    }
    return parrot.norwegianBlueSpeed(), nil
  default:
    return 0, errors.New("should be unreachable") 
  }
}

func (parrot mixedParrot) computeBaseSpeedForVoltage(voltage float64) float64 {
  return math.Min(24.0, voltage*parrot.europeanSpeed()) 
}

func (parrot mixedParrot) loadFactor() float64 {
  return LoadFactor
}

func (parrot mixedParrot) europeanSpeed() float64 {
  return EuropeanSpeed
}

func (parrot mixedParrot) africanSpeed() float64{
	return math.Max(0, parrot.europeanSpeed()-parrot.loadFactor()*float64(parrot.NumberOfCoconuts))
} 

func (parrot mixedParrot) norwegianBlueSpeed() float64{
	return parrot.computeBaseSpeedForVoltage(parrot.Voltage)
} 
