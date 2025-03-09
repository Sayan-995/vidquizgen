package types
type Question struct {
    QID              int     `json:"qid"`
    Title            *string  `json:"title"`
    TitleSlug        *string  `json:"titleslug"`
    Difficulty       *string  `json:"difficulty"`
    AcceptanceRate   *float64 `json:"acceptancerate"`
    PaidOnly         *bool    `json:"paidonly"`
    TopicTags        *string  `json:"topictags"` 
    CategorySlug     *string  `json:"categoryslug"`
}
type Transcript struct{
    Text []struct{
        Content string `xml:",chardata"`
    }`xml:"text"`
}

var (
    SummerizationPrompt = 
    `Context: You are an expert in summarizing educational content, especially in Data Structures and Algorithms (DSA). Your goal is to create a concise yet informative summary that captures the most important concepts, explanations, and key takeaways from the given YouTube video transcript.

    Task: Summarize the transcript strictly in less than 1000 words while ensuring that all crucial DSA concepts, examples, and explanations are preserved. Remove any unnecessary filler words, introductions, or repetitive explanations.Also give answer in a simple string format,do not use tables or other things
    If a/multiple question is/are discussed then give answer in the following format
    Title:

    Statement:

    Hints:

    Topics:

    else generate the normal summery according to the topic

    here is the text
    
    Text: %s` 
)