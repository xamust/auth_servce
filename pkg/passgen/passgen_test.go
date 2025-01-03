package passgen

import (
	"github.com/go-faster/errors"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"testing"
)

func Test_randomString(t *testing.T) {
	tests := []struct {
		name       string
		opt        []Option
		err        error
		wantResult string
		wantErr    bool
	}{
		{
			name: "correct length",
			opt: []Option{
				WithPassLen(10),
			},
			wantErr: false,
		},
		{
			name: "zero length",
			opt: []Option{
				WithPassLen(0),
			},
			err:     ErrIncorrectLength,
			wantErr: true,
		},
		{
			name: "negative length",
			opt: []Option{
				WithPassLen(-1),
			},
			err:     ErrIncorrectLength,
			wantErr: true,
		},
		{
			name: "custom charset",
			opt: []Option{
				WithCharset("a"),
				WithPassLen(10),
			},
			wantResult: "aaaaaaaaaa",
			wantErr:    false,
		},
		{
			name: "default password",
			opt: []Option{
				WithDefaultPassword("Pa$$word"),
			},
			wantResult: "Pa$$word",
			wantErr:    false,
		},
		{
			name: "default password, with incorrect length",
			opt: []Option{
				WithDefaultPassword("Pa$$word"),
				WithPassLen(3),
			},
			wantResult: "Pa$$word",
			wantErr:    false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			opt := NewOptions(test.opt...)
			got, err := randomString(opt)
			if test.wantErr {
				if err == nil {
					t.Errorf("expected error: %v, got nil", test.err)
				}
				if !errors.Is(err, test.err) {
					t.Errorf("expected error: %v, got: %v", test.err, err)
				}
				return
			}
			if strings.TrimSpace(test.wantResult) != "" && test.wantResult != got {
				t.Errorf("expected: %v, got: %v", test.wantResult, got)
			}

			if opt.passLen != len(got) {
				t.Errorf("expected: %v, got: %v", opt.passLen, len(got))
			}

			if err != nil {
				t.Errorf("expected nil, got: %v", err)
			}

			t.Logf("randomString() = %v", got)
		})
	}
}

func TestGeneratePassword(t *testing.T) {
	tests := []struct {
		name       string
		opt        []Option
		err        error
		wantResult string
		wantErr    bool
	}{
		{
			name: "correct passwd",
			opt: []Option{
				WithPassLen(10),
			},
			wantErr: false,
		},
		{
			name: "zero length",
			opt: []Option{
				WithPassLen(0),
			},
			err:     ErrIncorrectLength,
			wantErr: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			opt := NewOptions(test.opt...)
			password, hashedPassword, err := GeneratePassword(test.opt...)
			if test.wantErr {
				if err == nil {
					t.Errorf("expected error: %v, got nil", test.err)
				}
				if !errors.Is(err, test.err) {
					t.Errorf("expected error: %v, got: %v", test.err, err)
				}
				return
			}
			if err != nil {
				t.Errorf("expected nil, got: %v", err)
			}

			if opt.passLen != len(password) {
				t.Errorf("expected %v, got: %v", opt.passLen, len(password))
			}

			if strings.TrimSpace(hashedPassword) == "" {
				t.Error("expected not empty")
			}

			if err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
				t.Errorf("expected nil, got: %v", err)
			}
			t.Logf("password: %v, hashedPassword: %v", password, hashedPassword)
		})
	}

}

func TestGeneratePasswords(t *testing.T) {
	tests := []struct {
		name       string
		opt        []Option
		err        error
		wantResult string
		wantErr    bool
	}{
		{
			name: "correct passwd",
			opt: []Option{
				WithPassLen(10),
				WithPassCount(5),
			},
			wantErr: false,
		},
		{
			name: "zero length",
			opt: []Option{
				WithPassLen(0),
			},
			err:     ErrIncorrectLength,
			wantErr: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			opt := NewOptions(test.opt...)
			passwords, hashedPasswords, err := GeneratePasswords(test.opt...)
			if test.wantErr {
				if err == nil {
					t.Errorf("expected error: %v, got nil", test.err)
				}
				if !errors.Is(err, test.err) {
					t.Errorf("expected error: %v, got: %v", test.err, err)
				}
				return
			}
			if err != nil {
				t.Errorf("expected nil, got: %v", err)
			}

			if opt.passCount != len(passwords) {
				t.Errorf("expected %d passwords, got %d", opt.passCount, len(passwords))
			}
			if opt.passCount != len(hashedPasswords) {
				t.Errorf("expected %d hashedPasswords, got %d", opt.passCount, len(hashedPasswords))
			}
			for i := 0; i < opt.passCount; i++ {
				t.Logf("password: %v, hashedPassword: %v", passwords[i], hashedPasswords[i])
				if err = bcrypt.CompareHashAndPassword([]byte(hashedPasswords[i]), []byte(passwords[i])); err != nil {
					t.Fatal(err)
				}
			}
		})
	}
}
