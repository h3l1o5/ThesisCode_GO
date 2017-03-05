package CombinatorialAuction

import (
	model "ThesisCode/Model"
	"fmt"
	"math/rand"
	"os"
)

type auction struct {
	totalBidder             int
	totalGood               int
	chance                  float64
	instanceRangeOfEachGood int
	warehouseOriginal       map[int]int
	warehouseNow            map[int]int
	bidders                 []*model.Bidder
}

/*public functions*/
func CreateAuction(totalBidder int, totalGood int, chance float64, instanceRangeOfEachGood int) []*model.Bidder{
	auct := new(auction)
	auct.totalBidder = totalBidder
	auct.totalGood = totalGood
	auct.chance = chance
	auct.instanceRangeOfEachGood = instanceRangeOfEachGood
	auct.createBidders()
	auct.createGoods()
	auct.letBiddersChooseGood()
	auct.letBiddersKnowTheirCompetitors()
	auct.letBiddersDecideTheirBid()
	auct.processInGameBasedWay("LOS02")
	
	return auct.bidders
}

func (a *auction) GetBidders() []*model.Bidder {
	return a.bidders
}

/*private functions*/
func (a *auction) createBidders() {
	a.bidders = make([]*model.Bidder, a.totalBidder)
	for i := 0; i < a.totalBidder; i++ {
		a.bidders[i] = model.CreateNewBidder(i)
	}
}

func (a *auction) createGoods() {
	a.warehouseOriginal = make(map[int]int)
	a.warehouseNow = make(map[int]int)
	switch a.instanceRangeOfEachGood {
	case 0:
		fmt.Println("good's instance range must greater than 0")
		os.Exit(2)
	case 1:
		for i := 0; i < a.totalGood; i++ {
			a.warehouseOriginal[i] = 1
			a.warehouseNow[i] = 1
		}
	default:
		for i := 0; i < a.totalGood; i++ {
			randomNumber := rand.Intn(a.instanceRangeOfEachGood) + 1
			a.warehouseOriginal[i] = randomNumber
			a.warehouseNow[i] = randomNumber
		}
	}
}

func (a *auction) letBiddersChooseGood() {
	for i := 0; i < a.totalBidder; i++ {
		bidder := a.bidders[i]
		chooseNothing := true
		for chooseNothing {
			for j := 0; j < a.totalGood; j++ {
				if rand.Intn(1000) < int(a.chance*10) {
					chooseNothing = false
					bidder.SetBundleInstance(j, rand.Intn(a.warehouseOriginal[j])+1)
				} else {
					bidder.SetBundleInstance(j, 0)
				}
			}
		}
	}
}

func (a *auction) letBiddersKnowTheirCompetitors() {
	for i := 0; i < a.totalBidder; i++ {
		bidder := a.bidders[i]
		for j := i + 1; j < a.totalBidder; j++ {
			bidder2 := a.bidders[j]
			for k := 0; k < a.totalGood; k++ {
				if bidder.GetBundleInstance(k) > 0 && bidder2.GetBundleInstance(k) > 0 {
					bidder.SetCompetitor(bidder2.GetID(), true)
					bidder2.SetCompetitor(bidder.GetID(), true)
					break
				} else {
					bidder.SetCompetitor(bidder2.GetID(), false)
					bidder2.SetCompetitor(bidder.GetID(), false)
				}
			}
		}
	}
}

func (a *auction) letBiddersDecideTheirBid() {
	priceOfGood := make([]float64, a.totalGood)
	for i := 0; i < a.totalGood; i++ {
		priceOfGood[i] = rand.Float64()*10 + 10 // range: 10-20
	}

	for i := 0; i < a.totalBidder; i++ {
		bidder := a.bidders[i]
		bid := 0.0
		for j := 0; j < a.totalGood; j++ {
			if bidder.GetBundleInstance(j) > 0 {
				valuationOfThisGood := rand.NormFloat64()*priceOfGood[i]/10 + priceOfGood[i] // normal dist.
				// for one more instance take in each type of good, +5% valuation
				bid += valuationOfThisGood * float64(bidder.GetBundleInstance(j)) * (1.0 + float64(bidder.GetBundleInstance(j)/20))
			}
		}
		bidder.SetBid(bid)
	}
}
