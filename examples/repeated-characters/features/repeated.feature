Feature: Counting characters

  Scenario: Counting repeated same characters
    Given the word "AAAA"
    When I count the letters
    Then I should get "4A"

  Scenario: Counting repeated different characters
    Given the word "AABBBCC"
    When I count the letters
    Then I should get "2A3B2C"

  Scenario: Counting unrepeated characters
    Given the word "ABC"
    When I count the letters
    Then I should get "ABC"

  Scenario: Counting single character
    Given the word "A"
    When I count the letters
    Then I should get "A"

  Scenario: Counting repeated and unrepeated mixed characters
    Given the word "AAABCC"
    When I count the letters
    Then I should get "3AB2C"
