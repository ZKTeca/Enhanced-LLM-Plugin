package google

import "github.com/agi-cn/llmplugin/llm"

type Option func(g *Google)

// WithSummarizer 总结内容
fun