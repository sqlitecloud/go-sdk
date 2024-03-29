package test

import "flag"

const testConnectionStringLocalhost = "sqlitecloud://admin:admin@localhost:8860?tls=skip"

const testUsername = "admin"
const testPassword = "admin"

var testConnectionString string = testConnectionStringLocalhost

func init() {
	flag.StringVar(&testConnectionString, "server", testConnectionStringLocalhost, "Connection String")
}

func contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}
