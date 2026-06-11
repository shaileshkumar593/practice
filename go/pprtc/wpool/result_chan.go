type Job struct {
	ID int
}

type Result struct {
	JobID int
	Value int
}

func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		value := job.ID * job.ID // example computation
		results <- Result{JobID: job.ID, Value: value}
	}
}

func main() {
	jobs := make(chan Job, 10)
	results := make(chan Result, 10)
	var wg sync.WaitGroup

	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	for j := 1; j <= 10; j++ {
		jobs <- Job{ID: j}
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		fmt.Printf("Result: job %d => %d\n", res.JobID, res.Value)
	}
}

/*
	 Collects and processes worker results safely.
👉 Common in map-reduce, pipeline, and data processing systems.
*/