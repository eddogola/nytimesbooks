type Lists struct {
	Status       string `json:"status"`
	Copyright    string `json:"copyright"`
	NumResults   int    `json:"num_results"`
	LastModified string `json:"last_modified"`
	Results      []struct {
		ListName         string `json:"list_name"`
		DisplayName      string `json:"display_name"`
		BestsellersDate  string `json:"bestsellers_date"`
		PublishedDate    string `json:"published_date"`
		Rank             int    `json:"rank"`
		RankLastWeek     int    `json:"rank_last_week"`
		WeeksOnList      int    `json:"weeks_on_list"`
		Asterisk         int    `json:"asterisk"`
		Dagger           int    `json:"dagger"`
		AmazonProductURL string `json:"amazon_product_url"`
		ISBNs            []struct {
			ISBN10 string `json:"isbn10"`
			ISBN13 string `json:"isbn13"`
		} `json:"isbns"`
		BookDetails []struct {
			Title           string `json:"title"`
			Description     string `json:"description"`
			Contributor     string `json:"contributor"`
			Author          string `json:"author"`
			ContributorNote string `json:"contributor_note"`
			Price           int    "price"
			AgeGroup        string `json:"age_group"`
			Publisher       string `json:"publisher"`
			PrimaryISBN13   string `json:"primary_isbn13"`
			PrimaryISBN10   string `json:"primary_isbn10"`
		} `json:"book_details"`
		Reviews []struct {
			BookReviewLink     string `json:"book_review_link"`
			FirstChapterLink   string `json:"first_chapter_link"`
			SundayReviewLink   string `json:"sunday_review_link"`
			ArticleChapterLink string `json:"article_chapter_link"`
		} `json:"reviews"`
	} `json:"results"`
}

type ListsByDate struct {
	Status       string `json:"status"`
	Copyright    string `json:"copyright"`
	NumResults   int    `json:"num_results"`
	LastModified string `json:"last_modified"`
	Results      struct {
		ListName         string `json:"list_name"`
		BestsellersDate  string `json:"bestsellers_date"`
		PublishedDate    string `json:"published_date"`
		DisplayName      string `json:"display_name"`
		NormalListEndsAt int    `json:"normal_list_ends_at"`
		Updated          string `json:"updated"`
		Books            []struct {
			Rank               int    `json:"rank"`
			RankLastWeek       int    `json:"rank_last_week"`
			WeeksOnList        int    `json:"weeks_on_list"`
			Asterisk           int    `json:"asterisk"`
			Dagger             int    `json:"dagger"`
			PrimaryISBN13      string `json:"primary_isbn13"`
			PrimaryISBN10      string `json:"primary_isbn10"`
			Publisher          string `json:"publisher"`
			Description        string `json:"description"`
			Price              int    "price"
			Title              string `json:"title"`
			Author             string `json:"author"`
			Contributor        string `json:"contributor"`
			ContributorNote    string `json:"contributor_note"`
			BookImage          string `json:"book_image"`
			AmazonProductURL   string `json:"amazon_product_url"`
			AgeGroup           string `json:"age_group"`
			BookReviewLink     string `json:"book_review_link"`
			FirstChapterLink   string `json:"first_chapter_link"`
			SundayReviewLink   string `json:"sunday_review_link"`
			ArticleChapterLink string `json:"article_chapter_link"`
			ISBNs              []struct {
				ISBN10 string `json:"isbn10"`
				ISBN13 string `json:"isbn13"`
			} `json:"isbns"`
		} `json:"books"`
		Corrections []struct{} `json:"corrections"`
	} `json:"results"`
}
