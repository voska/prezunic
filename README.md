# prezunic

CLI for [Prezunic](https://www.prezunic.com.br), a supermarket chain in Rio de Janeiro.
Built for humans and AI agents: data to stdout, hints to stderr, stable exit codes,
structured output.

Built on [vtexkit](https://github.com/voska/vtexkit), a shared library for Brazilian VTEX
storefronts.

## Install

```sh
brew install voska/tap/prezunic
```

or `make build` for `bin/prezunic`.

## Logging in is different here

Prezunic has classic password login disabled and uses a custom OAuth provider a CLI
cannot drive. Use the emailed access code:

```sh
prezunic auth code send --email you@example.com
prezunic auth code verify --code 123456
```

`prezunic auth login` will refuse. That is the store, not a bug.

## Use

```sh
prezunic doctor
prezunic search "leite integral" --limit 5
prezunic cart add 11613 --qty 2
prezunic delivery windows
prezunic checkout --window 0            # preview — places nothing
prezunic checkout --window 0 --confirm  # places the order
```

```
$ prezunic search "leite integral" --limit 3
11613      Leite Líquido UHT Glória Integral c/ Tampa 1l        R$6,99  un
11766      Leite UHT Itambé Integral 1 Litro                    R$6,99  un
3565       Leite Líquido Piracanjuba Integral 1 Litro           R$6,99  un
```

Subscriptions are enabled at this store, so `prezunic subs` works — see
`prezunic subs --help`.

## Price-checking against Zona Sul

Both CLIs emit the same JSON shape, so comparing a basket is a shell one-liner:

```sh
for sku in "leite integral" "arroz" "cafe"; do
  p=$(prezunic search "$sku" --json --results-only | jq -r '.[0].price')
  z=$(zonasul  search "$sku" --json --results-only | jq -r '.[0].price')
  printf "%-18s prezunic=%-8s zonasul=%s\n" "$sku" "$p" "$z"
done
```

Prices are integer centavos in `--json`, so they compare and sum without float error.
Note the caveat: matching by search term compares the *top hit* at each store, which is
not necessarily the same product — treat it as a signal, not a like-for-like quote.

## What lives here

Only the store descriptor. Everything else — VTEX client, auth, search, cart, checkout,
subscriptions, output modes, exit codes — is in vtexkit.

## License

MIT
