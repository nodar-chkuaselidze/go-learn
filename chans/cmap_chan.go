package main

type ReadChan struct {
	ChanBack chan int
	Key      int
}

type cMapChanInt struct {
	store   map[int]int
	writeCh chan [2]int
	delCh   chan int
	readCh  chan ReadChan
	done    chan struct{}
}

func (c *cMapChanInt) watch() {
	for {
		select {
		case w := <-c.writeCh:
			c.store[w[0]] = w[1]
		case d := <-c.delCh:
			delete(c.store, d)
		case r := <-c.readCh:
			r.ChanBack <- c.store[r.Key]
			close(r.ChanBack)
		case <-c.done:
			return
		}
	}
}

func (c *cMapChanInt) set(key, val int) {
	c.writeCh <- [2]int{key, val}
}

func (c *cMapChanInt) get(key int) int {
	reader := ReadChan{
		ChanBack: make(chan int, 1),
		Key:      key,
	}

	c.readCh <- reader

	r := <-reader.ChanBack
	return r
}

func (c *cMapChanInt) del(key int) {
	c.delCh <- key
}

func newCapChanInt() *cMapChanInt {
	cmap := &cMapChanInt{
		store:   make(map[int]int),
		writeCh: make(chan [2]int),
		delCh:   make(chan int),
		readCh:  make(chan ReadChan),
	}

	go cmap.watch()

	return cmap
}
