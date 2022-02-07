Feature: Set up a market, with an opening auction, then uncross the book. Make sure opening auction can end if we have remaingin volume in the uncrossing range


  Background:

    Given the following assets are registered:
      | id  | decimal places |
      | ETH | 5              |
    And the simple risk model named "my-simple-risk-model":
      | long                   | short                  | max move up | min move down | probability of trading |
      | 0.08628781058136630000 | 0.09370922348428490000 | -1          | -1            | 0.2                    |
    And the fees configuration named "my-fees-config":
      | maker fee | infrastructure fee |
      | 0.004     | 0.001              |
    And the markets:
      | id        | quote name | asset | risk model           | margin calculator         | auction duration | fees           | price monitoring | oracle config          | decimal places |
      | ETH/DEC20 | ETH        | ETH   | my-simple-risk-model | default-margin-calculator | 1                | my-fees-config | default-none     | default-eth-for-future | 2              |

  Scenario: set up 2 parties with balance
    # setup accounts
    Given the parties deposit on asset's general account the following amount:
      | party  | asset | amount        |
      | party1 | ETH   | 1000000000000 |
      | party2 | ETH   | 1000000000000 |
      | party3 | ETH   | 1000000000000 |

    # place orders and generate trades - slippage 100
    When the parties place the following orders:
      | party  | market id | side | volume | price    | resulting trades | type       | tif     | reference |
      | party2 | ETH/DEC20 | buy  | 1      | 9500000  | 0                | TYPE_LIMIT | TIF_GTC | t2-b-1    |
      | party1 | ETH/DEC20 | buy  | 1      | 10000000 | 0                | TYPE_LIMIT | TIF_GFA | t1-b-1    |
      | party2 | ETH/DEC20 | sell | 2      | 10000000 | 0                | TYPE_LIMIT | TIF_GTC | t2-s-1    |

    Then the opening auction period ends for market "ETH/DEC20"

    And the following trades should be executed:
      | buyer   | price    | size | seller  |
      | party1  | 10000000 | 1    | party2  |
    And the mark price should be "10000000" for the market "ETH/DEC20"
