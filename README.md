# Form3 Take Home Exercise

## Instructions
This exercise has been designed to be completed in 4-8 hours. The goal of this exercise is to write a client library
in Go to access our fake [account API](http://api-docs.form3.tech/api.html#organisation-accounts) service.

### Should
- Client library should be written in Go
- Document your technical decisions
- Implement the `Create`, `Fetch`, `List` and `Delete` operations on the `accounts` resource. Note that filtering of the List operation is not required, but you should support paging
- Ensure your solution is well tested to the level you would expect in a commercial environment. Make sure your tests are easy to read.
- If you encounter any problems running the fake accountapi we would encourage you to do some debugging first,
before reaching out for help

#### Docker-compose
 - Add your solution to the provided docker-compose file
 - We should be able to run `docker-compose up` and see your tests run against the provided account API service

### Please don't
- Use a code generator to write the client library
- Use a library for your client (e.g: go-resty). Only test libraries are allowed.
- Implement an authentication scheme

## How to submit your exercise
- Include your name in the README. If you are new to Go, please also mention this in the README so that we can consider this when reviewing your exercise
- Create a private [GitHub](https://help.github.com/en/articles/create-a-repo) repository, copy the `docker-compose` from this repository
- [Invite](https://help.github.com/en/articles/inviting-collaborators-to-a-personal-repository) @form3tech-interviewer-1 to your private repo
- Let us know you've completed the exercise using the link provided at the bottom of the email from our recruitment team

## License
Copyright 2019-2020 Form3 Financial Cloud

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.

## TODO
- session creation?
- rename hadlers to middleware

- validate response e.g. json
- validate response body is not empty

- implement
- jsnon schema validator middleware
- retry middleware
- proper error validation
- escape values from the client. make sure cannot attack on different ways
- make it similar to AWS-SDK-GO
- validate response on deletion ?? NO_CONTENT=204
- error handling (e.g. record not exists, missing values, no query parameter)
- when not found, just say about it

### Endpoints
- /v1/health
- GET    /v1/organisation/accounts/:id
- DELETE /v1/organisation/accounts/:id
- GET    /v1/organisation/accounts
- POST   /v1/organisation/accounts

###

country
bank_id
bic

```
Country code: GB
Bank ID: required, 6 characters, UK sort code
BIC: required
Bank ID Code: required, has to be GBDSC
Account Number: optional, 8 characters, generated if not provided
IBAN: Generated if not provided
```

## Test

-
not able to create account with same id
- add errors as from here https://github.com/aws/aws-sdk-go/tree/master/aws

 // defer r.HTTPResponse.Body.Close()



// TODO: is this really required here?
  // Method interface {
  //   Get(path string) (*Response, error)
  //   Post(path, body io.Reader) (*Response, error)
  //   Delete(id string) (*Response, error)
  // }
