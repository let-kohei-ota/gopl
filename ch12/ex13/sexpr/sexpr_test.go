// Copyright 2017 budougumi0617 All Rights Reserved.

// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package sexpr

import (
	"fmt"
	"reflect"
	"testing"
)

// Test verifies that encoding and decoding a complex data value
// produces an equal result.
//
// The test does not make direct assertions about the encoded output
// because the output depends on map iteration order, which is
// nondeterministic.  The output of the t.Log statements can be
// inspected by running the test with the -v flag:
//
// 	$ go test -v gopl.io/ch12/sexpr
//
func Test(t *testing.T) {
	type Movie struct {
		Title    string `sexpr:"title"`
		Subtitle string
		Year     int `sexpr:"year"`
		Actor    map[string]string
		Oscars   []string
		Sequel   *string
	}
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}

	// Encode it
	data, err := Marshal(strangelove)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = %s\n", data)
	fmt.Println(string(data))

	// Decode it
	var movie Movie
	if err := Unmarshal(data, &movie); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	t.Logf("Unmarshal() = %+v\n", movie)

	// Check equality.
	if !reflect.DeepEqual(movie, strangelove) {
		t.Fatal("not equal")
	}
}

func TestMarshal(t *testing.T) {
	type Interface interface{}
	type Record struct {
		B    bool
		F32  float32
		F64  float64 `sexpr:"f64"`
		C64  complex64
		C128 complex128
		I    Interface `sexpr:"i"`
	}
	tests := []struct {
		r    Record
		want string
	}{
		{
			Record{true, 2.5, 0, 1 + 2i, 2 + 3i, Interface(5)},
			`((B t) (F32 2.5) (f64 0) (C64 #C(1 2)) (C128 #C(2 3)) (i ("github.com/budougumi0617/gopl/ch12/ex13/sexpr.Interface" 5)))`,
		},
		{
			Record{false, 0, 1.5, 0, 1i, Interface(0)},
			`((B nil) (F32 0) (f64 1.5) (C64 #C(0 0)) (C128 #C(0 1)) (i ("github.com/budougumi0617/gopl/ch12/ex13/sexpr.Interface" 0)))`,
		},
	}
	for _, test := range tests {
		data, err := Marshal(test.r)
		s := string(data)
		if err != nil {
			t.Errorf("Marshal(%s): %s", s, err)
		}
		if s != test.want {
			t.Errorf("Marshal(%#v) got %s, wanted %s", test.r, s, test.want)
		}
	}
}

func TestUnmarshal(t *testing.T) {
	type Interface interface{}
	type Record struct {
		B    bool
		F32  float32 `sexpr:"f_32"`
		F64  float64
		C64  complex64 `sexpr:"c_64"`
		C128 complex128
		I    Interface `sexpr:"i"`
	}
	Interfaces["sexpr.Interface"] = reflect.TypeOf(int(0))
	tests := []struct {
		s    string
		want Record
	}{
		{
			`((B t) (f_32 2.5) (F64 1.5) (i ("sexpr.Interface" 5)))`,
			Record{true, 2.5, 1.5, 0, 0, Interface(5)},
		},
		{
			`((B nil) (f_32 0) (F64 1.5) (i ("sexpr.Interface" 0)))`,
			Record{false, 0, 1.5, 0, 0, Interface(0)},
		},
	}
	for _, test := range tests {
		var r Record
		err := Unmarshal([]byte(test.s), &r)
		if err != nil {
			t.Errorf("Unmarshal(%q): %s", test.s, err)
		}
		if !reflect.DeepEqual(r, test.want) {
			t.Errorf("Unmarshal(%q) got %#v, wanted %#v", test.s, r, test.want)
		}
	}
}
