package pkg

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type IntegerField struct {
	Name  string
	Value int
}

type BlobField struct {
	Name string
	Data []byte
}

type Field struct {
	Type string
	Int  *IntegerField
	Blob *BlobField
}

func ReadFields(filename string) ([]Field, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var fields []Field
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, ":i ") {
			parts := strings.SplitN(line[3:], " ", 2)
			if len(parts) != 2 {
				return nil, fmt.Errorf("invalid integer field format: %s", line)
			}
			value, err := strconv.Atoi(parts[1])
			if err != nil {
				return nil, fmt.Errorf("invalid integer value: %s", parts[1])
			}
			fields = append(fields, Field{
				Type: "integer",
				Int: &IntegerField{
					Name:  parts[0],
					Value: value,
				},
			})
		} else if strings.HasPrefix(line, ":b ") {
			parts := strings.SplitN(line[3:], " ", 2)
			if len(parts) != 2 {
				return nil, fmt.Errorf("invalid blob field format: %s", line)
			}
			size, err := strconv.Atoi(parts[1])
			if err != nil {
				return nil, fmt.Errorf("invalid blob size: %s", parts[1])
			}
			data := make([]byte, size)
			_, err = io.ReadFull(reader, data)
			if err != nil {
				return nil, fmt.Errorf("failed to read blob data: %v", err)
			}
			_, err = reader.ReadString('\n')
			if err != nil {
				return nil, fmt.Errorf("failed to read trailing newline for blob: %v", err)
			}
			fields = append(fields, Field{
				Type: "blob",
				Blob: &BlobField{
					Name: parts[0],
					Data: data,
				},
			})
		} else {
			return nil, fmt.Errorf("unknown field type: %s", line)
		}
	}

	return fields, nil
}

func WriteFields(filename string, fields []Field) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, field := range fields {
		if field.Type == "integer" {
			_, err := writer.WriteString(fmt.Sprintf(":i %s %d\n", field.Int.Name, field.Int.Value))
			if err != nil {
				return err
			}
		} else if field.Type == "blob" {
			_, err := writer.WriteString(fmt.Sprintf(":b %s %d\n", field.Blob.Name, len(field.Blob.Data)))
			if err != nil {
				return err
			}
			_, err = writer.Write(field.Blob.Data)
			if err != nil {
				return err
			}
			_, err = writer.WriteString("\n")
			if err != nil {
				return err
			}
		}
	}
	return writer.Flush()
}
