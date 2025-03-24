package main

import (
        "fmt"
        "sync"
        "time"
)

func doch(x int, wg *sync.WaitGroup) {
        defer wg.Done()
        fmt.Println("Starting step",x)
        time.Sleep(2 * time.Second)
        fmt.Println("Finished step",x)
}

func main() {
        var wg sync.WaitGroup

        wg.Add(4)
        go doch(1, &wg)
        go doch(2, &wg)
        go doch(3, &wg)
        go doch(4, &wg)

        wg.Wait()
        fmt.Println("All steps complete.")

}