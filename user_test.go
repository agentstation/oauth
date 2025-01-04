package oauth

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUser_Validate(t *testing.T) {
	tests := []struct {
		name    string
		user    User
		wantErr bool
		want    User
	}{
		{
			name:    "empty user should fail due to missing email",
			user:    User{},
			wantErr: true,
		},
		{
			name: "user with only email should pass and use email as firstname",
			user: User{
				Email: "test@example.com",
			},
			wantErr: false,
			want: User{
				Email:     "test@example.com",
				FirstName: "test@example.com",
				LastName:  "",
			},
		},
		{
			name: "existing firstname should remain unchanged",
			user: User{
				Email:     "test@example.com",
				FirstName: "John",
				LastName:  "",
			},
			wantErr: false,
			want: User{
				Email:     "test@example.com",
				FirstName: "John",
				LastName:  "",
			},
		},
		{
			name: "should parse full name into first and last name",
			user: User{
				Email: "test@example.com",
				Name:  "John Smith",
			},
			wantErr: false,
			want: User{
				Email:     "test@example.com",
				Name:      "John Smith",
				FirstName: "John",
				LastName:  "Smith",
			},
		},
		{
			name: "should not override existing lastname when parsing full name",
			user: User{
				Email:    "test@example.com",
				Name:     "John Smith",
				LastName: "Existing",
			},
			wantErr: false,
			want: User{
				Email:     "test@example.com",
				Name:      "John Smith",
				FirstName: "John",
				LastName:  "Existing",
			},
		},
		{
			name: "should use nickname when shorter than full name",
			user: User{
				Email:    "test@example.com",
				Name:     "Jonathan Smith",
				NickName: "Jon",
			},
			wantErr: false,
			want: User{
				Email:     "test@example.com",
				Name:      "Jonathan Smith",
				NickName:  "Jon",
				FirstName: "Jonathan",
				LastName:  "Smith",
			},
		},
		{
			name: "should not use nickname when longer than full name",
			user: User{
				Email:    "test@example.com",
				Name:     "Jon",
				NickName: "Jonathan",
			},
			wantErr: false,
			want: User{
				Email:     "test@example.com",
				Name:      "Jon",
				NickName:  "Jonathan",
				FirstName: "Jon",
				LastName:  "",
			},
		},
		{
			name: "should use full name when no other options available",
			user: User{
				Email: "test@example.com",
				Name:  "JohnSmith",
			},
			wantErr: false,
			want: User{
				Email:     "test@example.com",
				Name:      "JohnSmith",
				FirstName: "John",
				LastName:  "Smith",
			},
		},
		{
			name: "should preserve all other fields during validation",
			user: User{
				RawData:           map[string]interface{}{"key": "value"},
				Provider:          "github",
				Email:             "test@example.com",
				Name:              "John Smith",
				Description:       "Test description",
				UserID:            "123",
				AvatarURL:         "http://example.com/avatar",
				Location:          "New York",
				AccessToken:       "token123",
				AccessTokenSecret: "secret123",
				RefreshToken:      "refresh123",
				ExpiresAt:         time.Now(),
				IDToken:           "idtoken123",
			},
			wantErr: false,
			want: User{
				RawData:           map[string]interface{}{"key": "value"},
				Provider:          "github",
				Email:             "test@example.com",
				Name:              "John Smith",
				FirstName:         "John",
				LastName:          "Smith",
				Description:       "Test description",
				UserID:            "123",
				AvatarURL:         "http://example.com/avatar",
				Location:          "New York",
				AccessToken:       "token123",
				AccessTokenSecret: "secret123",
				RefreshToken:      "refresh123",
				ExpiresAt:         time.Now(),
				IDToken:           "idtoken123",
			},
		},
		{
			name: "should handle unparseable name",
			user: User{
				Email: "test@example.com",
				Name:  "!@#$%^&*",
			},
			wantErr: false,
			want: User{
				Email:     "test@example.com",
				Name:      "!@#$%^&*",
				FirstName: "!@#$%^&*",
				LastName:  "",
			},
		},
		{
			name: "should handle name with special characters",
			user: User{
				Email: "test@example.com",
				Name:  "Jean-Pierre D'Artagnan",
			},
			wantErr: false,
			want: User{
				Email:     "test@example.com",
				Name:      "Jean-Pierre D'Artagnan",
				FirstName: "Jean-Pierre",
				LastName:  "D'Artagnan",
			},
		},
		{
			name: "should handle single character name",
			user: User{
				Email: "test@example.com",
				Name:  "J",
			},
			wantErr: false,
			want: User{
				Email:     "test@example.com",
				Name:      "J",
				FirstName: "J",
				LastName:  "",
			},
		},
		{
			name: "should use nickname when firstname is empty",
			user: User{
				Email:    "test@example.com",
				NickName: "Jon",
			},
			wantErr: false,
			want: User{
				Email:     "test@example.com",
				NickName:  "Jon",
				FirstName: "Jon",
				LastName:  "",
			},
		},
		{
			name: "should parse name but not override existing lastname",
			user: User{
				Email:    "test@example.com",
				Name:     "John James",
				LastName: "Smith",
			},
			wantErr: false,
			want: User{
				Email:     "test@example.com",
				Name:      "John James",
				FirstName: "John",
				LastName:  "Smith",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.user.Validate()

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)

			// For the test case with all fields, we need to compare ExpiresAt separately
			if tt.name == "should preserve all other fields during validation" {
				assert.WithinDuration(t, tt.want.ExpiresAt, tt.user.ExpiresAt, time.Second)
				// Set ExpiresAt to zero for both to allow deep equal comparison
				tt.want.ExpiresAt = time.Time{}
				tt.user.ExpiresAt = time.Time{}
			}

			assert.Equal(t, tt.want, tt.user)
		})
	}
}

func TestParseFullName(t *testing.T) {
	tests := []struct {
		name      string
		fullName  string
		wantFirst string
		wantLast  string
	}{
		{
			name:      "simple name",
			fullName:  "John Smith",
			wantFirst: "John",
			wantLast:  "Smith",
		},
		{
			name:      "single name",
			fullName:  "John",
			wantFirst: "",
			wantLast:  "",
		},
		{
			name:      "complex name",
			fullName:  "John James Smith",
			wantFirst: "John",
			wantLast:  "Smith",
		},
		{
			name:      "empty name",
			fullName:  "",
			wantFirst: "",
			wantLast:  "",
		},
		{
			name:      "gibberish name",
			fullName:  "!@#$%^&*",
			wantFirst: "",
			wantLast:  "",
		},
		{
			name:      "name with special characters",
			fullName:  "Jean-Pierre D'Artagnan",
			wantFirst: "Jean-Pierre",
			wantLast:  "D'Artagnan",
		},
		{
			name:      "single character",
			fullName:  "J",
			wantFirst: "",
			wantLast:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			first, last := parseFullName(tt.fullName)
			assert.Equal(t, tt.wantFirst, first)
			assert.Equal(t, tt.wantLast, last)
		})
	}
}
