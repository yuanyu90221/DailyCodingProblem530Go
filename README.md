# DailyCodingProblem530Go

## 問題描述

定義兩個字串間的修改距離為從其中一個字串組成另一個字串所需要新增刪改的最少字元數。舉例來說， "kitten" 與 "sitting" 的修改距離為 3：替換 "k" 為 "s" ，替換 "e" 為 "i" ，以及新增 "g" 。

給定兩個字串，算出兩個字串的修改距離。

## 解題思路

所謂兩個字串的修改距離就是看兩個字串有有多少字元不同。

首先算出兩個字串長度。

然後以較短長度的字串去遍歷找出每個對應位置字元與較長字串的字元不同的個數 n

最後把 n + 兩個字串的長度差就是所謂兩個字串的修改距離。

## 程式碼

```golang
package diff_string

func GetStringDistance(str1, str2 string) int {
	str1Len := len(str1)
	str2Len := len(str2)
	diff := 0
	min := str1Len
	max := str2Len
	target := str1
	compareStr := str2
	if str2Len < str1Len {
		min = str2Len
		max = str1Len
		target = str2
		compareStr = str1
	}

	for idx := range target {
		if target[idx] != compareStr[idx] {
			diff++
		}
	}
	return diff + (max - min)
}

```