package main

import (
	auction "ThesisCode/CombinatorialAuction"
	"fmt"
	"math/rand"
	"time"
)

func init() {

}

func main() {
	TOTAL_PLAYER := 10
	TOTAL_GOOD_TYPE := 500
	CHANCE := 5.0
	INSTANCE_RANGE_OF_EACH_GOOD := 10
	TOTAL_ITERATION := 1000

	revenue := 0.0
	for i := 0; i < TOTAL_ITERATION; i++ {
		rand.Seed(time.Now().UnixNano())
		bidders := auction.CreateAuction(TOTAL_PLAYER, TOTAL_GOOD_TYPE, CHANCE, INSTANCE_RANGE_OF_EACH_GOOD)
		for index := range bidders {
			if bidders[index].GetDecision() {
				revenue += bidders[index].GetPayment()
			}
		}
	}

	fmt.Println("revenue: ", revenue/float64(TOTAL_ITERATION))

}
