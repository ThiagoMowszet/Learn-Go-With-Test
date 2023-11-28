/* 

Rules of TEST with Go

1. It needs to be in a file with a name like xxx_test.go
2. The test function must start with the word Test
3. The test function takes one argument only t *testing.T
4. In order to use the *testing.T type, you need to import "testing"

*/

package main

import "testing"

const englishHelloPrefix = "Hello, "

func TestHello(t *testing.T) {
    got := Hello()
    want := "Hello, world!"

    // NOTE: why we use %q -> https://dev.to/mcaci/using-q-inside-fmt-printf-in-go-2i23
    if got != want {
        // t.Errorf("got %q want %q", got, want)
        t.Errorf("\nGOT: %q\nWANT:%q", got, want)
    }
}

func TestHelloWithParameter(t *testing.T) {
    got := HelloWithParameter("Thiago")
    want := "Hello, Thiago"

    if got != want {
        t.Errorf("\nGOT: %q\nWANT:%q", got, want)
    }
}

func TestHelloEnglishPrefix(t *testing.T) {
    got := HelloEnglishPrefix("Thiago")
    want := englishHelloPrefix + "Thiago"

    if got != want {
        t.Errorf("\nGOT: %q\nWANT: %q", got, want)
    }
} 


// Refactoring
func TestHelloWithParameter2(t *testing.T) {
    t.Run("saying hello to people", func(t *testing.T) {
        got := HelloWithParameter("Chris")
        want := "Hello, Chris"
        assertCorrectMessage(t, got, want)
    })

    t.Run("empty string defaults to 'world'", func(t *testing.T) {
        got := HelloWithParameter("")
        want := "Hello, "
        assertCorrectMessage(t, got, want)
    })
}


// We need to pass in t *testing.T so that we can tell the test code to fail when we need to.
func assertCorrectMessage(t testing.TB, got, want string) {
    // NOTE: By doing this when it fails the line number reported will be in our function call rather than inside our test helper.
    t.Helper()
    // t.Helper()

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}


func TestHello2(t *testing.T) {
    t.Run("In spanish", func(t *testing.T) {
        got := Hello2("Elodie", "Spanish")
        want := ("Hola, Elodie")
        assertCorrectMessage(t, got, want)
    })

    t.Run("Empty string default to World", func(t *testing.T){
        got := Hello2("", "")
        want := "Hello, World"
        assertCorrectMessage(t, got, want)
    })
    
    t.Run("In French", func(t *testing.T){
        got := Hello2("", "French")
        want := "Bonjour, World"
        assertCorrectMessage(t, got, want)
    })

}

