package services;

import (
    "github.com/tidwall/gjson"
    knowledgeModels "github.com/markdingemanse/loveless/models/knowledge"
    helpers "github.com/markdingemanse/loveless/services/helpers"
    // configs "github.com/markdingemanse/loveless/configs"
);

// ################### REDIFY API INTERACTION ###################### //

func HttpRedify(url string, userAgent string, key string) (string) {
    req, err := helpers.GetBaseRequest(url, userAgent);

    if (err != nil) {
        helpers.HandleError(err)
    }

    client := helpers.GetClient();
    resp, err := helpers.HandleRequest(client, req);

    if (err != nil) {
        helpers.HandleError(err)
    }

    hash  := helpers.ParseResponse(resp, key);

    return GetTitle(hash);
}

func GetTitle(hash gjson.Result) (string) {
    return hash.Array()[0].String();
}

// ################### REDIFY DB INTERACTION ###################### //

// Get the most recent Redify record from the db.
func DbRedify() (string) {
    uri := helpers.DatabaseUrl();
    db := helpers.OpenDbConnection(uri);

    model := knowledgeModels.CreateModel();
    helpers.SelectTable("rss", db).Last(&model);

    return model.GetFirstPost();
}

// Update the Redify db with a new record.
func UpdateRedify(recentPost string) (bool) {
    uri := helpers.DatabaseUrl();
    db := helpers.OpenDbConnection(uri);

    redify := knowledgeModels.Rss{FirstPost: recentPost};
    helpers.SelectTable("rss", db).Create(&redify)

    return true;
}
