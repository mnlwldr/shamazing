package shamazing

import (
	"testing"
)

func TestFindLongestString(t *testing.T) {
	want := "dbdc"
	got := FindLongestString("82e58c167ca53584582374c960b7456a8980dbdc")

	if got != want {
		t.Fatalf("FindLongestString(cf17d3d047e1981947bb5d1fc842483481454f08), want %v got %v", want, got)
	}
}

func TestFindLongestInteger(t *testing.T) {
	var want int64 = 4588172
	got, err := FindLongestInteger("e8c0cd00bc18bd927163ae080d4588172c7a83e3")

	if got != want {
		t.Fatalf("FindLongestString(e8c0cd00bc18bd927163ae080d4588172c7a83e3), want %v got %v err %v", want, got, err)
	}
}
