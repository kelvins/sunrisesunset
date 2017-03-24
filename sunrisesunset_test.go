package sunrisesunset

import (
	"testing"
	"time"
)

func TestGetSunriseSunset(t *testing.T) {

  date := time.Date(2017, 3, 23, 0, 0, 0, 0, time.UTC)

	// Test invalid parameters

	// Table tests
	var invalidParameters = []struct {
		latitude      float64
		longitude     float64
		utcOffset     float64
		date          time.Time
		expectedError string
	}{
		{     -95.0, -46.704082,  -3.0, date, "Latitude invalid"},
		{     100.0, -46.704082,  -3.0, date, "Latitude invalid"},
		{-23.545570,     -185.0,  -3.0, date, "Longitude invalid"},
		{-23.545570,      190.0,  -3.0, date, "Longitude invalid"},
		{-23.545570, -46.704082, -15.0, date, "UTC offset invalid"},
		{-23.545570, -46.704082,  18.0, date, "UTC offset invalid"},
		{-23.545570, -46.704082,  -3.0, time.Date(1000, 1, 1, 0, 0, 0, 0, time.UTC), "Date invalid"},
		{-23.545570, -46.704082,  -3.0, time.Date(3000, 1, 1, 0, 0, 0, 0, time.UTC), "Date invalid"},
	}

	// Test with all values in the table
	for _, pair := range invalidParameters {
		_, _, err := GetSunriseSunset(pair.latitude, pair.longitude, pair.utcOffset, pair.date)
		if err == nil {
			t.Error(
				"Expect an error",
			)
		}
	}

	// Test with valid values

	// Table tests
	var tTests = []struct {
		latitude  float64
		longitude float64
		utcOffset float64
		date      time.Time
		sunrise   time.Time
		sunset    time.Time
	}{
		{-23.545570, -46.704082, -3.0, date, time.Date(1, 1, 1, 6, 11, 44, 0, time.UTC), time.Date(1, 1, 1, 18, 14, 27, 0, time.UTC)}, // Sao Paulo - Brazil
		{36.7201600, -4.4203400,  1.0, date, time.Date(1, 1, 1, 7, 16, 45, 0, time.UTC), time.Date(1, 1, 1, 19, 32, 10, 0, time.UTC)}, // Málaga - Spain
		{ 28.613084,  77.209168,  5.5, date, time.Date(1, 1, 1, 6, 21, 45, 0, time.UTC), time.Date(1, 1, 1, 18, 34, 07, 0, time.UTC)}, // Nova Delhi - India
    { 32.755701, -96.797296, -5.0, date, time.Date(1, 1, 1, 7, 26, 34, 0, time.UTC), time.Date(1, 1, 1, 19, 41, 07, 0, time.UTC)}, // Dallas - United States of America
	}

	// Test with all values in the table
	for _, pair := range tTests {
		sunrise, sunset, err := GetSunriseSunset(pair.latitude, pair.longitude, pair.utcOffset, pair.date)

		if err != nil {
			t.Error(
				"Expect: nil",
				"Received: ", err,
			)
		}
		if !sunrise.Equal(pair.sunrise) {
			t.Error(
				"Expected: ", pair.sunrise,
				"Received: ", sunrise,
			)
		}
		if !sunset.Equal(pair.sunset) {
			t.Error(
				"Expected: ", pair.sunset,
				"Received: ", sunset,
			)
		}
	}
}
