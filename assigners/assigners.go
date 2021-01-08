package assigners

import "errors"

// CreditAssigner ...
type CreditAssigner interface {
	Assign(investment int32) (int32, int32, int32, error)
}

// Assigner is a CreditAssigner implementation.
type Assigner struct {
	ID            int   `json:"id"`
	Investment    int32 `json:"investment"`
	Success       bool  `json:"success"`
	CreditType300 int32 `json:"credit_type_300"`
	CreditType500 int32 `json:"credit_type_500"`
	CreditType700 int32 `json:"credit_type_700"`
}

// Assign returns the credit types values based on the investment.
func (a *Assigner) Assign(investment int32) (n300, n500, n700 int32, err error) {
	if investment < 300 {
		err = errors.New("inaccurate or insufficient investment")
		return
	}

	n700 = investment / int32(700)
	n500 = investment / int32(500)
	n300 = investment / int32(300)
	r := 0
	var posibilities []map[int]int
	posibility := make(map[int]int)

	// Calculate posibilities
	for i := 0; i <= int(n300); i++ {
		for j := 0; j <= int(n500); j++ {
			for k := 0; k <= int(n700); k++ {
				posibility[700] = k

				r = (700 * k) + (500 * j) + (300 * i)
				if r == int(investment) {
					posibility[300] = i
					posibility[500] = j
					posibility[700] = k

					posibilities = append(posibilities, posibility)
					posibility = make(map[int]int)
				}
			}
		}
	}

	if len(posibilities) == 0 {
		err = errors.New("inaccurate or insufficient investment")
		n700 = 0
		n500 = 0
		n300 = 0
		return
	}

	// Search the best option
	for _, p := range posibilities {
		if p[700] > 0 && p[500] > 0 && p[300] > 0 {
			n700 = int32(p[700])
			n500 = int32(p[500])
			n300 = int32(p[300])

			a.CreditType300 = n300
			a.CreditType500 = n500
			a.CreditType700 = n700
			a.Success = true
			return
		}
	}

	// Return the first option if there's not a good option
	n700 = int32(posibilities[0][700])
	n500 = int32(posibilities[0][500])
	n300 = int32(posibilities[0][300])

	a.CreditType300 = n300
	a.CreditType500 = n500
	a.CreditType700 = n700
	a.Success = true
	return
}

// Stats represents the assignments stadistics.
type Stats struct {
	TotalAsgmtsDone    int `json:"total_asignments_done"`
	TotalAsgmtsSuccess int `json:"total_asignments_success"`
	TotalAsgmtsFail    int `json:"total_asignments_fail"`

	AverageInvestmentSuccessful float64 `json:"average_investment_successful"`
	AverageInvestmentFail       float64 `json:"average_investment_fail"`
}
