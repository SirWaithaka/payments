# Payments SDK

This is  a Go SDK for common Kenyan payment apis, which includes payment processors, and digital wallets.

## 🚧 Project Under Development 🚧

**DISCLAIMER:** This project is currently under heavy development. Features may be incomplete, unstable, or subject to
significant changes without notice. Use at your own risk.

_Last updated: Aug 16, 2025_

## Roadmap
- [x] Daraja
    - [x] C2B Stk
    - [x] B2C
    - [x] B2B
    - [x] Transaction Status
    - [x] Account Balance
    - [ ] Reversal
    - [X] Org Name check
- [x] Quikk
    - [x] C2B Stk
    - [x] B2C
    - [x] B2B
    - [x] Transaction Status
    - [x] Account Balance
    - [ ] Refund
- [ ] Tanda
    - [x] Payment Requests
    - [x] Transaction Status
- [ ] JamboPay
- [ ] Airtel Money
- [ ] Pesalink

## Payments SDK
The sdk is a Go wrapper around the listed payments apis, which provides a simple interface to interact with them.

It has the following unique features:
- **Request Hooks**: powerful request hook design which unlocks the ability to extend the sdk with custom hooks that
  that meet unique business cases. Build custom hooks to intercept and modify requests before they are sent as well as hooks
  to intercept and modify responses.
- **Request Retrier**: ready to use retrier with exponential backoff and jitter.

### Installation
Use go get.
```bash
go get github.com/SirWaithaka/payments
```

Then import the payments sdk package into your code
```go
import "github.com/SirWaithaka/payments"
```

### Usage and Documentation
Please see examples for usage.
- [Simple Request](https://github.com/SirWaithaka/payments/blob/main/examples/simple/main.go)
- [Daraja C2B Request](https://github.com/SirWaithaka/payments/blob/main/examples/daraja/main.go)
