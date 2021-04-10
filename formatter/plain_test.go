package formatter

import "testing"

func TestPlainFormatStruct(t *testing.T) {
	movie := struct {
		Title string
		Year  int
	}{
		"Halloween",
		1978,
	}
	c := &Config{
		Format: "plain",
	}
	out, _ := OutputData(movie, c)
	got := string(out)

	want := "{Title:Halloween Year:1978}"
	if got != want {
		t.Fatalf(`values not equal ("%s" != "%s")`,
			got,
			want,
		)
	}
}

func TestPlainFormatStructList(t *testing.T) {
	movies := []struct {
		Title string
		Year  int
	}{
		{
			"Halloween",
			1978,
		},
		{
			"Phantasm",
			1979,
		},
	}
	c := &Config{
		Format: "plain",
	}
	out, _ := OutputData(movies, c)
	got := string(out)

	want := "[{Title:Halloween Year:1978} {Title:Phantasm Year:1979}]"
	if got != want {
		t.Fatalf(`values not equal ("%s" != "%s")`,
			got,
			want,
		)
	}
}
