package httpserver

import (
	"errors"
	"fmt"
)

type options struct {
	ip   string
	port int
}

func (o *options) getAddr() string {
	return fmt.Sprintf("%s:%d", o.ip, o.port)
}

// 설정 구조체를 업데이트하는 함수
type Option func(o *options) error

func WithIp(ip string) Option {
	return func(o *options) error {
		if ip == "" {
			o.ip = "" // default is localhost
		} else {
			o.ip = ip
		}
		return nil
	}
}

func WithPort(port int) Option {
	return func(o *options) error {
		if port < 0 {
			return errors.New("port can't be less than zero")
		} else if port == 0 {
			o.port = 3000
		} else {
			o.port = port
		}
		return nil
	}

}
