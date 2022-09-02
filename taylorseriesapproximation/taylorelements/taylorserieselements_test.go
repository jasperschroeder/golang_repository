package taylorelements

import (
	"testing"
	"reflect"
)

// Not the cleanest structure but just here to try out different testing procedures.

func TestFactorial(t *testing.T) {
	for i := 0; i <= 10; i++ {
		t.Logf("Testing factorial with %d", i)

		t.Log("Test that output is non-negative")
		if factorial(i) < 0 {
			t.Fatalf("Negative value found for factorial() with %d",i)
		}

		t.Log("Check that output is not nil")
		if reflect.TypeOf(factorial(i)) == nil {
			t.Fatalf("Output is nil for factorial() with %d", i)
		}
	}

}

func TestFactorialSubTests(t *testing.T) {
	for i := 0; i <= 10; i++ {
		t.Run("Non-negative", func(t *testing.T) {if factorial(i) < 0 {t.Fatalf("Negative value found for factorial() with %d", i)}
		})
		
		t.Run("Non-nil", func(t *testing.T) {if reflect.TypeOf(factorial(i)) == nil {t.Fatalf("Nil type detected for factorial(%d)", i)}	
		})
	}
}

func TestSkip(t *testing.T) {
	for i := 0; i <= 1000; i++ {
		// Because computers suck at math, we want to stop it if factorial is evaluated for an integer bigger than 20
		if i > 20 {
			t.Skip("Skipping as number supplied is bigger than 20")
		}
		if factorial(i) < 0 {
			t.Errorf("Negative value for factorial(%d)", i)
		}
	}
}

func TestCleanLogs(t *testing.T) {
	t.Cleanup(func() {
		t.Log("Cleaning up everything. That's enough testing for today. Goodbye!")
	})
	t.Log("Running some test")
}