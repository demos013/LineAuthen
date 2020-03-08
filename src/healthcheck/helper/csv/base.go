package csv

import (
	"bufio"
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strings"
)

func ReplaceSoloCarriageReturns(data io.Reader) io.Reader {
	return crlfReplaceReader{
		rdr: bufio.NewReader(data),
	}
}

// crlfReplaceReader wraps a reader
type crlfReplaceReader struct {
	rdr *bufio.Reader
}

// Read implements io.Reader for crlfReplaceReader
func (c crlfReplaceReader) Read(p []byte) (n int, err error) {
	if len(p) == 0 {
		return
	}

	for {
		if n == len(p) {
			return
		}

		p[n], err = c.rdr.ReadByte()
		if err != nil {
			return
		}

		// any time we encounter \r & still have space, check to see if \n follows
		// if next char is not \n, add it in manually
		if p[n] == '\r' && n < len(p) {
			if pk, err := c.rdr.Peek(1); (err == nil && pk[0] != '\n') || (err != nil && err.Error() == io.EOF.Error()) {
				n++
				p[n] = '\n'
			}
		}

		n++
	}
	return
}

// ReadFileCSV read data file csv
func ReadFileCSV(path string) ([][]string, error) {
	s := strings.Split(path, "/")
	fileName := s[len(s)-1]
	s = strings.Split(fileName, ".")
	fileType := s[len(s)-1]

	if fileType != "csv" {
		return nil, errors.New("Wrong of file type")
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	defer file.Close()

	reader := csv.NewReader(ReplaceSoloCarriageReturns(file))
	reader.Comma = ','
	reader.ReuseRecord = true
	reader.TrimLeadingSpace = true
	reader.FieldsPerRecord = -1

	rawCSVdata, err := reader.ReadAll()
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return rawCSVdata, nil
}
