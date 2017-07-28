package main

type stateMachine struct{
	currentState rune
	c            chan int
	s            chan rune
}

func (sm *stateMachine) send(i int) {
	sm.c <- i
}

func (sm stateMachine) state() rune {
	return <- sm.s
}

func (sm *stateMachine) loop() {
	for {
		select {
		// when we get a new state, process it:
		case v := <- sm.c:
			sm.currentState = sm.nextState(v)
		// when someone wants the current state, send it:
		case sm.s <- sm.currentState:
		}
	}
}

func (sm *stateMachine) nextState(operation int) rune {
	switch {
	case sm.currentState == 'A' && operation == 0:
		return 'A'
	case sm.currentState == 'A' && operation == 1:
		return 'B'
	case sm.currentState == 'B' && operation == 0:
		return 'C'
	case sm.currentState == 'B' && operation == 1:
		return 'A'
	case sm.currentState == 'C' && operation == 0:
		return 'B'
	case sm.currentState == 'C' && operation == 1:
		return 'A'
	}
	return 'X'
}

func newStateMachine() stateMachine {
	sm := stateMachine{
		c: make(chan int),
		s: make(chan rune),
		currentState: 'A',
	}

	go sm.loop()

	return sm
}

func main() {
	sm := newStateMachine()
	sm.send(1)
	sm.send(1)
	sm.send(1)
	sm.send(0)
	println(string(sm.state()))
}