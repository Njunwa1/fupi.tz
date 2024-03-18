package db

import "go.mongodb.org/mongo-driver/bson"

var (
	pipeline = []bson.M{
		{
			"$group": bson.M{
				"_id": "$urlid",
				"clicksByBrowser": bson.M{
					"$push": bson.M{
						"browser":     "$browser",
						"totalClicks": 1,
						"date":        "$createdat", // Adding date to clicksByBrowser
					},
				},
				"clicksByOS": bson.M{
					"$push": bson.M{
						"os":          "$os",
						"totalClicks": 1,
						"date":        "$createdat", // Adding date to clicksByOS
					},
				},
				"clicksByCountry": bson.M{
					"$push": bson.M{
						"device":      "$country",
						"totalClicks": 1,
						"date":        "$createdat", // Adding date to clicksByDevice
					},
				},
			},
		},
		{
			"$project": bson.M{
				"_id":             0,
				"urlid":           "$_id",
				"clicksByBrowser": 1,
				"clicksByOS":      1,
				"clicksByCountry": 1,
			},
		},
	}
)
