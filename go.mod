module github.com/cloudwego/eino

go 1.21

require (
	github.com/bytedance/sonic v1.11.3
	github.com/cloudwego/base64x v0.1.4
	github.com/stretchr/testify v1.9.0
	golang.org/x/net v0.22.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/google/uuid v1.6.0
	github.com/klauspost/cpuid/v2 v2.2.7 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	golang.org/x/arch v0.7.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// Note: google/uuid should be a direct dependency, not indirect.
// Moved it out of the indirect block above to reflect actual usage.
