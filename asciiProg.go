package main

import (
	"bytes"
	"strings"
)

func AsciiProg(input string, style string) string {
	result := ""
	content := getContent(style)
	var contentTable [][]byte
	if style == "thinkertoy" {
		contentTable = bytes.Split(content, []byte{13, 10})
	} else {
		contentTable = bytes.Split(content, []byte{10})

	}
	for _, c := range input {
		if c == 'é' || c == 'è' || c == 'ê' {
			input = strings.Replace(input, string(c), "e", -1)
		}
		if c == 'ç' {
			input = strings.Replace(input, string(c), "c", -1)
		}
		if c == 'à' || c == 'â' || c == 'ä' || c == 'á' {
			input = strings.Replace(input, string(c), "a", -1)
		}
		if c == 'ô' || c == 'ö' || c == 'ò' || c == 'ó' {
			input = strings.Replace(input, string(c), "o", -1)
		}
		if c == 'ù' || c == 'û' || c == 'ü' || c == 'ú' {
			input = strings.Replace(input, string(c), "u", -1)
		}
		if c == 'î' {
			input = strings.Replace(input, string(c), "i", -1)
		}
	}
	for _, c := range input {
		if (c < 20 || c >= 126) && c != 13 && c != 10 {
			input = "ERROR 500"
		}
	}
	input = strings.ReplaceAll(input, "\r\n", "\n")
	textToTransform := strings.Split(input, "\n")
	for k := 0; k < len(textToTransform); k++ {
		lines := make([]string, 8)
		for i := 0; i < 8; i++ {
			for j := 0; j < len(textToTransform[k]); j++ {
				startChar := ((int(textToTransform[k][j]) - 32) * 9) + 1

				lines[i] += string(contentTable[startChar+i])
			}
		}
		if lines[0] != "" {
			result += strings.Join(lines, "\n")
		}
		if k != len(textToTransform) {
			result += "\n"
		}
	}
	return result
}
