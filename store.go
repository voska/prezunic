// Package prezunic describes the Prezunic storefront
// (www.prezunic.com.br), a supermarket chain in Rio de Janeiro.
package prezunic

import "github.com/voska/vtexkit/store"

// Store is the Prezunic descriptor.
//
// Probed live on 2026-08-20. The one thing worth knowing about this store is
// its login. Prezunic has classic password authentication disabled and offers
// a custom OAuth provider, "Prezunic Login", that VTEX ID cannot drive
// generically. The only way in from a CLI is an emailed access code:
//
//	prezunic auth code send --email you@example.com
//	prezunic auth code verify --code 123456
//
// `prezunic auth login` will refuse, and that is correct rather than a bug.
//
// Finding this took a fix in vtexkit itself. A scoped authentication-start
// call — which is what surfaces the custom provider — reports accessKey=false
// here, so the CLI refused to log in at all. Unscoped, the same store reports
// accessKey=true, and accesskey/send genuinely delivers a code. vtexkit v0.5.1
// falls back to the unscoped probe, which is why no driver is needed and this
// descriptor stays three fields.
//
// Everything else is stock: the account derives correctly from the base URL,
// Intelligent Search REST and the catalog API both answer, and subscriptions
// are enabled (RNS returns 401 rather than 404), so `prezunic subs` works.
//
// No MinOrder: unlike Zona Sul, which enforces R$100 with no API field for it,
// nothing observed at Prezunic advertises a minimum. If checkout starts
// refusing small orders with a store rule, that is the field to add.
//
// No Quirks and no Wishlist hashes — no evidence for either.
var Store = store.Store{
	Name:        "prezunic",
	DisplayName: "Prezunic",
	BaseURL:     "https://www.prezunic.com.br",
}
