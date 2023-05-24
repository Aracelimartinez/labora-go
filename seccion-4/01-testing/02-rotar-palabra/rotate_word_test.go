package rotateWord

import(
	"testing"
)

type rotateTest struct {
	letters, expected string
}

type rotateTimesTest struct {
	letters, expected string
	times int
}

//Rotate right Tests

var rotateRightVar = []rotateTest{
{"ABCD", "BCDA"},
{"AHORA", "HORAA"},
{"SIEMPRE", "IEMPRES"},
{"NARANJA", "ARANJAN"},
}

func TestRotateRight(t *testing.T)  {
	for _, test := range rotateRightVar {
		if output := rotateRight(test.letters); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}

func BenchmarkRotateRight(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rotateRight("BCA")
	}
}

//Rotate left Tests

var rotateLeftVar = []rotateTest{
	{"ABCD", "DABC"},
	{"AHORA", "AAHOR"},
	{"SIEMPRE", "ESIEMPR"},
	{"NARANJA", "ANARANJ"},
}

func TestRotateLeft(t *testing.T)  {
	for _, test := range rotateLeftVar {
		if output := rotateLeft(test.letters); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}

func BenchmarkRotateLeft(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rotateLeft("ABC")
	}
}

// Rotate right times tests

var rotateRightTimesVar = []rotateTimesTest{
	{"ABCD", "BCDA", 1},
	{"AHORA", "ORAAH", 2},
	{"SIEMPRE", "MPRESIE", 3},
	{"NARANJA", "NJANARA", 4},
}

func TestRotateRightTimes(t *testing.T)  {
	for _, test := range rotateRightTimesVar {
		if output := rotateRightTimes(test.letters, test.times); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}

func BenchmarkRotateRightTimes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rotateRightTimes("SIEMPRE", 4)
	}
}

	// Rotate left times tests

var rotateLeftTimesVar = []rotateTimesTest{
	{"ABCD", "DABC", 1},
	{"AHORA", "RAAHO", 2},
	{"SIEMPRE", "PRESIEM", 3},
	{"NARANJA", "ANJANAR", 4},
}

func TestRotateleftTimes(t *testing.T)  {
	for _, test := range rotateLeftTimesVar {
		if output := rotateLeftTimes(test.letters, test.times); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}

func BenchmarkRotateLeftTimes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rotateRightTimes("SIEMPRE", 4)
	}
}
