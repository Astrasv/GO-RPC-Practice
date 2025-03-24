package main
import(
	"fmt"
	"sync"
)

func fibo(n int, ch chan int, wg *sync.WaitGroup){
	defer wg.Done()

	var a int
	var b int
	var i int
	a=0
	b=1
	for i = 0;i<n;i++{
		ch <- a
		a,b = b,a+b
	}
	return
}



func main(){
	var wg sync.WaitGroup
	var ch = make(chan int)


	wg.Add(1)
	go fibo(10,ch,&wg)
	go func(){
		wg.Wait()
		close(ch)
	}()

	fmt.Println("Normal print",<-ch)
	for res := range ch{
		fmt.Println(res)
	}

}