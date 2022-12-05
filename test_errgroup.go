package main

import (
	"errors"
	"log"
	"time"

	"golang.org/x/sync/errgroup"
)

func test_errgroup() {
	var eg errgroup.Group
	for i := 0; i < 10; i++ {
		n := i
		eg.Go(func() error {
			return do_test2(n)
		})
	}
	// errがある場合、最初のerrを返す。
	if err := eg.Wait(); err != nil {
		//
	}
}

func do_test2(n int) error {
	if n%2 == 0 {
		return errors.New("err")
	}

	time.Sleep(1 * time.Second)
	log.Printf("%d called", n)

	return nil
}
