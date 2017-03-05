package Model

/*Bidder ...*/
type Bidder struct {
	id            int
	decision      bool
	bid           float64
	competitors   map[int]bool // map[bidder's id]isCompetitor
	bundle        map[int]int  // map[good's type]amount
	priority      float64
	payment       float64
	criticalValue float64
}

func CreateNewBidder(id int) *Bidder {
	b := new(Bidder)
	b.id = id
	b.competitors = make(map[int]bool)
	b.bundle = make(map[int]int)
	return b
}

func (b Bidder) GetID() int {
	return b.id
}

func (b *Bidder) SetID(id int) {
	b.id = id
}

func (b Bidder) GetDecision() bool {
	return b.decision
}

func (b *Bidder) SetDecision(decision bool) {
	b.decision = decision
}

func (b Bidder) GetBid() float64 {
	return b.bid
}

func (b *Bidder) SetBid(bid float64) {
	b.bid = bid
}

func (b Bidder) GetPriority() float64 {
	return b.priority
}

func (b *Bidder) SetPriority(priority float64) {
	b.priority = priority
}

func (b Bidder) GetPayment() float64 {
	return b.payment
}

func (b *Bidder) SetPayment(payment float64) {
	b.payment = payment
}

func (b Bidder) GetCriticalValue() float64 {
	return b.criticalValue
}

func (b *Bidder) SetCriticalValue(criticalValue float64) {
	b.criticalValue = criticalValue
}

func (b Bidder) GetBundleInstance(goodType int) int {
	return b.bundle[goodType]
}

func (b Bidder) GetBundle() map[int]int {
	return b.bundle
}

func (b *Bidder) SetBundleInstance(goodType int, amount int) {
	b.bundle[goodType] = amount
}

func (b Bidder) IsCompetitor(bidderID int) bool {
	return b.competitors[bidderID]
}

func (b Bidder) GetAllCompetitor() map[int]bool {
	return b.competitors
}

func (b *Bidder) SetCompetitor(bidderID int, iscompetitor bool) {
	b.competitors[bidderID] = iscompetitor
}

func (b Bidder) GetBundleInstancesCount() int {
	count := 0
	for key := range b.bundle {
		count += b.bundle[key]
	}
	return count
}

func (b Bidder) GetcompetitorCount() int {
	count := 0
	for key := range b.competitors {
		if b.competitors[key] == true {
			count++
		}
	}
	return count
}
