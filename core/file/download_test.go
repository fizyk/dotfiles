package file

import (
	"bytes"
	"errors"
	"github.com/segmentio/ksuid"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

type httpSuite struct {
	suite.Suite
}

// Custom type that allows setting the func that our Mock Do func will run instead
type MockGetType func(url string) (*http.Response, error)

// MockClient is the mock client
type MockClient struct {
	MockGet MockGetType
}

// Overriding what the Do function should "do" in our MockClient
func (m *MockClient) Get(url string) (*http.Response, error) {
	return m.MockGet(url)
}

func TestHTTPSuite(t *testing.T) {
	suite.Run(t, new(httpSuite))
}

func removeFile(fileName string) error {
	if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
		return nil
	}
	return os.Remove(fileName)
}

func (s *httpSuite) TestCreateFile() {
	// create a new reader with that JSON
	var output string = ksuid.New().String()
	var fileName string = "file.test"
	r := ioutil.NopCloser(bytes.NewReader([]byte(output)))
	Client = &MockClient{
		MockGet: func(uri string) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       r,
			}, nil
		},
	}
	err := DownloadFile("http://example.com", fileName)
	defer removeFile(fileName)
	s.NoError(err)
	outputBytes, err := ioutil.ReadFile(fileName)
	s.NoError(err)
	s.Equal(string(outputBytes), output)

}
