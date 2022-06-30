package service

import (
	"testing"
)

func TestCheckIBANforCountryCode(t *testing.T) {
	iban := "IN13RZBR0000060007134800"
	_, err := CheckIBAN(iban)
	got := err.Error()
	want := "country code IN not supported"
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}

	iban2 := "RO13RZBR0000060007134800"
	got2, _ := CheckIBAN(iban2)
	want2 := true
	if got2 != want2 {
		t.Errorf("got %t, wanted %t", got2, want2)
	}
}
func TestCheckIBANforLength(t *testing.T) {
	iban := "RO13RZBR00000600071348"
	_, err := CheckIBAN(iban)
	got := err.Error()
	want := "IBAN length for RO is 24"
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
	iban2 := "RO13RZBR0000060007134800"
	got2, _ := CheckIBAN(iban2)
	want2 := true
	if got2 != want2 {
		t.Errorf("got %t, wanted %t", got2, want2)
	}
}
func TestGetNumbers(t *testing.T) {
	got := intMapper()["Z"]
	want := "35"
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestGetCountryCode(t *testing.T) {
	got := getCountryCode("GB82WEST12345698765432")
	want := "GB"
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestTheSwap(t *testing.T) {
	got := theSwap("GB82WEST12345698765432")
	want := "WEST12345698765432GB82"
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestReplace(t *testing.T) {
	got := Replace("WEST12345698765432GB82")
	want := "3214282912345698765432161182"
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestMod97(t *testing.T) {
	got := Mod97("3214282912345698765432161182")
	want := 1
	if got != int64(want) {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestCheckIBAN(t *testing.T) {
	iban := "GB82 WEST 1234 5698 7654 32"
	b, _ := CheckIBAN(iban)
	got := b
	want := true
	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

}
