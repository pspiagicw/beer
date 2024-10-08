package types

//go:generate stringer -type=Type

type Type int

const (
	Integer Type = iota
	Float
	String
	Boolean
	Void

	Illegal
)
