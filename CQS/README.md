# CQS原則

Golangによる例を示す。

## 基礎事項

- CQS = Command-Query Separation:コマンドクエリ分離
- メソッドを2つの種類だけに限定しよう
    - クエリ：結果を返し、システムの状態を変更しない（副作用がない）
    - コマンド：システムの状態を変更し、値を返さない
- > CQS: 「あらゆるメソッドは、アクションを実行するコマンドか、呼び出し元にデータを返すクエリかのいずれかであって、両方を行ってはならない。これは、質問をすることで回答を変化させてはならないということだ。」
    - 確かにクエリを投げるたびに結果が変わったら怖くて使いたくなくなる

## 例

```

type Printer struct{
    text string
}

func (printer *Printer) print() (n int, err error) { //クエリ
    return fmt.Print(printer.text)
}

func (printer *Printer) setText(text string){ //コマンド
    printer.text = text
}

```
とてもシンプルなもの、クエリを何回叩いても副作用が無いため結果は変わらない。

## 他の定義とか

- > 副作用をコマンドに局所化し、クエリの世界を広げることにより、ソフトウェアの信頼性・再利用性・拡張性を向上させることができる。


## 参考文献

- [コマンド・クエリという考え方](https://qiita.com/pakkun/items/7dc5a9b6bc57225a3673)
- [PHPではじめるCQRS](https://speakerdeck.com/dnskimo/phpdehazimerucqrs?slide=16)