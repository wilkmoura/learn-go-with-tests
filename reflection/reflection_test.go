package reflection

import (
	"reflect"
	"testing"
)

// reflections is the ability of a program to examin its own structure, particularly through types.
type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"Struct with non string fields",
			struct {
				Name string
				Age  int
			}{"Chris", 30},
			[]string{"Chris"},
		},
		{
			"Nested fields",
			Person{
				"Wilkinson",
				Profile{30, "Fortaleza"},
			},
			[]string{"Wilkinson", "Fortaleza"},
		},
		{
			"Pointers to things",
			&Person{
				"Wilkinson",
				Profile{30, "Fortaleza"},
			},
			[]string{"Wilkinson", "Fortaleza"},
		},
		{
			"Slices",
			[]Profile{
				{30, "Fortaleza"},
				{33, "Brasília"},
			},
			[]string{"Fortaleza", "Brasília"},
		},
		{
			"Arrays",
			[2]Profile{
				{30, "Fortaleza"},
				{33, "Brasília"},
			},
			[]string{"Fortaleza", "Brasília"},
		},
		{
			"Maps",
			map[string]string{
				"Foo": "Bar",
				"Baz": "Boz",
			},
			[]string{"Bar", "Boz"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{33, "Ceará"}, Profile{30, "London"}
		}

		var got []string
		want := []string{"Ceará", "London"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
