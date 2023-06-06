package coloring

import (
	. "GraphColoring/graph"
	"fmt"
	"math"
	"sync"
)

var CONF_PER_WORKER = 10000

func FindColoring(g IGraph, k int) (Result, error) {

	var res Result

	nodesNumber := len(g.Nodes())
	configurationsNumber := int(math.Pow(float64(k), float64(nodesNumber))) - 1
	lastInterval, err := ToBaseArray(configurationsNumber, k)
	if err != nil {
		fmt.Println("error getting the last interval: ", err)
		return MakeResult(false, nil), err
	}

	digitsNumber := len(lastInterval)
	workersNumber := int(math.Ceil(float64(configurationsNumber) / float64(CONF_PER_WORKER)))

	done := make(chan Result, workersNumber)
	defer close(done)
	stop := make(chan bool)

	var wg sync.WaitGroup
	wg.Add(workersNumber)

	var mu sync.Mutex

	for i := 0; i < workersNumber; i++ {
		intervalStart, err := ToBaseArrayFixed(i*CONF_PER_WORKER, k, digitsNumber)
		if err != nil {
			fmt.Println("error in converting interval start to base k")
			return MakeResult(false, nil), err
		}
		intervalEnd, err := ToBaseArrayFixed(int(math.Min(float64((i+1)*CONF_PER_WORKER-1), float64(configurationsNumber))), k, digitsNumber)
		if err != nil {
			fmt.Println("error in converting interval end to base k")
			return MakeResult(false, nil), err
		}

		go FindColoringWorker(g, intervalStart, intervalEnd, k, digitsNumber, done, stop, &wg, &mu)
	}

	wg.Wait()

	select {
	case res = <-done:
	default:
		res = MakeResult(false, nil)
	}

	return res, nil
}
