Feature: Allow markets to be specified with a smaller number of decimal places than the underlying settlement asset

    Background:
        Given the following network parameters are set:
            | name                                          | value |
            | market.stake.target.timeWindow                | 24h   |
            | market.stake.target.scalingFactor             | 1     |
            | market.liquidity.bondPenaltyParameter         | 0.2   |
            | market.liquidity.targetstake.triggering.ratio | 0.1   |
        And the following assets are registered:
            | id  | decimal places |
            | ETH | 5              |
            | USD | 2              |
        And the average block duration is "1"
        And the log normal risk model named "log-normal-risk-model-1":
            | risk aversion | tau | mu | r | sigma |
            | 0.000001      | 0.1 | 0  | 0 | 1.0   |
        #risk factor short: 3.5569036
        #risk factor long: 0.801225765

        And the fees configuration named "fees-config-1":
            | maker fee | infrastructure fee |
            | 0.004     | 0.001              |
        And the price monitoring named "price-monitoring-1":
            | horizon | probability | auction extension |
            | 360000  | 0.99        | 300               |
        And the markets:
            | id        | quote name | asset | risk model              | margin calculator         | auction duration | fees          | price monitoring   | oracle config          | decimal places | position decimal places |
            | USD/DEC22 | USD        | ETH   | log-normal-risk-model-1 | default-margin-calculator | 1                | fees-config-1 | price-monitoring-1 | default-usd-for-future | 5              | -1                      |
        And the parties deposit on asset's general account the following amount:
            | party  | asset | amount    |
            | party0 | ETH   | 5000000   |
            | party1 | ETH   | 100000000 |
            | party2 | ETH   | 100000000 |
            | party3 | ETH   | 100000000 |
            | lpprov | ETH   | 100000000 |

    Scenario: 001, test negative PDP when trading mode is auction (0019-MCAL-010)

        Given  the parties submit the following liquidity provision:
            | id  | party  | market id | commitment amount | fee   | side | pegged reference | proportion | offset | lp type    |
            | lp7 | party0 | USD/DEC22 | 1000              | 0.001 | sell | ASK              | 100        | 20     | submission |
            | lp7 | party0 | USD/DEC22 | 1000              | 0.001 | buy  | BID              | 100        | -20    | amendment  |
            | lp6 | lpprov | USD/DEC22 | 4000              | 0.001 | sell | ASK              | 100        | 20     | submission |
            | lp6 | lpprov | USD/DEC22 | 4000              | 0.001 | buy  | BID              | 100        | -20    | amendment  |

        And the parties place the following orders:
            | party  | market id | side | volume | price | resulting trades | type       | tif     | reference   |
            | party1 | USD/DEC22 | buy  | 1      | 1000  | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-2a  |
            | party2 | USD/DEC22 | sell | 1      | 1000  | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-3a |
            | party0 | USD/DEC22 | buy  | 1      | 900   | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-1a  |
            | party0 | USD/DEC22 | sell | 1      | 1100  | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-4a |

        Then the market data for the market "USD/DEC22" should be:
            | target stake | supplied stake |
            | 35569        | 5000           |
        # target stake= vol * mark price * rf = 1*10*1000*3.5569036*10 = 35569
        And the opening auction period ends for market "USD/DEC22"
        And the trading mode should be "TRADING_MODE_OPENING_AUCTION" for the market "USD/DEC22"
        And the mark price should be "0" for the market "USD/DEC22"

        Then the parties should have the following account balances:
            | party  | asset | market id | margin | general  | bond |
            | party0 | ETH   | USD/DEC22 | 89635  | 4909365  | 1000 |
            | party1 | ETH   | USD/DEC22 | 19218  | 99980782 | 0    |
            | party2 | ETH   | USD/DEC22 | 85368  | 99914632 | 0    |

        And the parties should have the following margin levels:
            | party  | market id | maintenance | search | initial | release |
            | party0 | USD/DEC22 | 74696       | 82165  | 89635   | 122831  |
            | party1 | USD/DEC22 | 16016       | 17616  | 19218   | 41640   |
            | party2 | USD/DEC22 | 71140       | 78254  | 85368   | 184964  |


    Scenario: 002, test negative PDP when trading mode is continous (0003-MTMK-013, 0019-MCAL-010, 0029-FEES-014)
        Given the parties submit the following liquidity provision:
            | id  | party  | market id | commitment amount | fee   | side | pegged reference | proportion | offset | lp type    |
            | lp2 | party0 | USD/DEC22 | 35569             | 0.001 | sell | ASK              | 500        | 20     | submission |
            | lp2 | party0 | USD/DEC22 | 35569             | 0.001 | buy  | BID              | 500        | 20     | amendment  |

        And the parties place the following orders:
            | party  | market id | side | volume | price | resulting trades | type       | tif     | reference  |
            | party1 | USD/DEC22 | buy  | 5      | 8     | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-1  |
            | party1 | USD/DEC22 | buy  | 1      | 9     | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-1  |
            | party1 | USD/DEC22 | buy  | 10     | 10    | 0                | TYPE_LIMIT | TIF_GTC | buy-ref-2  |
            | party2 | USD/DEC22 | sell | 10     | 10    | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-3 |
            | party2 | USD/DEC22 | sell | 1      | 10    | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-1 |
            | party2 | USD/DEC22 | sell | 5      | 11    | 0                | TYPE_LIMIT | TIF_GTC | sell-ref-2 |

        When the opening auction period ends for market "USD/DEC22"
        Then the trading mode should be "TRADING_MODE_CONTINUOUS" for the market "USD/DEC22"
        And the auction ends with a traded volume of "10" at a price of "10"

        And the market data for the market "USD/DEC22" should be:
            | mark price | trading mode            | horizon | min bound | max bound | target stake | supplied stake | open interest |
            | 10         | TRADING_MODE_CONTINUOUS | 360000  | 8         | 13        | 3556         | 35569          | 10            |
        # target stake = 10*10*10*3.5569036=3556

        And the parties should have the following account balances:
            | party  | asset | market id | margin | general  | bond  |
            | party0 | ETH   | USD/DEC22 | 632133 | 4332298  | 35569 |
            | party1 | ETH   | USD/DEC22 | 1778   | 99998222 | 0     |
            | party2 | ETH   | USD/DEC22 | 6830   | 99993170 | 0     |

        And the parties should have the following margin levels:
            | party  | market id | maintenance | search | initial | release |
            | party0 | USD/DEC22 | 526778      | 579455 | 632133  | 737489  |
            | party1 | USD/DEC22 | 1482        | 1630   | 1778    | 2074    |
            | party2 | USD/DEC22 | 5692        | 6261   | 6830    | 7968    |

        #risk factor short: 3.5569036
        #risk factor long: 0.801225765
        # Margin_maintenance_party0 = 1481*10*3.5569036*10+1206*10*0.801225765*10 = 623405

        #margin for party0 = 712*10*10*3.5569036+791*10*10*0.801225765=316629
        And the following trades should be executed:
            | buyer  | price | size | seller |
            | party1 | 10    | 10   | party2 |

        Then the parties should have the following profit and loss:
            | party  | volume | unrealised pnl | realised pnl |
            | party1 | 10     | 0              | 0            |
            | party2 | -10    | 0              | 0            |

        Then the order book should have the following volumes for market "USD/DEC22":
            | side | price | volume |
            | sell | 13    | 1481   |
            | sell | 11    | 5      |
            | sell | 10    | 1      |
            | buy  | 9     | 1      |
            | buy  | 8     | 1211   |

        And the parties place the following orders:
            | party  | market id | side | volume | price | resulting trades | type       | tif     | reference |
            | party3 | USD/DEC22 | sell | 1      | 9     | 1                | TYPE_LIMIT | TIF_GTC | buy-ref-3 |

        And the following trades should be executed:
            | seller | price | size | buyer  |
            | party3 | 9     | 1    | party1 |

        Then the parties should have the following profit and loss:
            | party  | volume | unrealised pnl | realised pnl |
            | party1 | 11     | -100           | 0            |
            | party2 | -10    | 100            | 0            |
            | party3 | -1     | 0              | 0            |
            | party0 | 0      | 0              | 0            |

        #MTM with price change from 10 to 9, party1 has long position of volume 10, price 10 ->9, MTM -1*10*10*1=-100; party2 has short position of volume 10, price 10 ->9, MTM 10*10*1=100;
        And the parties should have the following account balances:
            | party  | asset | market id | margin | general  | bond  |
            | party0 | ETH   | USD/DEC22 | 568920 | 4395511  | 35569 |
            | party1 | ETH   | USD/DEC22 | 1678   | 99998223 | 0     |
            | party2 | ETH   | USD/DEC22 | 6930   | 99993170 | 0     |
        # Margin_maintenance_party0 = 1481*10*3.5569036*9+1206*10*0.801225765*9 = 561064
        And the parties should have the following margin levels:
            | party  | market id | maintenance | search | initial | release |
            | party0 | USD/DEC22 | 474100      | 521510 | 568920  | 663740  |
            | party1 | USD/DEC22 | 1264        | 1390   | 1516    | 1769    |
            | party2 | USD/DEC22 | 5322        | 5854   | 6386    | 7450    |




