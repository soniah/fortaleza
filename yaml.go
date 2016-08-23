package fortaleza

import (
	"io"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// Problem: YAML data requires dummy grandtotals, etc
// next time do tests like GoSNMP, provide correct interface

func ImportCatalogue(r io.Reader) (CatalogueT, error) {
	var err error
	bs, err := ioutil.ReadAll(r)
	if err != nil {
		return CatalogueT{}, err
	}
	var products CatalogueT
	err = yaml.Unmarshal(bs, &products)
	if err != nil {
		return CatalogueT{}, err
	}
	return products, nil
}

func ImportPricings(r io.Reader) (PricingsT, error) {
	var err error
	bs, err := ioutil.ReadAll(r)
	if err != nil {
		return PricingsT{}, err
	}
	var pricings PricingsT
	err = yaml.Unmarshal(bs, &pricings)
	if err != nil {
		return PricingsT{}, err
	}
	return pricings, nil
}

func ImportCart(r io.Reader) (CartT, error) {
	var err error
	bs, err := ioutil.ReadAll(r)
	if err != nil {
		return CartT{}, err
	}
	var cart CartT
	err = yaml.Unmarshal(bs, &cart)
	if err != nil {
		return CartT{}, err
	}
	return cart, nil
}
