package submissions

type Foo struct {
	/*wgA sync.WaitGroup
	  wgB sync.WaitGroup*/
	doneOne chan bool
	doneTwo chan bool
	//doneThree chan bool
}

func NewFoo() *Foo {
	var foo = Foo{}

	/*foo.wgA.Add(1)
	  foo.wgB.Add(1)*/

	foo.doneOne = make(chan bool)
	foo.doneTwo = make(chan bool)
	//foo.doneThree = make(chan bool)

	return &foo
}

func (f *Foo) First(printFirst func()) {
	// Do not change this line
	printFirst()
	//f.wgA.Done()
	f.doneOne <- true
}

func (f *Foo) Second(printSecond func()) {
	//f.wgA.Wait()
	<-f.doneOne
	/// Do not change this line
	printSecond()
	f.doneTwo <- true
	//f.wgB.Done()
}

func (f *Foo) Third(printThird func()) {
	<-f.doneTwo
	//f.wgB.Wait()
	// Do not change this line
	printThird()
}
