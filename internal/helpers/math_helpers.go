package helpers

import "math"

func DegreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func Round(value float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return math.Round(value*shift) / shift
}

// Calculate cartesian coordinates for a point in polar coordinates
func CalculatePosition(distance, angularVelocity, day float64) (float64, float64) {
	angle := angularVelocity * day
	angleRadians := DegreesToRadians(angle)

	x := distance * math.Cos(angleRadians)
	y := distance * math.Sin(angleRadians)

	// Round x and y to 1 decimal place
	x = Round(x, 1)
	y = Round(y, 1)

	return x, y
}

// Calculate the slope of a point relative to the origin
func CalculateSlope(x, y float64) *float64 {
	if x != 0 {
		slope := y / x
		return &slope
	}
	return nil // Vertical line, no slope
}

func SlopesApproximatelyEqual(p1, p2 *float64, tolerance float64) bool {
	if p1 == nil || p2 == nil {
		return false
	}
	return math.Abs(*p1-*p2) <= tolerance
}

func CalculateSlopeBetweenPoints(x1, y1, x2, y2 float64) *float64 {
	if x2-x1 != 0 {
		slope := (y2 - y1) / (x2 - x1)
		return &slope
	}
	return nil // Vertical line, no slope
}

// Check if three points are aligned with the origin
func ArePointsAligned(slopes []*float64) bool {
	// Check if all slopes are nil or if all non-nil slopes are equal
	var firstSlope *float64
	for _, slope := range slopes {
		if slope != nil {
			if firstSlope == nil {
				firstSlope = slope
			} else if *slope != *firstSlope {
				return false // Slopes are not equal
			}
		}
	}
	return true // All slopes are either nil or equal
}

// Calculate the area of a triangle given three points
func CalculateArea(x1, y1, x2, y2, x3, y3 float64) float64 {
	return math.Abs((x1*(y2-y3) + x2*(y3-y1) + x3*(y1-y2)) / 2.0)
}
