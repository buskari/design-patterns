package main

import (
	"fmt"
	"sync"
)

var (
	singleInstance *single
	once           sync.Once
	mutex          sync.Mutex
)

type single struct{}

func GetSingleInstance() *single {
	// to avoid data racing
	mutex.Lock()
	defer mutex.Unlock()

	// ensure that the code inside the "Do" method will be executed only one time
	once.Do(func() {
		fmt.Println("Creating instance...")
		singleInstance = &single{}
	})

	return singleInstance
}

// Executing this code, you can see that the instance address will be the same every time the method GetSingleInstance is called
func main() {
	/**
		As we gonna use go routines to simulate data racing, we gonna need the WaitGroup to ensure
		that the program does not finish before all routines has been finished
	**/
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			singleInstance := GetSingleInstance()
			fmt.Printf("singleInstance Address: %p\n", singleInstance)
			wg.Done()
		}()
	}

	wg.Wait()
}
