# taxcalculator
Tax Calculator

## User Stories
1. Create tax object
    - Input: Name, Tax Code, Price
    - Tax Code is only 3 type:
        - 1: Food & Beverages
        - 2: Tobacco
        - 3: Entertainment
2. Get tax bill
    - Type, Refundable, Price, Tax, and Amount should be calculated on the fly
    - Tax Codes by Type:
        - 1: 10% of price
        - 2: 10 + (2% of price)
        - 3: if 0 < price < 100, tax is free. If price >= 100, tax = 1% of (price - 100)

## API
By running docker-compose up, we can access the application API at localhost:3000. And here is the API details

1. Create tax object
    - Endpoint: `/object`
    - Method: `POST`
    - Request :
    ```
      {
        "name": "TESTING",
        "tax_code": 1,
        "price": 1000
      }
    ```
    - Response if success:
    `status code`: `201`
    ```
      {
        "data": {
            "id": 1,
            "name": "TESTING",
            "tax_code": 1,
            "price": 1000,
            "type": "Food & Beverages",
            "refundable": true,
            "tax": 100,
            "Amount": 1100
        },
        "error": null
      }
    ```
    - Response if failed:
    `status code`: `400`
    ```
      {
        "data": null,
        "error": [
            "Name should not be empty",
            "Price must be greater than zero",
            "Invalid Tax Code"
        ]
      }
    ```

2. Get tax object
    - Endpoint: `/object`
    - Method: `GET`
    - Request : ```-```
    - Response if success:
    `status code`: `200`
    ```
      {
        "data": [
          {
            "id": 1,
            "name": "susisssssss",
            "tax_code": 1,
            "price": 88,
            "type": "Food & Beverages",
            "refundable": true,
            "tax": 8.8,
            "Amount": 96.8
          }
        ],
        "error": null
      }
    ```
    - Response if data not found:
    `status code`: `200`
    ```
      {
        "data": null,
        "error": null
      }
    ```
## Database Design
The database just store data of object, so there is just one table exists.
Database driver use PostgreSQL.
Database name is `tax`.
The database contain one table that is named `tax_object`.
Here is the `tax_object`'s table contains:

| Column        | Type          | Note                  |
| :------------ |:--------------| :---------------------|
| ID            | serial        | Primary key           |
| name          | string        | object name           |
| tax_code      | int           | to be mapped to type  |
| price         | float64       | price                 |
