package main

import "testing"

func TestGetUrlFromHtml(t *testing.T) {
	tests := []struct {
		name          string
		inputURL      string
		inputBody     string
		expected      []string
	}{
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
					<html>
						<body>
							<a href="/path/one">
								<span>Boot.dev</span>
							</a>
							<a href="https://other.com/path/one">
								<span>Boot.dev</span>
							</a>
						</body>
					</html>
					`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:     "absolute and relative URLs",
			inputURL: "http://localhost:3000",
			inputBody: `
					<html>
						<body>
							<a href="/path/one">
								<span>Boot.dev</span>
							</a>
							<a href="https://other.com/path/one">
								<span>Boot.dev</span>
							</a>
						</body>
					</html>
					`,
			expected: []string{"http://localhost:3000/path/one", "https://other.com/path/one"},
		},
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
					<html>
						<body>
							<a href="/path/one?query=1">
								<span>Boot.dev</span>
							</a>
							<a href="https://other.com/path/one">
								<span>Boot.dev</span>
							</a>
						</body>
					</html>
					`,
			expected: []string{"https://blog.boot.dev/path/one?query=1", "https://other.com/path/one"},
		},
		{
			name: "dumb ai test",
			inputURL: "https://blog.boot.dev",
			inputBody: `<body>
							<div>
								<span>
									<a href="/deep">Deep Link</a>
								</span>
							</div>
						</body>`,
			expected: []string{"https://blog.boot.dev/deep"},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			for i, url := range actual {
				if url != tc.expected[i] {
					t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
				}			
			}
		})
	}
}