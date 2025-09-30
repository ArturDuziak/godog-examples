Feature: Account deposit

  Scenario: New account
    Given I have a new account
    Then the account balance must be 0

  Scenario: Deposit money into  account
    Given I have an account with 0
    When I deposit 10.0
    Then the account balance must be 10

  Scenario: Withdraw money from account
    Given I have an account with 10
    When I withdraw 5.0
    Then the account balance must be 5

  Scenario: Withdraw money from account with insufficient balance
    Given I have an account with 15
    When I withdraw 30 
    Then the transaction should fail
