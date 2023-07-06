package scanner

import (
	"net"
	"strconv"
	"sync"
)

func (sj *scanJob) StartWorkers() {
	var wg sync.WaitGroup
	for i := 0; i < sj.Poolsize; i++ {
		wg.Add(1)
		go sj.worker(&wg)
	}

	go func() {
		wg.Wait()
		close(sj.Results)
	}()
}

func (sj *scanJob) worker(wg *sync.WaitGroup) {
	addr := sj.Address
	protocol := sj.Protocol

	for port := range sj.ports {
		conn, err := net.Dial(protocol, addr+":"+strconv.Itoa(port))
		if err != nil {
			sj.Results <- Result{
				Status:   Closed,
				base:     addr,
				port:     port,
				ScanAddr: addr + ":" + strconv.Itoa(port),
			}
		} else {
			sj.Results <- Result{
				Status:   Open,
				base:     addr,
				port:     port,
				ScanAddr: addr + ":" + strconv.Itoa(port),
			}
			conn.Close()
		}
	}

	wg.Done()
}
