package ch9

// 需要缓存的方法
type Func func(key string) (interface{}, error)

// 缓存
type Memo struct {
	requests chan request
}

// 一次获取缓存的请求
type request struct {
	key      string
	response chan<- result
}

// 对result的进一步包装，并发读取同一个result时可以通过ready来避免重复计算
type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

// 并发计算
func (e *entry) call(f Func, key string) {
	e.res.value, e.res.err = f(key)
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	// 等待计算完成
	<-e.ready

	// 准备好了就可以把result传递到response中
	response <- e.res
}

// Func计算的结果
type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string, done <-chan struct{}) (interface{}, error) {

	response := make(chan result)

	memo.requests <- request{key, response}

	res := <-response

	return res.value, res.err

}

func (memmo *Memo) server(f Func) {
	// 实际的缓存内容都限制在这个goroutine中
	cache := make(map[string]*entry)

	for req := range memmo.requests {
		e := cache[req.key]
		if e == nil {
			// 第一次调用，需要计算结果
			e := &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key)
		}
		// 响应计算结果
		go e.deliver(req.response)
	}
}
