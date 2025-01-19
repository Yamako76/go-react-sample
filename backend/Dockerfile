# Goの公式イメージをベースに使用
FROM golang:1.23 AS build

# 作業ディレクトリの設定
WORKDIR /app

# Goモジュールの依存関係をコピー
COPY go.mod go.sum ./
RUN go mod download

# ソースコードをコピー
COPY . .

# アプリケーションのビルド
RUN go build -o main .

# 実行用イメージ
FROM golang:1.23

WORKDIR /app

COPY --from=build /app/main .

EXPOSE 8080

CMD ["./main"]
