package fetch

// PingResponse returns API heartbeat status code
type PingResponse struct {
	Status int
}

// BuyerViewCategoryResult is a collection of OwnerViewCategory types
type BuyerViewCategoryResult struct {
	Count        int  `json:"count"`
	HasMore      bool `json:"hasMore"`
	Items        []*BuyerViewCategory
	Limit        int `json:"limit"`
	Offset       int `json:"offset"`
	TotalResults int `json:"totalResults"`
}

// BuyerViewCategory describes an owner view specific BK category object
type BuyerViewCategory struct {
	CategoryType            string        `json:"categoryType"`
	ID                      int           `json:"id"`
	IsForNavigationOnlyFlag bool          `json:"isForNavigationOnlyFlag"`
	IsLeafFlag              bool          `json:"isLeafFlag"`
	IsPublicFlag            bool          `json:"isPublicFlag"`
	Links                   []interface{} `json:"links"`
	Name                    string        `json:"name"`
	OwnershipType           string        `json:"ownershipType"`
	ParentCategory          struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"parentCategory"`
	Partner struct {
		ID int `json:"id"`
	} `json:"partner"`
	PathFromRoot struct {
		Ids   []int    `json:"ids"`
		Names []string `json:"names"`
	} `json:"pathFromRoot"`
	PriceFloor float32 `json:"priceFloor"`
	Stats      struct {
		Reach int `json:"reach"`
	} `json:"stats"`
	Status   string `json:"status"`
	Vertical struct {
		Name string `json:"name"`
	} `json:"vertical"`
}

/* OwnerViewCategoryResult Not in use yet, @alvaromuir 04.10.2016 */

// OwnerViewCategoryResult is a collection of BuyerViewCategory types
type OwnerViewCategoryResult struct {
	Count        int  `json:"count"`
	HasMore      bool `json:"hasMore"`
	Items        []*OwnerViewCategory
	Limit        int `json:"limit"`
	Offset       int `json:"offset"`
	TotalResults int `json:"totalResults"`
}

// OwnerViewCategory describes a buyer view specific BK category object
type OwnerViewCategory struct {
	CategoryType                    string        `json:"categoryType"`
	Description                     string        `json:"description"`
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
	PriceFloor float32 `json:"priceFloor"`
	SoftFloor  int     `json:"softFloor"`
	SortOrder  int     `json:"sortOrder"`
	Status     string  `json:"status"`
	Vertical   struct {
		Name string `json:"name"`
	} `json:"vertical"`
	VisibilityStatus string `json:"visibilityStatus"`
}

// SiteResult is a collection of Site types
type SiteResult struct {
	TotalCount int `json:"total_count"`
	Sites      []*Site
}

// Site describes a BK site object
type Site struct {
	AllowedBuyers []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"allowed_buyers"`
	AnalyticsOnly             string      `json:"analytics_only"` // REVIEW: should be boo
	BlockedCountries          []string    `json:"blocked_countries"`
	CreatedAt                 string      `json:"created_at"`                  // REVIEW: should be time
	UpdatedAt                 string      `json:"updated_at"`                  // REVIEW: should be time
	DataTransferBoostEnabled  string      `json:"data_transfer_boost_enabled"` // REVIEW: should be bool
	DataTransferBoostInterval int         `json:"data_transfer_boost_interval"`
	DataTransferLimit         int         `json:"data_transfer_limit"`
	GroupID                   int         `json:"group_id"`
	ID                        int         `json:"id"`
	Labels                    interface{} `json:"labels"`
	Name                      string      `json:"name"`
	Notes                     string      `json:"notes"`
	Partner                   struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"partner"`
	PrivateData      string `json:"private_data"`
	TransactionScope string `json:"transaction_scope"`
	Type             int    `json:"type"`
}
