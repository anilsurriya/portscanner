package scanner

func (sj *scanJob) StartJob() []Result {
	sj.Generate()
	sj.StartWorkers()

	out := make([]Result, sj.UpperLimit-sj.LowerLimit)
	for result := range sj.Results {
		out = append(out, result)
	}

	return out
}
