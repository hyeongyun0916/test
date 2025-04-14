package main

import (
	"errors"
	"fmt"
)

func log(msg string) {
	fmt.Println(msg)
}

func main() {
	err := reconcile()
	if err != nil {
		log("Reconciler error " + err.Error()) // https://github.com/kubernetes-sigs/controller-runtime/blob/v0.18.4/pkg/internal/controller/controller.go#L324
	}
}

func reconcile() error {
	err := func1()
	if err != nil {
		log("Error in func1: " + err.Error())
	}
	return err
}

func func1() error {
	err := func2()
	if err != nil {
		log("Error in func2: " + err.Error())
	}
	return err
}

func func2() error {
	return errors.New("an error occurred in func2")
}
