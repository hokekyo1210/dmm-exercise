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

こちらがわかりやすい, https://github.com/ngiasim/SOLID/blob/master/3-LSP.php

```
// Bad Example

class Bird
{
        public function  Fly()
        {
            return “I can Fly”;
        }
}

class Parrot extends Bird
{
        public function  Fly()
        {
            return “I can Fly”;
        }
}

class Ostrich extends Bird
{
        public function  Fly()
        {
            return “I can Fly”;
            // But ostrich can't fly
        }
}



// Good Example

class Bird{
	// Methods related to birds
}

class FlyingBirds extends Bird
{
        public function  Fly()
        {
            Return “I can Fly”;
        }
}

class Parrot extends FlyingBirds{
        public function  Fly()
        {
            Return “I can Fly”;
        }
}

class Ostrich extends Bird{
      // Methods related to birds
}


```

## 参考文献

- [Liskov Substitution Principle in Java | Baeldung](https://www.baeldung.com/java-liskov-substitution-principle)
- [リスコフの置換原則(LSP)](http://harumi.sakura.ne.jp/wordpress/2019/05/27/%E3%83%AA%E3%82%B9%E3%82%B3%E3%83%95%E3%81%AE%E7%BD%AE%E6%8F%9B%E5%8E%9F%E5%89%87lsp/)
- [SOLIDの原則(+α) まとめ](https://confl.arms.dmm.com/pages/viewpage.action?pageId=619080761)