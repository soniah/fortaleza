package fortaleza

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var cartTest1 string = `
grandtotal: 0
rows:
- code: ult_small
  quantity: 3
  completed: false
  total: 0
- code: ult_large
  quantity: 1
  completed: false
  total: 0
`

var cartTest2 string = `
grandtotal: 0
rows:
- code: ult_small
  quantity: 2
  completed: false
  total: 0
- code: ult_large
  quantity: 4
  completed: false
  total: 0
`

var cartTest3 string = `
grandtotal: 0
rows:
- code: ult_small
  quantity: 1
  completed: false
  total: 0
- code: ult_medium
  quantity: 2
  completed: false
  total: 0
`

var cartTest4 string = `
grandtotal: 0
rows:
- code: ult_small
  quantity: 1
  completed: false
  total: 0
- code: 1gb
  quantity: 1
  completed: false
  total: 0
- code: promo
  quantity: 0
  completed: false
  total: 0
`

func TestCalculations(t *testing.T) {
	// catalogue
	c := strings.NewReader(catalogue)
	catalogData, err := ImportCatalogue(c)
	assert.Nil(t, err)

	// pricings
	p := strings.NewReader(pricings)
	pricingsData, err := ImportPricings(p)
	assert.Nil(t, err)

	// cart1
	r1 := strings.NewReader(cartTest1)
	cart1, err := ImportCart(r1)
	assert.Nil(t, err)
	cart1 = AllPricing(cart1, catalogData, pricingsData)
	fmt.Printf("%s\n", cart1)

	// cart2
	r2 := strings.NewReader(cartTest2)
	cart2, err := ImportCart(r2)
	assert.Nil(t, err)
	cart2 = AllPricing(cart2, catalogData, pricingsData)
	fmt.Printf("%s\n", cart2)

	// cart3
	r3 := strings.NewReader(cartTest3)
	cart3, err := ImportCart(r3)
	assert.Nil(t, err)
	cart3 = AllPricing(cart3, catalogData, pricingsData)
	fmt.Printf("%s\n", cart3)

	// cart4
	r4 := strings.NewReader(cartTest4)
	cart4, err := ImportCart(r4)
	assert.Nil(t, err)
	cart4 = AllPricing(cart4, catalogData, pricingsData)
	fmt.Printf("%s\n", cart4)
}
