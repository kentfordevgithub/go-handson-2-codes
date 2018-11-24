package main

func main() {
	ch := make(chan int)

	v := 1
	go func() {
		ch <- (int)(v += 1)
	}()
	<-ch
}
