@api
Feature: Set quota
  As a user
  I want to set quota to different users
  So that users can only take up a certain amount of storage space

  Background:
    Given these users have been created with default attributes and without skeleton files:
      | username |
      | Alice    |
      | Brian    |


  Scenario Outline: admin sets personal space quota of user with different role
    Given the administrator has given "Alice" the role "Admin" using the settings api
    And the administrator has given "Brian" the role "<userRole>" using the settings api
    When user "Alice" changes the quota of the "Brian Murphy" space to "100" owned by user "Brian"
    Then the HTTP status code should be "200"
    And the JSON data of the response should match
    """
    {
      "type": "object",
      "required": [
        "quota"
      ],
      "properties": {
        "quota": {
          "type": "object",
          "required": [
            "total"
          ],
          "properties": {
            "total" : {
              "type": "number",
              "enum": [100]
            }
          }
        }
      }
    }
    """
    Examples:
      | userRole    |
      | Admin       |
      | Space Admin |
      | User        |
      | User Light  |


  Scenario Outline: non-admin user tries to set the personal space quota of other users
    Given the administrator has given "Alice" the role "<role>" using the settings api
    And the administrator has given "Brian" the role "<userRole>" using the settings api
    When user "Alice" changes the quota of the "Brian Murphy" space to "100" owned by user "Brian"
    Then the HTTP status code should be "403"
    Examples:
      | role        | userRole    |
      | Space Admin | Admin       |
      | Space Admin | Space Admin |
      | Space Admin | User        |
      | Space Admin | User Light  |
      | User        | Admin       |
      | User        | Space Admin |
      | User        | User        |
      | User        | User Light  |
      | User Light  | Admin       |
      | User Light  | Space Admin |
      | User Light  | User        |
      | User Light  | User Light  |


  Scenario Outline: admin or space admin user sets a quota of a project space
    Given the administrator has given "Alice" the role "Space Admin" using the settings api
    And the administrator has given "Brian" the role "<userRole>" using the settings api
    And user "Alice" has created a space "Project Jupiter" of type "project" with quota "20"
    When user "Brian" changes the quota of the "Project Jupiter" space to "100" owned by user "Alice"
    Then the HTTP status code should be "200"
    And the JSON data of the response should match
    """
    {
      "type": "object",
      "required": [
        "name",
        "quota"
      ],
      "properties": {
        "name": {
          "type": "string",
          "enum": ["Project Jupiter"]
        },
        "quota": {
          "type": "object",
          "required": [
            "total"
          ],
          "properties": {
            "total" : {
              "type": "number",
              "enum": [100]
            }
          }
        }
      }
    }
    """
    Examples:
      | userRole    |
      | Admin       |
      | Space Admin |


  Scenario Outline: normal or user light user tries to set quota of a space
    Given the administrator has given "Alice" the role "Space Admin" using the settings api
    And the administrator has given "Brian" the role "<userRole>" using the settings api
    And user "Alice" has created a space "Project Jupiter" of type "project" with quota "20"
    And user "Alice" has shared a space "Project Jupiter" with settings:
      | shareWith | Brian       |
      | role      | <spaceRole> |
    When user "Brian" changes the quota of the "Project Jupiter" space to "100"
    Then the HTTP status code should be "403"
    Examples:
      | userRole   | spaceRole |
      | User       | viewer    |
      | User       | editor    |
      | User       | manager   |
      | User Light | viewer    |
      | User Light | editor    |
      | User Light | manager   |


  Scenario: admin user can set their own personal space quota
    Given the administrator has given "Alice" the role "Admin" using the settings api
    When user "Alice" changes the quota of the "Alice Hansen" space to "100" owned by user "Alice"
    Then the HTTP status code should be "200"
    And the JSON data of the response should match
    """
    {
      "type": "object",
      "required": [
        "quota"
      ],
      "properties": {
        "quota": {
          "type": "object",
          "required": [
            "total"
          ],
          "properties": {
            "total" : {
              "type": "number",
              "enum": [100]
            }
          }
        }
      }
    }
    """


  Scenario Outline: non-admin user tries to set their own personal space quota
    Given the administrator has given "Alice" the role "<userRole>" using the settings api
    When user "Alice" changes the quota of the "Alice Hansen" space to "100" owned by user "Alice"
    Then the HTTP status code should be "403"
    Examples:
      | userRole    |
      | Space Admin |
      | User        |
      | User Light  |
