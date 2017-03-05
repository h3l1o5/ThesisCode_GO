package Model

import (
	"math"
	"os"
)

func (b *Bidder) CalculateCriticalValue(winnerDeterminationAlgo string, bidders []*Bidder, wareHouse map[int]int) {
	if b.GetDecision() {
		winner := b
		temp := -1.0

		// find all competitors who will become winner if current winner was not win.
		winner.SetDecision(false)
		for index := range bidders {
			ID := bidders[index].GetID()
			if winner.IsCompetitor(ID) && !bidders[index].GetDecision() {
				if bidders[index].MakeNewDecision(wareHouse, bidders) && bidders[index].GetPriority() > temp {
					temp = bidders[index].GetPriority()
				}
			}
		}
		winner.SetDecision(true)

		if temp == -1.0 {
			winner.SetCriticalValue(0)
			winner.SetPayment(0)
		} else {
			switch winnerDeterminationAlgo {
			case "LOS02":
				winner.SetCriticalValue(temp * math.Sqrt(float64(winner.GetBundleInstancesCount())))
				winner.SetPayment(temp * math.Sqrt(float64(winner.GetBundleInstancesCount())))
			case "ours":
				winner.SetCriticalValue(temp * math.Sqrt(float64(winner.GetBundleInstancesCount()*(winner.GetcompetitorCount()+1))))
				winner.SetPayment(temp * math.Sqrt(float64(winner.GetBundleInstancesCount()*(winner.GetcompetitorCount()+1))))
			default:
				os.Exit(2)
			}
		}
	} else {
		loser := b
		temp := 99999999.9

		// find the neighbor who is winner and has highest priority.
		for index := range bidders {
			ID := bidders[index].GetID()
			if loser.IsCompetitor(ID) && bidders[index].GetDecision() {
				bidders[index].SetDecision(false)
				if loser.MakeNewDecision(wareHouse, bidders) {
					if bidders[index].GetPriority() < temp {
						temp = bidders[index].GetPriority()
					}
				}
				bidders[index].SetDecision(true)
			}
		}

		switch winnerDeterminationAlgo {
		case "LOS02":
			loser.SetCriticalValue(temp * math.Sqrt(float64(loser.GetBundleInstancesCount())))
		case "ours":
			loser.SetCriticalValue(temp * math.Sqrt(float64(loser.GetBundleInstancesCount()*(loser.GetcompetitorCount()+1))))
		default:
			os.Exit(2)
		}

		loser.SetPayment(0)
	}
}
