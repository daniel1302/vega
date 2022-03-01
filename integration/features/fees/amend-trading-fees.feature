Feature: Fees calculations

Scenario: Testing fees get collected when amended order trades
    
    Given the fees configuration named "fees-config-1":
      | maker fee | infrastructure fee |
      | 0.005     | 0.002              |
    And the price monitoring updated every "1000" seconds named "price-monitoring":
      | horizon | probability | auction extension |
      | 1       | 0.99        | 3                 |
    
    When the simple risk model named "simple-risk-model-1":
      | long | short | max move up | min move down | probability of trading |
      | 0.2  | 0.1   | 100          | -100         | 0.1                    |

    And the markets:
      | id        | quote name | asset | risk model          | margin calculator         | auction duration | fees          | price monitoring | oracle config          | maturity date        |
      | ETH/DEC21 | ETH        | ETH   | simple-risk-model-1 | default-margin-calculator | 2                | fees-config-1 | price-monitoring | default-eth-for-future | 2019-12-31T23:59:59Z |

    # setup accounts
    Given the parties deposit on asset's general account the following amount:
      | party   | asset | amount    |
      | aux1     | ETH   | 100000000 |
      | aux2     | ETH   | 100000000 |
      | trader3a | ETH   | 10000     |
      | trader3b | ETH   | 10000     |
      | trader4  | ETH   | 1250      |

    Then the parties place the following orders:
      | party  | market id | side | volume | price | resulting trades | type       | tif     |
      | aux1    | ETH/DEC21 | buy  | 1      | 1000  | 0                | TYPE_LIMIT | TIF_GTC |
      | aux2    | ETH/DEC21 | sell | 1      | 1000  | 0                | TYPE_LIMIT | TIF_GTC |
      | aux1    | ETH/DEC21 | buy  | 1      | 920   | 0                | TYPE_LIMIT | TIF_GTC |
      | aux2    | ETH/DEC21 | sell | 1      | 1080  | 0                | TYPE_LIMIT | TIF_GTC |

    Then the opening auction period ends for market "ETH/DEC21"
    And the market data for the market "ETH/DEC21" should be:
      | mark price | trading mode            | 
      | 1000       | TRADING_MODE_CONTINUOUS | 
   
    When the parties place the following orders:
      | party   | market id  | side | volume | price | resulting trades | type       | tif     |
      | trader3a | ETH/DEC21 | buy  | 2      | 1002  | 0                | TYPE_LIMIT | TIF_GTC |
      | trader3b | ETH/DEC21 | buy  | 1      | 1002  | 0                | TYPE_LIMIT | TIF_GTC |
      | trader4  | ETH/DEC21 | sell | 4      | 1002  | 2                | TYPE_LIMIT | TIF_GTC |

    Then the market data for the market "ETH/DEC21" should be:
      | mark price | trading mode            |  
      | 1002       | TRADING_MODE_CONTINUOUS |

    Then the following trades should be executed:
    

      # | buyer   | price | size | seller  | maker   | taker   |
      # | trader3 | 1002  | 3    | trader4 | trader3 | trader4 |
      # TODO to be implemented by Core Team
      | buyer    | price | size | seller  |
      | trader3a | 1002  | 2    | trader4 |
      | trader3b | 1002  | 1    | trader4 |
        
    # For trader3a-
    # trade_value_for_fee_purposes for trader3a = size_of_trade * price_of_trade = 2 * 1002 = 2004
    # infrastructure_fee = fee_factor[infrastructure] * trade_value_for_fee_purposes = 0.002 * 2004 = 4.008 = 5 (rounded up to nearest whole value)
    # maker_fee =  fee_factor[maker]  * trade_value_for_fee_purposes = 0.005 * 2004 = 10.02 = 11 (rounded up to nearest whole value)
    # liquidity_fee = fee_factor[liquidity] * trade_value_for_fee_purposes = 0.001 * 2004 = 2.004 = 3 (rounded up to nearest whole value)

    # For trader3b -
    # trade_value_for_fee_purposes = size_of_trade * price_of_trade = 1 * 1002 = 1002
    # infrastructure_fee = fee_factor[infrastructure] * trade_value_for_fee_purposes = 0.002 * 1002 = 2.004 = 3 (rounded up to nearest whole value)
    # maker_fee =  fee_factor[maker]  * trade_value_for_fee_purposes = 0.005 * 1002 = 5.01 = 6 (rounded up to nearest whole value)
    # liquidity_fee = fee_factor[liquidity] * trade_value_for_fee_purposes = 0.001 * 1002 = 1.002 = 2 (rounded up to nearest whole value)

    And the following transfers should happen:
      | from    | to       | from account            | to account                       | market id | amount | asset |
      | trader4 | market   | ACCOUNT_TYPE_GENERAL    | ACCOUNT_TYPE_FEES_MAKER          | ETH/DEC21 | 11     | ETH   |
      | trader4 | market   | ACCOUNT_TYPE_GENERAL    | ACCOUNT_TYPE_FEES_MAKER          | ETH/DEC21 |  6     | ETH   |
      | trader4 |          | ACCOUNT_TYPE_GENERAL    | ACCOUNT_TYPE_FEES_INFRASTRUCTURE |           |  8     | ETH   |
      | market  | trader3a | ACCOUNT_TYPE_FEES_MAKER | ACCOUNT_TYPE_GENERAL             | ETH/DEC21 | 11     | ETH   |  
      | market  | trader3b | ACCOUNT_TYPE_FEES_MAKER | ACCOUNT_TYPE_GENERAL             | ETH/DEC21 |  6     | ETH   |
     
    # total_fee = maker_fee + infrastructure_fee + liquidity_fee =  11 + 6 + 8 = 25
    # Trader3a margin + general account balance = 10000 + 11 ( Maker fees) = 10011
    # Trader3b margin + general account balance = 10000 + 6 ( Maker fees) = 10006
    # Trader4  margin + general account balance = 10000 - (11+6) ( Maker fees) - 8 (Infra fee) = 99975

    # TODO: Check why margin doesn't go up after the trade WHEN the liquidity provision order gets included (seems to work fine without LP orders) (expecting first commented out values) but getting second value in other cases
    Then the parties should have the following account balances:
      | party      | asset | market id | margin | general |
      | trader3a    | ETH   | ETH/DEC21 | 678    | 9333    | 
      | trader3b    | ETH   | ETH/DEC21 | 339    | 9667    |
      | trader4     | ETH   | ETH/DEC21 | 621    | 604     |
   
   # Placing second set of orders
    When the parties place the following orders:
      | party   | market id | side | volume | price | resulting trades | type       | tif     | reference      |
      | trader3a | ETH/DEC21 | buy  | 2      | 1000  | 0                | TYPE_LIMIT | TIF_GTC | trader3a-buy-1 |
      | trader4  | ETH/DEC21 | sell | 4      | 1002  | 0                | TYPE_LIMIT | TIF_GTC | trader4-sell-2 |

    Then the parties should have the following account balances:
      | party   | asset | market id | margin | general |
      | trader3a | ETH   | ETH/DEC21 | 1159   | 8852    |
      | trader4  | ETH   | ETH/DEC21 | 1102   | 123     |

      # reducing size
      And the parties amend the following orders:
      | party  | reference      | price | size delta | tif     |
      | trader4 | trader4-sell-2 | 1000  | 0          | TIF_GTC |

    # matching the order now
    Then the following trades should be executed:
      # | buyer   | price | size | seller  | maker   | taker   |
      # | trader3 | 1002  | 3    | trader4 | trader3 | trader4 |
      # TODO to be implemented by Core Team
      | buyer    | price | size | seller  |
      | trader3a | 1000  | 2    | trader4 |
      
      # checking if continuous mode still exists
    Then the market data for the market "ETH/DEC21" should be:
      | mark price | trading mode            |  
      | 1000       | TRADING_MODE_CONTINUOUS |

    #  Then the parties should have the following account balances:
    #   | party      | asset | market id | margin | general |
    #   | trader3a    | ETH   | ETH/DEC21 | 1159   | 8852    |
    #   | trader4     | ETH   | ETH/DEC21 | 1102   |  123    |

   And debug transfers

    And the following transfers should happen:
      | from    | to       | from account            | to account                       | market id | amount | asset |
      | trader4 | market   | ACCOUNT_TYPE_GENERAL    | ACCOUNT_TYPE_FEES_MAKER          | ETH/DEC21 | 10     | ETH   |
      | trader4 |          | ACCOUNT_TYPE_GENERAL    | ACCOUNT_TYPE_FEES_INFRASTRUCTURE |           |  4     | ETH   |
      | market  | trader3a | ACCOUNT_TYPE_FEES_MAKER | ACCOUNT_TYPE_GENERAL             | ETH/DEC21 |   10   | ETH   |  

     Then the parties should have the following account balances:
      | party    | asset | market id | margin | general |
      | trader3a | ETH   | ETH/DEC21 | 1344   | 8673    |
      | trader4  | ETH   | ETH/DEC21 | 1108   |  109    |