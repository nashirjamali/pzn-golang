package channel_test

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)

	// Mengirim data ke channel
	// channel <- "Nashir"

	// Ngambil data dari channel
	// data := <-channel

	// Menutup channel
	// defer close(channel)

	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Nashir Jamali"
		fmt.Println("Selesai mengirim data")
	}()

	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}
