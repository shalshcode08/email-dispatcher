# Email Dispatcher

A concurrent email dispatching system built in Go that sends personalized emails to multiple recipients using worker pools and channels.

## Overview

Email Dispatcher is a high-performance Go application that demonstrates concurrent programming patterns for sending bulk emails. It uses Go's goroutines and channels to efficiently process and send emails to multiple recipients in parallel, with customizable HTML email templates.

## Features

- **Concurrent Processing**: Uses worker pool pattern with 5 concurrent workers
- **Channel-based Communication**: Producer-consumer pattern for efficient email distribution
- **Template Support**: HTML email templates with personalized content
- **CSV Input**: Reads recipient data from CSV files
- **SMTP Integration**: Sends emails via SMTP protocol
- **Development-Friendly**: Includes Mailpit integration for testing without sending real emails

## Project Structure

```
email-dispatcher/
├── main.go          # Entry point with worker pool orchestration
├── producer.go      # CSV reader and recipient loader
├── consumer.go      # Email worker implementation
├── email.tmpl       # HTML email template
├── emails.csv       # Recipient data (name, email)
├── go.mod           # Go module definition
├── info.md          # Mailpit setup instructions
└── README.md        # This file
```

## Architecture

### Components

1. **Producer (`producer.go`)**
   - Reads recipient data from `emails.csv`
   - Parses CSV and creates `Recipient` structs
   - Sends recipients to a channel for processing

2. **Consumer (`consumer.go`)**
   - Worker goroutines that receive recipients from the channel
   - Executes email templates with recipient data
   - Sends emails via SMTP
   - Includes configurable delay between sends (50ms)

3. **Main Orchestrator (`main.go`)**
   - Initializes the recipient channel
   - Spawns producer goroutine
   - Creates worker pool (5 workers)
   - Manages synchronization with WaitGroup
   - Provides template execution functionality

4. **Email Template (`email.tmpl`)**
   - HTML template with personalized placeholders
   - Uses Go's `html/template` package
   - Supports `{{.Name}}` and `{{.Email}}` variables

## Prerequisites

- Go 1.23.5 or higher
- Docker (for Mailpit, optional but recommended for testing)

## Setup

### 1. Clone the Repository

```bash
git clone https://github.com/shalshcode08/email-dispatcher.git
cd email-dispatcher
```

### 2. Set Up Mailpit (Recommended for Testing)

Mailpit is a local SMTP server with a web interface for viewing emails without actually sending them.

```bash
docker run -d \
  --restart unless-stopped \
  --name=mailpit \
  -p 8025:8025 \
  -p 1025:1025 \
  axllent/mailpit
```

- SMTP Server: `localhost:1025`
- Web Interface: `http://localhost:8025`

### 3. Prepare Your Email List

Edit `emails.csv` with your recipient data:

### Run the Application

```bash
go run .
```

Or build and run:

```bash
go build -o email-dispatcher
./email-dispatcher
```

## Configuration

### Adjust Worker Count

In `main.go`, modify the `workerCount` variable:

```go
workerCount := 5  // Change to desired number of concurrent workers
```

### SMTP Settings

In `consumer.go`, update SMTP configuration:

```go
smtpHost := "localhost"  // Your SMTP host
smtpPort := "1025"       // Your SMTP port
```

## How It Works

1. **Initialization**: Main function creates a buffered channel for recipients
2. **Producer Start**: A goroutine reads `emails.csv` and sends recipients to the channel
3. **Worker Pool**: 5 worker goroutines start consuming from the channel
4. **Email Processing**: Each worker:
   - Receives a recipient from the channel
   - Executes the email template with recipient data
   - Sends the email via SMTP
   - Logs the operation
5. **Synchronization**: WaitGroup ensures all emails are sent before program exits
6. **Channel Closure**: When CSV is fully processed, the channel closes and workers finish

## Error Handling

The application handles:
- File reading errors (CSV not found)
- CSV parsing errors
- Template execution errors
- SMTP connection/sending errors

Errors are logged with worker ID and recipient information for debugging.

## Performance

- **Concurrent Workers**: 5 parallel workers (configurable)
- **Throttling**: 50ms delay between sends per worker
- **Throughput**: ~100 emails/second with default settings
- **Scalability**: Adjust worker count based on SMTP server limits

## Development

### Run Tests

```bash
go test ./...
```

### Format Code

```bash
go fmt ./...
```
