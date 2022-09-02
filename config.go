package main

import (
	"bufio"
	"os"
	"strings"
)

// OtpConfig reads otps from a config file
type OtpConfig struct {
	otps Otps
	path string
}

func (o *OtpConfig) init(path string) error {

	o.path = path
	o.otps.resetCron()

	input, err := o.readLines(path)
	if err != nil {
		return err
	}

	for _, s := range input {
		name, key, found := strings.Cut(s, ":")
		if !found {
			continue
		}
		o.otps.add(name, key)
	}
	o.otps.refresh()

	return err
}

func (o *OtpConfig) readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
