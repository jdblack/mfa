package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"strings"
	"time"
)

//Otp is an actaul Otp object
type Otp struct {
	name  string
	key   string
	token string
}

func (o *Otp) refresh() error {

	key, err := base32.StdEncoding.DecodeString(strings.ToUpper(o.key))
	if err != nil {
		return err
	}
	bs := make([]byte, 8)
	binary.BigEndian.PutUint64(bs, uint64(time.Now().Unix()/30))

	hash := hmac.New(sha1.New, key)
	hash.Write(bs)
	h := hash.Sum(nil)

	offset := (h[19] & 15)

	var header uint32
	r := bytes.NewReader(h[offset : offset+4])
	err = binary.Read(r, binary.BigEndian, &header)

	if err != nil {
		return err
	}
	h12 := (int(header) & 0x7fffffff) % 1000000

	o.token = fmt.Sprintf("%06d", h12)
	return nil
}
