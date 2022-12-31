package simpledi

import (
	"reflect"
)

type provider struct {
	// typ
	// specify type of the provider, which we assume is a function
	typ reflect.Type

	// target constructor
	target reflect.Value

	// provides
	// list of values which a provider givbe us
	provides []output

	// inputs
	// list number of inputs and required inputs
	inputs []input

	// isCalled
	// check if this function already called or not
	isCalled bool
}

func (p *provider) readytogo(inputs []input) bool {
	if len(p.inputs) == 0 {
		return true
	}

	return sourandset(inputs, p.inputs)
}

func (p *provider) call(inputs []input) []output {
	if p.isCalled {
		return p.provides
	}

	funcInputs := make([]reflect.Value, len(p.inputs))
	i := 0
	for _, in := range inputs {
		for _, input := range p.inputs {
			if input.typ == in.typ {
				funcInputs[i] = in.value
				i++
			}
		}
	}

	outputValues := p.target.Call(funcInputs)
	for i, o := range outputValues {
		p.provides[i].value = o
	}

	p.isCalled = true

	return p.provides
}

// WithProvider
// the proivider option tells container we have a constructor function which wish to provide some items
// for our container.
func WithProvider(function interface{}) Option {
	p := newprovider(function)
	return func(c *container) error {
		if c.providers == nil {
			c.providers = make([]provider, 0)
		}

		c.providers = append(c.providers, p)
		return nil
	}
}

func newprovider(function interface{}) provider {
	functionType := reflect.TypeOf(function)
	if functionType.Kind() != reflect.Func {
		panic(DiInvalidFunction.Error()) // panic is writen here inorder to clarify usage interface.
	}

	p := provider{
		typ:      reflect.TypeOf(function),
		target:   reflect.ValueOf(function),
		provides: make([]output, functionType.NumOut()),
		inputs:   make([]input, functionType.NumIn()),
	}
	for i := 0; i < functionType.NumIn(); i++ {
		in := input{
			name: functionType.In(i).Name(),
			typ:  functionType.In(i),
		}

		p.inputs[i] = in
	}

	for i := 0; i < functionType.NumOut(); i++ {
		ou := output{
			name: functionType.Out(i).Name(),
			typ:  functionType.Out(i),
		}

		p.provides[i] = ou
	}

	return p
}
