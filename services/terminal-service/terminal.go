package terminalservice

import (
	"io"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/creack/pty"
	"golang.org/x/term"
)

type Terminal struct {
	pid      int
	ptmx     *os.File
	ch       chan os.Signal
	incoming chan []byte
	outgoing chan []byte
}

func NewTerminal() *Terminal {
	// create a new terminal
	c := exec.Command("/bin/bash")

	// Start the command with a pty
	ptmx, err := pty.Start(c)
	if err != nil {
		log.Println(err, "Failed to start pty")
	}

	// Handle pty size.
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGWINCH)
	go func() {
		for range ch {
			if err := pty.InheritSize(os.Stdin, ptmx); err != nil {
				log.Printf("error resizing pty: %s", err)
			}
		}
	}()
	ch <- syscall.SIGWINCH // Initial resize.

	// Set stdin in raw mode.
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer func() { _ = term.Restore(int(os.Stdin.Fd()), oldState) }() // Best effort.

	terminal := &Terminal{
		pid:      c.Process.Pid,
		ptmx:     ptmx,
		ch:       ch,
		incoming: make(chan []byte),
		outgoing: make(chan []byte),
	}

	go terminal.handleIO()

	return terminal
}

func (t *Terminal) handleIO() {
	go func() {
		for data := range t.incoming {
			_, _ = t.ptmx.Write(data)
		}
	}()

	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := t.ptmx.Read(buf)
			if err != nil {
				if err == io.EOF {
					close(t.outgoing)
					return
				}
				log.Printf("error reading from pty: %s", err)
				continue
			}
			t.outgoing <- buf[:n]
		}
	}()
}

func (t *Terminal) Close() error {
	close(t.incoming)
	return t.ptmx.Close()
}
