# web_go

* directory構成は`$GOPATH/web_go/`になっている
  * `go install web_go/first_webapp` で、`${GOPATH}/src/web_go/`配下からfirst_webappを探し、該当するソースをコンパイルしてくれる
  * コンパイルされたバイナリは、`$GOPATH/bin/` に格納される