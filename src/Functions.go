package nekomaid

import (
	"fmt"
	"net/http"
	"encoding/json"
)

/**
 * 以 site / author_id 取得作者下的所有作品
 */
func GetByAuthor (site string, authorId string) (result SearchResult, err error) {
	url := fmt.Sprintf(BASEURL + "artwork.json?site=%s&author_id=%s", site, authorId)
	resp, err := http.Get(url)

	if err != nil {
		resp.Body.Close()
		return result, err
	}
	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&result)

	return result, err
}