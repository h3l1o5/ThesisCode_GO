package Model

func (b Bidder) MakeNewDecision(wareHouse map[int]int, bidders []*Bidder) bool {
	bundleOfThisBidder := b.GetBundle()
	for typeOfGood := range bundleOfThisBidder {
		if bundleOfThisBidder[typeOfGood] > 0 {
			instanceCountOfThisGood := wareHouse[typeOfGood]
			totalNeededOfThisGood := bundleOfThisBidder[typeOfGood]

			for index := range bidders {
				ID := bidders[index].GetID()
				if b.IsCompetitor(ID) && bidders[index].GetDecision() {
					competitor := bidders[index]
					if competitor.GetPriority() > b.GetPriority() {
						totalNeededOfThisGood += competitor.GetBundleInstance(typeOfGood)
					}
				}
			}
			if totalNeededOfThisGood > instanceCountOfThisGood {
				return false
			}
		}
	}
	return true
}
