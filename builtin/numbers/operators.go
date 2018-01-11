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
			return false, errors.New("+ only accepts numbers")
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

// NumberSubtract adds 1 or more numbers
func NumberSubtract(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	var total *numbers.Number
	var err error

	if args.Length() == 1 {
		num, ok := args.Car.(*numbers.Number)
		if !ok {
			return nil, errors.New("- only accepts numbers")
		}

		return numbers.New(num.Kind).Subtract(num)
	}

	err = args.Iter(func(obj types.Object, index interface{}) (bool, error) {
		num, ok := obj.(*numbers.Number)
		if !ok {
			return false, errors.New("- only accepts numbers")
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

// NumberMultiply multiply 1 or more numbers
func NumberMultiply(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	var total *numbers.Number
	var err error

	err = args.Iter(func(obj types.Object, index interface{}) (bool, error) {
		num, ok := obj.(*numbers.Number)
		if !ok {
			return false, errors.New("* only accepts numbers")
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

// NumberDivide divide 1 or more numbers
func NumberDivide(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	var total *numbers.Number
	var err error

	if args.Length() == 1 {
		num, ok := args.Car.(*numbers.Number)
		if !ok {
			return nil, errors.New("- only accepts numbers")
		}

		otherNum := numbers.New(num.Kind)
		otherNum.SetInt64Value(1)

		return otherNum.Divide(num)
	}

	err = args.Iter(func(obj types.Object, index interface{}) (bool, error) {
		num, ok := obj.(*numbers.Number)
		if !ok {
			return false, errors.New("/ only accepts numbers")
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

// NumberModulo modulo 2 numbers
func NumberModulo(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	num1, ok := args.Car.(*numbers.Number)
	if !ok {
		return nil, errors.New("% only accepts numbers")
	}

	num2, ok := args.Cdr.(*cons.Cons).Car.(*numbers.Number)
	if !ok {
		return nil, errors.New("% only accepts numbers")
	}

	return num1.Modulo(num2)
}

// NumberMax get max of 1 or more numbers
func NumberMax(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	var max *numbers.Number
	var err error

	err = args.Iter(func(obj types.Object, index interface{}) (bool, error) {
		num, ok := obj.(*numbers.Number)
		if !ok {
			return false, errors.New("MAX only accepts numbers")
		}

		if max == nil {
			max = num
		} else {
			max, err = max.Max(num)
			if err != nil {
				return false, err
			}
		}

		return false, nil
	})

	if err != nil {
		return nil, err
	}

	return max, nil
}

// NumberMin get min of 1 or more numbers
func NumberMin(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	var min *numbers.Number
	var err error

	err = args.Iter(func(obj types.Object, index interface{}) (bool, error) {
		num, ok := obj.(*numbers.Number)
		if !ok {
			return false, errors.New("MIN only accepts numbers")
		}

		if min == nil {
			min = num
		} else {
			min, err = min.Min(num)
			if err != nil {
				return false, err
			}
		}

		return false, nil
	})

	if err != nil {
		return nil, err
	}

	return min, nil
}

// NumberGreaterThan check if each argument is strictly greater than the following argument
func NumberGreaterThan(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	var prev *numbers.Number
	var err error

	greaterThan := true

	err = args.Iter(func(obj types.Object, index interface{}) (bool, error) {
		num, ok := obj.(*numbers.Number)
		if !ok {
			return false, errors.New("> only accepts numbers")
		}

		if prev != nil {
			greaterThan, err = prev.GreaterThan(num)
			if err != nil {
				return false, err
			}

			if !greaterThan {
				return true, nil
			}
		}

		prev = num

		return false, nil
	})

	if err != nil {
		return nil, err
	}

	if greaterThan {
		return types.T, nil
	}

	return types.NIL, nil
}

// NumberGreaterThanOrEqual check if each argument is strictly greater or equal than the following argument
func NumberGreaterThanOrEqual(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	var prev *numbers.Number
	var err error

	greaterThanOrEqual := true

	err = args.Iter(func(obj types.Object, index interface{}) (bool, error) {
		num, ok := obj.(*numbers.Number)
		if !ok {
			return false, errors.New(">= only accepts numbers")
		}

		if prev != nil {
			greaterThanOrEqual, err = prev.GreaterThanOrEqual(num)
			if err != nil {
				return false, err
			}

			if !greaterThanOrEqual {
				return true, nil
			}
		}

		prev = num

		return false, nil
	})

	if err != nil {
		return nil, err
	}

	if greaterThanOrEqual {
		return types.T, nil
	}

	return types.NIL, nil
}

// NumberLesserThan check if each argument is strictly lesser than the following argument
func NumberLesserThan(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	var prev *numbers.Number
	var err error

	lesserThan := true

	err = args.Iter(func(obj types.Object, index interface{}) (bool, error) {
		num, ok := obj.(*numbers.Number)
		if !ok {
			return false, errors.New("< only accepts numbers")
		}

		if prev != nil {
			lesserThan, err = prev.LesserThan(num)
			if err != nil {
				return false, err
			}

			if !lesserThan {
				return true, nil
			}
		}

		prev = num

		return false, nil
	})

	if err != nil {
		return nil, err
	}

	if lesserThan {
		return types.T, nil
	}

	return types.NIL, nil
}

// NumberLesserThanOrEqual check if each argument is strictly lesser than or equal to the following argument
func NumberLesserThanOrEqual(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	var prev *numbers.Number
	var err error

	lesserThanOrEqual := true

	err = args.Iter(func(obj types.Object, index interface{}) (bool, error) {
		num, ok := obj.(*numbers.Number)
		if !ok {
			return false, errors.New("<= only accepts numbers")
		}

		if prev != nil {
			lesserThanOrEqual, err = prev.LesserThanOrEqual(num)
			if err != nil {
				return false, err
			}

			if !lesserThanOrEqual {
				return true, nil
			}
		}

		prev = num

		return false, nil
	})

	if err != nil {
		return nil, err
	}

	if lesserThanOrEqual {
		return types.T, nil
	}

	return types.NIL, nil
}

// CreateBuiltinNumberGreaterThan creates a function object
func CreateBuiltinNumberGreaterThan() *functions.BuiltinFunction {
	return functions.NewBuiltinFunction(NumberGreaterThan, 1, true)
}

// CreateBuiltinNumberGreaterThanOrEqual creates a function object
func CreateBuiltinNumberGreaterThanOrEqual() *functions.BuiltinFunction {
	return functions.NewBuiltinFunction(NumberGreaterThanOrEqual, 1, true)
}

// CreateBuiltinNumberLesserThan creates a function object
func CreateBuiltinNumberLesserThan() *functions.BuiltinFunction {
	return functions.NewBuiltinFunction(NumberLesserThan, 1, true)
}

// CreateBuiltinNumberLesserThanOrEqual creates a function object
func CreateBuiltinNumberLesserThanOrEqual() *functions.BuiltinFunction {
	return functions.NewBuiltinFunction(NumberLesserThanOrEqual, 1, true)
}

// CreateBuiltinNumberAdd creates a function object
func CreateBuiltinNumberAdd() *functions.BuiltinFunction {
	return functions.NewBuiltinFunction(NumberAdd, 1, true)
}

// CreateBuiltinNumberSubtract creates a function object
func CreateBuiltinNumberSubtract() *functions.BuiltinFunction {
	return functions.NewBuiltinFunction(NumberSubtract, 1, true)
}

// CreateBuiltinNumberMultiply creates a function object
func CreateBuiltinNumberMultiply() *functions.BuiltinFunction {
	return functions.NewBuiltinFunction(NumberMultiply, 1, true)
}

// CreateBuiltinNumberDivide creates a function object
func CreateBuiltinNumberDivide() *functions.BuiltinFunction {
	return functions.NewBuiltinFunction(NumberDivide, 1, true)
}

// CreateBuiltinNumberModulo creates a function object
func CreateBuiltinNumberModulo() *functions.BuiltinFunction {
	return functions.NewBuiltinFunction(NumberModulo, 2, true)
}

// CreateBuiltinNumberMax creates a function object
func CreateBuiltinNumberMax() *functions.BuiltinFunction {
	return functions.NewBuiltinFunction(NumberMax, 1, true)
}

// CreateBuiltinNumberMin creates a function object
func CreateBuiltinNumberMin() *functions.BuiltinFunction {
	return functions.NewBuiltinFunction(NumberMin, 1, true)
}
