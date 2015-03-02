# [Overseer](https://alveary-overseer.herokuapp.com/)

![Travis-CI](https://travis-ci.org/alveary/overseer.svg) by [Travis-CI](https://travis-ci.org/alveary/overseer)

Simplified service discovery, with included health check.

(Because of running that service on heroku in the moment in an architecture, where we don't have any
access to the single dyno entities in the health check, we will assume sigle dyno instances for now.)

## run test suite

Overseer is using Godep as dependency manager,
so just run `godep go test ./...` to run all included tests.

Install `reflex` when you want to run the tests all the time and reload on file changes

```sh
go get github.com/cespare/reflex
```

With reflex enabled just run the `test-all` file.

```sh
./test-all
```
