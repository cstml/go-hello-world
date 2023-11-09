package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"regexp"
	"sync"
	"testing"
)

// Generic abstraction of a test generator that captures the output of a
// function `f` and compares it to an expected regex pattern.
func testOut(t *testing.T, pattern string, f func()) {
	// Open Pipe
	reader, writer, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	// Store og outputs + change them with the pipe
	ogErr, ogOut := os.Stdout, os.Stderr
	os.Stdout = writer
	os.Stderr = writer
	log.SetOutput(writer)

	// call function
	f()

	// Read what is in the buffer
	out := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		var buf bytes.Buffer
		wg.Done()
		io.Copy(&buf, reader)
		out <- buf.String()
	}()
	wg.Wait()
	writer.Close()

	// Change back the outs/ins.
	log.SetOutput(ogOut)
	os.Stdout = ogOut
	os.Stderr = ogErr

	// read the value & match
	val := <-out
	match, err := regexp.MatchString(pattern, val)

	if err != nil {
		t.Error("Did not find expected string.")
		return
	}

	if match != true {
		t.Error("No match found.")
		return
	}

	return
}

func Test_SayHello(t *testing.T) {
	testOut(t, ".*!", func() { SayHello("Bob") })
}

func Test_Main(t *testing.T) {
	testOut(t, "Hello World!", main)
}
