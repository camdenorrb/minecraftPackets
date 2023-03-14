package nbt

import "testing"

type Hello struct {
	Hello string
	Thing string
}

func FuzzName(f *testing.F) {
	f.Add(Hello{Hello: "Hello", Thing: "Thing"})
	f.Fuzz(func(t *testing.T, hello Hello) {
		if hello.Hello != "Hello" {
			t.Errorf("Expected Hello to be Hello, got %s", hello.Hello)
		}
		if hello.Thing != "Thing" {
			t.Errorf("Expected Thing to be Thing, got %s", hello.Thing)
		}
	})
}
