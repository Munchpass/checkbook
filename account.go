// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package checkbook

import (
	"github.com/Munchpass/checkbook/option"
)

// AccountService contains methods and other services that help with interacting
// with the checkbook API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountService] method instead.
type AccountService struct {
	Options []option.RequestOption
	Bank    AccountBankService
	Card    AccountCardService
	Interac AccountInteracService
	Paypal  AccountPaypalService
	Vcc     AccountVccService
	Venmo   AccountVenmoService
	Wallet  AccountWalletService
	Wire    AccountWireService
}

// NewAccountService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAccountService(opts ...option.RequestOption) (r AccountService) {
	r = AccountService{}
	r.Options = opts
	r.Bank = NewAccountBankService(opts...)
	r.Card = NewAccountCardService(opts...)
	r.Interac = NewAccountInteracService(opts...)
	r.Paypal = NewAccountPaypalService(opts...)
	r.Vcc = NewAccountVccService(opts...)
	r.Venmo = NewAccountVenmoService(opts...)
	r.Wallet = NewAccountWalletService(opts...)
	r.Wire = NewAccountWireService(opts...)
	return
}
