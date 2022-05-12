package util

import "testing"

func TestStringToTime(t *testing.T) {
	var convertedTime = StringToTime("2019-01-01T10:10:10")
	if convertedTime.Year() != 2019 {
		t.Errorf("Espero que o ano seja 2019")
	}

	if convertedTime.Month() != 1 {
		t.Errorf("Espero que o mês seja 1")
	}

	if convertedTime.Day() != 1 {
		t.Errorf("Espero que o dia seja 1")
	}

	if convertedTime.Hour() != 10 {
		t.Errorf("Espero que a hora seja 10")
	}

}
