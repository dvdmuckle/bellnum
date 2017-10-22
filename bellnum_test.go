package main

import "testing"

func TestBellsum(t *testing.T) {
	tables := []struct {
		x    int
		bell int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 5},
		{4, 15},
		{5, 52},
		{6, 203},
		{7, 877},
		{8, 4140},
		{9, 21147},
		{10, 115975},
		{11, 678570},
		{12, 4213597},
		{13, 27644437},
		{14, 190899322},
		{15, 1382958545},
		{16, 10480142147},
		{17, 82864869804},
		{18, 682076806159},
		{19, 5832742205057},
	}
	for _, table := range tables {
		bell := Bellnum(table.x)
		if bell != table.bell {
			t.Errorf("Bell of %d was incorrect, got %d, want %d", table.x, bell, table.bell)
		}
	}
}
