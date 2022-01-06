package service

import (
	"errors"
	"math"
)

type Calculator struct {
	Num1 float64
	NUm2 float64
}

func (c *Calculator) Add() (*float64, error) {
	if math.Signbit(c.Num1) || math.Signbit(c.NUm2) {
		return nil, errors.New("negative number detected")
	}
	result := c.Num1 + c.NUm2
	return &result, nil
}

func (c *Calculator) Sub() (*float64, error) {
	if math.Signbit(c.Num1) || math.Signbit(c.NUm2) {
		return nil, errors.New("negative number detected")
	}
	result := c.Num1 - c.NUm2
	return &result, nil
}
