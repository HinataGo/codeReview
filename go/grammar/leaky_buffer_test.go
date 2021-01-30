package grammar

// 并发编程工具甚至可以使非并发思想更容易表达。
// 这是从RPC包中抽象出来的示例。
// 客户端goroutine循环从某个来源（可能是网络）接收数据。
// 为了避免分配和释放缓冲区，它会保留一个空闲列表，并使用一个缓冲的通道来表示它。
// 如果通道为空，则会分配一个新的缓冲区。
// 一旦消息缓冲区准备就绪，它将被发送到serverChan上的服务器。
var freeList = make(chan *Buffer, 100)
var serverChan = make(chan *Buffer)

func client() {
	for {
		var b *Buffer
		// Grab a buffer if available; allocate if not.
		select {
		case b = <-freeList:
			// Got one; nothing more to do.
		default:
			// None free, so allocate a new one.
			b = new(Buffer)
		}
		load(b)         // Read next message from the net.
		serverChan <- b // Send to server.
	}
}

// 服务器循环从客户端接收每个消息，对其进行处理，然后将缓冲区返回到空闲列表
func server() {
	for {
		b := <-serverChan // Wait for work.
		process(b)
		// Reuse buffer if there's room.
		select {
		case freeList <- b:
			// Buffer on free list; nothing more to do.
		default:
			// Free list full, just carry on.
		}
	}
}

// 客户端尝试从freeList检索缓冲区； 如果没有可用的，它将分配一个新的。
// 除非列表已满，否则服务器的send to freeList会将b放回到空闲列表中，在这种情况下，
// 缓冲区将被放置在地板上以由垃圾收集器回收。
// （select语句中的默认子句在没有其他情况下准备就绪时执行，这意味着selects永不阻塞。）
//  此实现仅依靠缓冲的通道和垃圾收集器进行簿记，仅用几行即可构建无泄漏的无桶列表。
