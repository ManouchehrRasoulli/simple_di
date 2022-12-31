package simpledi

import (
	"reflect"
)

type invoker struct {
	// target function which we are goring to run
	target reflect.Value

	// inputs list of inputs which we have to provide into invoker
	inputs []input
}

func (i *invoker) readytogo(inputs []input) bool {
	if len(i.inputs) == 0 {
		return true
	}

	return sourandset(inputs, i.inputs)
}

func (iv invoker) call(inputs []input) error {
	funcInputs := make([]reflect.Value, len(iv.inputs))
	i := 0
	for _, in := range inputs {
		for _, input := range iv.inputs {
			if input.typ == in.typ {
				funcInputs[i] = in.value
				i++
			}
		}
	}

	outputValues := iv.target.Call(funcInputs)
	if len(outputValues) == 1 { // got error
		er, _ := outputValues[0].Interface().(error)
		return er
	}

	return nil
}

// WithInvoker
// provider inputs for followign function and if requested run in asyncronous manner
// TODO :: provide na option to stop async process inorder to prevent leakage.
func WithInvoker(function interface{}, async bool) Option {
	i := newinvoker(function)
	return func(c *container) error {
		c.invoker = &i
		return nil
	}
}

func newinvoker(function interface{}) invoker {
	functionType := reflect.TypeOf(function)
	if functionType.Kind() != reflect.Func {
		panic(DiInvalidFunction.Error())
	}

	if functionType.NumOut() > 1 {
		panic(DiInvalidInvokerFunction.Error())
	}

	iok := invoker{
		target: reflect.ValueOf(function),
		inputs: make([]input, functionType.NumIn()),
	}
	for i := 0; i < functionType.NumIn(); i++ {
		in := input{
			name: functionType.In(i).Name(),
			typ:  functionType.In(i),
		}

		iok.inputs[i] = in
	}

	return iok
}
