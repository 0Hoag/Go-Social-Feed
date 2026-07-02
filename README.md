# 🛡️ CryptoCheck (Go-Social-Feed)

> **A personal project** — a crypto news feed combined with a smart-contract safety scanner.
> Originally started as a small internal-style social feed (hence the repo name `go-social-feed`),
> it grew into a full **crypto platform**: an automated news aggregator + an on-chain
> **Trust Score** scanner for smart contracts, with a Go backend and a Next.js frontend.

Built to practice production-grade backend patterns: **Clean Architecture**, multi-service
workers, MongoDB, RabbitMQ, i18n, AI integration (Gemini), and Docker deployment.

---

## ✨ What it does

CryptoCheck is made of three main parts:

### 1. 📰 Crypto News Feed
- A background **worker** crawls articles from **CoinDesk** and **CoinTelegraph** on a cron schedule.
- Each article is **translated to Vietnamese** (Google Translate, with an optional Gemini pipeline),
  de-duplicated, and published as a post.
- New posts are pushed to a **Telegram channel** automatically.
- Users can browse the feed, create posts, comment, react, and follow other users.

### 2. 🔍 Smart-Contract Scanner
- Fetches verified contract source from block explorers across **multiple chains**
  (Ethereum, BSC, Base, Arbitrum, Polygon) via the Etherscan-family APIs.
- Computes a **Trust Score (0–100)** by detecting:
  - ⚠️ **Risk patterns** — honeypots, hidden mint, upgradeable proxy, `selfdestruct`,
    `tx.origin`, high tax/fees, blacklist functions, mutable fees…
  - ✅ **Safe patterns** — OpenZeppelin, ReentrancyGuard, Timelock, MultiSig, DAO governance, EIP-712…
- Uses a **regex rule engine** with an optional **Gemini AI** deep-analysis layer
  (graceful fallback to regex when AI is unavailable).
- Ships as a **CLI**, an **HTTP API**, and a dedicated **Telegram bot** where users
  send a contract address and get a report back.

### 3. 💻 Web Frontend (Next.js)
- News feed, article pages, user profiles, auth, and a scanner UI.
- Live market data via **TradingView** / lightweight-charts and **DexScreener**,
  plus an AI analysis chat.

---

## 🧰 Tech Stack

| Category | Technologies |
|-----------|--------------|
| **Backend** | Go 1.24, Gin, Clean Architecture |
| **Frontend** | Next.js 16, React 19, TypeScript, Tailwind CSS, Zustand |
| **Database** | MongoDB |
| **Message Queue** | RabbitMQ |
| **AI** | Google Gemini (`generative-ai-go`) |
| **Crawling** | Colly |
| **Integrations** | Etherscan/BSCScan/BaseScan/Arbiscan/PolygonScan, DexScreener, Telegram Bot API |
| **i18n** | go-i18n (vi / en / ja) |
| **Docs** | Swagger (swaggo) |
| **Deployment** | Docker Compose, Caddy |

---

## 🧩 Project Structure

```bash
go-social-feed/
├── cmd/
│   ├── api/                # HTTP API server (feed, auth, scanner endpoints)
│   ├── worker/             # Cron worker: crawl → translate → post → Telegram
│   ├── bot/                # Telegram scanner bot (send address → Trust Score)
│   ├── scanner/            # CLI: scan a single contract address
│   └── list_models/        # Dev helper: list available Gemini models
├── config/                 # Env-based configuration
├── docs/                   # Swagger documentation
├── internal/
│   ├── adapters/           # External clients: etherscan, gemini, dexscreener, telegram
│   ├── core/scanner/       # Trust Score engine (regex rules + AI)
│   ├── crawler/            # News crawler manager + site scrapers (coindesk, cointelegraph)
│   ├── processor/          # Article processing / translation (simple + gemini)
│   ├── auth/ users/ post/ comment/ follow/ scanner/   # Feature modules
│   │                       #   each: delivery/http → usecase → repository/mongo
│   ├── delivery/rabbitmq/  # Producer & consumer (notifications)
│   ├── middleware/         # Auth, permission, locale, internal-key
│   ├── models/             # Domain models
│   ├── seeder/             # Idempotent DB seeding on startup
│   └── httpserver/         # Server wiring & route mapping
├── pkg/                    # Shared libs: jwt, mongo, rabbitmq, i18n, encrypter, log, paginator
├── frontend/               # Next.js web app
├── deployment/             # Dockerfiles, docker-compose, Caddyfile
├── go.mod
└── Makefile
```

> Backend feature modules follow **Clean Architecture**:
> `delivery/http` (transport) → `usecase` (business logic) → `repository/mongo` (data).

---

## ⚙️ Getting Started

### 1️⃣ Clone
```bash
git clone https://github.com/0Hoag/go-social-feed.git
cd go-social-feed
```

### 2️⃣ Configure environment
Copy the example and fill in your values:
```bash
cp .env.example .env
```

Key variables (see `.env.example` for the full list):

| Variable | Purpose |
|----------|---------|
| `MONGODB_URI`, `MONGODB_DATABASE` | MongoDB connection |
| `RABBITMQ_URL` | RabbitMQ connection |
| `JWT_SECRET` | JWT signing secret |
| `ENCRYPT_KEY` | Field encryption key (**must be 32 bytes**) |
| `APP_PORT`, `API_MODE` | HTTP server port & mode |
| `GEMINI_API_KEY` | Gemini AI (optional — scanner falls back to regex) |
| `TELEGRAM_BOT_TOKEN`, `TELEGRAM_CHAT_ID` | News-feed Telegram notifications |
| `SCANNER_BOT_TOKEN` | Scanner Telegram bot |
| `ETHERSCAN_API_KEY` … `POLYGONSCAN_API_KEY` | Block-explorer API keys per chain |

### 3️⃣ Run with Docker (full stack)
Brings up MongoDB, RabbitMQ, the API, the worker, the frontend, and Caddy:
```bash
cd deployment
docker-compose up --build
```

### 4️⃣ Run locally (Go)
```bash
go mod tidy

# API server (Swagger + REST)
make run-api          # or: go run cmd/api/main.go

# News worker (crawl → translate → post → Telegram)
go run cmd/worker/main.go

# Telegram scanner bot
go run cmd/bot/main.go

# One-off CLI scan of a contract
go run cmd/scanner/main.go --addr 0x...
```

### 5️⃣ Frontend
```bash
cd frontend
npm install
npm run dev            # http://localhost:3000
```

---

## 🔗 Access

| Service | URL |
|---------|-----|
| API base | `http://localhost:8080/api/v1` |
| Swagger UI | `http://localhost:8080/swagger/index.html` |
| Frontend | `http://localhost:3000` |

Regenerate API docs after changing handlers:
```bash
make swagger
```

---

## 📝 Notes

- This is an **old personal project** built for learning and portfolio purposes — not a company product.
- The Gemini pipeline for news translation is optional; by default the worker uses a free
  Google Translate path, and the scanner falls back to its regex engine when no AI key is set.
- The scanner's Trust Score is a **heuristic, not an audit** — always do your own research (DYOR)
  before interacting with any smart contract.
