Feature: Counting characters

  Scenario Outline: Counting characters
    Given the word "<word>"
    When I count the letters
    Then I should get "<expected>"

    Examples:
      | word    | expected | description               |
      | AAAA    |       4A | identical characters      |
      | AABBBCC |   2A3B2C | multiple groups           |
      | ABC     | ABC      | no repeated characters    |
      | A       | A        | single character          |
      | AAABCC  |    3AB2C | mixed repeated and single |
