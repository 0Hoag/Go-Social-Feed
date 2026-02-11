# 🚀 Deployment Guide

Using Docker Compose, you can deploy the entire stack (Frontend, Backend API, Worker, Database, Message Queue) with a single command.

## Prerequisites

- [Docker](https://docs.docker.com/get-docker/) installed.
- [Docker Compose](https://docs.docker.com/compose/install/) installed (usually comes with Docker).

## 1. Setup Environment Variables

Create a `.env` file in the `deployment/` directory (or root, depending on where you run compose from).

**Copy the contents below into `.env`:**

```env
# Security (CHANGE THESE For Production)
JWT_SECRET=super_secret_jwt_key_please_change
ENCRYPT_KEY=12345678901234567890123456789012 # Must be 32 bytes

# API Keys (Required for features)
GEMINI_API_KEY=your_gemini_key
ETHERSCAN_API_KEY=your_etherscan_key
BSCSCAN_API_KEY=your_bscscan_key
BASESCAN_API_KEY=your_basescan_key
ARBITRUMSCAN_API_KEY=your_arbitrum_key
POLYGONSCAN_API_KEY=your_polygon_key

# Telegram Bot (For Worker Notifications)
TELEGRAM_BOT_TOKEN=your_bot_token
TELEGRAM_CHAT_ID=your_chat_id
BOT_USER_ID=your_user_id
```

## 2. Build and Run

Navigate to the `deployment/` directory and run:

```bash
# Build and start in background
docker-compose -f deployment/docker-compose.yml up -d --build
```

*(Note: If running from root, use `docker-compose -f deployment/docker-compose.yml up -d --build`)*

## 3. Verify Deployment

- **Frontend**: [http://localhost:3000](http://localhost:3000)
- **Backend API**: [http://localhost:8080/health](http://localhost:8080/health) (or similar endpoint)
- **RabbitMQ UI**: [http://localhost:15672](http://localhost:15672) (User: `user`, Pass: `password`)

## 4. Stopping

```bash
docker-compose -f deployment/docker-compose.yml down
```

## Notes for Production

1.  **Domain Setup**: You will need Nginx or Traefik as a reverse proxy to handle SSL (HTTPS) and route `yourdomain.com` to port 3000/8080.
2.  **Passwords**: Change default MongoDB and RabbitMQ passwords in `docker-compose.yml` and updating `MONGODB_URI` / `RABBITMQ_URL` env vars accordingly.
3.  **Persistence**: The `mongodb_data` volume ensures DB data persists across restarts.
