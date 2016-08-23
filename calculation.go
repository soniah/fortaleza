package fortaleza

// AllPricing updates the card with the application of all conditions and rules
func AllPricing(cart CartT, catalog CatalogueT, pricings PricingsT) CartT {
	for _, pricing := range pricings {
		cart = OnePricing(cart, catalog, pricing)
	}
	return cart
}

// OnePricing updates the cart with the application of one condition and rule
func OnePricing(cart CartT, catalogue CatalogueT, pricing PricingT) CartT {

	// an important part of the logic is that each cart row is only
	// processed once, by testing the Completed boolean

	var resultRows []CartRowT
	grandtotal := 0

	for _, cartrow := range cart.Rows {

		// only process each row once
		if cartrow.Completed {
			resultRows = append(resultRows, cartrow)
			continue
		}

		price := cataloguePrice(catalogue, cartrow.Code)
		discount := pricing.Rule.Discount

		// n-for-m deal? bulk discount?
		if cartrow.Code == pricing.Condition.Code && cartrow.Quantity == pricing.Condition.Quantity {

			newprice := pricing.Rule.NewPrice

			// n-for-m deal?
			if discount != 0 {
				newrow := CartRowT{
					Code:      cartrow.Code,
					Quantity:  cartrow.Quantity,
					Completed: true,
					Total:     cartrow.Quantity * applyDiscount(price, discount),
				}
				resultRows = append(resultRows, newrow)
			} else {

				// bulk discount
				newrow := CartRowT{
					Code:      cartrow.Code,
					Quantity:  cartrow.Quantity,
					Completed: true,
					Total:     newprice * cartrow.Quantity,
				}
				resultRows = append(resultRows, newrow)

			}

		} else if cartrow.Code == pricing.Condition.Code && pricing.Rule.Extra != "" {
			// bundle?

			cartrow.Total = price * cartrow.Quantity
			cartrow.Completed = true
			resultRows = append(resultRows, cartrow)

			newrow := CartRowT{
				Code:      pricing.Rule.Extra,
				Quantity:  1,
				Completed: true,
				Total:     0,
			}
			resultRows = append(resultRows, newrow)

		} else if pricing.Condition.Promo != "" {
			// promo code?
			//
			// this makes the weak assumption that there's only one promo code

			// add up grand total
			for _, row := range resultRows {
				grandtotal += row.Total
			}

			// discount grand total
			discount = pricing.Rule.TotalDiscount
			grandtotal = applyDiscount(grandtotal, discount)

			// break due to promo being sorted to end of conditions
			break

		} else if cartrow.Code != pricing.Condition.Code {
			// no match - just copy across
			resultRows = append(resultRows, cartrow)

		} else {
			// else just update price
			cartrow.Total = price * cartrow.Quantity
			cartrow.Completed = true
			resultRows = append(resultRows, cartrow)
		}
	}

	if grandtotal == 0 {
		for _, row := range resultRows {
			grandtotal += row.Total
		}
	}
	resultcart := CartT{grandtotal, resultRows}
	return resultcart
}

// from the catalogue, find the price for the given code
func cataloguePrice(catalogue CatalogueT, code string) int {
	for _, product := range catalogue {
		if product.Code == code {
			return product.Price
		}
	}
	return 0
}

func applyDiscount(price int, discount int) int {
	discountF := float64(100-discount) / 100
	priceF := float64(price)
	return int(discountF * priceF)
}
