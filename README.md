# 📬 smtp2zoho

[![Docker Pulls](https://img.shields.io/docker/pulls/lacinf/smtp2zoho?style=flat-square)](https://hub.docker.com/r/lacinf/smtp2zoho)
[![Docker Image Version (latest by date)](https://img.shields.io/docker/v/lacinf/smtp2zoho?sort=date&label=version&style=flat-square)](https://hub.docker.com/r/lacinf/smtp2zoho/tags)

**Minimal SMTP server that receives emails and forwards them to the [Zoho Mail API](https://www.zoho.com/mail/help/api/overview.html)**.

Built for `arm64` systems — compatible with Oracle Cloud Ampere, AWS Graviton, Raspberry Pi 4+, and more.

> This project is a customized version of [`alash3al/smtp2http`](https://github.com/alash3al/smtp2http), adapted to serve as a lightweight SMTP bridge for Zoho Mail users.

---

## ✨ Features

- 📩 Accepts emails via SMTP (plain or HTML)
- 🚀 Sends `to`, `subject`, and `body` to the Zoho Mail API
- 🛡️ Supports both access token and refresh token OAuth flows
- 🔁 Automatically retrieves access tokens using a `refresh_token`
- 🧠 Designed for multi-client usage via container-level config
- 🧱 Lightweight Alpine base — ideal for ARM64 deployments

---

## 🐳 Usage (Docker)

### 🔧 Environment variables required:

| Variable               | Description                               | Example                                           |
|------------------------|-------------------------------------------|---------------------------------------------------|
| `ZOHO_API_URL`         | Zoho Mail API endpoint                    | `https://mail.zoho.com/api/accounts/123456789/messages` |
| `ZOHO_FROM_ADDRESS`    | Authorized sender address                 | `alerts@yourdomain.com`                          |
| `ZOHO_TOKEN`           | (Optional) Access token                   | `Zoho-oauthtoken abcdef123456`                   |
| `ZOHO_CLIENT_ID`       | (Optional) Client ID for token refresh    | `1000.xxxxxxx`                                   |
| `ZOHO_CLIENT_SECRET`   | (Optional) Client secret                  | `zzzzzzzzzzzz`                                   |
| `ZOHO_REFRESH_TOKEN`   | (Optional) Refresh token for auto-renewal | `1000.yyyyyyyyy`                                 |

---
<pre> ```bash #
docker run -p 1025:1025 \
  -e ZOHO_API_URL="https://mail.zoho.com/api/accounts/123456789/messages" \
  -e ZOHO_FROM_ADDRESS="alerts@yourdomain.com" \
  -e ZOHO_CLIENT_ID="your_client_id" \
  -e ZOHO_CLIENT_SECRET="your_client_secret" \
  -e ZOHO_REFRESH_TOKEN="your_refresh_token" \
  lacinf/smtp2zoho:0.2
``` </pre>
---

## ⚙️ Example: docker-compose.yml
<pre> ```yaml #
version: '3.8'

services:
  smtp2zoho:
    image: lacinf/smtp2zoho:0.2
    container_name: smtp2zoho
    ports:
      - "1025:1025"
    environment:
      ZOHO_API_URL: https://mail.zoho.com/api/accounts/123456789/messages
      ZOHO_FROM_ADDRESS: alerts@yourdomain.com
      ZOHO_CLIENT_ID: your_client_id
      ZOHO_CLIENT_SECRET: your_client_secret
      ZOHO_REFRESH_TOKEN: your_refresh_token
``` </pre>
---

## 🔐 Token Management

You can now use either:

- A short-lived `access_token` (via `ZOHO_TOKEN`), **or**
- A long-lived `refresh_token` along with `ZOHO_CLIENT_ID` and `ZOHO_CLIENT_SECRET` for automated renewal

If `ZOHO_TOKEN` is not defined, the service will automatically retrieve a new `access_token` using the refresh flow.

---

## 🗂 Supported Versions

| Tag   | Feature set                          |
|--------|--------------------------------------|
| `0.1`  | Static access token only             |
| `0.2`  | Automatic token refresh (recommended)|

---

## 💡 Notes

- This service acts as a **bridge**: it does not store or queue emails.
- It is **not a full MTA** — intended for use with alerting and automation systems like [Chatwoot](https://www.chatwoot.com/) and [Typebot](https://typebot.io/).
- Ensure your `fromAddress` is authorized in your Zoho account.

---

## 🧑‍💻 Author & Credits

- ✨ Project by [Fabiana Azevedo](https://github.com/fabianaadelaide) & ChatGPT (code + docs)
- 💡 Based on original work by [@alash3al](https://github.com/alash3al)

---

## 🖼️ License

MIT — Use freely, modify, share. Contributions welcome!
