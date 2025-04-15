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
- 🛡️ Uses fixed `fromAddress` from environment (ensures compliance with Zoho's authorized senders)
- 🧠 Designed for multi-client usage via container-level config
- 🧱 Lightweight Alpine base — ideal for ARM64 deployments

---

## 🐳 Usage (Docker)

### 🔧 Environment variables required:

| Variable              | Description                            | Example                                           |
|-----------------------|----------------------------------------|---------------------------------------------------|
| `ZOHO_API_URL`        | Zoho Mail API endpoint                 | `https://mail.zoho.com/api/accounts/123456789/messages` |
| `ZOHO_TOKEN`          | Authorization token (OAuth)            | `Zoho-oauthtoken abcdef123456`                   |
| `ZOHO_FROM_ADDRESS`   | Fixed sender (must be authorized)      | `alerts@yourdomain.com`                          |

---

### 🚀 Run with Docker

<pre> ```bash #
docker run -p 1025:1025 \
  -e ZOHO_API_URL="https://mail.zoho.com/api/accounts/123456789/messages" \
  -e ZOHO_TOKEN="Zoho-oauthtoken abcdef1234567890" \
  -e ZOHO_FROM_ADDRESS="alerts@yourdomain.com" \
  lacinf/smtp2zoho:0.1
``` </pre>

---

### ⚙️ Example: docker-compose.yml

<pre> ```yaml #
version: '3.8'
services:
  smtp2zoho:
    image: lacinf/smtp2zoho:0.1
    ports:
      - "1025:1025"
    environment:
      ZOHO_API_URL: https://mail.zoho.com/api/accounts/123456789/messages
      ZOHO_TOKEN: Zoho-oauthtoken abcdef1234567890
      ZOHO_FROM_ADDRESS: alerts@yourdomain.com
``` </pre>

---

---

## 💡 Notes

- This service acts as a **bridge**: it does not store or queue emails.
- It is **not a full MTA** — intended for use with alerting and automation systems like [Chatwoot](https://www.chatwoot.com/) and [Typebot](https://typebot.io/).
- Ensure your `fromAddress` is authorized in your Zoho account.

---

## 🧑‍💻 Author & Credits

- ✨ Project by [Marcio Lacerda](https://github.com/MarcioLacerda) & ChatGPT (code + docs)
- 💡 Based on original work by [@alash3al](https://github.com/alash3al)

---

## 🖼️ License

MIT — Use freely, modify, share. Contributions welcome!
