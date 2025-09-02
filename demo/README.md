# Demo Request Handler

This package handles enterprise demo form submissions with spam protection and email notifications.

## Features

- **Rate Limiting**: 5 requests per IP per hour, 100 total requests per hour
- **Honeypot Protection**: Hidden URL field to catch bots
- **Input Validation**: Server-side validation and sanitization
- **Email Notifications**: 
  - Thank you email to prospect
  - Notification email to admin with reply-to set
- **Exponential Backoff**: Automatic retry for email delivery failures

## Configuration

Set these environment variables in `.env`:

```bash
EMAILIT_API_KEY=your-api-key          # Required: EmailIt API key
EMAILIT_FROM_EMAIL=enterprise@blue.cc # Optional: From email address
EMAILIT_FROM_NAME=Blue Enterprise Team # Optional: From name
NOTIFICATION_EMAIL=manny@blue.cc      # Optional: Where to send notifications
```

## API Endpoint

`POST /api/demo-request`

### Request Body

```json
{
  "fullName": "John Doe",
  "email": "john@company.com",
  "company": "Acme Corp",
  "jobTitle": "VP Engineering",
  "companySize": "250-1000",
  "useCase": "project-management",
  "phone": "+1 555-1234",      // optional
  "message": "We need help...", // optional
  "url": ""                     // honeypot - must be empty
}
```

### Response

Success:
```json
{
  "success": true,
  "message": "Thank you for your interest! We'll contact you within 24 hours."
}
```

Error:
```json
{
  "success": false,
  "error": "Error message here"
}
```

## Testing

Run tests with:
```bash
go test ./demo -v
```

## Security Features

1. **Rate Limiting**: Prevents abuse by limiting requests per IP and globally
2. **Honeypot Field**: Hidden URL field that bots typically fill out
3. **Input Sanitization**: All inputs are HTML-escaped and length-limited
4. **Validation**: Email format, company size, and use case are validated
5. **Silent Bot Handling**: Bots receive success response but no emails are sent