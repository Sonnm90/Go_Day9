package mutex

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var mu sync.Mutex
var count = 0

func process(n int) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(rand.Int31n(2)) * time.Second)
		mu.Lock()
		temp := count
		temp++
		time.Sleep(time.Duration(rand.Int31n(2)) * time.Second)
		count = temp
		mu.Unlock()
	}
	fmt.Println("Count after i="+strconv.Itoa(n)+" Count:", strconv.Itoa(count))

}

func Demo() {
	for i := 1; i < 4; i++ {
		go process(i)
	}

	//time.Sleep(25 * time.Second)
	fmt.Println("Final count", count)
}
