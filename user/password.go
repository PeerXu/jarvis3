package user

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/rand"
	"strings"
)

type EncryptedPassword string

func genBase64RandomString(n int) string {
	buf := make([]byte, n)
	rand.Read(buf)
	s := base64.StdEncoding.EncodeToString(buf)
	return s
}

func genSalt() string {
	salt := genBase64RandomString(32)
	return salt
}

func genTokenString() string {
	token := genBase64RandomString(64)
	return token
}

func NewEncryptedPassword(password string) EncryptedPassword {
	salt := genSalt()
	return NewEncryptedPasswordWithSalt(password, salt)
}

func NewEncryptedPasswordWithSalt(password, salt string) EncryptedPassword {
	mac := hmac.New(sha256.New, []byte(salt))
	mac.Write([]byte(password))
	encryptedBytes := mac.Sum(nil)
	encryptedStr := base64.StdEncoding.EncodeToString(encryptedBytes)
	passwdStr := fmt.Sprintf("%s$%s", encryptedStr, salt)
	return EncryptedPassword(passwdStr)
}

func (ep EncryptedPassword) equal(o EncryptedPassword) bool {
	return ep == o
}

func (ep EncryptedPassword) split() (string, string) {
	s := string(ep)
	t := strings.Split(s, "$")
	encrypt, salt := t[0], t[1]
	return encrypt, salt
}

func (ep EncryptedPassword) Salt() string {
	_, salt := ep.split()
	return salt
}

func (ep EncryptedPassword) Validate(password string) bool {
	salt := ep.Salt()
	o := NewEncryptedPasswordWithSalt(password, salt)
	return ep.equal(o)
}

func (ep EncryptedPassword) String() string {
	return string(ep)
}
