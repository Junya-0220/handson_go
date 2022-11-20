# このソースコードは株式会社秀和システム「Go 言語 ハンズオン」を参考にしています。

## 変数の\_って何？

for の部分で、\_, v := range arr のように range を使っていました。
ここの\_とはなんのことですか？

Go ではアンダースコアで始まる変数は、他の変数とは少し違います。
これは使わなくても良い変数です。

//Hello World チュートリアルや他の例を完了したあなたは、基本的なユーザインタフェースを作成したことでしょう。このページでは、GUI の内容をコードから更新する方法について説明します。

最初のステップは、更新したいウィジェットを変数に代入することです。hello world チュートリアルでは、SetContent()に直接 widget.NewLabel を渡していましたが、これを更新するには、以下のように 2 つの行に変更します。

```
clock := widget.NewLabel("")
	w.SetContent(clock)
```

内容が変数に代入されると、SetText("new text")のような関数を呼び出すことができます。この例では、Time.Format.Label を使用して、ラベルの内容を現在の時刻に設定します。

```
	formatted := time.Now().Format("Time: 03:04:05")
	clock.SetText(formatted)
```
