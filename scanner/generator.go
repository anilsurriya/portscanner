package scanner

func (sj *scanJob) Generate() error {
	go func() {
		defer close(sj.ports)
		for port := sj.LowerLimit; port <= sj.UpperLimit; port++ {
			sj.ports <- port
		}
	}()

	sj.Status = Active
	return nil
}
