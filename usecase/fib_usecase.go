package usecase

import (
	"math/big"
)

type FibUsecase interface {
	GetFib(n uint) big.Int
}

type fibUsecase struct{}

func NewFibUsecase() FibUsecase {
	return &fibUsecase{}
}

func (fu *fibUsecase) GetFib(n uint) big.Int {
	return calcFibByDP(n)
}

/*
 * O(n)
 */
func calcFibByDP(n uint) big.Int {
	dp := make([]big.Int, n+1)

	dp[1] = *big.NewInt(1)
	if n == 1 {
		return dp[1]
	}
	dp[2] = *big.NewInt(1)
	if n == 2 {
		return dp[2]
	}

	for i := uint(3); i <= n; i++ {
		dp[i].Add(&dp[i-1], &dp[i-2])
	}

	return dp[n]
}

// 別のアルゴリズムを使う場合はメソッドを追加してGetFib内で切り替える
