Feature: CRUD API Get example

  Scenario: Get example object
    Given I have an API "https://api.restful-api.dev" with path "/objects"
    When I send a GET request with params "?id=1"
    Then status code should be 200
    And response body should be:
      """
      [{
        "id": "1",
        "name": "Google Pixel 6 Pro",
        "data": {
          "color": "Cloudy White",
          "capacity": "128 GB"
        }
      }]
      """

  Scenario: Get non existing object
    Given I have an API "https://api.restful-api.dev" with path "/objects"
    When I send a GET request with params "?id=1234567890"
    Then status code should be 200
    And response body should be:
      """
      []
      """
