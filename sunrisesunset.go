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

// Calculate the Eccent Earth Orbit based on the formula: 0.016708634 - julianCentury * (0.000042037 + 0.0000001267 * julianCentury)
// julianCentury - Julian century calculated by the calcJulianCentury function
// Return The Eccent Earth Orbit slice
func calcEccentEarthOrbit(julianCentury []float64) (eccentEarthOrbit []float64) {
	for index := 0; index < len(julianCentury); index++ {
		temp := 0.016708634 - julianCentury[index]*(0.000042037+0.0000001267*julianCentury[index])
		eccentEarthOrbit = append(eccentEarthOrbit, temp)
	}
	return
}

// Calculate the Sun Eq Ctr based on the formula: sin(deg2rad(geomMeanAnomSun))*(1.914602-julianCentury*(0.004817+0.000014*julianCentury))+sin(deg2rad(2*geomMeanAnomSun))*(0.019993-0.000101*julianCentury)+sin(deg2rad(3*geomMeanAnomSun))*0.000289;
// julianCentury - Julian century calculated by the calcJulianCentury function
// geomMeanAnomSun - Geom Mean Anom Sun calculated by the calcGeomMeanAnomSun function
// Return The Sun Eq Ctr slice
func calcSunEqCtr(julianCentury []float64, geomMeanAnomSun []float64) (sunEqCtr []float64) {
	if len(julianCentury) != len(geomMeanAnomSun) {
		return
	}

	for index := 0; index < len(julianCentury); index++ {
		temp := math.Sin(deg2rad(geomMeanAnomSun[index]))*(1.914602-julianCentury[index]*(0.004817+0.000014*julianCentury[index])) + math.Sin(deg2rad(2*geomMeanAnomSun[index]))*(0.019993-0.000101*julianCentury[index]) + math.Sin(deg2rad(3*geomMeanAnomSun[index]))*0.000289
		sunEqCtr = append(sunEqCtr, temp)
	}
	return
}

// Calculate the Sun True Long in degrees based on the formula: sunEqCtr + geomMeanLongSun
// sunEqCtr - Sun Eq Ctr calculated by the calcSunEqCtr function
// geomMeanLongSun - Geom Mean Long Sun calculated by the calcGeomMeanLongSun function
// Return The Sun True Long slice
func calcSunTrueLong(sunEqCtr []float64, geomMeanLongSun []float64) (sunTrueLong []float64) {
	if len(sunEqCtr) != len(geomMeanLongSun) {
		return
	}

	for index := 0; index < len(sunEqCtr); index++ {
		temp := sunEqCtr[index] + geomMeanLongSun[index]
		sunTrueLong = append(sunTrueLong, temp)
	}
	return
}

// Calculate the Sun App Long in degrees based on the formula: sunTrueLong-0.00569-0.00478*sin(deg2rad(125.04-1934.136*julianCentury))
// sunTrueLong - Sun True Long calculated by the calcSunTrueLong function
// julianCentury - Julian century calculated by the calcJulianCentury function
// Return The Sun App Long slice
func calcSunAppLong(sunTrueLong []float64, julianCentury []float64) (sunAppLong []float64) {
	if len(sunTrueLong) != len(julianCentury) {
		return
	}

	for index := 0; index < len(sunTrueLong); index++ {
		temp := sunTrueLong[index] - 0.00569 - 0.00478*math.Sin(deg2rad(125.04-1934.136*julianCentury[index]))
		sunAppLong = append(sunAppLong, temp)
	}
	return
}

// Calculate the Mean Obliq Ecliptic in degrees based on the formula: 23+(26+((21.448-julianCentury*(46.815+julianCentury*(0.00059-julianCentury*0.001813))))/60)/60
// julianCentury - Julian century calculated by the calcJulianCentury function
// Return the Mean Obliq Ecliptic slice
func calcMeanObliqEcliptic(julianCentury []float64) (meanObliqEcliptic []float64) {
	for index := 0; index < len(julianCentury); index++ {
		temp := 23.0 + (26.0+(21.448-julianCentury[index]*(46.815+julianCentury[index]*(0.00059-julianCentury[index]*0.001813)))/60.0)/60.0
		meanObliqEcliptic = append(meanObliqEcliptic, temp)
	}
	return
}

// Calculate the Obliq Corr in degrees based on the formula: meanObliqEcliptic+0.00256*cos(deg2rad(125.04-1934.136*julianCentury))
// meanObliqEcliptic - Mean Obliq Ecliptic calculated by the calcMeanObliqEcliptic function
// julianCentury - Julian century calculated by the calcJulianCentury function
// Return the Obliq Corr slice
func calcObliqCorr(meanObliqEcliptic []float64, julianCentury []float64) (obliqCorr []float64) {
	if len(meanObliqEcliptic) != len(julianCentury) {
		return
	}

	for index := 0; index < len(julianCentury); index++ {
		temp := meanObliqEcliptic[index] + 0.00256*math.Cos(deg2rad(125.04-1934.136*julianCentury[index]))
		obliqCorr = append(obliqCorr, temp)
	}
	return
}

// Calculate the Sun Declination in degrees based on the formula: rad2deg(asin(sin(deg2rad(obliqCorr))*sin(deg2rad(sunAppLong))))
// obliqCorr - Obliq Corr calculated by the calcObliqCorr function
// sunAppLong - Sun App Long calculated by the calcSunAppLong function
// Return the sun declination slice
func calcSunDeclination(obliqCorr []float64, sunAppLong []float64) (sunDeclination []float64) {
	if len(obliqCorr) != len(sunAppLong) {
		return
	}

	for index := 0; index < len(obliqCorr); index++ {
		temp := rad2deg(math.Asin(math.Sin(deg2rad(obliqCorr[index])) * math.Sin(deg2rad(sunAppLong[index]))))
		sunDeclination = append(sunDeclination, temp)
	}
	return
}

// Calculate the equation of time (minutes) based on the formula:
// 4*rad2deg(multiFactor*sin(2*deg2rad(geomMeanLongSun))-2*eccentEarthOrbit*sin(deg2rad(geomMeanAnomSun))+4*eccentEarthOrbit*multiFactor*sin(deg2rad(geomMeanAnomSun))*cos(2*deg2rad(geomMeanLongSun))-0.5*multiFactor*multiFactor*sin(4*deg2rad(geomMeanLongSun))-1.25*eccentEarthOrbit*eccentEarthOrbit*sin(2*deg2rad(geomMeanAnomSun)))
// multiFactor - The Multi Factor vector calculated in the calculate function
// geomMeanLongSun - The Geom Mean Long Sun vector calculated by the calcGeomMeanLongSun function
// eccentEarthOrbit - The Eccent Earth vector calculated by the calcEccentEarthOrbit function
// geomMeanAnomSun - The Geom Mean Anom Sun vector calculated by the calcGeomMeanAnomSun function
// Return the equation of time slice
func calcEquationOfTime(multiFactor []float64, geomMeanLongSun []float64, eccentEarthOrbit []float64, geomMeanAnomSun []float64) (equationOfTime []float64) {

	if len(multiFactor) != len(geomMeanLongSun) ||
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
