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
func rad2deg(radians float64) float64 {
	return radians * (180.0 / math.Pi)
}

// Convert degrees to radians
func deg2rad(degrees float64) float64 {
	return degrees * (math.Pi / 180.0)
}

// Calculate the equation of time (minutes) based on the formula:
// 4*rad2deg(multiFactor*sin(2*deg2rad(geomMeanLongSun))-2*eccentEarthOrbit*sin(deg2rad(geomMeanAnomSun))+4*eccentEarthOrbit*multiFactor*sin(deg2rad(geomMeanAnomSun))*cos(2*deg2rad(geomMeanLongSun))-0.5*multiFactor*multiFactor*sin(4*deg2rad(geomMeanLongSun))-1.25*eccentEarthOrbit*eccentEarthOrbit*sin(2*deg2rad(geomMeanAnomSun)))
// multiFactor - The Multi Factor vector calculated in the calculate function
// geomMeanLongSun - The Geom Mean Long Sun vector calculated by the calcGeomMeanLongSun function
// eccentEarthOrbit - The Eccent Earth vector calculated by the calcEccentEarthOrbit function
// geomMeanAnomSun - The Geom Mean Anom Sun vector calculated by the calcGeomMeanAnomSun function
// Return the equation of time slice
func calcEquationOfTime(multiFactor []float64, geomMeanLongSun []float64, eccentEarthOrbit []float64, geomMeanAnomSun []float64) (equationOfTime []float64) {

	if len(multiFactor) != len(geomMeanLongSun)  ||
		 len(multiFactor) != len(eccentEarthOrbit) ||
		 len(multiFactor) != len(geomMeanAnomSun) {
		return
	}

	for index := 0; index < len(multiFactor); index++ {
		a := multiFactor[index] * math.Sin(2.0*deg2rad(geomMeanLongSun[index]))
		b := 2.0 * eccentEarthOrbit[index] * math.Sin(deg2rad(geomMeanAnomSun[index]))
		c := 4.0 * eccentEarthOrbit[index] * multiFactor[index] * math.Sin(deg2rad(geomMeanAnomSun[index]))
		d := math.Cos(2.0 * deg2rad(geomMeanLongSun[index]))
		e := 0.5 * multiFactor[index] * multiFactor[index] * math.Sin(4.0*deg2rad(geomMeanLongSun[index]))
		f := 1.25 * eccentEarthOrbit[index] * eccentEarthOrbit[index] * math.Sin(2.0*deg2rad(geomMeanAnomSun[index]))
		temp := 4.0 * rad2deg(a-b+c*d-e-f)
		equationOfTime = append(equationOfTime, temp)
	}
	return
}

// Calculate the HaSunrise in degrees based on the formula: rad2deg(acos(cos(deg2rad(90.833))/(cos(deg2rad(latitude))*cos(deg2rad(sunDeclination)))-tan(deg2rad(latitude))*tan(deg2rad(sunDeclination))))
// latitude - The latitude defined by the user
// sunDeclination - The Sun Declination calculated by the calcSunDeclination function
// Return the HaSunrise slice
func calcHaSunrise(latitude float64, sunDeclination []float64) (haSunrise []float64) {
	for index := 0; index < len(sunDeclination); index++ {
		temp := rad2deg(math.Acos(math.Cos(deg2rad(90.833))/(math.Cos(deg2rad(latitude))*math.Cos(deg2rad(sunDeclination[index]))) - math.Tan(deg2rad(latitude))*math.Tan(deg2rad(sunDeclination[index]))))
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