package processor

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/mpv/kir/yamlparser"
)

func ProcessStdin() ([]string, error) {
	reader := bufio.NewReader(os.Stdin)
	var data []byte
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil && err != io.EOF {
			return nil, fmt.Errorf("error reading stdin: %v", err)
		}
		data = append(data, line...)
		if err == io.EOF {
			break
		}
	}
	return yamlparser.ProcessData(data)
}

func ProcessFile(filePath string) ([]string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}
	return yamlparser.ProcessData(data)
}
