# go-transaction-api
Simple banking simulation service with transfer and withdraw transaction capabilities

## Description
### [account-manager-service]
```
- /sign-in        : allow user to register their data
- /sign-up        : allow user to get their auth credentials (in this case, jwt token)
- /verify         : allow another service, in this case [payment-manager-service], to use this service authorization
- /accounts       : provide users their accounts list with pagination, can be filtered via account type
- /transactions   : provide users their transactions list with pagination, can be filtered via recipient and/or sender account number, and status
```
### [payment-manager-service]
```
- /transfer       : allow users to transfer from their account to another user account, with custom concurrency. transaction will automaticall succeeded after 30 seconds if user didn't withdraw during those 30 seconds period after transfer creation time, and recipient or sender account balance will be updated
- /withdraw       : allow users to withdraw their transaction within 30 senconds period after transfer labeled as "pending"
```

## Techstack
- Go, with Gin framework
- PostgreSQL
- Redis

## Database structure
![erd drawio (1)](https://github.com/n9mi/go-transaction-api/assets/113373725/1eeef6f7-35da-491e-bec3-e69bd5d011a5)

## Notes
- I implemented my own authentication with JWT and Redis

## TODO
- OpenApi.yaml documentation, but I already include the postman documentation in this repository
