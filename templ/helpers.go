package templ

import (
	"fmt"
	"htmxx/model"
)

func getIDName(basename string, id int64) string {
	return fmt.Sprintf("%s-%d", basename, id)
}

func getIDURL(basename string, id int64) string {
	return fmt.Sprintf("/%s/%d", basename, id)
}

func getStringURL(basename string, arg int64) string {
	return fmt.Sprintf("/%s/%d/", basename, arg)
}

func getProfileURL(username string) string {
	return fmt.Sprintf("/%s/", username)
}

func getLikesURL(username string) string {
	return fmt.Sprintf("/likes/%s/", username)
}

func getHXSelector(basename string, id int64) string {
	return fmt.Sprintf("#%s-%d", basename, id)
}

func getMaxID(tweets []*model.Tweet) int64 {
	if len(tweets) == 0 {
		return 0
	}
	finalID := tweets[len(tweets)-1].ID

	if finalID == 0 {
		return 0
	}
	return finalID - 1
}
