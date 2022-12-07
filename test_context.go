package main

import (
	"context"
	"fmt"
	"sync"
)

/*
並行処理のうち１つでも失敗すれば、残りをスキップする処理
*/

func doSomeThingPage(workerNum int) error {

	// 必要なコンテキストを生成する
	ctx := context.Background()
	cancelCtx, cancel := context.WithCancel(ctx)

	// 正常終了時にコンテキストのリリースを解放
	defer cancel()

	// 複数のゴルーチンからエラーメッセージを集約するためにチャネルを用意する
	errCh := make(chan error, workerNum)

	// workerNum分の並行処理を行う
	wg := sync.WaitGroup{}
	for i := 0; i < workerNum; i++ {
		i := i
		wg.Add(1)
		go func(num int) {
			defer wg.Done()

			// エラーが発生すれば、キャンセル処理を行い、エラーメッセージを送信する
			if err := doSomeThingWithContext(cancelCtx, num); err != nil {
				cancel()
				errCh <- err
			}
			return
		}(i)
	}

	// 並行処理の終了を待つ
	wg.Wait()

	// エラーチャネルに入ったメッセージを取り出す
	close(errCh)
	var errs []error
	for err := range errCh {
		errs = append(errs, err)
	}

	// エラーが発生していれば、最初のエラーを返す
	if len(errs) > 0 {
		return errs[0]
	}

	// 正常終了
	return nil
}

// コンテキストを利用した何らかの処理をする関数
func doSomeThingWithContext(ctx context.Context, num int) error {
	// 処理に入る前に、コンテキストの死活を確認する
	select {
	case <-ctx.Done():
		return ctx.Err()
	// コンテキストがまだキャンセルされていなければ、そのまま処理に進む
	default:
	}
	fmt.Println(num)
	return nil
}
