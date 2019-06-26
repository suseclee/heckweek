package main

//import "time"
import "runtime"
import "sync"

func main() {
    runtime.GOMAXPROCS(8)

    var wg sync.WaitGroup
    //doneCh := make(chan bool)

    go abcgen(wg)
    println("Main Function started")

    //<-doneCh



    println("%%%%%%%%%%%%%%%%%")
    ch := make(chan string)
    go abcGen(ch)
    go printer(ch)
    wg.Wait()
    //time.Sleep(3*time.Second)
}

func abcgen(wg sync.WaitGroup){

    for c := byte('a'); c <= byte('z'); c ++ {
        wg.Add(1)
        c := c
        go print("#" + string(c) + "#")
        wg.Done()
    }
    wg.Wait()
    //doneCh <- true

}

func abcGen(ch chan string){
    for c := byte('a'); c <= byte('z'); c ++ {
        ch <- string(c)
    }
    close(ch)
}

func printer(ch chan string) {

    //for {
    //   print(<-ch)
    //}

    for c := range ch {
        print(c)
    }
}
