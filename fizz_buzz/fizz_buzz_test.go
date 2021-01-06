package fizz_buzz

import "testing"

func TestFizzBuzz15(t *testing.T) {
	got := FizzBuzz(3, 5, 15, "Fizz", "Buzz")
	want :=
		`1,2,Fizz,4,Buzz,Fizz,7,8,Fizz,Buzz,11,Fizz,13,14,FizzBuzz`
	if got != want {
		t.Errorf("FizzBuzz \n got: \n%v \n want: \n%v", got, want)
	}
}

func TestFizzBuzz100(t *testing.T) {
	got := FizzBuzz(3, 5, 100, "Fizz", "Buzz")
	want :=
		`1,2,Fizz,4,Buzz,Fizz,7,8,Fizz,Buzz,11,Fizz,13,14,FizzBuzz,16,17,Fizz,19,Buzz,Fizz,22,23,Fizz,Buzz,26,Fizz,28,29,FizzBuzz,31,32,Fizz,34,Buzz,Fizz,37,38,Fizz,Buzz,41,Fizz,43,44,FizzBuzz,46,47,Fizz,49,Buzz,Fizz,52,53,Fizz,Buzz,56,Fizz,58,59,FizzBuzz,61,62,Fizz,64,Buzz,Fizz,67,68,Fizz,Buzz,71,Fizz,73,74,FizzBuzz,76,77,Fizz,79,Buzz,Fizz,82,83,Fizz,Buzz,86,Fizz,88,89,FizzBuzz,91,92,Fizz,94,Buzz,Fizz,97,98,Fizz,Buzz`
	if got != want {
		t.Errorf("FizzBuzz \n got: \n%v \n want: \n%v", got, want)
	}
}

func TestFizzBuzz15Custom(t *testing.T) {
	got := FizzBuzz(3, 5, 15, "Toto", "Tata")
	want :=
		`1,2,Toto,4,Tata,Toto,7,8,Toto,Tata,11,Toto,13,14,TotoTata`
	if got != want {
		t.Errorf("FizzBuzz \n got: \n%v \n want: \n%v", got, want)
	}
}
