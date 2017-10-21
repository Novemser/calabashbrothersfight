package main
import "execution"


type Level struct {
	id string
	name string
	shortDescription string
	longDescription string
	victoryText string
	threads execution.ThreadContext
	variables string

}