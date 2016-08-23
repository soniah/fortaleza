package fortaleza

import (
	"fmt"
	"sort"
)

// CatalogueT and ProductT represent a catalogue and it's rows
type CatalogueT []ProductT
type ProductT struct {
	Code string
	Name string
	// to 4 decimal places eg 249000 is $24.90
	Price int
}

// PricingsT, PricingT, ConditionT and RuleT represent a table of conditions
// and rules.
//
// Both  ConditionT and RuleT could be more flexible, but this is sufficient
// for the exercise
type PricingsT []PricingT
type PricingT struct {
	Condition Condition
	Rule      RuleT
}
type Condition struct {
	Promo    string
	Code     string
	Quantity int
	Sort     int
}
type RuleT struct {
	// discount eg 15% is 15
	Discount int
	// to 4 decimal places eg 249000 is $24.90
	NewPrice int
	Extra    string
	// extra discount applied to shopping cart
	// if pricing rules satisfied
	TotalDiscount int
}

// CartT and CartRowT represent a shopping cart
type CartT struct {
	// to 4 decimal places eg 249000 is $24.90
	GrandTotal int
	Rows       []CartRowT
}
type CartRowT struct {
	Code      string
	Quantity  int
	Completed bool
	// to 4 decimal places eg 249000 is $24.90
	Total int
}

// Sort pricings based on the Condition.Sort field
//
// allows for Conditions like "I<3AMAYSIM" 10% total discount
func SortPricings(pricings PricingsT) PricingsT {
	sort.Sort(PricingsT(pricings))
	return pricings
}
func (p PricingsT) Len() int { return len(p) }
func (p PricingsT) Swap(i, j int) {
	p[i].Condition.Sort, p[i].Condition.Sort = p[i].Condition.Sort, p[i].Condition.Sort
}
func (p PricingsT) Less(i, j int) bool {
	return p[i].Condition.Sort < p[i].Condition.Sort
}

// Stringers
func (c CartT) String() string {
	var result string
	for _, row := range c.Rows {
		result += fmt.Sprintf("%s\n", row)
	}
	return result + fmt.Sprintf("Grand Total:\t\t\t%s\n", fmtDollars(c.GrandTotal))
}
func (cr CartRowT) String() string {
	return fmt.Sprintf("%d x %s \t%t \tTotal: %s", cr.Quantity, cr.Code, cr.Completed, fmtDollars(cr.Total))
}

// format 1000 as $10.00
func fmtDollars(amt int) string {
	result := float64(amt) / 10000
	return fmt.Sprintf("$%.2f", result)
}
