package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header
		want    string
		wantErr bool
	}{
		{
			name:    "valid API key",
			headers: http.Header{"Authorization": []string{"ApiKey test-key-123"}},
			want:    "test-key-123",
			wantErr: false,
		},
		{
			name:    "no authorization header",
			headers: http.Header{},
			want:    "",
			wantErr: true,
		},
		{
			name:    "malformed header - missing scheme",
			headers: http.Header{"Authorization": []string{"test-key-123"}},
			want:    "",
			wantErr: true,
		},
		{
			name:    "malformed header - wrong scheme",
			headers: http.Header{"Authorization": []string{"Bearer test-key-123"}},
			want:    "",
			wantErr: true,
		},
		{
			name:    "malformed header - empty value",
			headers: http.Header{"Authorization": []string{"ApiKey"}},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.headers)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetAPIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
