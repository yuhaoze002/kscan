package ftp

import (
	"fmt"
	"github.com/jlaffaye/ftp"
	"time"
)

func Check(Host, Username, Password string, Port int) (bool, error) {
	conn, err := ftp.DialTimeout(fmt.Sprintf("%s:%d", Host, Port), 5*time.Second)
	if err != nil {
		return false, err
	}
	defer conn.Logout()
	err = conn.Login(Username, Password)
	if err != nil {
		return false, err
	}
	return true, nil
}
