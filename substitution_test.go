package main

import (
	"testing"
)

func TestSubstitutionEncrypt(t *testing.T) {
	var encodeTests = []struct {
		in  string
		out string
	}{
		{"aaaa", "eeee"},
		{"bbbb", "ffff"},
		{"abcd", "efgh"},
		{"test this", "xiwxxlmw"},
		{"wxyz", "ABCD"},
	}

	cipher := NewSubstitutionCipher()

	for _, v := range encodeTests {
		out, _ := cipher.Encrypt(v.in)
		if out != v.out {
			t.Errorf("Expected '%s', got '%s'", v.out, out)
		}
	}
}
func TestSubstitutionEncryptWithError(t *testing.T) {
	cipher := NewSubstitutionCipher()
	out, err := cipher.Encrypt("%$!")
	if err == nil {
		t.Errorf("Expected error but got '%s'", out)
	}

	out, err = cipher.Encrypt(`   ⌘`)
	if err == nil {
		t.Errorf("Expected error but got '%s'", out)
	}
}

func TestSubstitutionDecrypt(t *testing.T) {
	var encodeTests = []struct {
		in  string
		out string
	}{
		{"eeee", "aaaa"},
		{"ffff", "bbbb"},
		{"efgh", "abcd"},
		{"ABCD", "wxyz"},
		{"xiw xx lmw", "testthis"},
	}

	cipher := NewSubstitutionCipher()

	for _, v := range encodeTests {
		out, _ := cipher.Decrypt(v.in)
		if out != v.out {
			t.Errorf("Expected '%s', got '%s'", v.out, out)
		}
	}
}

func TestSubstitutionDecryptWithError(t *testing.T) {
	cipher := NewSubstitutionCipher()
	out, err := cipher.Decrypt("%$!")
	if err == nil {
		t.Errorf("Expected error but got '%s'", out)
	}

	out, err = cipher.Decrypt(`   ⌘`)
	if err == nil {
		t.Errorf("Expected error but got '%s'", out)
	}
}
