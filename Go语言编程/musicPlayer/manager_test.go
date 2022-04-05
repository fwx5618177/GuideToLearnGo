package library

import (
	"testing"
)

func TestOps(t *testing.T) {
	mm := NewMusicManager()

	if mm == nil {
		t.Error("new music manager failed.")
	}

	if mm.Len() != 0 {
		t.Error("not empty.")
	}

	m0 := &MusicEntry {
		"1", "Heart", "Pop", "1", "Mp3",
	}

	mm.Add(m0)

	if mm.Len() != 1 {
		t.Error("Add() failed.")
	}

	m, _ := mm.Find(m0.Name)

	if m == nil {
		t.Error("Find() failed.")
	}

	if m.Id != m0.Id || m.Artist != m0.Artist || m.Source != m0.Source || m.Type != m0.Type {
		t.Error("found item failed.")
	}

	m, err := mm.Get(0)
	if m == nil {
		t.Error("Get() failed.")

	}

	m, _ = mm.Remove(0)

	if m == nil || mm.Len() != 0 {
		t.Error("Remove() failed.", err)

	}
}
