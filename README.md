## 2024/05/13(雑記)

どこでも実行できる自作コマンドの作り方を知ったので、goの勉強も兼ねて作成したリポジトリ。

とりあえず、今日思いついて雑に設計してみた。
並行処理とかテストとか積極的に書いていきたい。

## 概要

習慣化をゲーム感覚で楽しめるCLIのアプリケーション

## 完成イメージ

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

## 実行できるコマンド(例)
- `keizoku git_commit +2`
- `keizoku new running {ランクキープの値} {ランクアップの値} ` 
- `keizoku delete running`
- `keizoku list`
- `keizoku asaktsu history`
- `keizoku help`

## コンセプト

- 手軽さ
	- ローカルだけでいつでもすぐに実行できる(別アプリとしてweb連携版(多機能で高い自由度)も作りたい)
	- 操作が分かりやすい
- ゲーミフィケーションの考え方を取り入れる
	- 初めはシンプルかつ誘導的に(オンボーディング)　-> 熟達しているという満足感と、達成感を与える
		- ex:ランクの仕組み,デフォルト値,チュートリアル,丁寧な説明
- 短期(1日~1週間)、中長期(1ヶ月~半年)両方の変化を実感できる

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