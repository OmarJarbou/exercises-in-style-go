package main

type TFQuarantine struct {
	functions_list []func(interface{}) interface{}
}

func (tfq *TFQuarantine) initializeTFQuarantine() {
	tfq.functions_list = []func(interface{}) interface{}{}
}

func (tfq *TFQuarantine) bind(function func(interface{}) interface{}) {
	tfq.functions_list = append(tfq.functions_list, function)
}

func (tfq *TFQuarantine) execute(input interface{}) interface{} {
	function_input := input
	for _, function := range tfq.functions_list {
		function_input = function(function_input)
	}
	output := function_input

	return output
}

// I DIDN'T IMPLEMENT AUTOMATIC CALLING FOR function_input IF IT WAS A FUNCTION
// AND SIMPLY DONE IT EXPLICITLY IN MAIN WHILE BINDING FOR CALLABLE FUNCTIONS
// THE REAL CAUSE:

// this function won't work with us; because even if it recieved a function
// Go does not allow function-type assertions between different signatures — even
// if the return type is "compatible" (like string being assignable to interface{} in value land).
// Function types in Go must match exactly in parameter and return types for type
// assertions or assignments to succeed.

// func (tfq *TFQuarantine) callable(value interface{}) interface{} {
// 	if fn, ok := value.(func() interface{}); ok { // if function_input is a function, then call it instead of passing it as a parameter
// 		return fn()
// 	}
// 	return value
// }

// this function also won't work because we cant call value() directly in go
// func (tfq *TFQuarantine) callable(value interface{}) interface{} {
// 	typ := reflect.TypeOf(value)

// 	// You can check if it's a function
// 	if typ.Out(0).Kind() == reflect.Func {
// 		return value() // it simply doesn't work in go
// 	}
// 	return value
// }
