package bench

import "testing"

func TestIsPrime(t *testing.T) {
	if !isPrime(7867) {
		t.Error("7867 is not prime")
	}
	if !isPrimeV2(7867) {
		t.Error("7867 is not prime")
	}
	if isPrime(10) {
		t.Error("10 is prime")
	}
	if isPrimeV2(10) {
		t.Error("10 is prime")
	}
}

//Benchmack][Name] (* testing.B){}

func BenchmackIsPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPrime(7867)
	}
}

func BenchmackIsPrimeV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPrimeV2(7867)
	}
}
func BenchmackIsPrimeV3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPrimeV3(7867)
	}
}
