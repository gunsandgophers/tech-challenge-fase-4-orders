Feature: Checkout order
  I need to be able to open an order

  Scenario: Open order with products
    Given there are a SANDWICH
    Given there are a DRINKS
    Given there are a SIDEDISHES
    Given there are a DESSETS
    When I add all items
    Then there should be one order with 4 items
