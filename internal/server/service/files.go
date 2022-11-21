package server

import (
	"bytes"
	"fmt"
)

func (g *GRPCServer) FileSave(fileName string, fileData bytes.Buffer) error {
	_, err := fileData.WriteTo(file)
	if err != nil {
		return fmt.Errorf("cannot write file: %w", err)
	}

	return nil
}
