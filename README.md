# Gin Example — Go Web Service ✅

簡短說明：這是一個以 **Gin** + **GORM** 建置的範例 Web 服務，內含 REST API、WebSocket 範例（含 ping 行為示範），預設使用 **SQLite**，並支援切換到 **PostgreSQL**。

---

## 主要功能 🔧
- HTTP server（監聽：`8081`）
- REST API（範例：`GET /api/v1/user/list` 回傳使用者清單）
- WebSocket 範例（頁面：`/template/websocket`，接點：`/websocket/v1/start`，可傳入 ping/ip 參數並回傳連線輸出）
- 使用 GORM 做資料庫管理（`AutoMigrate` + `seed`）
- 預設 SQLite（`./databases/sqlites/oprueba.db`），也可透過環境變數切換到 PostgreSQL

---

## 快速開始（本地開發） 🚀

Prerequisites
- Go >= 1.24
- (選用) `iperf3` / `ping` 可用於測試 ping / iperf 範例

安裝（參考 `install.sh`）
```bash
# 安裝 gvm 並安裝指定 Go 版本
bash install.sh
```

執行
```bash
go run main.go
# 或 build
go build -o gin-example
./gin-example
```
服務會監聽在 `:8081`

---

## 環境變數與 DB 設定 🗄️
- 預設使用 SQLite（路徑：`./databases/sqlites/oprueba.db`）
- 使用 PostgreSQL：設定 `database=postgres` 及 `portgresStr`（Postgres 連線字串）
- 資料表會在啟動時由 `databases` 套件使用 GORM `AutoMigrate()` 自動建立並由 `seed` 初始化

---

## API & WebSocket 範例 📡

- GET `/api/v1/user/list`
    - 說明：列出所有使用者
    - 回傳：JSON 陣列（每筆包含 `id`, `name`, `email`）

WebSocket
- 範例頁面：`GET /template/websocket`（提供簡易前端測試 UI）
- 啟動 WebSocket：`GET /websocket/v1/start`
    - 傳入 JSON 範例：
        ```json
        { "ip": "1.2.3.4", "target": 5 }
        ```
    - `ip`：IPv4（驗證 `ip4_addr`）
    - `target`：要 ping 的次數（int）
    - 行為：伺服器會執行 `internal/ping` 的 ping 流程，並透過 WebSocket 回傳執行輸出

---

## 專案結構（摘要） 📁
- `main.go` — 入口
- `gin/` — Web server 啟動、路由與 API
    - `gin/services` — 建立 `gin.Engine`
    - `gin/routes` — 路由定義（API、template、websocket）
    - `gin/api` — API handlers（`userAPI`, `pingAPI`, `template`）
- `internal/` — 業務邏輯（`user`, `ping`）
- `databases/` — DB 初始化與 seed（`sqlites/`, `postgresql/`）
- `model/` — GORM model（目前 `User`）
- `pkg/layout/` — WebSocket、共用 helper
- `validators/` — 自訂驗證
- `template/` — HTML 範例（WebSocket 示範）

---

## 開發提示 & 工具 🧰
- 推薦 VSCode extensions：`eamodio.gitlens`, `golang.go`, `alexcvzz.vscode-sqlite`
- 測試 WebSocket 可開啟 `http://localhost:8081/template/websocket`