# Telegram Vacancy Parser

A Go-based Userbot that parses messages from your Telegram groups to find job vacancies.

## Features
- Connects as a Telegram Userbot (monitors groups you are already in).
- Filters messages based on keywords (e.g., "vacancy", "golang", "remote").
- Logs potential job posts to the terminal.

## Prerequisites
- [Go](https://go.dev/) 1.24+
- Telegram API ID and Hash (get them from [my.telegram.org](https://my.telegram.org))

## Setup or Installation

1.  **Clone the repository** (if applicable)

2.  **Environment Variables**
    Copy the example environment file:
    ```bash
    cp .env.example .env
    ```
    Open `.env` and fill in your details:
    ```env
    APP_ID=your_app_id
    APP_HASH=your_app_hash
    PHONE=your_phone_number_with_country_code
    ```

## Usage

1.  **Build the project**:
    ```bash
    go build -o telegram-vacancy-parser
    ```

2.  **Run the parser**:
    ```bash
    ./telegram-vacancy-parser
    ```

3.  **Authentication**:
    - The first time you run it, you will be prompted to enter the login code sent to your Telegram account.
    - A session will be stored so you don't have to log in every time.

4.  **Testing**:
    Send a message to your "Saved Messages" containing keywords like:
    > "Looking for a Golang vacancy with good salary"

    You should see output similar to:
    ```text
    [VACANCY FOUND] ChatID: 123456789 | Content: Looking for a Golang vacancy with good salary
    ```

## Keywords
The parser looks for keywords defined in `parser/parser.go`, such as:
- vacancy, hiring, job
- developer, engineer
- golang, remote, salary
