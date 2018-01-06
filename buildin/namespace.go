package buildin

import (
	"errors"
	"fmt"

	goStrings "strings"

	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/interfaces/namespace"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/functions"
	"github.com/almerlucke/glisp/types/namespaces"
	"github.com/almerlucke/glisp/types/strings"
	"github.com/almerlucke/glisp/types/symbols"
)

func namespaceGetName(obj types.Object) (string, error) {
	if obj.Type() == types.Symbol {
		return obj.(*symbols.Symbol).Name, nil
	}

	if obj.Type() == types.String {
		return goStrings.ToUpper(string(obj.(strings.String))), nil
	}

	return "", errors.New("namespace expected a string or symbol")
}

func namespaceClauseContainsOnlyNames(c *cons.Cons) bool {
	nameOnly := true

	c.Iter(func(obj types.Object, index interface{}) (bool, error) {
		t := obj.Type()

		if t != types.Symbol && t != types.String {
			nameOnly = false
			return true, nil
		}

		return false, nil
	})

	return nameOnly
}

func importSymbolsFromOtherNamespace(ns namespace.Namespace, clause *cons.Cons, env environment.Environment) error {
	nsName, _ := namespaceGetName(clause.Car)
	otherNS := env.FindNamespace(nsName)

	if otherNS == nil {
		return fmt.Errorf("import-from undefined namespace %v", nsName)
	}

	if clause.Cdr.Type() != types.Cons {
		return nil
	}

	return clause.Cdr.(*cons.Cons).Iter(func(obj types.Object, index interface{}) (bool, error) {
		name, _ := namespaceGetName(obj)
		success := ns.Import(name, otherNS)

		if !success {
			return false, fmt.Errorf("import-from namespace %v unknown symbol %v", nsName, name)
		}

		return false, nil
	})
}

func shadowSymbolsInNamespace(ns namespace.Namespace, clause *cons.Cons, env environment.Environment) {
	clause.Iter(func(obj types.Object, index interface{}) (bool, error) {
		name, _ := namespaceGetName(obj)

		ns.Shadow(name)

		return false, nil
	})
}

func exportSymbolsInNamespace(ns namespace.Namespace, clause *cons.Cons, env environment.Environment) {
	clause.Iter(func(obj types.Object, index interface{}) (bool, error) {
		name, _ := namespaceGetName(obj)

		ns.Export(ns.Intern(name))

		return false, nil
	})
}

func useOtherNamespaces(ns namespace.Namespace, clause *cons.Cons, env environment.Environment) error {
	return clause.Iter(func(obj types.Object, index interface{}) (bool, error) {
		name, _ := namespaceGetName(obj)
		otherNS := env.FindNamespace(name)

		if otherNS == nil {
			return false, fmt.Errorf("undefined namespace %v", name)
		}

		ns.Use(otherNS)

		return false, nil
	})
}

// Namespace buildin function
func Namespace(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	name, err := namespaceGetName(args.Car)
	if err != nil {
		return nil, err
	}

	ns := namespaces.NewNamespace(name, true)

	if args.Cdr.Type() == types.Cons {
		err = args.Cdr.(*cons.Cons).Iter(func(obj types.Object, index interface{}) (bool, error) {
			if obj.Type() != types.Cons {
				return false, errors.New("illegal namespace clause")
			}

			clause := obj.(*cons.Cons)

			if !namespaceClauseContainsOnlyNames(clause) {
				return false, errors.New("namespace clause must contains only names")
			}

			name, _ := namespaceGetName(clause.Car)
			if name == "USE" {
				if clause.Cdr.Type() == types.Cons {
					err = useOtherNamespaces(ns, clause.Cdr.(*cons.Cons), env)
					if err != nil {
						return false, err
					}
				}
			} else if name == "EXPORT" {
				if clause.Cdr.Type() == types.Cons {
					exportSymbolsInNamespace(ns, clause.Cdr.(*cons.Cons), env)
				}
			} else if name == "IMPORT-FROM" {
				if clause.Cdr.Type() == types.Cons {
					err = importSymbolsFromOtherNamespace(ns, clause.Cdr.(*cons.Cons), env)
					if err != nil {
						return false, err
					}
				}
			} else if name == "SHADOW" {
				if clause.Cdr.Type() == types.Cons {
					shadowSymbolsInNamespace(ns, clause.Cdr.(*cons.Cons), env)
				}
			}

			return false, nil
		})

		if err != nil {
			return nil, err
		}
	}

	err = env.AddNamespace(ns)
	if err != nil {
		return nil, err
	}

	return ns, nil
}

// InNamespace buildin function
func InNamespace(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	name, err := namespaceGetName(args.Car)
	if err != nil {
		return nil, err
	}

	ns := env.FindNamespace(name)
	if ns == nil {
		return nil, fmt.Errorf("undefined namespace %v", name)
	}

	if !ns.CanIntern() {
		return nil, fmt.Errorf("namespace %v is locked", name)
	}

	env.ChangeCurrentNamespace(name)

	return ns, nil
}

// UseNamespace buildin function
func UseNamespace(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	name, err := namespaceGetName(args.Car)
	if err != nil {
		return nil, err
	}

	ns := env.FindNamespace(name)
	if ns == nil {
		return nil, fmt.Errorf("undefined namespace %v", name)
	}

	curNS := env.CurrentNamespace()

	curNS.Use(ns)

	return ns, nil
}

// CreateBuildinNamespace creates a buildin function object
func CreateBuildinNamespace() *functions.BuildinFunction {
	return functions.NewBuildinFunction(Namespace, 1, false)
}

// CreateBuildinInNamespace creates a buildin function object
func CreateBuildinInNamespace() *functions.BuildinFunction {
	return functions.NewBuildinFunction(InNamespace, 1, false)
}

// CreateBuildinUseNamespace creates a buildin function object
func CreateBuildinUseNamespace() *functions.BuildinFunction {
	return functions.NewBuildinFunction(UseNamespace, 1, false)
}
