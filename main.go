package main

func main() {
	even := make(chan bool)
	odd := make(chan bool)
	done := make(chan bool)
	//var wait sync.WaitGroup
	//wait.Add(2)
	go func() {
		for i := 0; i <= 10; i += 2 {
			<-even
			println(i)
			odd <- true
		}
		close(odd)
		close(even)
		done <- true
		// wait.Done()
	}()
	go func() {
		for i := 1; ; i += 2 {
			_, ok := <-odd
			if !ok {
				//wait.Done()
				return
			}
			println(i)
			select {
			case even <- true:
			case <-done:
				return
			}
		}
	}()
	even <- true
	//wait.Wait()
	<-done
}
