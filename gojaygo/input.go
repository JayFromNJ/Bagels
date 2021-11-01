package gojaygo

import (
	"bufio"
	"os"
	"strings"
)

type CLInput struct {
	CLReader bufio.Reader
}

func GetNewCLReader() CLInput {
	return CLInput{
		CLReader: *bufio.NewReader(os.Stdin),
	}
}

func (cli *CLInput) GetCLInput() (string, error) {
	toReturn, err := cli.CLReader.ReadString('\n')

	if err != nil {
		return "", err
	}

	toReturn = strings.TrimRight(toReturn, "\r\n")
	return toReturn, nil
}
