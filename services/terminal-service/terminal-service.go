package terminalservice

import (
	"log"
	"sync"
)

type TerminalService struct {
	mu        sync.Mutex
	terminals map[string]*Terminal
}

func NewTerminalService() *TerminalService {
	return &TerminalService{
		terminals: make(map[string]*Terminal),
	}
}

func (ts *TerminalService) CreateTerminal(id string) *Terminal {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	terminal := NewTerminal()
	ts.terminals[id] = terminal
	return terminal
}

func (ts *TerminalService) DeleteTerminal(id string) {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	if terminal, ok := ts.terminals[id]; ok {
		if err := terminal.Close(); err != nil {
			log.Printf("error closing terminal: %s", err)
		}
		delete(ts.terminals, id)
	}
}

func (ts *TerminalService) WriteTerminal(id string, data []byte) {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	if terminal, ok := ts.terminals[id]; ok {
		terminal.incoming <- data
	}

	// read from terminal
	if terminal, ok := ts.terminals[id]; ok {
		go func() {
			for data := range terminal.outgoing {
				log.Printf("TerminalService: %s", data)
			}
		}()
	}
}

// func (ts *TerminalService) ReadTerminal(id string) <-chan []byte {
//     ts.mu.Lock()
//     defer ts.mu.Unlock()

//     if terminal, ok := ts.terminals[id]; ok {
//         return terminal.outgoing
//     }
//     return nil
// }
