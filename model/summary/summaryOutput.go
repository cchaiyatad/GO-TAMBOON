package summary

import "fmt"

const (
	summaryFormat = `Summary:

%-25s: %-12.2f THB
%-25s: %-12.2f THB
%-25s: %-12.2f THB

%-25s: %-d times
%-25s: %.2f %%
%-25s: %-12.2f THB

`
	totalReceived  = "total received"
	successDonated = "successfully donated"
	faultyDonate   = "faulty donation"
	all            = "all"
	successRate    = "success rate"
	avg            = "average per person"

	topDonorHeaderFormat = "Top %d Donors:\n"
	topDonorFormat       = "%2d. %-30s : %12.2f THB\n"
)

func (summaries *Summary) PrintSummaries() {
	summaries.printStatus()
	summaries.printTopDonor()
}

func (summaries *Summary) printStatus() {
	fmt.Printf(summaryFormat,
		totalReceived,
		summaries.GetTotalReceived(),
		successDonated,
		summaries.GetAmountSuccess(),
		faultyDonate,
		summaries.GetAmountFail(),
		all,
		summaries.GetTotalCount(),
		successRate,
		summaries.GetSuccessRate(),
		avg,
		summaries.GetAvg(),
	)

}

func (summaries *Summary) printTopDonor() {
	fmt.Printf(topDonorHeaderFormat, len(summaries.donors))
	for i := range summaries.donors {
		fmt.Printf(topDonorFormat,
			i+1,
			summaries.donors[i].Name,
			convertAmountFormat(summaries.donors[i].Amount))
	}
}
