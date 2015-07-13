package signer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStub(t *testing.T) {
	assert.True(t, true, "This is good. Canary test passing")
}

// NewSigner(email string, privring_path string, signdir_path string) Signer
func TestNewSigner(t *testing.T) {
  ts := NewSigner("tarball-signer@example.org","secring.gpg","/bar")
	assert.NotNil(t, ts, "We are expecting a Signer object")
}

//
func TestSigner(t *testing.T) {
  ts := NewSigner("tarball-signer@example.org","secring.gpg","/bar")
	assert.NotNil(t,ts.Entity,"We are expecting an openpgp.Entity object")
	assert.Equal(t,ts.Path,"/bar","Path should be set right")
}

// TODO
// test SignIt method
