

`// A Context carries a deadline，cancellation signal，and request-scoped values
 // across API. Its methods are safe for simultaneous use by multiple goroutines
 // 一个 Context 可以在 API (无论是否是协程间) 之间传递截止日期、取消信号、请求数据。
 // Context 中的方法都是协程安全的。
 type Context interface {
     // Done returns a channel that is closed when this context is cancelled
     // or times out.
     // Done 方法返回一个 channel，当 context 取消或超时，Done 将关闭。
     Done() <-chan struct{}
 
     // Err indicates why this context was canceled, after the Done channel
     // is closed
     // 在 Done 关闭后，Err 可用于表明 context 被取消的原因
     Err() error
 
     // Deadline returns the time when this Context will be canceled, if any.
     // 到期则取消 context
     Deadline() (deadline time.Time, ok bool)
 
     // Value returns the value associated with key or nil if none
     Value(key interface{}) interface{}
 }
`