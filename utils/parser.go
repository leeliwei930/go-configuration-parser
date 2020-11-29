package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

// Parser type
type Parser struct {
	file            *os.File
	parseContent    *map[string]map[string]string
	lastReadSection string
}

// MakeParser create a parser instance
func MakeParser(filename string) (*Parser, error) {
	file, err := os.Open(filename)
	if err == nil {
		return &Parser{file, nil, ""}, err
	}

	defer file.Close()

	return nil, err
}

// HasFile determine file read is success
func (parser *Parser) HasFile() bool {
	return parser.file != nil
}

// PrintContent
func (parser *Parser) PrintContent() {
	for sectionKey, sectionContent := range *parser.parseContent {
		fmt.Printf("Section %s", sectionKey)
		fmt.Println()

		for key, value := range sectionContent {
			fmt.Printf("%s -> %s", key, value)
			fmt.Println()
		}
		fmt.Println()

	}
}

// Parse - perform data parsing
func (parser *Parser) Parse() {
	reader := bufio.NewScanner(parser.file)
	var parsedContent = make(map[string]map[string]string)
	keyVal := make(map[string]string)

	for reader.Scan() {
		var text string = reader.Text()
		isSectionText, sectionText := isSection(text)
		isKeyVal, key, val := isKeyVal(text)

		if isSectionText {
			parser.lastReadSection = sectionText
			keyVal = make(map[string]string)

		}
		if isKeyVal {
			keyVal[key] = val
		}
		parsedContent[parser.lastReadSection] = keyVal

	}
	parser.parseContent = &parsedContent
}

func isSection(text string) (bool, string) {
	var characterCount int = utf8.RuneCountInString(text)
	if characterCount >= 2 {
		if text[0] == '[' && text[characterCount-1] == ']' {
			return true, text[1 : characterCount-1]
		}
	}
	return false, ""
}

func isKeyVal(text string) (bool, string, string) {
	var characterCount int = utf8.RuneCountInString(text)
	if characterCount >= 3 && strings.Contains(text, "=") {
		spIndex := strings.Index(text, "=")
		return true, text[0:spIndex], text[spIndex+1:]
	}

	return false, "", ""
}

// GetParseContent - return the parsed content
func (parser *Parser) GetParseContent() *map[string]map[string]string {
	return parser.parseContent
}
