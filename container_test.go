package simpledi

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContainer(t *testing.T) {
	type sum struct {
		a int
		b int
	}
	type mult struct {
		a int
		b int
	}

	sumConstructor := func() sum {
		fmt.Printf("summation constructor called -->>\n")

		s := sum{
			a: 1,
			b: 2,
		}

		return s
	}

	mulConstructor := func() mult {
		fmt.Printf("multiplication constructor called -->>\n")

		m := mult{
			a: 9,
			b: 10,
		}

		return m
	}

	usage := func(a sum, b mult) {
		fmt.Printf("invoker called -->>\nsum is : %d\nmul is : %d\n", a.a+a.b, b.a*b.b)
	}

	c, err := New(
		WithProvider(sumConstructor),
		WithProvider(mulConstructor),
		WithInvoker(usage, false),
	)
	require.NoError(t, err, "all dependencies are satisfied !")

	err = c.Run()
	require.NoError(t, err)
	// Output
	// summation constructor called -->>
	// multiplication constructor called -->>
	// invoker called -->>
	// sum is : 3
	// mul is : 90
}

func TestContainer_Retry(t *testing.T) {
	type bar struct {
		value string
	}
	type foo struct {
		value string
	}

	fooConstructor := func(b bar) foo {
		fmt.Printf("foo constructor called -->> %s\n", b.value)

		f := foo{
			value: "foo called !",
		}

		return f
	}

	barConstructor := func() bar {
		fmt.Printf("bar constructor called -->>\n")

		b := bar{
			value: "bar called !",
		}

		return b
	}

	usage := func(b bar, f foo) {
		fmt.Printf("invoker called -->>\nfoo is : %s\nbar is : %s\n", f.value, b.value)
	}

	c, err := New(
		WithProvider(fooConstructor),
		WithProvider(barConstructor),
		WithInvoker(usage, false),
	)
	require.NoError(t, err, "all dependencies are satisfied !")

	err = c.Run()
	require.NoError(t, err)
	// Output
	// bar constructor called -->>
	// foo constructor called -->> bar called !
	// invoker called -->>
	// foo is : foo called !
	// bar is : bar called !
}

func TestConstructor_Cycle(t *testing.T) {
	type bar struct {
		value string
	}
	type foo struct {
		value string
	}

	fooConstructor := func(b bar) foo {
		fmt.Printf("foo constructor called -->> %s !!!!\n", b.value)

		f := foo{
			value: "foo called !",
		}

		return f
	}

	barConstructor := func(f foo) bar {
		fmt.Printf("bar constructor called -->> %s !!!!\n", f.value)

		b := bar{
			value: "bar called !",
		}

		return b
	}

	usage := func(b bar, f foo) {
		fmt.Printf("invoker called -->>\nfoo is : %s\nbar is : %s\n", f.value, b.value)
	}
	c, err := New(
		WithProvider(fooConstructor),
		WithProvider(barConstructor),
		WithInvoker(usage, false),
	)
	require.NoError(t, err, "all dependencies are satisfied !")

	err = c.Run()
	require.Error(t, err)
	fmt.Printf("got error %v\n", err)
	// Output
	// got error probably there are a cycle in given dependency tree, container cant't proceed !
}
