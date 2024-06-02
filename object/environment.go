package object

// NewEnclosedEnvironment creates a new environment that is enclosed within the given outer environment.
// The new environment will have access to the variables and functions defined in the outer environment.
func NewEnclosedEnvironment(outer *Enviroment) *Enviroment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

func NewEnvironment() *Enviroment {
	s := make(map[string]Object)
	return &Enviroment{store: s, outer: nil}
}

type Enviroment struct {
	store map[string]Object
	outer *Enviroment
}

// Get retrieves an Object from the Environment by name. If the Object is not found in the
// current Environment, it will recursively search the outer Environment, if one exists.
// Returns the Object and a boolean indicating whether the Object was found.
func (e *Enviroment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

func (e *Enviroment) Set(name string, value Object) Object {
	e.store[name] = value
	return value
}
