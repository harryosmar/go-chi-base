package encrypt_test

import (
	"github.com/harryosmar/go-chi-base/core/encrypt"
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestEncrypt(t *testing.T) {
	type args struct {
		plainPassword string
	}

	var testData = []struct {
		name string
		args args
	}{
		{
			name: "#testcase 1",
			args: args{
				plainPassword: "makcik",
			},
		},
	}
	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			ec := encrypt.MakeEncryptPassword()
			s, err := ec.Encrypt(tt.args.plainPassword)
			if err != nil {
				t.Error(err)
				return
			}
			log.Infof("encypted password %s", s)

			err = ec.Validate("$2a$04$KkEewy/XI4amB6EkZ78ku.J4hyQcpHQO7.XBBII6ujYZ8FzKoJpKq", "makcik")
			if err != nil {
				t.Error(err)
			}
		})
	}
}
