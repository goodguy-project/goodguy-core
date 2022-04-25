package subscribe

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestGetEmailBody(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		body, err := getEmailBody()
		if err != nil {
			t.Error(err)
		}
		err = ioutil.WriteFile("email-basic.html", []byte(body), 0644)
		if err != nil {
			t.Error(err)
		}
		fmt.Println(body)
	})
}
