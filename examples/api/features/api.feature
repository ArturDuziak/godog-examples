Feature: CRUD API Example

  Scenario: Example Update by object id
    Given I have an API "https://api.restful-api.dev" with path "/objects"
    When I send a PUT request with params "/ff8081819782e69e019953d3b01471d1" and following body:
      """
      {
      "name": "Apple MacBook Pro 16",
      "data": {
      "year": 2019,
      "price": 2049.99,
      "CPU model": "Intel Core i9",
      "Hard disk size": "1 TB",
      "color": "silver"
      }
      }
      """
    Then status code should be 200
    And value "$.name" should equal "Apple MacBook Pro 16"
