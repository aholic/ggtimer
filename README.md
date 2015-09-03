# ggtimer

ggtimer is a cancelable timer(ticker) in golang. A timer(ticker) in standard lib of golang is not cancelable. That is, when you are blocked on a channel of a timer(ticker), you are not able to stop it. The timer.Stop() doesn't close the channel. The following code will result an error.

    t := time.NewTimer(3 * time.Second)
    done := make(chan bool, 1)

    go func() {
        <-t.C
        fmt.Println("after 3 seconds")
        done <- true
    }()
    
    t.Stop()
    <-done
    
The implementation of ggtimer is easy, for more information, refer to [here](http://cstdlib.com/tech/2015/08/17/golang-timer/)
