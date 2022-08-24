package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/robfig/cron/v3"
)

// OtpConfig reads otps from a config file
type OtpConfig struct {
	otps []*Otp
	path string
	cron *cron.Cron
	lock sync.Mutex
}

func (o *OtpConfig) refresh() {
	o.lock.Lock()
	defer o.lock.Unlock()
	fmt.Println("refreshing")
	for _, otp := range o.otps {
		otp.refresh()
	}
}

func (o *OtpConfig) get() []Otp {
	o.lock.Lock()
	defer o.lock.Unlock()
	res := []Otp{}

	for _, otp := range o.otps {
		res = append(res, *otp)
	}
	return res
}

func (o *OtpConfig) resetCron() {
	if o.cron != nil {
		o.cron.Stop()
		o.cron = nil
	}
	o.cron = cron.New(cron.WithSeconds())
	o.cron.AddFunc("*/30 * * * * *", func() { o.refresh() })
	o.cron.Start()
}

func (o *OtpConfig) init(path string) error {

	o.lock.Lock()
	o.otps = []*Otp{}
	o.path = path
	o.resetCron()

	input, err := o.readLines(path)
	if err != nil {
		o.lock.Unlock()
		return err
	}

	for _, s := range input {
		name, key, found := strings.Cut(s, ":")
		if found == false {
			continue
		}
		o.otps = append(o.otps, &Otp{name: name, key: key})
	}
	o.lock.Unlock()
	o.refresh()

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
