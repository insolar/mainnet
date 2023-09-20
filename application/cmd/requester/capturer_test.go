package main

import (
	"bufio"
	"io"
	"os"
	"time"
)

// Capturer has flags whether capture stdout/stderr or not.
type Capturer struct {
	captureStdout bool
	captureStderr bool
}

// CaptureStdout captures stdout.
func CaptureStdout(f func(), d time.Duration) (string, error) {
	capturer := &Capturer{captureStdout: true}
	return capturer.capture(f, d)
}

// CaptureStderr captures stderr.
func CaptureStderr(f func(), d time.Duration) (string, error) {
	capturer := &Capturer{captureStderr: true}
	return capturer.capture(f, d)
}

// CaptureOutput captures stdout and stderr.
func CaptureOutput(f func(), d time.Duration) (string, error) {
	capturer := &Capturer{captureStdout: true, captureStderr: true}
	return capturer.capture(f, d)
}

func (capturer *Capturer) capture(fn func(), duration time.Duration) (string, error) {
	r, w, err := os.Pipe()
	if err != nil {
		return "", err
	}

	err = r.SetReadDeadline(time.Now().Add(duration))
	if err != nil {
		return "", err
	}

	if capturer.captureStdout {
		stdout := os.Stdout
		os.Stdout = w
		defer func() {
			os.Stdout = stdout
		}()
	}

	if capturer.captureStderr {
		stderr := os.Stderr
		os.Stderr = w
		defer func() {
			os.Stderr = stderr
		}()
	}

	defer w.Close()

	fn()

	getPipeResultFunction := func() interface{} {
		defer r.Close()
		reader := bufio.NewReader(r)
		line, _, e := reader.ReadLine()
		if e == io.EOF {
			return result{string(line), nil}
		} else if e != nil {
			return result{"", e}
		}
		return result{string(line), e}
	}

	res, err := waitForFunction(getPipeResultFunction, duration*2)
	if err != nil {
		return "", err
	}
	return res.(result).str, res.(result).err
}

type result struct {
	str string
	err error
}
