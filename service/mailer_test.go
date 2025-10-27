package service

import (
	"net/http"
	"strings"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_GetRandomWord(t *testing.T) {
    word := GetRandomWord()
    if strings.TrimSpace(word) == "" {
        t.Error("Expected non-empty word")
    }
}

// func Test_SendEmail(t *testing.T) {

//     emails := []string{"post@pon.com"}
//     sendTests := []struct {
//         word string
//         emails []string
//     }{
//         {"lorem", nil},
//         {"", emails},
//     }
    
// }

func Test_GetAppParticipants(t *testing.T) {

    json := `"participants" : ["post@pon.com", "post1@pon.com", "post2@pon.com"]`

    httpmock.Activate()
    defer httpmock.DeactivateAndReset()
    
    httpmock.RegisterResponder(
        http.MethodGet, 
        "https://example.com",
    //     func(req *http.Request) (*http.Response, error) {
    //        resp, err := httpmock.NewJsonResponse(200, map[string][]interface{}{
    //           "participants" : {"post@pon.com", "post1@pon.com", "post2@pon.com"},
    //        })
    //        return resp, err
    //    },
        httpmock.NewStringResponder(200, json),
    )

    value, _ := GetAppParticipants("https://example.com")
    if value != nil {
        t.Errorf("Expected post@pon.com, post1@pon.com, post2@pon.com, got %s", value)
    }
 
}
