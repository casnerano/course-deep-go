package interfaces

import (
	"errors"
	"sync"
)

var ErrTypeNotRegistered = errors.New("type not registered")

type dep struct {
	constructor func() any
	singleton   bool
	instance    any
}

func (d *dep) getInstance() any {
	if d.singleton {
		if d.instance == nil {
			d.instance = d.constructor()
		}

		return d.instance
	}

	return d.constructor()
}

type Container struct {
	mu   sync.RWMutex
	deps map[string]dep
}

func NewContainer() *Container {
	return &Container{
		deps: make(map[string]dep),
	}
}

func (c *Container) RegisterType(name string, constructor func() any) {
	c.register(name, constructor, false)
}

func (c *Container) RegisterSingletonType(name string, constructor func() any) {
	c.register(name, constructor, true)
}

func (c *Container) register(name string, constructor func() any, singleton bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.deps[name] = dep{
		constructor: constructor,
		singleton:   singleton,
	}
}

func (c *Container) Resolve(name string) (interface{}, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	resolved, ok := c.deps[name]
	if !ok {
		return nil, ErrTypeNotRegistered
	}

	return resolved.getInstance(), nil
}
