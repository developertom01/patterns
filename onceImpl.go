package patterns

type once chan struct{}

func NewOnce() once {
	o := make(chan struct{}, 1)
	o <- struct{}{}
	return o
}

func (o once) Do(f func()) {
	_, ok := <-o

	if !ok {
		return
	}
	f()
	close(o)
}
