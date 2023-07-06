package scanner

type Status int

const (
	Active = iota
	Pending
	ErrorStopped
	Open
	Closed
	TimeOut
)

type Result struct {
	ScanAddr string
	base     string
	port     int
	Status   Status
}

type scanJob struct {
	Protocol   string
	Address    string
	ports      chan int
	LowerLimit int
	UpperLimit int
	Poolsize   int
	Status     Status
	Results    chan Result
}

func NewJob(protocol, address string, lowerlimit, upperlimit int, poolsize int) *scanJob {
	return &scanJob{
		Protocol:   protocol,
		Address:    address,
		ports:      make(chan int),
		LowerLimit: lowerlimit,
		UpperLimit: upperlimit,
		Poolsize:   poolsize,
		Status:     Pending,
		Results:    make(chan Result),
	}
}

func (r *Result) String() string {
	ret := r.ScanAddr + "\t State: "
	switch r.Status {
	case Open:
		ret += "Open"
	case Closed:
		ret += "Closed"
	case TimeOut:
		ret += "Client timed out"
	}
	return ret
}
