package main

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	dict := Dict{"test": "this is just a test"}

	t.Run("Known word", func(t *testing.T) {
		got, _ := dict.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})

	t.Run("Unknown word", func(t *testing.T) {
		_, err := dict.Search("batata")
		assertError(t, err, ErrNotFoundWord)
	})
}

func TestAdd(t *testing.T) {
	t.Run("New Word", func(t *testing.T) {
		dict := Dict{}
		key := "new key"
		value := "new value"
		err := dict.Add(key, value)
		assertError(t, err, nil)
		assertDefinition(t, dict, key, value)
	})

	t.Run("Known Word", func(t *testing.T) {
		key := "new key"
		value := "new value"
		dict := Dict{key: value}
		err := dict.Add(key, "new new value")
		assertError(t, err, ErrExistentWord)
		assertDefinition(t, dict, key, value)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("New Word", func(t *testing.T) {
		key := "new key"
		value := "new value"
		dict := Dict{}
		err := dict.Update(key, value)
		fmt.Print(err)
		assertError(t, err, ErrNonexistentWord)
	})

	t.Run("Known Word", func(t *testing.T) {
		key := "new key"
		value := "new value"
		dict := Dict{key: value}
		newDef := "new new value"
		err := dict.Update(key, newDef)
		assertError(t, err, nil)
		assertDefinition(t, dict, key, newDef)
	})
}

func TestDelete(t *testing.T) {
	key := "test"
	dict := Dict{key: "value"}

	dict.Delete(key)

	_, err := dict.Search(key)
	if err != ErrNotFoundWord {
		t.Errorf("It is expected that %s be deleted", key)
	}
}

func assertDefinition(t *testing.T, dict Dict, key, value string) {
	got, err := dict.Search(key)
	if err != nil {
		t.Fatal("it was not possible to find the added word")
	}

	if got != value {
		t.Errorf("Got: %s, Want %s", got, value)
	}
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got: %s, Want: %s", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("Got: %s, Want: %s", got, want)
	}
}
