package numbers

import (
	"errors"

	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/functions"
	"github.com/almerlucke/glisp/types/numbers"
)

// NumberAdd adds 1 or more numbers
func NumberAdd(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	var total *numbers.Number
	var err error

	err = args.Iter(func(obj types.Object, index interface{}) (bool, error) {
		num, ok := obj.(*numbers.Number)
		if !ok {
			return false, errors.New("+ accepts only numbers")
		}

		if total == nil {
			total = num
		} else {
			total, err = total.Add(num)
			if err != nil {
				return false, err
			}
		}

		return false, nil
	})

	if err != nil {
		return nil, err
	}

	return total, nil
}

// CreateBuiltinNumberAdd creates a function object
func CreateBuiltinNumberAdd() *functions.BuiltinFunction {
	return functions.NewBuiltinFunction(NumberAdd, 1, true)
}

// NumberSubtract adds 1 or more numbers
func NumberSubtract(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	var total *numbers.Number
	var err error

	if args.Length() == 1 {
		num, ok := args.Car.(*numbers.Number)
		if !ok {
			return nil, errors.New("- accepts only numbers")
		}

		return numbers.New(num.Kind).Subtract(num)
	}

	err = args.Iter(func(obj types.Object, index interface{}) (bool, error) {
		num, ok := obj.(*numbers.Number)
		if !ok {
			return false, errors.New("- accepts only numbers")
		}

		if total == nil {
			total = num
		} else {
			total, err = total.Subtract(num)
			if err != nil {
				return false, err
			}
		}

		return false, nil
	})

	if err != nil {
		return nil, err
	}

	return total, nil
}

// CreateBuiltinNumberSubtract creates a function object
func CreateBuiltinNumberSubtract() *functions.BuiltinFunction {
	return functions.NewBuiltinFunction(NumberSubtract, 1, true)
}

// NumberMultiply multiply 1 or more numbers
func NumberMultiply(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	var total *numbers.Number
	var err error

	err = args.Iter(func(obj types.Object, index interface{}) (bool, error) {
		num, ok := obj.(*numbers.Number)
		if !ok {
			return false, errors.New("* accepts only numbers")
		}

		if total == nil {
			total = num
		} else {
			total, err = total.Multiply(num)
			if err != nil {
				return false, err
			}
		}

		return false, nil
	})

	if err != nil {
		return nil, err
	}

	return total, nil
}

// CreateBuiltinNumberMultiply creates a function object
func CreateBuiltinNumberMultiply() *functions.BuiltinFunction {
	return functions.NewBuiltinFunction(NumberMultiply, 1, true)
}

// NumberDivide divide 1 or more numbers
func NumberDivide(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	var total *numbers.Number
	var err error

	if args.Length() == 1 {
		num, ok := args.Car.(*numbers.Number)
		if !ok {
			return nil, errors.New("- accepts only numbers")
		}

		otherNum := numbers.New(num.Kind)
		otherNum.SetInt64Value(1)

		return otherNum.Divide(num)
	}

	err = args.Iter(func(obj types.Object, index interface{}) (bool, error) {
		num, ok := obj.(*numbers.Number)
		if !ok {
			return false, errors.New("/ accepts only numbers")
		}

		if total == nil {
			total = num
		} else {
			total, err = total.Divide(num)
			if err != nil {
				return false, err
			}
		}

		return false, nil
	})

	if err != nil {
		return nil, err
	}

	return total, nil
}

// CreateBuiltinNumberDivide creates a function object
func CreateBuiltinNumberDivide() *functions.BuiltinFunction {
	return functions.NewBuiltinFunction(NumberDivide, 1, true)
}

// NumberModulo modulo 2 numbers
func NumberModulo(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	num1, ok := args.Car.(*numbers.Number)
	if !ok {
		return nil, errors.New("% accepts only numbers")
	}

	num2, ok := args.Cdr.(*cons.Cons).Car.(*numbers.Number)
	if !ok {
		return nil, errors.New("% accepts only numbers")
	}

	return num1.Modulo(num2)
}

// CreateBuiltinNumberModulo creates a function object
func CreateBuiltinNumberModulo() *functions.BuiltinFunction {
	return functions.NewBuiltinFunction(NumberModulo, 2, true)
}
