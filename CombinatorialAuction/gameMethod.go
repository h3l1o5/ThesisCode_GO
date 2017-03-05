package CombinatorialAuction

func (a *auction) processInGameBasedWay(winnerDeterminationAlgo string) {
	var done bool

	for index := range a.bidders {
		a.bidders[index].CalculatePriority(winnerDeterminationAlgo)
	}

	for {
		done = true
		for i := 0; i < a.totalBidder; i++ {
			bidder := a.bidders[i]
			newDecision := bidder.MakeNewDecision(a.warehouseOriginal, a.bidders) // TODO
			if bidder.GetDecision() != newDecision {
				if newDecision == false {
					for j := 0; j < a.totalGood; j++ {
						a.warehouseNow[j] += bidder.GetBundleInstance(j)
					}
				}
				if newDecision == true {
					for j := 0; j < a.totalGood; j++ {
						a.warehouseNow[j] -= bidder.GetBundleInstance(j)
					}
				}
				bidder.SetDecision(newDecision)
				done = false
				break
			}
		}
		if done {
			break
		}
	}

	for index := range a.bidders {
		a.bidders[index].CalculateCriticalValue(winnerDeterminationAlgo, a.bidders, a.warehouseOriginal)
	}
	//	checkGoodStore()                                // TODO
}
