// GOALS
// 1. Write HTTP endpoint that returns "hello world" ✅
// 2. Deploy as google Cloud Function ✅
// -- what's up with go.mod? what's the significance of "module ... helloworld"
// 3. Update to write request payload to Google Firestore
// 4. Deploy updated version as google Cloud Function

// starter: https://cloud.google.com/functions/docs/create-deploy-gcloud
package helloworld

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("ingest", ingest)
}

//	{
//		data: number | string | boolean'
//		tags: Array<string> | undefined
//	}
type RequestBody struct {
	Data interface{} `json:"data"`
	Tags []string    `json:"tags,omitempty"`
}

// ingest is an HTTP Cloud Function.
func ingest(w http.ResponseWriter, r *http.Request) {
	// Use the application default credentials
	// ctx := context.Background()
	// conf := &firebase.Config{ProjectID: projectID}
	// app, err := firebase.NewApp(ctx, conf)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// client, err := app.Firestore(ctx)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer client.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var req RequestBody
	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Fprint(w, req.Data)
}
