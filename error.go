package simpledi

import "errors"

var (
	DiInvalidFunction         = errors.New("got invalid type, you have to pass function type.")
	DiInvalidInvokerFunction  = errors.New("got invalid function type for invoke, invoke function return an optional error value.")
	DiInvalidInvokeReturnType = errors.New("got invalid invoke return type")
	DiNeedInvoke              = errors.New("you have to provide invoke inorder to run continer !")
	DiCycleDetected           = errors.New("probably there are a cycle in given dependency tree, container cant't proceed !")
	DiInvokeNotSatisfied      = errors.New("you can't run a container wihtoud providing requested attribitues !")
)
