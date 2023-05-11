FROM golang:1.17

WORKDIR /app

# go.mod ファイルを生成するために go mod init を実行
RUN go mod init fomgit
# 依存関係のパッケージをダウンロード
RUN go mod download

COPY . .

ENTRYPOINT ["tail"]
CMD ["-f", "/dev/null"]