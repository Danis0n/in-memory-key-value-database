package compute

import (
	"strings"
)

const (
	foundLetterEvent     = 0
	foundWhiteSpaceEvent = 1
	eventId              = 2
)

const (
	initialState    = 0
	wordState       = 1
	whiteSpaceState = 2
	invalidState    = 3
	stateId         = 4
)

type transition struct {
	action     func(byte) int
	postAction func()
}

type StateMachine struct {
	transitions [stateId][eventId]transition
	sb          strings.Builder
	tokens      []string
	state       int
}

func newStateMachine() *StateMachine {

	machine := &StateMachine{
		state: initialState,
	}

	machine.transitions = [stateId][eventId]transition{
		initialState: {
			foundLetterEvent:     transition{action: machine.appendLetterJump},
			foundWhiteSpaceEvent: transition{action: machine.skipWhiteSpaceJump},
		},
		wordState: {
			foundLetterEvent:     transition{action: machine.appendLetterJump},
			foundWhiteSpaceEvent: transition{action: machine.skipWhiteSpaceJump, postAction: machine.addTokenAction},
		},
		whiteSpaceState: {
			foundLetterEvent:     transition{action: machine.appendLetterJump},
			foundWhiteSpaceEvent: transition{action: machine.skipWhiteSpaceJump},
		},
		invalidState: {},
	}

	return machine
}

func (sm *StateMachine) Parse(query string) ([]string, error) {
	for i := 0; i < len(query); i++ {
		symbol := query[i]
		if isWhiteSpace(symbol) {
			sm.proceedEvent(foundWhiteSpaceEvent, symbol)
		} else if isLetter(symbol) {
			sm.proceedEvent(foundLetterEvent, symbol)
		} else {
			return nil, errInvalidSymbol
		}
	}

	sm.proceedEvent(foundWhiteSpaceEvent, ' ')
	return sm.tokens, nil
}

func (sm *StateMachine) proceedEvent(eventId int, symbol byte) {
	t := sm.transitions[sm.state][eventId]
	sm.state = t.action(symbol)

	if t.postAction != nil {
		t.postAction()
	}
}

func (sm *StateMachine) appendLetterJump(b byte) int {
	sm.sb.WriteByte(b)
	return wordState
}

func (sm *StateMachine) skipWhiteSpaceJump(b byte) int {
	return whiteSpaceState
}

func (sm *StateMachine) addTokenAction() {
	sm.tokens = append(sm.tokens, sm.sb.String())
	sm.sb.Reset()
}

func isLetter(symbol byte) bool {
	return (symbol >= 'a' && symbol <= 'z') ||
		(symbol >= 'A' && symbol <= 'Z') ||
		(symbol >= '0' && symbol <= '9') ||
		(symbol == '_')
}

func isWhiteSpace(symbol byte) bool {
	return symbol == '\t' || symbol == '\n' || symbol == ' '
}
