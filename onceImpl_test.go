package patterns

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOnce(t *testing.T) {

	once := NewOnce()
	//Give the res channel  capacity receive more
	resChan := make(chan int, 2)
	doneCh := make(chan struct{}, 5)

	defer func() {
		close(resChan)
		close(doneCh)
	}()

	for i := 0; i < 5; i++ {
		go func(ch chan int, done chan struct{}, i int) {
			once.Do(func() {
				ch <- i
			})
			done <- struct{}{}
		}(resChan, doneCh, i)
	}

	tom := []int{}
	for i := 0; i < 5; i++ {
		select {
		case v := <-resChan:
			tom = append(tom, v)
		case <-doneCh:
		}
	}

	assert.Equal(t, 1, len(tom))
}
