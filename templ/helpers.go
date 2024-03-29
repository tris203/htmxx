package templ

import (
	"fmt"
"htmxx/model"
)

func getIDName(basename string, id int) string {
	return fmt.Sprintf("%s-%d", basename, id)
}

func getIDURL(basename string, id int) string {
	return fmt.Sprintf("/%s/%d", basename, id)
}

func getStringURL(basename string, arg int) string {
	return fmt.Sprintf("/%s/%d", basename, arg)
}

func getProfileURL(username string) string {
	return fmt.Sprintf("/%s", username)
}

func getHXSelector(basename string, id int) string {
	return fmt.Sprintf("#%s-%d", basename, id)
}

func getMaxID(tweets []*model.Tweet) int {
	if len(tweets) == 0 {
		return 0
	}
	return tweets[len(tweets)-1].ID
}
