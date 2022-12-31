package simpledi

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProviderOption(t *testing.T) {
	type foo struct {
		a int
		b int
	}
	type bar struct {
		sum int
	}
	f := func(f foo) bar {
		b := bar{
			sum: f.a + f.b,
		}
		fmt.Printf("function called --> output value is %d\n", b.sum)
		return b
	}

	p := newprovider(f)
	require.NotNil(t, p)

	inputs := make([]input, 0)
	s := foo{
		a: 12,
		b: 2,
	}

	inputs = append(inputs, input{
		name:  "foo",
		typ:   reflect.TypeOf(s),
		value: reflect.ValueOf(s),
	})

	ready := p.readytogo(inputs)
	require.True(t, ready, "must return true")

	outputs := p.call(inputs)
	fmt.Println(outputs)
	require.NotEmpty(t, outputs)
}
