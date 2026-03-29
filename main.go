package main

import (
"fmt"
"log"
"math/rand"
""sync"
"time"
)

// Worker simulates a single GPU node in a distributed training setup.
func worker(id int, dataChannel <-chan int, resultChannel chan<- float64, wg *sync.WaitGroup) {
defer wg.Done()
log.Printf("Worker %d started\n", id)

for data := range dataChannel {
Simulate processing a batch of data
gTime := time.Duration(rand.Intn(500)) * time.Millisecond
gTime)

Simulate gradient calculation and model update
:= rand.Float64() * float64(data) / 100.0
tf("Worker %d processed data %d, loss: %.4f\n", id, data, loss)
nel <- loss
}

log.Printf("Worker %d finished\n", id)
}

func main() {
rand.Seed(time.Now().UnixNano())
numWorkers := 4
numDataBatches := 20

dataChannel := make(chan int, numDataBatches)
resultChannel := make(chan float64, numDataBatches)
var wg sync.WaitGroup

log.Printf("Starting distributed training with %d workers and %d data batches\n", numWorkers, numDataBatches)

// Start workers
for i := 1; i <= numWorkers; i++ {
worker(i, dataChannel, resultChannel, &wg)
}

// Send data to workers
for i := 1; i <= numDataBatches; i++ {
nel <- i
}
close(dataChannel)

// Wait for all workers to finish
wg.Wait()
close(resultChannel)

// Collect results (simulated losses)
totalLoss := 0.0
for loss := range resultChannel {
+= loss
}

fmt.Printf("\nDistributed training completed. Total simulated loss: %.4f\n", totalLoss)
fmt.Printf("Average simulated loss per batch: %.4f\n", totalLoss/float64(numDataBatches))
}
