package main

import "testing"

func TestComplexity(t *testing.T) {
	testcases := []struct {
		name       string // テストケース名
		code       string // テスト対象のGoのコード
		complexity int    // 期待する循環複雑度
	}{
		// ここにケースを追加する
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			// a := GetAST(t, testing.T)

			// c := Count(a)

			// if c != testcase.complexity {
			// 	t.Errorf("got=%d, want=%d", c, testcase.complexity)
			// }
		})
	}
}
