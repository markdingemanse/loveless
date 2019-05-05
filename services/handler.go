package services;

import (
    "fmt"

    configs "github.com/markdingemanse/loveless/configs"
    tasks "github.com/markdingemanse/loveless/services/tasks"
    helpers "github.com/markdingemanse/loveless/services/helpers"
)

// Gets the needed function to run dynamicly based on a k / v array.
func Functions(key string) func() {
    return map[string] func() {
        "rss": rss}[key];
}

// Handle the rss task.
func rss() {
    configs.Info("[RSS_HANDLER] Starting the process of redify.");
    redifycsv := tasks.DBRedifyURI();

    for _, element := range redifycsv {
        configs.Info("[RSS_HANDLER] Running redify for the option : [" + element.Slug + "]");
        recentPost := tasks.HttpRedify(element.Uri, "loveless", "data.children.#.data.title");
        registerdPost := tasks.DbRedify();
        helpers.Communicate("test")
        if (recentPost == registerdPost) {
            configs.Info("[RSS_HANDLER] There are no new post on redit for the option: [" + element.Slug + "]");
            return;
        }

        configs.Info("[RSS_HANDLER] There seems to be a new post on redit for the option: [" + element.Slug + "]");
        succes := tasks.UpdateRedify(recentPost);

        if (!succes) {
            fmt.Println("[RSS_HANDLER] Error occured when trying to insert new redify reccord");
        }


    }
}
