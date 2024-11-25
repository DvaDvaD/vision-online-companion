package main

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/portierglobal/vision-online-companion/api/gen"
	"github.com/portierglobal/vision-online-companion/database/data"
)

func (app *application) generateToken() string {
	bytes := make([]byte, 32)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

func (app *application) convertDataIssueToGenIssue(dataIssue data.Issue) gen.Issue {
	return gen.Issue{
		Copy:        int(dataIssue.Copy),
		Description: &dataIssue.Description.String,
		Number:      dataIssue.Number,
	}
}

// func (app *application) convertGenIssueToDataIssue(genIssue gen.Issue, signRequestID string) data.Issue {
// 	return data.Issue{
// 		SignRequestID: signRequestID,
// 		Number:        genIssue.Number,
// 		Copy:          int32(genIssue.Copy),
// 		Description:   pgtype.Text{String: *genIssue.Description, Valid: genIssue.Description != nil},
// 	}
// }
