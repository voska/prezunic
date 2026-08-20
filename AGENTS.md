# Prezunic CLI

Go CLI for ordering from Prezunic (`www.prezunic.com.br`), a Rio de Janeiro supermarket
chain. Useful alongside the Zona Sul CLI for price comparison.

All logic lives in **`github.com/voska/vtexkit`**. This repo holds only the descriptor.

## Login is access-code only

Classic password login is disabled at this store and its OAuth provider ("Prezunic Login")
is custom. `auth login` refuses by design. Use:

```
prezunic auth code send --email <email>
prezunic auth code verify --code <code>
```

This needed a vtexkit fix (v0.5.1): a scoped authentication-start call reports
`accessKey=false` here even though VTEX ID accepts an access key, so the probe now falls
back to an unscoped call.

## Build & test

`make build` `make test` `make lint` `make vet` `make ci`

## Commits

Conventional Commits (`feat:`, `fix:`, `chore:`, `docs:`, `test:`, `refactor:`).
