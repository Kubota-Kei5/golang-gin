package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"

	utils "golang-gin/integration/integutils"
	"golang-gin/models"
)

type AlbumTestSuite struct {
	suite.Suite
}

func TestAlbumSuite(t *testing.T) {
	suite.Run(t, new(AlbumTestSuite))
}

func (suite *AlbumTestSuite) TestAlbumCreateGetDelete() {
	// Create
	endpoint := utils.GetEndpoint("/album")
	body, _ := json.Marshal(map[string]string{"Title": "test"})
	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(req)
	defer response.Body.Close()
	suite.Assert().Nil(err)
	suite.Assert().Equal(http.StatusCreated, response.StatusCode)
	bodyBytes, _ := io.ReadAll(response.Body)
	var album models.Album
	json.Unmarshal(bodyBytes, &album)

	// Get
	endpoint = utils.GetEndpoint(fmt.Sprintf("/album/%d", album.ID))
	req, err = http.NewRequest(http.MethodGet, endpoint, nil)
	req.Header.Set("Content-Type", "application/json")
	response, err = client.Do(req)
	suite.Assert().Nil(err)
	suite.Assert().Equal(http.StatusOK, response.StatusCode)

	// Delete
	endpoint = utils.GetEndpoint(fmt.Sprintf("/album/%d", album.ID))
	req, err = http.NewRequest(http.MethodDelete, endpoint, nil)
	req.Header.Set("Content-Type", "application/json")
	response, err = client.Do(req)
	suite.Assert().Nil(err)
	suite.Assert().Equal(http.StatusNoContent, response.StatusCode)
}