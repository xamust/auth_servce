package passgen

import (
	"crypto/rand"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/sync/errgroup"
	"strings"
)

var (
	ErrIncorrectLength = fmt.Errorf("length must be greater than 0")
)

func randomString(opt *Options) (string, error) {
	if opt.passLen <= 0 {
		return "", ErrIncorrectLength
	}
	if strings.TrimSpace(opt.password) != "" {
		return opt.password, nil
	}
	result := make([]byte, opt.passLen)
	if _, err := rand.Read(result); err != nil {
		return "", err
	}
	for i := 0; i < opt.passLen; i++ {
		result[i] = opt.charset[int(result[i])%len(opt.charset)]
	}
	return string(result), nil
}

func GeneratePassword(options ...Option) (string, string, error) {
	opt := NewOptions(options...)
	password, err := randomString(opt)
	if err != nil {
		return "", "", err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), opt.cost)
	if err != nil {
		return "", "", err
	}
	return password, string(hashedPassword), nil
}

func GeneratePasswords(options ...Option) ([]string, []string, error) {
	var g errgroup.Group
	opt := NewOptions(options...)
	passwords := make([]string, opt.passCount)
	hashedPasswords := make([]string, opt.passCount)

	for i := 0; i < opt.passCount; i++ {
		index := i
		g.Go(func() error {
			password, hashedPassword, err := GeneratePassword(options...)
			if err != nil {
				return err
			}
			passwords[index] = password
			hashedPasswords[index] = hashedPassword
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return nil, nil, err
	}
	return passwords, hashedPasswords, nil
}
