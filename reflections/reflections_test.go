package reflections

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Job string
	Age int
}

func TestWalk(t *testing.T) {
	testcases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string",
			struct {
				Name string
			}{"long"},
			[]string{"long"},
		},
		{
			"struct with two string",
			struct {
				Name string
				City string
			}{"long", "Ho Chi Minh"},
			[]string{"long", "Ho Chi Minh"},
		},
		{
			"struct with one is not string",
			struct {
				Name string
				Age  int
			}{"long", 32},
			[]string{"long"},
		},
		{
			"struct with nested",
			Person{
				Name: "long",
				Profile: Profile{
					Age: 32,
					Job: "developer",
				},
			},
			[]string{"long", "developer"},
		},
		{
			"pointer struct with nested",
			&Person{
				Name: "long",
				Profile: Profile{
					Age: 32,
					Job: "developer",
				},
			},
			[]string{"long", "developer"},
		},
		{
			"slice of struct with nested",
			[]Profile{
				{"developer", 32},
				{"teacher", 30},
			},
			[]string{"developer", "teacher"},
		},
		{
			"array of struct with nested",
			[2]Profile{
				{"developer", 32},
				{"teacher", 30},
			},
			[]string{"developer", "teacher"},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			var got []string
			Walk(testcase.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, testcase.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, testcase.ExpectedCalls)
			}
		})
	}

	t.Run("Maps", func(t *testing.T) {
		aMap := map[string]interface{}{
			"name": "long",
			"job":  "developer",
		}
		var got []string
		Walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "long")
		assertContains(t, got, "developer")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)
		go func() {
			aChannel <- Profile{
				Job: "developer",
				Age: 32,
			}
			aChannel <- Profile{
				Job: "teacher",
				Age: 30,
			}
			close(aChannel)
		}()

		var got []string
		want := []string{"developer", "teacher"}
		Walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunc := func() (Profile, Profile) {
			return Profile{
					Job: "developer",
					Age: 32,
				}, Profile{
					Job: "teacher",
					Age: 30,
				}
		}

		var got []string
		want := []string{"developer", "teacher"}
		Walk(aFunc, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	isContain := false
	for _, val := range haystack {
		if val == needle {
			isContain = true
			break
		}
	}
	if !isContain {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}
