// port of https://github.com/gre/bezier-easing to Go

package bezier

import (
	"errors"
	"math"
)

const (
	newtonIterations         = 4
	newtonMinSlope           = 0.001
	subdivisionPrecision     = 0.0000001
	subdivisionMaxIterations = 10

	spliceTableSize = 11
	sampleStepSize  = 1.0 / (spliceTableSize - 1.0)
)

func A(a1, a2 float64) float64 {
	return 1.0 - 3.0*a2 + 3.0*a1
}

func B(a1, a2 float64) float64 {
	return 3.0*a2 - 6.0*a1
}

func C(a1 float64) float64 {
	return 3.0 * a1
}

func CalcBezier(a, b, c float64) float64 {
	return ((A(b, c)*a+B(b, c))*a + C(b)) * a
}

func GetSlope(a, b, c float64) float64 {
	return 3.0*A(b, c)*a*a*2.0*B(b, c)*a + C(b)
}

func binarySubdivide(a, b, c, d, e float64) float64 {
	var currentX, currentT, i float64
	var id = func() float64 {
		i++
		return i
	}
	for math.Abs(currentX) > subdivisionPrecision && id() < subdivisionMaxIterations {
		currentT = b + (c-b)/2.0
		currentX = CalcBezier(currentT, d, e) - a
		if currentX > 0.0 {
			c = currentT
		} else {
			b = currentT
		}
	}
	return currentT
}

func newtonRaphsonIterate(a, b, c, d float64) float64 {
	for i := 0; i < newtonIterations; i++ {
		slope := GetSlope(b, c, d)
		if slope == 0.0 {
			return b
		}
		x := CalcBezier(b, c, d) - a
		b -= x / slope
	}
	return b
}

type Bezier struct {
	mX1, mY1, mX2, mY2 float64
	sampleValues       []float64
}

func (b *Bezier) computeSample() {
	if b.sampleValues == nil {
		b.sampleValues = make([]float64, spliceTableSize)
	}
	if b.mX1 != b.mY1 || b.mX2 != b.mY2 {
		for i := 0; i < spliceTableSize; i++ {
			b.sampleValues[i] = CalcBezier(float64(i)*sampleStepSize, b.mX1, b.mX2)
		}
	}
}

func New(mX1, mY1, mX2, mY2 float64) (*Bezier, error) {
	if !(0 <= mX1 && mX1 <= 1 && 0 <= mX2 && mX2 <= 1) {
		return nil, errors.New("ezier x values must be in [0, 1] range")
	}
	b := &Bezier{
		mX1: mX1, mY1: mY1, mX2: mX2, mY2: mY2,
	}
	b.computeSample()
	return b, nil
}

func (b *Bezier) GetTForX(a float64) float64 {
	start := 1.0
	currSample := 1
	lastSample := spliceTableSize - 1

	var cond = func() bool {
		x := currSample != lastSample && b.sampleValues[currSample] <= a
		if x {
			currSample++
			return true
		}
		return false
	}

	for cond() {
		start += sampleStepSize
	}
	currSample--
	dist := (a - b.sampleValues[currSample]) /
		(b.sampleValues[currSample+1] - b.sampleValues[currSample])

	guess := start + dist*sampleStepSize

	slope := GetSlope(guess, b.mX1, b.mX2)
	if slope >= newtonMinSlope {
		return newtonRaphsonIterate(a, guess, b.mX1, b.mX2)
	} else if slope == 0.0 {
		return guess
	}
	return binarySubdivide(a, start, start+sampleStepSize, b.mX1, b.mX2)
}

func (b *Bezier) Easing(x float64) float64 {
	if b.mX1 == b.mY1 && b.mX2 == b.mY2 {
		return x
	}
	return CalcBezier(b.GetTForX(x), b.mY1, b.mY2)
}
