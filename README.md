# Clean Architecture in Golang

## App features
### Use Cases
- Get information related to an employee and their salary
- Get the salary from an employee in a desired currency

### Storage
There are implementations in:
- MySql
- MongoDb
- In memory

### External APIs
Exchange rates can be retrieved from:
- Local currency converter
- Free Currency Api, token required. [Link](https://app.freecurrencyapi.com/)

## Structure

![img.png](docs/assets/structure.png "Structure Diagram")

## Things to improve / Things missing
- Improve error handling
- Add more unit tests
- Add integration and functional tests
- Add profile/config manager