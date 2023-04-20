package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func sumSequentialNumbers(num []int) int {
	var sum int
	for _, value := range num {
		sum += value
	}
	return sum
}

func sumNumbersConcurrency(num []int, wg *sync.WaitGroup, sumChannel chan int) {
	defer wg.Done()
	var sum int
	for _, value := range num {
		sum += value
	}

	sumChannel <- sum
}

func main() {
	numbers := make([]int, 2000000)

	for i := 0; i < len(numbers); i++ {
		numbers[i] = i + 1
	}

	start := time.Now()
	//Secuencial
	result := sumSequentialNumbers(numbers)
	fmt.Printf("La suma de los primeros %d números es: %d \n", len(numbers), result)
	duration := time.Since(start)
	fmt.Println("La duración de la suma en forma secuencial es: ", duration)

	//Concurrente
	runtime.GOMAXPROCS(runtime.NumCPU())

	var wg sync.WaitGroup
	sumChannel := make(chan int)
	lenght := len(numbers) / 2
	wg.Add(2)

	start2 := time.Now()
	go sumNumbersConcurrency((numbers[:lenght]), &wg, sumChannel)
	go sumNumbersConcurrency((numbers[lenght:]), &wg, sumChannel)
	sum1, sum2 := <-sumChannel, <-sumChannel
	total := sum1 + sum2
	fmt.Printf("La suma de los primeros %d números es: %d \n", len(numbers), total)
	wg.Wait()
	durationC := time.Since(start2)
	fmt.Println("La duración de la suma en forma concurrente es: ", durationC)
}
