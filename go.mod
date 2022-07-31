module github.com/MaestroJolly/go-be-apis-scaffold

go 1.18

replace github.com/MaestroJolly/go-be-apis-scaffold/src/services/greetings => ./src/services/greetings

replace github.com/MaestroJolly/go-be-apis-scaffold/src/routes => ./src/routes

replace github.com/MaestroJolly/go-be-apis-scaffold/src/utils => ./src/utils

replace github.com/MaestroJolly/go-be-apis-scaffold/src/services/articles => ./src/services/articles

require (
	github.com/MaestroJolly/go-be-apis-scaffold/src/middlewares v0.0.0-00010101000000-000000000000
	github.com/MaestroJolly/go-be-apis-scaffold/src/routes v0.0.0-00010101000000-000000000000
	github.com/MaestroJolly/go-be-apis-scaffold/src/services/greetings v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.8.0
)

require (
	github.com/MaestroJolly/go-be-apis-scaffold/src/services/articles v0.0.0-00010101000000-000000000000 // indirect
	github.com/MaestroJolly/go-be-apis-scaffold/src/utils v0.0.0-00010101000000-000000000000 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
)

replace github.com/MaestroJolly/go-be-apis-scaffold/src/middlewares => ./src/middlewares
