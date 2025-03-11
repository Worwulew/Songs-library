package http

import (
	"encoding/json"
	"fmt"
	"github.com/Worwulew/Songs-library/internal/model"
	"github.com/pkg/errors"
	"net/http"
)

func FetchSongDetails(song *model.Song) (*model.Song, error) {
	apiURL := fmt.Sprintf("https://api.example.com/info?group=%s&song=%s", song.Group, song.SongTitle)
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch data from external API")
	}
	defer resp.Body.Close()

	var details model.SongDetail
	err = json.NewDecoder(resp.Body).Decode(&details)
	if err != nil {
		return nil, err
	}

	song.ReleaseDate = details.ReleaseDate
	song.Text = details.Text
	song.Link = details.Link

	return song, nil
}
