package simpledi

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInvoker_NoError(t *testing.T) {
	type bar struct {
		sum int
	}

	f := func(b bar) {
		fmt.Printf("invoke called --> value is %d\n", b.sum)
	}

	iv := newinvoker(f)
	inputs := make([]input, 0)
	b := bar{
		sum: 13,
	}

	inputs = append(inputs, input{
		name:  "bar",
		typ:   reflect.TypeOf(b),
		value: reflect.ValueOf(b),
	})

	require.True(t, iv.readytogo(inputs))
	err := iv.call(inputs)
	require.NoError(t, err)
	// Output
	// invoke called --> value is 13
}

func TestInvoker_Error(t *testing.T) {
	type bar struct {
		sum int
	}

	f := func(b bar) error {
		fmt.Printf("invoke called --> value is %d\n", b.sum)
		return errors.New("some things happenned")
	}

	iv := newinvoker(f)
	inputs := make([]input, 0)
	b := bar{
		sum: 13,
	}

	inputs = append(inputs, input{
		name:  "bar",
		typ:   reflect.TypeOf(b),
		value: reflect.ValueOf(b),
	})

	require.True(t, iv.readytogo(inputs))
	err := iv.call(inputs)
	require.Error(t, err, "somethings happednned error have to be returned")
	fmt.Println(err)
	// Output
	// invoke called --> value is 13
	// some things happenned
}
