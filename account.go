package webfinger

import (
	"errors"
	"strings"
)

type account struct {
	Name     string
	Hostname string
}

func (a *account) ParseString(str string) (err error) {
	items := strings.Split(str, "@")
	if strings.HasPrefix(str, "acct:") {
		a.Name = items[0][5:]
	} else {
		a.Name = items[0]
	}

	if len(items) < 2 {
		//TODO: this might not be required
		err = errors.New("No domain on account")
		return
	}

	a.Hostname = strings.Split(items[1], "/")[0]

	return
}
