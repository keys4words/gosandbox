package main

func sender(sendChan chan<- int) {
	sendChan <- 10
}
func main() {
	sendChan := make(chan<- int)
	go sender(sendChan)
	// <-sendChan
}
