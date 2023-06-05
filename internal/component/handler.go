package component

import (
	"sync"
)

var registers sync.Map

type Binder interface {
	Func(args ...any)
}

type A struct {
}

func (a *A) Func(args ...any) {

}

func Register(b Binder) {
	//a := b.(*A)
	//a.Func()
	registers.Store("a", b)
}

func Push(method string, args any) {

	registers.Range(func(key, value any) bool {
		//if key ...
		binder := value.(Binder)
		binder.Func()
		return true
	})
}
