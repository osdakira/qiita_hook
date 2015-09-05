# direnv で GOPATH を制御しないバージョン

ライブラリは、共通ディレクトリに入る

```
export GOPATH="$HOME/.gopath"
export PATH=$PATH:$GOPATH/bin
go get github.com/tools/godep
go get github.com/zenazn/goji

```


# ↓↓↓ Version 1

# インストールした手順

```
brew install go
brew install direnv
go get github.com/zenazn/goji
go run src/github.com/osdakira/qiita_hook/main.go
```

# godeps うまく動かない

```
go get github.com/tools/godep
```

# goenv 使いたいけど面倒すぎ

```
# https://bitbucket.org/ymotongpoo/goenv
mkdir -p $HOME/.goenv
echo 'export GOENVTARGET=$HOME/.goenv/' >> ~/.bashrc
exec $SHELL -l
curl -L https://bitbucket.org/ymotongpoo/goenv/raw/master/shellscripts/fast-install.sh | bash
rm goenv
exec $SHELL -l
```
