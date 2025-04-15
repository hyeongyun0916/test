package main

import (
	"errors"
	"fmt"
)

type accl struct {
	// spec   string
	status struct {
		health string
	}
}

func updateStatus(a *accl, status string) error {
	a.status.health = status
	// update via k8s client
	return errors.New("k8s api error")
}

func getAccl() *accl {
	// This function returns a new instance of accl.
	return &accl{}
}

func main() {
	// This is the main function where the program starts executing.
	// You can add your code here to perform specific tasks or call other functions.
	fmt.Println("Hello, World!")

	var err error
	accl := getAccl()

	defer func() {
		// 1. func1과 func2의 에러가 다름에도 같은 처리
		// 2. 위에서 아래로 흐르는 로직이 길어진다면 위에서 defer했다는 사실을 까먹지는 않을까?
		//		open(); defer close(); 를 쓰는 이유는 까먹지 않기 위해서
		if err != nil {
			// 3. 여기서의 에러가 남지 않음.
			_ = updateStatus(accl, err.Error())
		}
	}()

	err = func1()
	if err != nil {
		return
	}

	err = func2()
	if err != nil {
		return
	}
}

func func1() error {
	fmt.Println("json.Marshal error")
	return nil
}

func func2() error {
	fmt.Println("update ft but k8s api error")
	return nil
}
