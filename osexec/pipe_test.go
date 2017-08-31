package main_test

import (
	"testing"
)

import (
	"io"
	// "io/ioutil"
	// "os"
	"bytes"
	"os/exec"
	"log"
)

type NoopPipe struct {
	log log.Logger
}
func (p *NoopPipe) Read(p []byte) (n int, err error) {
	p.log.Printf("NoopPipe.Read %s", string(p))
	return 1, nil
}

func (p *NoopPipe) Close() error {
	p.log.Print("NoopPipe.Close")
	return nil
}



func TestStdin(t *testing.T) {
	ls := exec.Command("ls")
	o, err := ls.StdoutPipe()
	if err != nil {
		t.Fatalf("Couldn't get stdout pipe %#v", err)
	}

	var pp io.ReadCloser

	p, err := exec.LookPath("peco")
	if err == nil {
		p1, perr := ls.
	}


	peco := exec.Command("peco")

	var pecobuf bytes.Buffer
	peco.Stdin = o
	peco.Stdout = &pecobuf

	if err := ls.Start(); err != nil {
		t.Fatalf("Failed to start ls command %#v", err)
	}

	if err := peco.Start(); err != nil {
		t.Fatalf("Failed to start peco command %#v", err)
	}

	if err := ls.Wait(); err != nil {
		t.Fatalf("Failed to wait %#v", err)
	}

	if err := peco.Wait(); err != nil {
		t.Fatalf("Failed to wait on peco %#v", err)
	}

	po := pecobuf.Bytes()

	t.Logf("peco output %s", string(po))
	t.Log("Done")
}
