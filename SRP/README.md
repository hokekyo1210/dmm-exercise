# Solid原則 - SRP

Golangによる例で示す。

## 基礎事項

- SRP = Single Responsibility Principle
- 「1つのクラスは1つの責任を持つべきで、1つしか持ってはいけない。」 (書籍「Hands-On Software Architecture with Golang」)
- たとえば。。。utilsというパッケージ名の場合、雑多な機能の溜まり場となってしまう恐れがあるため、このような名前は避けるべきである

## 例

悪いパターン
```
type EmployeeWorkManage interface {
    calculatePay()
    saveEmployee()
}
```
- EmployeeWorkManageインタフェースが「従業員の管理」と「データの保存」複数の責任を持ってしまっている
- 今後の拡張で従業員関係の機能のたまり場になってしまう可能性がある

改善後
```
type EmployeeWorkManage interface {
    calculatePay()
}
type EmployeeDataManage interface {
    saveEmployee()
}
```
- インタフェースを分割することで責任を分散
- コード変更時の影響範囲も限定される

## 他の定義とか

- > あるモジュールやクラスや関数などを、ある1つのアクターに対する責任だけで済むように適切に分割しましょう
    - アクターとは
        - ウェブサービスにアクセスするユーザー
        - 管理者
        - 自動処理など
- > クラスを変更する理由は複数存在してはいけない(アジャイルソフトウェア開発の奥義)
- > The trick of implementing SRP in our software is knowing the responsibility of each class.
    - それぞれのクラスの責務をどう置くかは開発者に委ねられる
    - やりすぎは良くないかもしれない

## 参考文献

- [SOLID原則を理解するためにGo言語で書いてみた](https://qiita.com/MAKOTODA/items/976f47ea036e35c7538d)
- [よくわかるSOLID原則1: S（単一責任の原則）](https://note.com/erukiti/n/n67b323d1f7c5)
- [単一責任の原則（SRP）](https://qiita.com/gomi_ningen/items/02c42e2487d035f9c3c8)
- [Single Responsibility Principle in Java | Baeldung](https://www.baeldung.com/java-single-responsibility-principle)