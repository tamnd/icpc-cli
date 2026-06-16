package icpc

import (
	"testing"

	"github.com/tamnd/any-cli/kit"
)

func TestDomainInfo(t *testing.T) {
	info := Domain{}.Info()
	if info.Scheme != "icpc" {
		t.Errorf("Scheme = %q, want icpc", info.Scheme)
	}
	if len(info.Hosts) == 0 {
		t.Error("Hosts should not be empty")
	}
	if info.Identity.Binary != "icpc" {
		t.Errorf("Binary = %q, want icpc", info.Identity.Binary)
	}
}

func TestClassify(t *testing.T) {
	cases := []struct {
		in      string
		wantTyp string
		wantID  string
		wantErr bool
	}{
		{"2024", "results", "2024", false},
		{"2019", "results", "2019", false},
		{"4234", "problem", "4234", false},
		{"100", "problem", "100", false},
		{"", "", "", true},
		{"not-a-number", "", "", true},
	}
	for _, tc := range cases {
		typ, id, err := Domain{}.Classify(tc.in)
		if tc.wantErr {
			if err == nil {
				t.Errorf("Classify(%q) expected error, got (%q,%q,nil)", tc.in, typ, id)
			}
			continue
		}
		if err != nil {
			t.Errorf("Classify(%q) error = %v", tc.in, err)
			continue
		}
		if typ != tc.wantTyp || id != tc.wantID {
			t.Errorf("Classify(%q) = (%q,%q), want (%q,%q)", tc.in, typ, id, tc.wantTyp, tc.wantID)
		}
	}
}

func TestLocate(t *testing.T) {
	cases := []struct {
		typ     string
		id      string
		want    string
		wantErr bool
	}{
		{"results", "2024", "https://icpc.global/worldfinals", false},
		{"problem", "100", "https://icpcarchive.ecs.baylor.edu/index.php?option=com_onlinejudge&Itemid=8&page=show_problem&problem=100", false},
		{"bogus", "x", "", true},
	}
	for _, tc := range cases {
		got, err := Domain{}.Locate(tc.typ, tc.id)
		if tc.wantErr {
			if err == nil {
				t.Errorf("Locate(%q,%q) expected error", tc.typ, tc.id)
			}
			continue
		}
		if err != nil {
			t.Errorf("Locate(%q,%q) error = %v", tc.typ, tc.id, err)
			continue
		}
		if got != tc.want {
			t.Errorf("Locate(%q,%q) = %q, want %q", tc.typ, tc.id, got, tc.want)
		}
	}
}

func TestHostWiring(t *testing.T) {
	h, err := kit.Open()
	if err != nil {
		t.Fatal(err)
	}
	got, err := h.ResolveOn("icpc", "2024")
	if err != nil {
		t.Fatalf("ResolveOn: %v", err)
	}
	if got.String() != "icpc://results/2024" {
		t.Errorf("ResolveOn = %q, want icpc://results/2024", got.String())
	}
}
