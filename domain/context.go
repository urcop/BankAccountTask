package domain

type Context interface {
	MakeContext() Context

	Connection() Connection
}
