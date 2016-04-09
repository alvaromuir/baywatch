package serve

// not in use yet - 04.08.2016, @alvaromuir

// Categories is a json container for taxonomy endpoint results
type Categories struct {
	TotalCount int `json:"total_count"`
	Details    []*CategoryDetail
}

// CategoryDetail describes a BK taxonomy object
type CategoryDetail struct {
	Count   int  `json:"count"`
	HasMore bool `json:"hasMore"`
	Items   []struct {
		CategoryType                    string        `json:"categoryType"`
		ID                              int           `json:"id"`
		IsCountableFlag                 bool          `json:"isCountableFlag"`
		IsExplicitPriceFloorFlag        bool          `json:"isExplicitPriceFloorFlag"`
		IsForNavigationOnlyFlag         bool          `json:"isForNavigationOnlyFlag"`
		IsIncludeForAnalyticsFlag       bool          `json:"isIncludeForAnalyticsFlag"`
		IsLeafFlag                      bool          `json:"isLeafFlag"`
		IsMutuallyExclusiveChildrenFlag bool          `json:"isMutuallyExclusiveChildrenFlag"`
		IsPublicFlag                    bool          `json:"isPublicFlag"`
		Links                           []interface{} `json:"links"`
		Name                            string        `json:"name"`
		NamespaceID                     int           `json:"namespaceId"`
		OwnershipType                   string        `json:"ownershipType"`
		ParentCategory                  struct {
			ID int `json:"id"`
		} `json:"parentCategory"`
		Partner struct {
			ID int `json:"id"`
		} `json:"partner"`
		PathFromRoot struct {
			Ids   []int    `json:"ids"`
			Names []string `json:"names"`
		} `json:"pathFromRoot"`
		PriceFloor int    `json:"priceFloor"`
		SoftFloor  int    `json:"softFloor"`
		SortOrder  int    `json:"sortOrder"`
		Status     string `json:"status"`
		Vertical   struct {
			Name string `json:"name"`
		} `json:"vertical"`
		VisibilityStatus string `json:"visibilityStatus"`
	} `json:"items"`
	Limit       int `json:"limit"`
	Offset      int `json:"offset"`
	TotalResult int `json:"totalResults"`
}
