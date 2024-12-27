package gosafe

import (
	"context"
	"testing"
)

func TestNoPanicGo(t *testing.T) {
	group, _ := GoGroupWithContext(context.Background())
	group.Go(func() error {
		return nil
	})
	err := group.Wait()
	if err != nil {
		t.Errorf("Expected no error, got err: %v", err)
	}
}

func TestPanicGo(t *testing.T) {
	group, _ := GoGroupWithContext(context.Background())
	group.Go(func() error {
		panic("never gonna give you up, never gonna let you down")
	})
	err := group.Wait()
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
}

func TestNoPanicTryGo(t *testing.T) {
	group, _ := GoGroupWithContext(context.Background())
	ok := group.TryGo(func() error {
		return nil
	})
	if !ok {
		t.Errorf("Expected true, but got false.")
	}
	err := group.Wait()
	if err != nil {
		t.Errorf("Expected no error, got err: %v", err)
	}
}

func TestPanicTryGo(t *testing.T) {
	group, _ := GoGroupWithContext(context.Background())
	ok := group.TryGo(func() error {
		panic("never gonna give you up, never gonna let you down")
	})
	if !ok {
		t.Errorf("Expected true, but got false.")
	}
	err := group.Wait()
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
}
