package main

import (
	"fmt"
	"sync"

	"github.com/robfig/cron/v3"
)

type Otps struct {
	otps []*Otp
	lock sync.Mutex
	cron *cron.Cron
}

func (o *Otps) refresh() {
	o.lock.Lock()
	defer o.lock.Unlock()
	fmt.Println("refreshing")
	for _, otp := range o.otps {
		otp.refresh()
	}
}

func (o *Otps) get() []Otp {
	o.lock.Lock()
	defer o.lock.Unlock()
	res := []Otp{}

	for _, otp := range o.otps {
		res = append(res, *otp)
	}
	return res
}

func (o *Otps) resetCron() {
	if o.cron != nil {
		o.cron.Stop()
		o.cron = nil
	}
	o.cron = cron.New(cron.WithSeconds())
	o.cron.AddFunc("*/30 * * * * *", func() { o.refresh() })
	o.cron.Start()
}

func (o *Otps) add(name string, key string) {
	o.otps = append(o.otps, &Otp{name: name, key: key})
}
