package testy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFooMeOnceNil(t *testing.T) {
	assert.Equal(t, nil, fooMeOnceNil(), "Test did not return nil")
}

func TestFooMeOnceArray(t *testing.T) {
	assert.Equal(t, nil, fooMeOnceArray(), "Test did not return nil")
}

func TestFooMeOncePointerWithNil(t *testing.T) {
	assert.Equal(t, nil, fooMeOncePointerWithNil(), "Test did not return nil")
}

func TestFooMeOncePointerToArrayWithNil(t *testing.T) {
	assert.Equal(t, nil, fooMeOncePointerToArrayWithNil(), "Test did not return nil")
}

func TestPassFooMeOnceNil(t *testing.T) {
	assert.Equal(t, []foo(nil), fooMeOnceNil(), "Test did not return nil")
}

func TestPassFooMeOncePointerWithNil(t *testing.T) {
	assert.Equal(t, (*foo)(nil), fooMeOncePointerWithNil(), "Test did not return nil")
}

func TestPassFooMeOncePointerToArrayWithNil(t *testing.T) {
	assert.Equal(t, (*[]foo)(nil), fooMeOncePointerToArrayWithNil(), "Test did not return nil")
}
