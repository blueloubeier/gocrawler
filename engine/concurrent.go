package engine

type ConcurrentEngine struct{
	Scheduler Scheduler
	WorkerCount int
	ItemChan chan interface{}
}
type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)
}
func (e *ConcurrentEngine) Run(seeds ...Request){

	in := make(chan Request)
	out := make(chan ParseResult)
	e.Scheduler.ConfigureMasterWorkerChan(in)

	for i := 0; i < e.WorkerCount; i++{
		createWorker(in, out)
	}
	for _,r := range seeds{
		e.Scheduler.Submit(r)
	}

	for {
		result := <- out
		for _, item := range result.Items {
			//log.Printf("Got item #%d: %v", itemCount,item)
			go func(){ e.ItemChan <- item}()
		}
		for _, request := range result.Request{
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult){
	go func(){
		for{
			request := <- in
			result, err :=worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}