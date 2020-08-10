// Using simple graphics library from https://github.com/fogleman/gg

package main

import (
	"github.com/fogleman/gg"
	"log"
)

const (
	w = 10000
	h = 10000

	N = 1000
	r = h / N / 5
)

var (
	// Primes up to 1000
	primes = []int{2,3,5,7,11,13,17,19,23,29,31,37}//,41,43,47,53,59,61,67,71,73,79,83,89,97,101}//,103,107,109,113,127,131,137,139,149,151,157,163,167,173,179,181,191,193,197,199,211,223,227,229,233,239,241,251,257,263,269,271,27,281,283,293,307,311,313,317,331,337,347,349,353,359,367,373,379,383,389,397,401,409,419,421,431,433,439,443,449,457,461,463,467,479,487,491,499,503,509,521,523,541,547,557,563,569,571,577,587,593,599,601,607,613,617,619,631,641,643,647,653,659,661,673,677,683,691,701,709,719,727,733,739,743,751,757,761,769,773,787,797,809,811,821,823,827,829,839,853,857,859,863,877,881,883,887,907,911,919,929,937,941,947,953,967,971,977,983,991,997}
)

func calcOmega(x int) int {
	if x == 0 {
		return -999999
	}

	omega := 0
	for _, p := range primes {
		for x % p == 0 {
			x /= p
			omega++
		}
	}
	return omega
}

func setMarkerWidth(dc *gg.Context, x int) {
	if x % 10 == 0 {
		dc.SetLineWidth(h / N / 5 * 4)
	} else if x % 5 == 0 {
		dc.SetLineWidth(h / N / 5 * 2)
	} else {
		dc.SetLineWidth(h / N / 5 * 1)
	}
}

func drawGrid(dc *gg.Context) {
	dc.SetRGBA(0, 0, 1, 1)

	for x := 0; x <= N; x++ {
		setMarkerWidth(dc, x)

		px := float64(x * w) / N
		dc.DrawLine(px, h-0, px, h-r)
		dc.Stroke()
	}

	for y := 0; y <= N; y++ {
		setMarkerWidth(dc, y)

		py := float64(y * h) / N
		dc.DrawLine(0, h-py, r, h-py)
		dc.Stroke()
	}

	dc.NewSubPath()
	dc.MoveTo(0, 0)
	dc.LineTo(w, h)
	dc.ClosePath()
}

func main() {
	dc := gg.NewContext(w, h)

	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)

	drawGrid(dc)


	dc.SetRGB(1, 0, 0)

	for x := 0; x <= N; x++ {
		log.Printf("x = %d", x)
		for y:= 0; y <= N; y++ {
			normSq := x * x + y * y

			twiceOmega := calcOmega(normSq)
			//omega := twiceOmega / 2.0     // because norm is the square root of the norm, so we need to halve the omega

			if twiceOmega == 2 {
				px := float64(x * w) / N
				py := h - float64(y * h) / N
				dc.DrawCircle(px, py, r)

				dc.Fill()
			}
		}
	}

/*
	w := 0.1
	for i := 100; i <= 900; i += 20 {
		x := float64(i)
		dc.DrawLine(x+50, 0, x-50, 1000)
		dc.SetLineWidth(w)
		dc.Stroke()
		w += 0.1
	}
*/	dc.SavePNG("out.png")
}