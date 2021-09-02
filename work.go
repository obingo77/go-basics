 // Example provided with help from Jason Waldrip.
 // Package work manages a pool of goroutines to perform work.
 package work

 import "sync"

 // Worker must be implemented by types that want to use
 // the work pool.
 type Worker interface {
 Task()

 // Pool provides a pool of goroutines that can execute any Worker
 // tasks that are submitted.
 type Pool struct {
 work chan Worker
 wg sync.WaitGroup
 }

 // New creates a new work pool.
 func New(maxGoroutines int) *Pool {
 p := Pool{
 tasks: make(chan Worker),
 }

 p.wg.Add(maxGoroutines)
 for i := 0; i < maxGoroutines; i++ {
 go func() {
 for w := range p.work {
 w.Task()
 }
 p.wg.Done()
 }()
 }

 return &p
 }

 // Run submits work to the pool.
 func (p *Pool) Run(w Worker) {
 p.work <- w
 }

 // Shutdown waits for all the goroutines to shutdown.
 func (p *Pool) Shutdown() {
 close(p.tasks)
 p.wg.Wait()
 }
/*The work package in listing 7.28 starts off with the declaration of an interface named
Worker and a struct named Pool.
  Worker must be implemented by types that want to use
 */ the work pool.
 type Worker interface {
 Task()
 }

 // Pool provides a pool of goroutines that can execute any Worker
 // tasks that are submitted.
 type Pool struct {
 work chan Worker
 wg sync.WaitGroup
 }