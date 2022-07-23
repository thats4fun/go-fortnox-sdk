# go-fortnox-sdk

Fortnox client SDK

```
package main

import (
	"context"
	"fmt"
	"log"

	fortnox "github.com/thats4fun/go-fortnox-sdk"
)

func main() {
	client := fortnox.NewClient(
		fortnox.WithAuthOpts("token", "secret"),
	)

	order, err := client.GetOrder(context.Background(), 1)
	if err != nil {
		log.Fatalln(err, "Error getting order")
	}

	fmt.Println(order)
}

```

# Tests

### [Integration Tests]:

1. Setup envs:

- FORTNOX_AUTH_CODE
- FORTNOX_ACCESS_TOKEN
- FORTNOX_CLIENT_SECRET

2.

### [Unit Tests]:

- Run

```
go test ./... -v -count=1
```

# API Implementation

| Name                                   | Implemented? | Tested?   | 
|----------------------------------------|-------------|-----------| 
| accountcharts                          | ✓           | ❌         |  
| absencetransactions                    | ✓           | ❌         |   
| accounts                               | ✓           | ❌         |  
| archive                                | ✓           | ❌         |  
| articlefileconnections                 | ✓           | ❌         |  
| articles                               | ✓           | ❌         |  
| assetfileconnections                   | ✓           | ❌         |  
| assets/types                           | ✓           | ❌         |  
| assets                                 | ✓           | ❌         |  
| attendancetransactions                 | ✓           | ❌         |  
| companyinformation                     | ✓           | ❌         |  
| settings/company                       | ✓           | ❌         |  
| contractaccruals                       | ✓           | ❌         |  
| contracttemplates                      | ✓           | ❌         |  
| contacts                               | ✓            | ❌         |  
| costcenters                            | ✓           | ❌         |  
| currencies                             | ✓           | ❌         |  
| customerreferences                     | ✓           | ❌         |  
| customers                              | ✓           | ❌         |  
| employees                              | ❌           | ❌         |  
| euvatlimitregulation                   | ✓           | ❌         |  
| expenses                               | ❌           | ❌         |  
| financialyears                         | ✓           | ❌         |  
| inbox                                  | ❌           | ❌         |  
| invoiceaccruals                        | ✓           | ❌         |  
| invoicepayments                        | ✓           | ❌         |  
| invoices                               | ✓           | ❌         |  
| labels                                 | ✓           | ❌         |  
| lockedperiod                           | ❌           | ❌         |  
| me                                     | ✓           | ❌         |  
| modesofpayments                        | ❌           | ❌         |  
| offers                                 | ❌           | ❌         |  
| orders                                 | ✓           | ❌         |  
| predefinedaccounts                     | ❌           | ❌         |  
| pricelist                              | ❌           | ❌         |  
| prices                                 | ❌           | ❌         |  
| printtemplates                         | ❌           | ❌         |  
| projects                               | ✓           | ❌         |  
| salarytransactions                     | ❌           | ❌         |  
| scheduletimes                          | ❌           | ❌         |  
| sie                                    | ❌           | ❌         |  
| supplierinvoiceaccruals                | ❌           | ❌         |  
| supplierinvoiceexternalurlconnections  | ❌           | ❌         |  
| supplierinvoicefileconnections         | ❌           | ❌         |  
| supplierinvoicepayments                | ❌           | ❌         |  
| supplierinvoices                       | ❌           | ❌         |  
| suppliers                              | ❌           | ❌         |  
| taxreductions                          | ❌           | ❌         |  
| termsofdeliveries                      | ❌           | ❌         |  
| termsofpayments                        | ❌           | ❌         |  
| emailsenders/trusted                   | ❌           | ❌         |  
| units                                  | ✓           | ❌         |  
| voucherfileconnections                 | ❌           | ❌         |  
| voucherseries                          | ✓           | ❌         |  
| vouchers                               | ❌           | ❌         |  
| wayofdeliveries                        | ✓           | ❌         |  
