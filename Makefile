SHELL := /bin/sh

# Dockerコンテナをビルドしてバックグラウンドで起動
.PHONY: setup
setup:
	docker compose up -d --build

# Dockerコンテナをビルド
.PHONY: build
build:
	docker compose build

# Dockerコンテナをバックグラウンドで起動
.PHONY: up
up:
	docker compose up -d

# Dockerコンテナを停止
.PHONY: stop
stop:
	docker compose stop

# Dockerコンテナを停止して削除
.PHONY: down
down:
	docker compose down

# Dockerコンテナを停止してボリュームも削除
.PHONY: down-v
down-v:
	docker compose down --volumes

# フロントエンドコンテナにシェルで接続
.PHONY: frontend
frontend:
	docker compose exec frontend sh

# データベースコンテナにMySQLクライアントで接続
.PHONY: sql
sql:
	docker compose exec db mysql -u root -p

# Dockerコンテナのステータスを表示
.PHONY: ps
ps:
	docker compose ps

# Goコードをフォーマット
.PHONY: fmt
fmt:
	cd backend && go fmt ./...

# Goコードを静的解析
.PHONY: lint
lint:
	cd backend && go vet ./...
	cd backend && staticcheck ./...

# Goコードをテスト
.PHONY: test
test:
	cd backend && go test ./...
