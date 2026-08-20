---
name: prezunic
description: >-
  Order groceries from Prezunic (www.prezunic.com.br), a Rio de Janeiro
  supermarket chain, using the `prezunic` CLI. Also the price check against
  Zona Sul — both CLIs emit comparable JSON.
allowed-tools: Bash, Read
---

# prezunic

Order from Prezunic using the `prezunic` CLI.

## Login is access-code only

This store has classic password login disabled and uses a custom OAuth
provider a CLI cannot drive. **`prezunic auth login` will refuse — that is the
store, not a bug.** Do not retry it. Use:

```bash
prezunic auth code send --email <email>
prezunic auth code verify --code <code>
```

The code arrives by email, so the user has to supply it. Ask, then wait.

## Always start here

```bash
prezunic doctor
```

Exit 0 means ordering will work. Any other exit code means it will not. Each
failed line prints the exact fix. Do that fix, or report it to the user. Do not
retry the command that failed.

## The flow

```bash
prezunic search "leite integral" --limit 5   # 1. find a SKU — first column
prezunic cart add 11613 --qty 2              # 2. add it
prezunic delivery windows                    # 3. pick a window number
prezunic checkout --window 0                 # 4. preview — places nothing
prezunic checkout --window 0 --confirm       # 5. order
```

Between steps 4 and 5: **show the preview to the user and get explicit
approval.** `--confirm` spends real money.

```bash
prezunic cart show
prezunic cart update 0 --qty 3    # index from 'cart show'
prezunic cart remove 0
prezunic cart clear
```

**Search** takes Portuguese terms. **Cart** never needs a `--seller`.

## Price comparison with Zona Sul

This is the main reason Prezunic is in the fleet. Both CLIs emit the same JSON
shape and prices are integer centavos, so they compare and sum without float
error.

```bash
for q in "leite integral" "arroz" "cafe"; do
  p=$(prezunic search "$q" --json --results-only | jq -r '.[0].price')
  z=$(zonasul  search "$q" --json --results-only | jq -r '.[0].price')
  printf "%-18s prezunic=%-8s zonasul=%s\n" "$q" "$p" "$z"
done
```

**State the caveat whenever you report a comparison:** matching by search term
compares the *top hit* at each store, which is often a different brand or size.
It is a signal, not a like-for-like quote. If the user is deciding based on the
number, confirm the two products actually match on name and unit first.

Do not silently switch a user's usual store on price alone — Zona Sul remains
the default for groceries unless the user says otherwise.

## Subscriptions

```bash
prezunic subs                  # status, frequency, next delivery
prezunic subs <id>
prezunic subs pause <id>
prezunic subs resume <id>
prezunic subs skip <id>        # skip the next delivery only
prezunic subs unskip <id>
```

Prefer `skip` to `pause` — skipping drops one delivery and the schedule
continues. There is no `cancel`: VTEX has no transition out of `CANCELED`, so
tell the user to cancel on the website.

## Output for scripts and agents

```bash
prezunic search arroz --json
prezunic search arroz --json --select sku,name,price
prezunic search arroz --plain          # tab-separated
```

Prices in `--json` are integer centavos: `699` is R$6,99.

Data goes to stdout; progress and errors go to stderr.

## Exit codes

| Code | Meaning | What to do |
|---|---|---|
| 0 | success | continue |
| 2 | bad arguments | fix the command; do not retry it unchanged |
| 3 | empty result | tell the user nothing matched |
| 4 | login required | `prezunic auth code send --email <email>` |
| 5 | not found | the SKU is wrong; search again |
| 7, 8 | temporary | wait, then retry once |
| 9 | store rule refused | read the message; it names the rule |
| 10 | not set up | `prezunic doctor` and follow the fixes |

Full table: `prezunic exit-codes --json`

## Rules

- Never run `--confirm` without the user approving that exact cart and total.
- If a command fails twice the same way, stop and report it. Do not loop.
- `prezunic doctor` diagnoses anything unexpected; its output names the fix.
