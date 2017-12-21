package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type customReader struct {
	f     io.Reader
	slice []byte
	tmp   [1]byte
}

func (c *customReader) ReadByte() (byte, error) {
	c.slice = c.tmp[:1]
	_, err := io.ReadFull(c.f, c.slice)
	if err != nil {
		return 0, err
	}
	return c.slice[0], nil
}

func (c *customReader) Read(p []byte) (int, error) {
	return c.f.Read(p)
}

type customFileReader struct {
	r     *customReader
	slice []byte
	tmp   [256]byte
}

func (reader *customFileReader) Read(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	if len(reader.slice) == 0 {
		blockLenB, err := reader.r.ReadByte()
		if err != nil {
			return 0, err
		}

		blockLen, _ := strconv.Atoi(string(blockLenB))
		if blockLen == 0 {
			return 0, io.EOF
		}
		reader.slice = reader.tmp[0:blockLen]
		if _, err = io.ReadFull(reader.r, reader.slice); err != nil {
			return 0, err
		}
	}
	n := copy(p, reader.slice)
	reader.slice = reader.slice[n:]
	return n, nil
}

func main() {
	fileName := "./io.txt"
	r, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	cr := &customReader{f: r}

	cfr := &customFileReader{r: cr}
	buf := make([]byte, 4)
	for {
		n, err := cfr.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(buf[:n]))
	}
}
