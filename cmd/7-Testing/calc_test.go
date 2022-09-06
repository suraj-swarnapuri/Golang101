package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddition(t *testing.T){
	c := Calc{}
	assert.Equal(t, c.add(1,2), 3)
}

func TestSubtraction(t *testing.T){
	c := Calc{}
	assert.Equal(t, c.sub(1,2), -1)
}

func TestMultiplication(t *testing.T){
	c := Calc{}
	assert.Equal(t, c.mul(1,2), 2)
}

func TestDivision(t *testing.T){
	c := Calc{}
	actual,_ := c.div(4,2)
	assert.Equal(t, actual, 2)
}

func TestDivisionByZero(t *testing.T){
	c := Calc{}
	_,err := c.div(4,0)
	assert.Error(t, err)
	assert.True(t, err.Error() == "Division by zero")
}