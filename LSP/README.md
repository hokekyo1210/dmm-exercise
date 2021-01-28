# Solid原則 - LSP

たいていの高級言語でこの原則に違反するコードを書くことが難しい。
とりあえずすぐに再現できそうなPHPで例を示す。

## 基礎事項

- LSP = Liskov Substitution Principle：リスコフの置換原則
- > Subtypes must be substitutable for their base types.
    - (Javaなどで言うなら)サブクラスは、そのスーパークラスで代用可能でなければならない。
- > S が T の派生型であれば、プログラム内で T 型のオブジェクトが使われている箇所は全て S 型のオブジェクトで置換可能
が言える。
    - このとき, Tはクラスでもインターフェースでも仕様書でも良いが, SはTの仕様に従った実装になっていなければならず, 各S毎に呼び出し方がTと違っていてはならない。

## 例

```
<?php
interface Printer
{
    public function print(); //なんらか標準出力する
}

class PrinterA implements Printer
{
    public function print()
    {
        print "PrinterA";
    }
}

class PrinterB implements Printer
{
    public function print()
    {
        return "PrinterB" //仕様違反
    }
}

```

PHPの場合がわかりやすく, 上記のようにPrinterを実装したクラスのprint()の実装の動きが統一されていないと, 外から呼ぶときに値が返ってきているか確認してから使う羽目になってしまう。
こうならないよう仕様を統一しましょうね, という原則だ。

## 参考文献

- [Liskov Substitution Principle in Java | Baeldung](https://www.baeldung.com/java-liskov-substitution-principle)
- [リスコフの置換原則(LSP)](http://harumi.sakura.ne.jp/wordpress/2019/05/27/%E3%83%AA%E3%82%B9%E3%82%B3%E3%83%95%E3%81%AE%E7%BD%AE%E6%8F%9B%E5%8E%9F%E5%89%87lsp/)
- [SOLIDの原則(+α) まとめ](https://confl.arms.dmm.com/pages/viewpage.action?pageId=619080761)