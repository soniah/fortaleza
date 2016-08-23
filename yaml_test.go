package fortaleza

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var catalogue = `
---
- code: ult_small
  name: Unlimited 1GB
  price: 249000
- code: ult_medium
  name: Unlimited 2GB
  price: 299000
- code: ult_large
  name: Unlimited 4GB
  price: 449000
- code: 1gb
  name: Unlimited 2GB
  price: 99000
- code: promo
  name: "1<3AMAYSIM"
  price: 0
`

func TestImportCatalogue(t *testing.T) {
	r := strings.NewReader(catalogue)
	products, err := ImportCatalogue(r)

	assert.Nil(t, err)
	assert.Equal(t, products[0].Code, "ult_small")
	assert.Equal(t, products[1].Name, "Unlimited 2GB")
	assert.Equal(t, products[1].Price, 299000)
}

var pricings string = `
- condition:
    promo: ""
    code: ult_small
    quantity: 3
    sort: 10
  rule:
    discount: 33
    newprice: 0
    extra: ""
    totaldiscount: 0

- condition:
    promo: ""
    code: 1gb
    quantity: 0
    sort: 10
  rule:
    discount: 0
    newprice: 0
    extra: ""
    totaldiscount: 0

- condition:
    promo: ""
    code: ult_large
    quantity: 3
    sort: 20
  rule:
    discount: 0
    newprice: 399000
    extra: ""
    totaldiscount: 0

- condition:
    promo: ""
    code: ult_medium
    quantity: 0
    sort: 30
  rule:
    discount: 0
    newprice: 0
    extra: 1gb
    totaldiscount: 0

- condition:
    promo: "promo"
    code: ""
    quantity: 0
    sort: 99
  rule:
    discount: 0
    newprice: 0
    extra: ""
    totaldiscount: 10
`

func TestImportPricings(t *testing.T) {
	r := strings.NewReader(pricings)
	result, err := ImportPricings(r)

	assert.Nil(t, err)
	assert.Equal(t, result[0].Condition.Promo, "")
	assert.Equal(t, result[1].Condition.Sort, 10)
	assert.Equal(t, result[1].Rule.Discount, 0)
}

var cart string = `
grandtotal: 42
rows:
- code: c10
  quantity: 11
  completed: false
  total: 13
- code: c20
  quantity: 22
  completed: false
  total: 23
- code: c30
  quantity: 33
  completed: false
  total: 33
`

func TestImportCart(t *testing.T) {
	r := strings.NewReader(cart)
	result, err := ImportCart(r)

	assert.Nil(t, err)
	assert.Equal(t, result.Rows[0].Code, "c10")
	assert.Equal(t, result.Rows[2].Quantity, 33)
}
