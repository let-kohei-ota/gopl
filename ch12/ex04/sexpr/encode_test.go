// Copyright 2016 budougumi0617 All Rights Reserved.

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
		Title, Subtitle string
		Year            int
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
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
	t.Logf("Marshal() = \n%s\n", data)

	// Decode it
	var movie Movie
	if err = Unmarshal(data, &movie); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	t.Logf("Unmarshal() = %+v\n", movie)

	// Check equality.
	if !reflect.DeepEqual(movie, strangelove) {
		t.Fatal("not equal")
	}

}

func TestFloat(t *testing.T) {
	var tests = []struct {
		num32 float32
		num64 float64
	}{
		{12.3, 3.21},
		{0.0, 10000},
	}
	for _, test := range tests {
		actual32, err := Marshal(test.num32)
		if err != nil {
			t.Fatalf("return err %v", err.Error())
		}
		actual64, err := Marshal(test.num64)
		if err != nil {
			t.Fatalf("return err %v", err.Error())
		}

		if string(actual32) != fmt.Sprintf("%4.4f", test.num32) {
			t.Errorf("Result = %v, Expected %v", actual32, test.num32)
		}
		if string(actual64) != fmt.Sprintf("%4.4f", test.num64) {
			t.Errorf("Result = %v, Expected %v", actual64, test.num64)
		}
	}
}

func TestComplex(t *testing.T) {
	var tests = []struct {
		num64  complex64
		num128 complex128
	}{
		{12.0 - 3i, 3.2 + 1i},
		{0.0 - 0i, 10000 + 0i},
	}
	for _, test := range tests {
		actual64, err := Marshal(test.num64)
		if err != nil {
			t.Fatalf("return err %v", err.Error())
		}
		actual128, err := Marshal(test.num128)
		if err != nil {
			t.Fatalf("return err %v", err.Error())
		}

		if string(actual64) != fmt.Sprintf("#C(%4.4f %4.4f)", real(test.num64), imag(test.num64)) {
			t.Errorf("Result = %s, Expected %v", actual64, test.num64)
		}
		if string(actual128) != fmt.Sprintf("#C(%4.4f %4.4f)", real(test.num128),
			imag(test.num128)) {
			t.Errorf("Result = %s, Expected %v", actual128, test.num128)
		}
	}
}

func TestInterface(t *testing.T) {

	type Interface interface{}
	type Outer struct {
		i Interface
	}
	type Inner struct {
		i []int
		f []float32
	}
	var tests = []struct {
		input Interface
		want  string
	}{
		{Inner{[]int{1, 2, 3}, []float32{6, 7, 8}},
			"(sexpr.Interface sexpr.Inner\n ((i (1\n      2\n      3))\n  (f (6.0000\n      7.0000\n      8.0000))))"},
		{nil, "(sexpr.Interface nil)"},
	}
	for _, test := range tests {
		get, err := Marshal(&(test.input))
		if err != nil {
			t.Fatalf("return err %v", err.Error())
		}
		if string(get) != test.want {
			t.Errorf("Result = \n%v, Expected \n%v", string(get), test.want)
		}
	}
}

func TestBool(t *testing.T) {
	var tests = []struct {
		b    bool
		want string
	}{
		{true, "t"},
		{false, "nil"},
	}
	for _, test := range tests {
		actual, err := Marshal(test.b)
		if err != nil {
			t.Fatalf("return err %v", err.Error())
		}
		if string(actual) != test.want {
			t.Errorf("Result = %s, Expected %v", actual, test.want)
		}
	}
}
