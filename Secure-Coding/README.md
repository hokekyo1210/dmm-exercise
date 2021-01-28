# セキュアコーディング

[セキュリティ部 Secure-Coding-Guideline](https://git.dmm.com/pages/security/Secure-Coding-Guideline/)を参考にした。

## SQLインジェクション

### 基礎事項

- SQL文を取り扱う際, ユーザーから入力された値をSQL文に直接埋め込んでしまうようなコードになっていると, 悪意あるクエリが実行される可能性がある(以下, PHPで例を示す)
    - 例: `$sql = "SELECT name, nickname, birthday, last_login_date FROM user WHERE id = '${id}'";`
    - 攻撃者が`id`パラメータに`1'+OR+'a'='a`といった入力を行うと, SQL文の実行結果として全ユーザの情報が返ってしまう
        - 実際に実行されるクエリ: `SELECT name, nickname, birthday, last_login_date FROM user WHERE id = '1' OR 'a'='a'`
- 対策として, 以前はユーザーから入力された値をエスケープするような手法(例えばmysql_real_escape_string関数など)が用いられていたが, 現在ではプレースホルダを用いる方法が一般的になっている
    - プレースホルダ: `$sql = "SELECT name, nickname, birthday, last_login_date FROM user WHERE id = ?";`
    - `?`の部分がユーザーから入力された値で置換される
    - PHPにおいては[PDOStatementクラス](https://www.php.net/manual/ja/class.pdostatement.php)を用いる


## OSコマンド・インジェクション

### 基礎事項

- SQLインジェクションのシェル版, 攻撃者が任意のOSコマンドを実行出来てしまう可能性がある
    - PHP: system関数, exec関数などを使う際
    - golang: [exec.Command](https://golang.org/pkg/os/exec/#Command)など
- 影響としては...
    - サーバ内のファイルの情報漏洩、改ざん、削除
    - 他サイトへの攻撃の踏み台としての利用などなど
- 対策
    - 入力値をエスケープする
        - PHPの例: escapeshellarg()
    - 入力値に正規表現などでバリデーションをかける
    - OSコマンドを使用する実装をやめる


### 例(PHP)

```
$cmd = "nslookup " . $target;
$output=null;
exec($cmd, $output);
print_r($output);
```

このとき, 攻撃者から$targetとして「www.dmm.com;cat /etc/hosts」が入力されると,実際に実行されるコマンドは「nslookup www.dmm.com;cat /etc/hosts」となり, webサーバーの/etc/hostsが攻撃者に返ってしまう可能性がある。

## パストラバーサル

### 基礎事項

- ユーザーから入力された文字列からパスを組み立ててファイルを出力するようなアプリケーションにおいて, 意図しないファイルにアクセス出来てしまう脆弱性
- 対策
    - ファイル名にバリデーションをかけてディレクトリを指定出来ないようにする
        - もしくはPHPのbasename()関数などを利用してファイル名のみを抽出する
    - 外部からファイル名を指定できる仕様を避ける
        - ファイル名を固定にする
        - ユーザID等を利用してファイルを指定する
    

### 例(PHP)

```
$json = file_get_contents('/var/www/html/files/' . $target);
print($json)
```

このとき, 攻撃者から$targetとして「../../../../../../etc/passwd」が入力されると, /etc/passwdの中身が攻撃者に返ってしまう。

## XSS: クロスサイトスクリプティング

### 基礎事項

- 例えばPHPにおいて, ユーザーから入力された文字列からHTMLを生成する場合に不正なスクリプトが埋め込まれ, 結果的に他ユーザーがそのスクリプトを実行してしまう脆弱性
    - 正規ページに似せた偽ページを他ユーザーに見せることでフィッシングを行う
    - クッキー情報を盗むスクリプトの実行などなど
- 対策
    - 正規表現などを用いた入力値のバリデーション
    - HTMLに出力される場所でエスケープを行う
        - PHP: htmlspecialchars()

### 例(PHP)

```
<p><?php echo $user['1111']['name']; ?></p>
```

このとき, userID:1111が攻撃者であり, 自分の名前を`<script>任意の攻撃用スクリプト</script>`などとすると, 最終的なHTMLは

```
<p><script>任意の攻撃用スクリプト</script></p>
```

となり, 他のユーザーのブラウザで任意の攻撃用スクリプトが実行されてしまう。
この例はいわゆる蓄積型XSSと呼ばれる脆弱性だが, 他にも他ユーザーを攻撃用のスクリプトが仕込まれた罠サイトに誘導して攻撃を行う反射型XSSなどがある。

## CSRF: Cross-Site Request Frogeries

### 基礎事項

- 攻撃者が用意した罠サイトに他ユーザを誘導し, 意図しないスクリプトをCSRF未対策のサイトに対して実行できる脆弱性
    - 決済
    - メールアドレスやパスワードの変更などなど
- 対策
    - 攻撃者が知り得ないトークンを送らないと重要な処理を実行出来ないようにする
    - 重要な処理の前にパスワードの再入力を求める
        - ただしあらゆる処理に再入力を求めるとUXが下がる


## HTTPヘッダインジェクション

### 基礎事項

- 不正なクエリパラメータやリクエストパラメータが載ったURLを他ユーザに踏ませることで, レスポンスヘッダやレスポンスボディに任意の文字列を追加することができる脆弱性
    - レスポンスボディに任意の文字列を追加できるため, 攻撃者は任意のコンテンツが埋め込める
    - 悪意あるWebページへリダイレクト
    - 任意のCookieが生成出来るためセッションIDが固定されるなどなど
- 以下の画像がとてもわかりやすい
    ![改行コードに要注意！ HTTP ヘッダインジェクションの概要と対策 より](https://yamory.io/blog/images/posts/about-http-header-injection/flow-of-setting-the-cookie-intended-by-the-attacker.png)
- 対策
    - 外部からのパラメータをHTTPレスポンスヘッダに出力しないようにする

## オープンリダイレクト

### 基礎事項

- Webアプリケーションのリダイレクトを利用して, 悪意あるURLを他ユーザーに踏ませることで任意のドメインにリダイレクトさせることが出来てしまう脆弱性
    - 例えば「https://hoge.com/login/path=https://google.com」 というようにログイン成功後にpath以下のURLにリダイレクトさせるWebアプリケーションがあったとき, リダイレクト先のURLを改変したURLを他ユーザーに踏ませることが出来る
    - 罠サイトに誘導
    - マルウェアをダウンロードさせるなど
- 対策
    - リダイレクト先を固定のURLにする
    - リダイレクト先をURL毎に一意に設定したID等で指定