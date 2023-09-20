
package calculator

import (
	"context"
	"fmt"

	"github.com/mnogu/go-calculator"
)

const (
	pluginName         = `Calculator`
	pluginDesc         = `A calculator, capable of performing mathematical calculations, where the input is a description of a mathematical expression and the return is the result of the calculation. For example: the input is: one plus two, the return is three.`
	pluginInputExample = `1+2`
)

type Calculator struct{}

func NewCalculator() *Calculator {

	return &Calculator{}
}