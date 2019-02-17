package services;

import (
    "fmt"

    configs "github.com/markdingemanse/loveless/configs"
    tasks "github.com/markdingemanse/loveless/services/tasks"
)

// Gets the needed function to run dynamicly based on a k / v array.
func Functions(key string) func() {
    return map[string] func() {
        "rss": rss}[key];
}

// Handle the rss task.
// func rss() {
//
//     redifycsv := tasks.DBRedifyURI();
//
//     fmt.Println(redifycsv);
//
//     recentPost := tasks.HttpRedify("https://www.reddit.com/r/symphonicmetal/new.json?sort=new&limit=1", "loveless", "data.children.#.data.title");
//     registerdPost := tasks.DbRedify();
//
//     // configs.Info("[RSS_HANDLER] finished rss most recent API post is: " + recentPost);
//     configs.Info("[RSS_HANDLER] finished rss most recent REGISTERED post is: " + registerdPost);
//         configs.Info(recentPost + " : " + registerdPost);
//     if (recentPost == registerdPost) {
//
//         return;
//     }
//
//     configs.Info("[RSS_HANDLER] There seems to be a new post on redit for the option: [symphonicmetal] my love <3");
//     succes := tasks.UpdateRedify(recentPost);
//
//     if (!succes) {
//         fmt.Println("[RSS_HANDLER] Error occured when trying to insert new redify reccord");
//     }
// }

func rss() {
    configs.Info("[RSS_HANDLER] Starting the process of redify.");
    redifycsv := tasks.DBRedifyURI();

    for _, element := range redifycsv {
        configs.Info("[RSS_HANDLER] Running redify for the option : [" + element.Slug + "]");
        recentPost := tasks.HttpRedify(element.Uri, "loveless", "data.children.#.data.title");
        registerdPost := tasks.DbRedify();

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
