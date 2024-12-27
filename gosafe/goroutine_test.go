package gosafe

import "testing"

func TestGoNoPanic(t *testing.T) {
	errChan := Go(func() error {
		return nil
	})

	err := <-errChan
	if err != nil {
		t.Errorf("Expected no error, got err: %v", err)
	}
}

func TestGoPanic(t *testing.T) {
	errChan := Go(func() error {
		panic("never gonna give you up, never gonna let you down")
	})

	err := <-errChan
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
}
