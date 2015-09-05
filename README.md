# heroku

```
heroku create -b https://github.com/kr/heroku-buildpack-go.git
echo "web: $(basename `pwd`)" > Procfile
```

# godeps は階層が重要？

```
mkdir -p src/github.com/osdakira/qiita_webhook
cp main.go src/github.com/osdakira/qiita_webhook
cd src/github.com/osdakira/qiita_webhook
godeps save
```

# インストールした手順

```
brew install go
brew install direnv
go get github.com/zenazn/goji
# go run main.go

go get github.com/bitly/go-simplejson

go get github.com/pilu/fresh
fresh # shotgun みたいな webserver 起動

go get github.com/tnantoka/chatsworth
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
