package delay_init

import (
	"errors"
)

var g_init_funcs []func() error

func AddFunc(initor func() error) error {
	g_init_funcs = append(g_init_funcs, initor)
	return nil
}

func Add(initor interface{}) error {

	if obj, ok := initor.(interface{ Init() error }); ok {
		return AddFunc(func() error {
			return obj.Init()
		})
	}

	if obj, ok := initor.(interface{ Init() }); ok {
		return AddFunc(func() error {
			obj.Init()
			return nil
		})
	}

	if obj, ok := (initor.(func() error)); ok {
		return AddFunc(obj)
	}

	if obj, ok := (initor.(func())); ok {
		return AddFunc(func() error {
			obj()
			return nil
		})
	}

	panic("invalid delay init function")
	return errors.New("invalid delay init function")
}

var g_init_done = false

func Init() error {
	if g_init_done {
		return nil
	}

	g_init_done = true

	if len(g_init_funcs) <= 0 {
		return nil
	}

	for _, f := range g_init_funcs {
		if f != nil {
			err := f()
			if err != nil {
				return err
			}
		}
	}

	return nil
}
