package main

type Action interface {
	Run() error
}
