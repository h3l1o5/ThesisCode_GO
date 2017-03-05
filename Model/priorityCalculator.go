package Model

import (
	"math"
	"os"
)

func (b *Bidder) CalculatePriority(winnerDeterminationAlgo string) {
	switch winnerDeterminationAlgo {
	case "LOS02":
		b.SetPriority(b.GetBid() / math.Sqrt(float64(b.GetBundleInstancesCount())))
	case "ours":
		b.SetPriority(b.GetBid() / math.Sqrt(float64(b.GetBundleInstancesCount()*(b.GetcompetitorCount()+1))))
	default:
		os.Exit(2)
	}
}
