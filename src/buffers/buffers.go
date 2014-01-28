package buffers

var (
	largeMin = 1024
)

var (
	smallBuffers = make(chan []byte, 32)
	largeBuffers = make(chan []byte, 32)
)

func GetBuffer(size int) []byte {
	var ch = smallBuffers
	if size > largeMin {
		ch = largeBuffers
	}

	var buf []byte
	select {
	case buf = <-ch:
	default:
	}

	if len(buf) < size {
		buf = make([]byte, size)
	}

	return buf[:size]
}

func PutBuffer(buf []byte) {
	buf = buf[:cap(buf)]

	if len(buf) <= 0 {
		return
	}

	var ch = smallBuffers
	if len(buf) < largeMin {
		ch = largeBuffers
	}

	select {
	case ch <- buf:
	default:
	}
}
