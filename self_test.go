package main

import (
	"io/ioutil"
	"strings"
	"testing"
)

// TestSelfPass makes sure all the code in this project passes spell checking.
func TestSelfPass(t *testing.T) {
	paths := gopaths(".")
	sp, err := NewSpellcheck("")
	if err != nil {
		t.Fatal(err)
	}
	defer sp.Close()
	cts, err := sp.Check(paths)
	rejects := 0
	for _, ct := range cts {
		if strings.Contains(ct.ctok.lit, "Reject") {
			// reject test cases will give false positives
			rejects++
			continue
		}
		t.Errorf("flagged %s", ct.ctok.lit)
	}
	if rejects == 0 {
		t.Errorf("expected rejects from tests")
	}
}

func gopaths(rootdir string) (ret []string) {
	fi, err := ioutil.ReadDir(rootdir)
	if err != nil {
		return ret
	}
	for _, f := range fi {
		newpath := rootdir + "/" + f.Name()
		if f.IsDir() {
			ret = append(ret, gopaths(newpath)...)
		} else if len(newpath) > 3 && newpath[len(newpath)-3:] == ".go" {
			ret = append(ret, newpath)
		}
	}
	return ret
}
