# High-Throughput HTTP Health Monitor

A concurrent CLI tool written in Go to benchmark the availability and latency of distributed web endpoints. 

Unlike traditional sequential health checkers, this tool utilizes Go's **concurrency primitives (Goroutines and Channels)** to process multiple requests in parallel, reducing total execution time from $O(N)$ to approximately $O(1)$ (limited only by the slowest response).

## 🚀 Features

* **Concurrency**: Spawns lightweight threads (Goroutines) for each URL check.
* **Thread Safety**: Uses the **CSP (Communicating Sequential Processes)** model. Data is shared by communicating via Channels, eliminating the need for complex mutex locks on shared memory.
* **Synchronization**: Implements `sync.WaitGroup` to manage worker lifecycles.
* **Non-Blocking Collection**: Uses a dedicated "Closer" goroutine to signal completion, allowing the main thread to process results as they arrive in real-time.
* **Race Detector Verified**: Validated with `go run -race` to ensure zero race conditions.

## 🛠️ Architecture

The system follows a **Fan-Out / Fan-In** pattern:

1.  **Fan-Out**: The main process iterates over the input list and spawns a worker Goroutine for each URL.
2.  **Process**: Each worker independently handles the HTTP request and latency calculation.
3.  **Fan-In**: All workers push their `Result` structs into a shared, buffered channel.
4.  **Aggregation**: The main function iterates over the channel, printing results as they arrive, and exits gracefully only when the "Closer" routine signals that all workers are done.

```mermaid
graph LR
    Main[Main Process] -->|Spawns| Worker1[Worker: Google]
    Main -->|Spawns| Worker2[Worker: GitHub]
    Main -->|Spawns| Worker3[Worker: Reddit]
    
    Worker1 -->|Send Result| Channel{Result Channel}
    Worker2 -->|Send Result| Channel
    Worker3 -->|Send Result| Channel
    
    Closer[Closer Routine] -.->|Wait & Close| Channel
    
    Channel -->|Stream| Output[Console Output]
