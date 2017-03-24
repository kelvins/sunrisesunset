/*
 * Package details
 */
package sunrisesunset

import (
//"time"
//"math"
//"errors"
)

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
