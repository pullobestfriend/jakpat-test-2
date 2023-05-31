# Jakpat Backend Test 2
This project is created for second backend test for Jakpat
This project was forked from https://github.com/meyiapir/books-api as a template

## The project satisfies these requirements:
- User can login & register
- User is either a seller or buyer
- User can see item details
- Seller can see all his items
- Seller can add items to his inventories
- Seller can modify his items
- Seller can delete his items (currently soft delete, denoted by item status)
- There's 5 types of order status (`waiting`, `on process`, `shipping`, `delivered`, `expired`)
- Check for order expiry if it's in waiting and passed expired time
- Seller can change order status 1 step at a time
- Buyer can add order
- Only related seller and buyer can see each orders
- Seller can see all orders related to him

## The project satisfies these requirements:
- Auto expire order with status waiting

### To get this project running:

for local run, create .env file with
```
DB_PASSWORD=xxxxxx
```
and run
```bash
make build
make run
```

If the application is launched for the first time, you need to apply migrations to the database:
```bash
migrate -path ./schema -database 'postgres://postgres:xxxxxx@localhost:5433/postgres?sslmode=disable' up
```
this requires https://github.com/golang-migrate/migrate

### Available API for this project:
[![Run in Postman](https://run.pstmn.io/button.svg)](https://god.gw.postman.com/run-collection/27681259-973483a5-6fbd-494a-a92d-fedd9fb91d2c?action=collection%2Ffork&source=rip_markdown&collection-url=entityId%3D27681259-973483a5-6fbd-494a-a92d-fedd9fb91d2c%26entityType%3Dcollection%26workspaceId%3D6fa7ddca-6237-4f5f-93da-e6ba16616e0c)

#### Notes
As this project was made with limited time available, there's a lot of things that could be improved, such as:
- Use DB transaction for function with multiple db changes
- Table name in database is still very simple and susceptible to attacks
- More unit test coverage
- Better unit testing with less gomock.Any() and more test cases
- Better code architechture
- Util and constant package
- More validation for inputs
- Better update functions (currently need to input data for all column)
- Better authorization security