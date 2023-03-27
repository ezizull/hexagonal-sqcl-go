package todo

import (
	useCaseTodo "skyshi-gethired.go/application/usecases/todo"
)

type Controller struct {
	BookService useCaseTodo.Service
}
