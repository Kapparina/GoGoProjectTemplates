package Hello

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestGreet(t *testing.T) {
	tests := []struct {
		name       string
		query      url.Values
		wantOutput string
		wantErr    bool
	}{
		{
			name:       "NoName",
			query:      url.Values{},
			wantOutput: "Hello, World!",
			wantErr:    false,
		},
		{
			name:       "WithName",
			query:      url.Values{"name": []string{"Alice"}},
			wantOutput: "Hello, Alice!",
			wantErr:    false,
		},
		{
			name:       "EmptyName",
			query:      url.Values{"name": []string{""}},
			wantOutput: "Hello, World!",
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/?"+tt.query.Encode(), nil)
			w := httptest.NewRecorder()

			err := greet(w, req)
			resp := w.Result()

			if (err != nil) != tt.wantErr {
				t.Fatalf("greet: unexpected error state, got err = %v, wantErr = %v", err, tt.wantErr)
			}

			body := strings.TrimSpace(w.Body.String())
			if body != tt.wantOutput {
				t.Errorf("greet: unexpected response, got = %q, want = %q", body, tt.wantOutput)
			}

			_ = resp.Body.Close()
		})
	}
}
