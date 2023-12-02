package utils

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func LoadInputFile(filepath string) []string {
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	var rows []string
	s := bufio.NewScanner(f)
	for s.Scan() {
		err = s.Err()
		if err != nil {
			log.Panic(err)
		}
		rows = append(rows, s.Text())
	}
	return rows
}

func LoadInputFileParts(filepath string) [][]string {
	rows := LoadInputFile(filepath)

	var parts [][]string
	var part []string
	for i := 0; i < len(rows); i++ {
		if rows[i] == "" {
			parts = append(parts, part)
			part = []string{}
			continue
		}
		part = append(part, rows[i])
	}
	parts = append(parts, part)
	return parts
}

func LoadInputFileChan(filepath string) chan string {
	rows := make(chan string)
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	s := bufio.NewScanner(f)
	for s.Scan() {
		err = s.Err()
		if err != nil {
			log.Panic(err)
		}
		rows <- s.Text()
	}
	close(rows)
	return rows
}

func LoadInputFromUrl(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []string{}, fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []string{}, fmt.Errorf("status error: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return []string{}, fmt.Errorf("read body: %v", err)
	}

	return strings.Split(string(data), "\n"), nil
}

func RegexMatchAll(rg *regexp.Regexp, input string) (paramsMap map[string]string) {
	match := rg.FindStringSubmatch(input)
	paramsMap = make(map[string]string)
	for i, name := range rg.SubexpNames() {
		if i > 0 && i <= len(match) {
			paramsMap[name] = match[i]
		}
	}
	return paramsMap
}
