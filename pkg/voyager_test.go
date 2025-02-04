package voyager_test

// import (
// 	"io/ioutil"
// 	"net/http/httptest"
// 	"testing"
//
// 	voyager "github.com/kiran94/graphql-voyager-go/pkg"
// 	"github.com/stretchr/testify/assert"
// )
//
// func TestNewVoyagerHandler(t *testing.T) {
//
// 	var cases = []struct {
// 		endpoint string
// 	}{
// 		{endpoint: "graphql"},
// 		{endpoint: "v1/graphql"},
// 		{endpoint: "v2/graphql"},
// 	}
//
// 	for _, currentCase := range cases {
// 		t.Run(currentCase.endpoint, func(t *testing.T) {
// 			t.Parallel()
//
// 			handler := voyager.NewVoyagerHandler(currentCase.endpoint)
//
// 			server := httptest.NewServer(handler)
// 			defer server.Close()
//
// 			client := server.Client()
// 			response, err := client.Get(server.URL)
// 			assert.Nil(t, err)
// 			defer response.Body.Close()
//
// 			rawResponse, err := ioutil.ReadAll(response.Body)
// 			assert.Nil(t, err)
//
// 			stringResponse := string(rawResponse)
// 			assert.Contains(t, stringResponse, currentCase.endpoint)
// 			assert.Contains(t, stringResponse, "GraphQLVoyager.init")
// 			assert.Contains(t, stringResponse, "fetch(url")
// 			assert.Contains(t, stringResponse, "introspectionQuery")
// 		})
// 	}
// }
