# fortaleza

Fortaleza is a program for calculating phone charges

# About

[Fortaleza](https://en.wikipedia.org/wiki/Fortaleza) is one of my favourite cities in Brazil.

# Viewing Results

    go test

# Comments

* *flexibility* is approached by using the `PricingT` type to hold
  pricing rules.

* I ran out of time, so the code in `OnePricing()` is messy and needs a
  lot of refactoring and unit tests. Also due to time constraints
  commenting and code layout could be improved.

* I made the mistake of using YAML to import testing data into the
  structs, which then meant I couldn't implement the cart interface
  correctly (ie `ShoppingCart.new, cart.add, cart.total, cart.items)

* sample *Pricings*, *Carts*  and *Catalogues*   are provided in the
  tests.

* I was unable to implement `Bulk Discount` for `ult_large`, so the
  price is slightly incorrect on line 2

* there's some small rounding errors, due to the discount being
  implemented as an `int` rather than a `float64`. However other
  rounding issues are avoided by using ints with 4 significant digits of
  cents ie `$1` is stored as `10000`

* I used the `gopkg.in/yaml.v2` library as it's not included in the
  GoLang standard library (even though *JSON* **is** included)
