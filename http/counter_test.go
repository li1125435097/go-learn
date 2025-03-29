// counter_test.go
package main

import (
    "sync"
    "testing"
    "fmt"
)

func TestConcurrentCounter(t *testing.T) {
    c := &Counter{}
    var wg sync.WaitGroup
    n := 1000

    wg.Add(n)
    for i := 0; i < n; i++ {
        go func() {
            c.Inc()
            wg.Done()
        }()
    }
    wg.Wait()

    if c.Value() != n {
        t.Errorf("Counter value = %d; want %d", c.Value(), n)
    }
}

func TestC(t *testing.T){
    fmt.Println("1325646")
}
