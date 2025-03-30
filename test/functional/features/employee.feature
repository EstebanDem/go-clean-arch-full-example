Feature: Employee features

  Scenario Outline: Get Employee and their salary
    Given an employee with id <id> and salary in <currency>
    When  their information is requested
    Then  the response status code should be <statusCode>
    And   the response body should be <response>

    Examples:
      | id                                   | currency | statusCode | response                        |
      | aa02193c-0592-4191-955f-eefdc04ea35d | ARS      | 200        | ./responses/employee_salary_ars |
