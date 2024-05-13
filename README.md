## 2024/05/13(雑記)

どこでも実行できる自作コマンドの作り方を知ったので、goの勉強も兼ねて作成したリポジトリ。

とりあえず、今日思いついて雑に設計してみた。
並行処理とかテストとか積極的に書いていきたい。

## イメージ

command:`keizoku asakatsu`
```
<< A2にrankアップしました！おめでとう！ >>
【asakatsu】 rank : A2
あと5日でランクダウン/リセット
ランクキープ/ランクアップまであと+3
1日スキップ：3/10
```

## 要件

- コマンドが実行されたタイミングで値を計算する必要がある

## 実行できるコマンド
- `keizoku git_commit +2`
- `keizoku new rannning 10 20`
- `keizoku delete ranning`
- `keizoku list`
- `keizoku help`

## データ構造
- オブジェクト：Routine,History
```
type Routine struct {
	Name           string // 半角英数字
	Description    string
	Rank           string
	RankUpAt       string
	RankDownAt     string
	RankUpCount    string
	RankKeepCount  string
	skipticketsNum int
	UpdateAt       string
}

func ratedContinuation() {
	ranks := map[string]int{
		"SS": 3,
		"S":  5,
		"A":  3,
		"B":  3,
		"C":  5,
		"D":  3,
		"E":  3,
		"F":  1,
	}
}
```