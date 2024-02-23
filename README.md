# payments

![GitHub language count](https://img.shields.io/github/languages/count/iuricode/README-template?style=for-the-badge)

Simplified payment processing system. Backend services are written in golang and will contain authentication using jwt, a transaction processing system meant to represent online payment flows, a merchant accounting system representing the offline flows like settlement and merchant risk, an authentication service using jwt and eventually downstream payment processor mocks.

## 💻 Pre-requisites
The authorization service uses EdDSA keys for JWT signing. To create those keys use the following commands, on OpenSSL :

openssl genpkey -algorithm ed25519 -out test-priv.pem
openssl pkey -in test-priv.pem -pubout -out test-pub.pem

## ☕ Running the system
TBA
