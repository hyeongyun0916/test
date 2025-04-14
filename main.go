package main

import (
	"fmt"

	"github.com/pkg/errors"
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
	return errors.Wrap(err, "reconcile failed")
}

func func1() error {
	err := func2()
	return errors.Wrap(err, "func1 failed")
}

func func2() error {
	return errors.New("an error occurred in func2")
}
