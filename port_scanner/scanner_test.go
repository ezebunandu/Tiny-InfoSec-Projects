package scanner_test

import (
	"scanner"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestScanner(t *testing.T) {
	assertError := func(t testing.TB, got error, want error) {
		t.Helper()
		if got == nil {
			t.Fatal("didn't get an error but wanted one")
		}
		if want != got {
			t.Errorf("wanted %s, got %s", want, got)
		}
	}
	t.Run("valid ip address and one port is open", func(t *testing.T) {
		t.Parallel()
		want := []int{443}
		got, _ := scanner.GetOpenPorts("209.216.230.240", []int{440, 444})
		if !cmp.Equal(want, got) {
			t.Errorf("want %#v, got %#v", want, got)
		}
	})

	t.Run("url port 80 open", func(t *testing.T) {
		t.Parallel()
		want := []int{80}
		got, _ := scanner.GetOpenPorts("www.stackoverflow.com", []int{79, 82})
		if !cmp.Equal(want, got) {
			t.Errorf("want #%v, got #%v", want, got)
		}
	})
	t.Run("url multiple ports open", func(t *testing.T) {
		t.Parallel()
		want := []int{22, 80}
		got, _ := scanner.GetOpenPorts("scanme.nmap.org", []int{20, 80})
		if !cmp.Equal(want, got) {
			t.Errorf("want #%v, got #%v", want, got)
		}
	})
	t.Run("invalid hostname returns error", func(t *testing.T) {
		t.Parallel()
		_, err := scanner.GetOpenPorts("scanme.nmap", []int{22, 42})
		assertError(t, err, scanner.ErrInvalidHost)
	})

}
