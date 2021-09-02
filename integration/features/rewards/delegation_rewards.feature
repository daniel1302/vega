Feature: Staking & Delegation 

  Background:
    Given the following network parameters are set:
      | name                                              |  value |
      | governance.vote.asset                             |  VEGA  |
      | validators.epoch.length                           |  10s   |
      | validators.delegation.minAmount                   |  10    |
      | reward.staking.delegation.payoutDelay             |  0s    |
      | reward.staking.delegation.delegatorShare          |  0.883 |
      | reward.staking.delegation.minimumValidatorStake   |  100   |
      | reward.staking.delegation.payoutFraction          |  0.5   |
      | reward.staking.delegation.maxPayoutPerParticipant | 100000 |
      | reward.staking.delegation.competitionLevel        |  1.1   |

    Given time is updated to "2021-08-26T00:00:00Z"
 
    And the validators:
      | id     | staking account balance |
      | node1  |         1000000         |
      | node2  |         1000000         |
      | node3  |         1000000         |
      | node4  |         1000000         |
      | node5  |         1000000         |
      | node6  |         1000000         |
      | node7  |         1000000         |
      | node8  |         1000000         |
      | node9  |         1000000         |
      | node10 |         1000000         |
      | node11 |         1000000         |
      | node12 |         1000000         |
      | node13 |         1000000         |

    #set up the self delegation of the validators
    Then the parties submit the following delegations:
      | party  | node id  | amount |
      | node1  |  node1   | 10000  | 
      | node2  |  node2   | 10000  |       
      | node3  |  node3   | 10000  | 
      | node4  |  node4   | 10000  | 
      | node5  |  node5   | 10000  | 
      | node6  |  node6   | 10000  | 
      | node7  |  node7   | 10000  | 
      | node8  |  node8   | 10000  | 
      | node9  |  node9   | 10000  | 
      | node10 |  node10  | 10000  | 
      | node11 |  node11  | 10000  | 
      | node12 |  node12  | 10000  | 
      | node13 |  node13  | 10000  | 

    And the global reward account gets the following deposits:
      | asset | amount |
      | VEGA  | 100000 | 
    
    And the parties deposit on asset's general account the following amount:
      | party  | asset  | amount |
      | party1 | VEGA   |  10000 |
      | party2 | VEGA   |  20000 |
      | party3 | VEGA   |  30000 |

    And the parties deposit on staking account the following amount:
      | party  | asset  | amount |
      | party1 | VEGA   | 10000  |  

    Then the parties submit the following delegations:
    | party  | node id  | amount |
    | party1 |  node1   |  100   | 
    | party1 |  node2   |  200   |       
    | party1 |  node3   |  300   |     

    #complete the first epoch for the self delegation to take effect
    Then time is updated to "2021-08-26T00:00:10Z"
    Then time is updated to "2021-08-26T00:00:11Z"

  Scenario: Parties withdraw from their staking account during an epoch once having active delegations - they should not get rewarded for those uncovered delegations 
    Desciption: Parties have active delegations on epoch 1 and withdraw stake from the staking account. They should only get rewarded for any delegation that still has cover 

    #party1 has a balance of 10k tokens in their staking account and an active delegation in this epoch of 600. By withdrawing 9850, 450 of their delegation needs to be revoked and they should only get rewarded for the 150 tokens
    #NB: the undelegation is done proportionally to the stake they have in each node, so for example party1 has 100, 200, 300 in nodes 1-3 respectively so 
    #after undelegation they will have 25, 50, 75 in nodes 1-3 respectively
    Given the parties withdraw from staking account the following amount:  
    | party  | asset  | amount |
    | party1 | VEGA   |  9850  |

    #advance to the end of the epoch
    When time is updated to "2021-08-26T00:00:21Z"

    #verify validator score 
    Then the validators should have the following val scores for epoch 1:
    | node id | validator score  | normalised score |
    |  node1  |      0.07703     |     0.07703      |    
    |  node2  |      0.07722     |     0.07722      |
    |  node3  |      0.07741     |     0.07741      | 
    |  node4  |      0.07683     |     0.07683      | 

    #node1 has 10k self delegation + 25 from party1
    #node2 has 10k self delegation + 50 from party1
    #node3 has 10k self delegation + 75 from party1
    #all other nodes have 10k self delegation 

    #party1 gets 0.07703 * 50000 * 0.883 * 25/10025 + 0.07722 * 50000 * 0.883 * 50/10050 + 0.07741 * 50000 * 0.883 * 75/10075
    #node1 gets: (1 - 0.883 * 25/10025) * 0.07703 * 50000
    #node2 gets: (1 - 0.883 * 50/10050) * 0.07722 * 50000
    #node3 gets: (1 - 0.883 * 75/10075) * 0.07741 * 50000
    #node4 - node13 get: 0.07683 * 50000

    And the parties receive the following reward for epoch 1:
    | party  | asset | amount |
    | party1 | VEGA  |  49    | 
    | node1  | VEGA  |  3842  | 
    | node2  | VEGA  |  3843  | 
    | node3  | VEGA  |  3845  | 
    | node4  | VEGA  |  3841  | 
    | node5  | VEGA  |  3841  | 
    | node6  | VEGA  |  3841  | 
    | node7  | VEGA  |  3841  | 
    | node8  | VEGA  |  3841  | 
    | node9  | VEGA  |  3841  | 
    | node10 | VEGA  |  3841  | 
    | node11 | VEGA  |  3841  | 
    | node12 | VEGA  |  3841  | 
    | node13 | VEGA  |  3841  | 

   