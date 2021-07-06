package books

const (
	// ListsEndpoint is the endpoint used to Get Best Sellers list. If no date is provided returns the latest list.
	ListsEndpoint = "/lists.json"

	// ListsByDateEndpoint is the endpoint to Get Best Sellers list by date.
	// the first placeholder is a date
	// the second placeholder is the list name
	ListsByDateEndpoint = "/lists/%v/%v.json"

	// HistoryEndpoint is the endpoint used to Get Best Sellers list history.
	HistoryEndpoint = "/lists/best-sellers/history.json"

	// NamesEndpoint is the endpoint used to Get Best Sellers list names.
	NamesEndpoint = "/lists/names.json"

	// OverviewEndpoint is the endpoint used to Get top 5 books for all the Best Sellers lists for specified date.
	OverviewEndpoint = "/lists/overview.json"

	// ReviewsEndpoint is the endpoint for Getting book reviews.
	ReviewsEndpoint = "/reviews.json"
)