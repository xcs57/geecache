package singleflight

import "sync"

// 正在进行或者结束的请求
type call struct {
	wg sync.WaitGroup

	val interface{}

	err error
}

// Group 管理不同的key的请求
type Group struct {
	mu sync.Mutex
	m  map[string]*call
}

// Do 针对同样的key,无论Do调用多少次,fn都只会调用一次
func (g *Group) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	g.mu.Lock()
	if g.m == nil {
		g.m = make(map[string]*call)
	}
	if c, ok := g.m[key]; ok {
		g.mu.Unlock()
		c.wg.Wait()
		return c.val, c.err
	}

	c := new(call)
	c.wg.Add(1)
	g.m[key] = c
	g.mu.Unlock()

	c.val, c.err = fn()
	c.wg.Done()

	g.mu.Lock()
	delete(g.m, key)
	g.mu.Unlock()

	return c.val, c.err

}
