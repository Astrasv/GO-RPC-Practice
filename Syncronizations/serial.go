package main

import (
        "fmt"
        "sync"
        "time"
)

func doch1(x int, ch chan bool, wg *sync.WaitGroup) {
        defer wg.Done()
        fmt.Println("Starting step",x)
        time.Sleep(2 * time.Second)
        fmt.Println("Finished step",x)
        ch <- true
}

func doch2(x int, ch chan bool, ch1 chan bool, wg *sync.WaitGroup) {
        defer wg.Done()
        val := <-ch1
        if val {
			fmt.Println("Starting step",x)
			time.Sleep(2 * time.Second)
			fmt.Println("Finished step",x)
                ch <- true
        }

}

func doch3(x int, ch chan bool, ch2 chan bool, wg *sync.WaitGroup) {
        defer wg.Done()
        val := <-ch2
        if val {
			fmt.Println("Starting step",x)
			time.Sleep(2 * time.Second)
			fmt.Println("Finished step",x)
                ch <- true
        }

}

func doch4(x int, ch chan bool, ch3 chan bool, wg *sync.WaitGroup) {
        defer wg.Done()
        val := <-ch3
        if val {
			fmt.Println("Starting step",x)
			time.Sleep(2 * time.Second)
			fmt.Println("Finished step",x)
                ch <- true
        }

}

func main() {
        var wg sync.WaitGroup
        var ch1 = make(chan bool, 1) 
        var ch2 = make(chan bool, 1)
        var ch3 = make(chan bool, 1)
        var ch4 = make(chan bool, 1)

        wg.Add(4)
        go doch1(1, ch1, &wg)
        go doch2(2, ch2, ch1, &wg)
        go doch3(3, ch3, ch2, &wg)
        go doch4(4, ch4, ch3, &wg)

        wg.Wait()
        fmt.Println("All steps complete.")

}