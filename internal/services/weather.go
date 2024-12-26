package services

import (
	"fmt"
	"math"

	"github.com/julians58/weather-predictor/internal/helpers"
	"github.com/julians58/weather-predictor/internal/models"
)

func CalculateWeather(day int, maxPerimeter *float64, rainBelowPeak *bool) (models.Weather, bool, error) {
	x1, y1 := helpers.CalculatePosition(5, -1, float64(day))
	x2, y2 := helpers.CalculatePosition(10, 5, float64(day))
	x3, y3 := helpers.CalculatePosition(20, -3, float64(day))

	slope1 := helpers.CalculateSlope(x1, y1)
	slope2 := helpers.CalculateSlope(x2, y2)
	slope3 := helpers.CalculateSlope(x3, y3)
	alignedWithOrigin := helpers.ArePointsAligned([]*float64{slope1, slope2, slope3})

	var weatherCondition string
	perimeter := 0.0
	var alignedBetweenThem bool = false
	var originInside bool = false
	var setPrevAsPeak bool = false
	if !alignedWithOrigin {
		slope12 := helpers.CalculateSlopeBetweenPoints(x1, y1, x2, y2)
		slope23 := helpers.CalculateSlopeBetweenPoints(x2, y2, x3, y3)

		alignedBetweenThem = helpers.SlopesApproximatelyEqual(slope12, slope23, 0.2)
	}
	if !alignedWithOrigin && !alignedBetweenThem {
		area := helpers.CalculateArea(x1, y1, x2, y2, x3, y3)
		perimeter = math.Sqrt(math.Pow(x2-x1, 2)+math.Pow(y2-y1, 2)) +
			math.Sqrt(math.Pow(x3-x2, 2)+math.Pow(y3-y2, 2)) +
			math.Sqrt(math.Pow(x1-x3, 2)+math.Pow(y1-y3, 2))

		originInside = math.Abs(area-helpers.CalculateArea(0, 0, x2, y2, x3, y3)-helpers.CalculateArea(x1, y1, 0, 0, x3, y3)-helpers.CalculateArea(x1, y1, x2, y2, 0, 0)) <= 1e-9
	}
	if alignedWithOrigin {
		weatherCondition = "Sequia"
	} else if alignedBetweenThem {
		weatherCondition = "Condiciones optimas de presion y temperatura"
	} else if originInside {
		weatherCondition = "Lluvia"
		if !*rainBelowPeak {
			if perimeter > *maxPerimeter {
				*maxPerimeter = perimeter

			} else if perimeter < *maxPerimeter {
				*rainBelowPeak = true
				setPrevAsPeak = true
			}
		}
	} else {
		weatherCondition = "Condiciones normales"
		*rainBelowPeak = false
		*maxPerimeter = 0.0
	}

	fmt.Printf("day: %d, condition: %s\n", day+1, weatherCondition)
	return models.Weather{
		Day:       day + 1,
		Condition: weatherCondition,
	}, setPrevAsPeak, nil
}
