package main

import "testing"

func TestHello(t *testing.T) {
	checkCorrectMessage := func(t *testing.T, result, expected string) {
		t.Helper()
		if result != expected {
			t.Errorf("Result '%s', expected '%s'", result, expected)
		}
	}

	t.Run("Say hello to the user", func(t *testing.T) {
		result := Hello("Chris", "")
		expected := "Hello, Chris!"
		checkCorrectMessage(t, result, expected)
	})

	t.Run("Say 'Hello, world!' when an empty string was provided", func(t *testing.T) {
		result := Hello("", "")
		expected := "Hello, world!"
		checkCorrectMessage(t, result, expected)
	})

	t.Run("In Portuguese", func(t *testing.T) {
		result := Hello("Elodie", "Portuguese")
		expected := "Olá, Elodie!"
		checkCorrectMessage(t, result, expected)
	})

	t.Run("In French", func(t *testing.T) {
		result := Hello("Zidane", "French")
		expected := "Bonjour, Zidane!"
		checkCorrectMessage(t, result, expected)
	})
}
