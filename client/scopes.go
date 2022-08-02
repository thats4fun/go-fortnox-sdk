package client

import "strings"

const urlEncodedDelim = "%20"

type Scope string

const (
	ScopeSalary             Scope = "salary"
	ScopeBookkeeping        Scope = "bookkeeping"
	ScopeArchive            Scope = "archive"
	ScopeFileConnection     Scope = "connectfile"
	ScopeArticle            Scope = "article"
	ScopeCompanyInformation Scope = "companyinformation"
	ScopeSettings           Scope = "settings"
	ScopeInvoice            Scope = "invoice"
	ScopeCostCenters        Scope = "costcenter"
	ScopeCurrency           Scope = "currency"
	ScopeCustomer           Scope = "customer"
	ScopeInbox              Scope = "inbox"
	ScopePayments           Scope = "payment"
	ScopeNoxFinanceInvoice  Scope = "noxfinansinvoice"
	ScopeOffer              Scope = "offer"
	ScopeOrder              Scope = "order"
	ScopePrice              Scope = "price"
	ScopePrint              Scope = "print"
	ScopeProject            Scope = "project"
	ScopeProfile            Scope = "profile"
	ScopeSupplierInvoice    Scope = "supplierinvoice"
	ScopeSupplier           Scope = "supplier"
)

type Scopes []Scope

var FullScope = Scopes{
	ScopeSalary,
	ScopeBookkeeping,
	ScopeArchive,
	ScopeFileConnection,
	ScopeArticle,
	ScopeCompanyInformation,
	ScopeSettings,
	ScopeInvoice,
	ScopeCostCenters,
	ScopeCurrency,
	ScopeCustomer,
	ScopeInbox,
	ScopePayments,
	ScopeNoxFinanceInvoice,
	ScopeOffer,
	ScopeOrder,
	ScopePrice,
	ScopePrint,
	ScopeProject,
	ScopeProfile,
	ScopeSupplierInvoice,
	ScopeSupplier,
}

func (s Scopes) toStrings() []string {
	res := make([]string, 0, len(s))
	for _, v := range s {
		res = append(res, string(v))
	}
	return res
}

func (s Scopes) toString() string {
	return strings.Join(s.toStrings(), urlEncodedDelim)
}
