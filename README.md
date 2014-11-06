# ゆるふわ Development Club 管理ツール

Under development...

## How to use

### Install

```
$ go get github.com/yurufuwa/yurufuwa/cmd/yurufuwa
$ yurufuwa -h
```

You need to install golang and mercurial. (e.g. `$ brew install go` and `$ brew install mercurial`)

### Update

```
$ go get -u github.com/yurufuwa/yurufuwa/cmd/yurufuwa
```

### 自分のリポジトリにコラボレータとして Yurufuwa メンバーを追加

```
$ yurufuwa collaborators add your/repos
```

### 自分のリポジトリにコラボレータとして追加した Yurufuwa メンバーを削除

```
$ yurufuwa collaborators remove your/repos
```

### 開催予定の meetup の情報を表示

```
$ yurufuwa meetups
```

## How to develop

### Check out and build

```
# recommend to use ghq!
# or: git clone https://github.com/yurufuwa/yurufuwa.git
ghq get yurufuwa/yurufuwa
cd $( ghq list -p -e yurufuwa )

# download dependencies
go get

# build binary
go build -o $GOPATH/bin/yurufuwa ./cmd/yurufuwa
```

### Running tests

* FIXME: Please add tests!!!
