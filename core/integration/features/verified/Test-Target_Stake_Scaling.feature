Feature: test the implementation of market.stake.target.scalingFactor

  Background:

    Given the log normal risk model named "log-normal-risk-model-1":
      | risk aversion | tau  | mu | r | sigma |
      | 0.000001      | 0.01 | 0  | 0 | 1.0   |
    #risk factor short: 0.6323374
    #risk factor long: 0.393276818
    And the fees configuration named "fees-config-1":
      | maker fee | infrastructure fee |
      | 0.004     | 0.001              |
    And the price monitoring named "price-monitoring-1":
      | horizon | probability | auction extension |
      | 3600    | 0.99        | 300               |
    And the markets:
      | id        | quote name | asset | risk model              | margin calculator         | auction duration | fees          | price monitoring   | oracle config          |
      | ETH/MAR22 | ETH        | USD   | log-normal-risk-model-1 | default-margin-calculator | 1                | fees-config-1 | price-monitoring-1 | default-eth-for-future |
    And the parties deposit on asset's general account the following amount:
      | party  | asset | amount    |
      | party0 | USD   | 500000000 |
      | party1 | USD   | 100000000 |
      | party2 | USD   | 100000000 |
      | party3 | USD   | 100000000 |

  Scenario: 002, LP first commit 50,000 which is less than required to end auction, LP then amend commit to 55,000
    Given the following network parameters are set:
      | name                                          | value |
      | market.stake.target.timeWindow                | 24h   |
      | market.stake.target.scalingFactor             | 1.5   |
      | market.liquidity.bondPenaltyParameter         | 0.2   |
      | market.liquidity.targetstake.triggering.ratio | 0.24  |

    And the average block duration is "1"

    And the parties submit the following liquidity provision:
      | id  | party  | market id | commitment amount | fee   | side | pegged reference | proportion | offset | lp type    |
      | lp1 | party0 | ETH/MAR22 | 50000             | 0.001 | sell | ASK              | 500        | 17     | submission |
      | lp1 | party0 | ETH/MAR22 | 50000             | 0.001 | buy  | BID              | 500        | 17     | amendment  |

    And the parties place the following orders:
      | party  | market id | side | volume | price | resulting trades | type       | tif     | reference  |
      | party1 | ETH/MAR22 | buy  | 1      | 900   | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-1  |
      | party1 | ETH/MAR22 | buy  | 1      | 990   | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-1  |
      | party1 | ETH/MAR22 | buy  | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-2  |
      | party2 | ETH/MAR22 | sell | 10     | 1000  | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-3 |
      | party2 | ETH/MAR22 | sell | 1      | 1010  | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-1 |
      | party2 | ETH/MAR22 | sell | 1      | 1100  | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-2 |

    And the parties submit the following liquidity provision:
      | id  | party  | market id | commitment amount | fee   | side | pegged reference | proportion | offset | lp type   |
      | lp1 | party0 | ETH/MAR22 | 55000             | 0.001 | sell | ASK              | 500        | 17     | amendment |
      | lp1 | party0 | ETH/MAR22 | 55000             | 0.001 | buy  | BID              | 500        | 17     | amendment |

    Then the opening auction period ends for market "ETH/MAR22"
    And the trading mode should be "TRADING_MODE_CONTINUOUS" for the market "ETH/MAR22"

    And the market data for the market "ETH/MAR22" should be:
      | mark price | trading mode            | horizon | min bound | max bound | target stake | supplied stake | open interest |
      | 1000       | TRADING_MODE_CONTINUOUS | 3600    | 973       | 1027      | 9484         | 55000          | 10            |

    Then the parties should have the following account balances:
      | party   | asset | market id | margin  | general   |
      | party0  | USD   | ETH/MAR22 | 1656013 | 498288987 |

    Then debug detailed orderbook volumes for market "ETH/MAR22"

    #LP margin(maintanance level)= 1000*125*0.6323374+1000*131*0.393276818=130562
    Then the parties should have the following margin levels:
      | party  | market id | maintenance | search | initial | release |
      | party0 | ETH/MAR22 | 79043       | 86947  | 94851   | 110660  |
    And the parties should have the following account balances:
      | party  | asset | market id | margin  | general   | bond  |
      | party0 | USD   | ETH/MAR22 | 1656013 | 498288987 | 55000 |

    And the parties place the following orders:
      | party  | market id | side | volume | price | resulting trades | type       | tif     |
      | party1 | ETH/MAR22 | buy  | 1     | 1000  | 0                | TYPE_LIMIT | TIF_GTC |
      | party2 | ETH/MAR22 | sell | 1     | 1000  | 1                | TYPE_LIMIT | TIF_GTC |

    # Excessive margin gets released now 
    Then the parties should have the following margin levels:
      | party  | market id | maintenance | search | initial | release |
      | party0 | ETH/MAR22 | 79043       | 86947  | 94851   | 110660  |
    And the parties should have the following account balances:
      | party  | asset | market id | margin | general   | bond  |
      | party0 | USD   | ETH/MAR22 | 94851  | 499850149 | 55000 |

