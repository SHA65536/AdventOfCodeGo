package helper

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type InputReader struct {
	file   *os.File
	reader *bufio.Reader
}

func NewInputReader(filename string) (*InputReader, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}

	return &InputReader{
		file:   file,
		reader: bufio.NewReader(file),
	}, nil
}

func (ir *InputReader) Close() error {
	return ir.file.Close()
}

func (ir *InputReader) ReadWord() (string, error) {
	var word string
	_, err := fmt.Fscan(ir.reader, &word)
	if err != nil {
		return "", err
	}
	return word, nil
}

func (ir *InputReader) ReadLine() (string, error) {
	line, err := ir.reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return "", err
	}
	return line, nil
}

func (ir *InputReader) ReadLines() ([]string, error) {
	var lines []string
	for {
		line, err := ir.reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return nil, err
		}
		lines = append(lines, strings.TrimSuffix(line, "\n"))
		if err == io.EOF {
			break
		}
	}
	return lines, nil
}

func (ir *InputReader) ReadNChars(n int) (string, error) {
	buf := make([]byte, n)
	_, err := io.ReadFull(ir.reader, buf)
	if err != nil && err != io.EOF {
		return "", err
	}
	return string(buf), nil
}

func (ir *InputReader) ReadAll() (string, error) {
	content, err := io.ReadAll(ir.reader)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (ir *InputReader) IterateLines(yield func(string) bool) {
	for {
		line, err := ir.reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return
		}
		if !yield(strings.TrimRight(line, "\n")) {
			return
		}
		if err == io.EOF {
			break
		}
	}
}

func (ir *InputReader) IterateWords(yield func(string) bool) {
	for {
		line, err := ir.reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return
		}
		for _, word := range strings.Split(line, " ") {
			if !yield(word) {
				return
			}
		}
		if !yield("\n") {
			return
		}
		if err == io.EOF {
			break
		}
	}
}

func MustConvNum(in string) int {
	res, _ := strconv.Atoi(in)
	return res
}
