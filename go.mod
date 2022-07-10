module pkgcentral.dev/go-be-apis-scaffold

go 1.18

replace pkgcentral.dev/go-be-apis-scaffold/src/services/greetings => ./src/services/greetings

require (
	pkgcentral.dev/go-be-apis-scaffold/src/routes v0.0.0-00010101000000-000000000000
	pkgcentral.dev/go-be-apis-scaffold/src/services/greetings v0.0.0-00010101000000-000000000000
)

replace pkgcentral.dev/go-be-apis-scaffold/src/routes => ./src/routes
