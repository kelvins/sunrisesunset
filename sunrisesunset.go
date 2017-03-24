/*
 * Package details
 */
package sunrisesunset

import (
  "math"
//"time"
//"errors"
)

// Convert radians to degrees
func rad2deg(radians float64) (float64) {
  return radians * (180.0/math.Pi)
}

// Convert degrees to radians
func deg2rad(degrees float64) (float64) {
  return degrees * (math.Pi/180.0)
}

// Calculate the HaSunrise in degrees based on the formula: rad2deg(acos(cos(deg2rad(90.833))/(cos(deg2rad(latitude))*cos(deg2rad(sunDeclination)))-tan(deg2rad(latitude))*tan(deg2rad(sunDeclination))))
// latitude - The latitude defined by the user
// sunDeclination - The Sun Declination calculated by the calcSunDeclination function
// Return the HaSunrise slice
func calcHaSunrise(latitude float64, sunDeclination []float64) (haSunrise []float64) {
    for index := 0; index < len(sunDeclination); index++ {
        temp := rad2deg(acos(cos(deg2rad(90.833))/(cos(deg2rad(latitude))*cos(deg2rad(sunDeclination[index])))-tan(deg2rad(latitude))*tan(deg2rad(sunDeclination[index]))))
        haSunrise = append(haSunrise, temp)
    }
    return
}


// Calculate the Solar Noon based on the formula: (720 - 4 * longitude - equationOfTime + utcOffset * 60) * 60
// longitude - The longitude is defined by the user
// equationOfTime - The Equation of Time slice is calculated by the calcEquationOfTime function
// utcOffset - The UTC offset is defined by the user
// Return the Solar Noon slice
func calcSolarNoon(longitude float64, equationOfTime []float64, utcOffset float64) (solarNoon []float64) {
	for index := 0; index < len(equationOfTime); index++ {
	  temp := (720.0 - 4.0*longitude - equationOfTime[index] + utcOffset*60.0) * 60.0
		solarNoon = append(solarNoon, temp)
	}
	return
}
