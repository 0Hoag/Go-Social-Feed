# 🌍 Real Server Deployment Guide (Vietnam VPS)

To deploy your website in Vietnam for low latency and easy payment, you need a **VPS (Virtual Private Server)**.

## 1. Choose a Vietnam VPS Provider

You need a Linux server (**Ubuntu 22.04** or **24.04**) with at least **2GB RAM** (4GB recommended).

*   **Recommended Options:**
    *   **BizFly Cloud** (Viettel IDC): Pay-as-you-go, stable, good support.
    *   **Vietnix**: Affordable, ~150k-200k VND/month for 2GB RAM.
    *   **CMC Cloud**: Enterprise grade for high stability.
    *   **Tenten / Long Van**: Reliable local providers.

**Action:** Go to their website, sign up, and "Create Server" (Ubuntu 24.04). You will get an **IP Address** (e.g., `103.1.2.3`).

## 2. Connect to Your Server

Open your Mac terminal:

```bash
ssh root@<YOUR_SERVER_IP>
# Enter the password sent to your email (or set during creation)
```

## 3. Install Docker on Server

Run these commands inside the server (copy-paste line by line):

```bash
# Update System
apt update && apt upgrade -y

# Install Docker
curl -fsSL https://get.docker.com -o get-docker.sh
sh get-docker.sh

# Install Docker Compose
apt install -y docker-compose-plugin
```

## 4. Upload Your Code

### Option A: Via Git (Best)
1.  Push your code to GitHub.
2.  On Server:
    ```bash
    git clone https://github.com/your-username/Go-Social-Feed.git
    cd Go-Social-Feed/deployment
    ```

### Option B: Copy from Mac (Simplest if no Git)
On your **Mac Terminal** (not inside server):
```bash
scp -r ~/Downloads/Go-Social-Feed/deployment root@<YOUR_SERVER_IP>:/root/app
```

## 5. Run the App

On the Server:

1.  Create `.env` file:
    ```bash
    nano .env
    # Paste your env variables here (API keys, etc.)
    # Press Ctrl+X, then Y, then Enter to save
    ```

2.  Start the app:
    ```bash
    docker compose -f docker-compose.yml up -d --build
    ```

## 6. Accessing the Website

*   **Frontend**: `http://<YOUR_SERVER_IP>:3000`
*   **Backend**: `http://<YOUR_SERVER_IP>:8080`

Now anyone with this IP can visit your site!

---

## 7. (Optional) Custom Domain (VN or Global)

If you have a domain (e.g., `cryptocheck.vn`):
1.  Go to your Domain Manager (Mắt Bão, PAVietnam, TenTen).
2.  Add an **A Record**:
    *   Name: `@`
    *   Value: `103.1.2.3` (Your VPS IP)
3.  Wait 5-30 mins. You can access via `http://cryptocheck.vn:3000`.

*(To remove the `:3000` port and get HTTPS/SSL, we need to setup **Nginx Reverse Proxy**. Let me know if you want that next!)*
