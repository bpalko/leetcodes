# Design a rate limiter for an API.

You're working on a Shopify service that handles high-volume webhooks from merchants. You need to prevent abuse by ensuring no client can send more than 100 requests per minute.

## 🔧 Requirements

  - Implement a RateLimiter class or module.

  - Each request is tied to a client ID (string).

  - If a client exceeds the rate limit, reject the request.

  - Your method should return true (accepted) or false (rejected).

  - You don’t need to worry about persistence or distributed systems.