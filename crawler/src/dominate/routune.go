package dominate

type Routune struct {
	channel chan int
}

func NewRoutune() *Routune {
	return &Routune{make(chan int, 50)}
}

func (self *Routune) SetMaxRoutineNum(num uint32) bool {
	self.channel = make(chan int, num)
	return true
}

func (self *Routune) Start(callback func()) bool {
	if nil == callback {
		return false
	}

	channl := self.channel
	for {
		channl <- 1
		go func() {
			defer func() {
				<-channl
			}()
			callback()
		}()
	}
	return true
}
