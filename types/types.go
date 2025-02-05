package types
type Question struct {
    QID              int     `json:"qid"`
    Title            string  `json:"title"`
    TitleSlug        string  `json:"titleslug"`
    Difficulty       string  `json:"difficulty"`
    AcceptanceRate   float64 `json:"acceptancerate"`
    PaidOnly         bool    `json:"paidonly"`
    TopicTags        string  `json:"topictags"` 
    CategorySlug     string  `json:"categoryslug"`
}
