# Git

## 特徴

2005年に登場。
ローカルにリポジトリのクローンを作成できる。そのため, リモートにアクセスできない環境でも, 履歴の調査や変更の記録といったほとんどの作業を行うことができる。

## 基本操作

- git init
    - Gitのリポジトリ作成
- git add
    - ファイル更新
        - `git add .` //すべてのファイル・ディレクトリ
        - `git add -n` //追加されるファイルを調べる
- git commit
    - 成果をコミットする
        - `git commit -m "first commit"`
- git reset
    - コミットの取り消し
        - `git reset --soft HEAD~2` // 最新のコミットから2件分をワークディレクトリの内容を保持し取り消す
        - `git reset --hard HEAD~2` // 最新のコミットから2件分のワークディレクトリの内容とコミットを取り消す
- git fetch
    - リモートリポジトリ(例えばgithubのリポジトリ)の最新の履歴の取得を行う
- git pull
    - リモートリポジトリのコミット内容をローカルリポジトリにマージする
        - `git pull origin master`
- git push 
    - ローカルリポジトリのコミット内容をリモートリポジトリにプッシュする


## GitとCVS比較

- リポジトリ
    - Git: ローカルとリモート両方に持てる
    - CVS: 中央集中管理型なのでリモートにしか持てない
    - Gitはリモートを汚さずにテストコード用のリポジトリをどんどん作れる
- ブランチ
    - Git: サクサク切る
    - CVS: めったに切らない

Gitを使用すると開発の規模に合わせて設定を行えるため、開発者同士が柔軟に作業を進めることができる。しかし、仕組みや使用方法が多少複雑。

## rebase

git rebaseを利用してmergeコンフリクトを解消した。

```
tsuchida-yuki1@PC-0001357 rebase % git init
Initialized empty Git repository in /Users/tsuchida-yuki1/go/src/git.dmm.com/tsuchida-yuki1/dmm-exercise/git/rebase/.git/
tsuchida-yuki1@PC-0001357 rebase % echo "v1" > ver.txt
tsuchida-yuki1@PC-0001357 rebase % git add .
tsuchida-yuki1@PC-0001357 rebase % git commit -m "v1"
[master (root-commit) 0ed003d] v1
 1 file changed, 1 insertion(+)
 create mode 100644 ver.txt
tsuchida-yuki1@PC-0001357 rebase % git checkout -b v2
Switched to a new branch 'v2'
tsuchida-yuki1@PC-0001357 rebase % echo "v2" > ver.txt
tsuchida-yuki1@PC-0001357 rebase % git add .          
tsuchida-yuki1@PC-0001357 rebase % git commit -m "v2" 
[v2 bc9c46a] v2
 1 file changed, 1 insertion(+), 1 deletion(-)
tsuchida-yuki1@PC-0001357 rebase % git checkout main
error: pathspec 'main' did not match any file(s) known to git
tsuchida-yuki1@PC-0001357 rebase % git branch
  master
* v2
tsuchida-yuki1@PC-0001357 rebase % git checkout master
Switched to branch 'master'
tsuchida-yuki1@PC-0001357 rebase % echo "v3" > ver.txt
tsuchida-yuki1@PC-0001357 rebase % git add .
tsuchida-yuki1@PC-0001357 rebase % git commit -m "v3"
[master 6259a69] v3
 1 file changed, 1 insertion(+), 1 deletion(-)
tsuchida-yuki1@PC-0001357 rebase % git checkout v2 
Switched to branch 'v2'
tsuchida-yuki1@PC-0001357 rebase % git rebase master
First, rewinding head to replay your work on top of it...
Applying: v2
Using index info to reconstruct a base tree...
M       ver.txt
Falling back to patching base and 3-way merge...
Auto-merging ver.txt
CONFLICT (content): Merge conflict in ver.txt
error: Failed to merge in the changes.
Patch failed at 0001 v2
hint: Use 'git am --show-current-patch' to see the failed patch
Resolve all conflicts manually, mark them as resolved with
"git add/rm <conflicted_files>", then run "git rebase --continue".
You can instead skip this commit: run "git rebase --skip".
To abort and get back to the state before "git rebase", run "git rebase --abort".
tsuchida-yuki1@PC-0001357 rebase % git rebase --skip                     
tsuchida-yuki1@PC-0001357 rebase % git checkout master
Switched to branch 'master'
tsuchida-yuki1@PC-0001357 rebase % git log
commit 6259a69a6130861b694c27243e69cb3253efff07 (HEAD -> master, v2)
Author: tsuchida-yuki1 <tsuchida-yuki1@dmm.com>
Date:   Wed Jan 27 19:49:36 2021 +0900

    v3

commit 0ed003db042d2d9f9393b636806ad6014a610de4
Author: tsuchida-yuki1 <tsuchida-yuki1@dmm.com>
Date:   Wed Jan 27 19:48:49 2021 +0900

    v1
```

