package mle_tracking

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	"context"
	"io"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

func TrackingHTTP(w http.ResponseWriter, r *http.Request) {
	var d struct {
		Company string `json:"company"`
		IsoCode string `json:"iso"`
		AWB string `json:"awb"`
	}
	// get secret projects/53494987374/secrets/Tracking-API-Key/versions/1
	secret := accessSecret(w, "projects/53494987374/secrets/Tracking-API-Key/versions/1")
	// validate api token
	if !validateApiKey(r) {
		fmt.Fprint(w, "Unauthorized")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		fmt.Fprint(w, "Error")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if d.AWB == "" {
		fmt.Fprint(w, "error no AWB received")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if d.Company == "" {
		fmt.Fprint(w, "error no company received")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if d.IsoCode == "" {
		fmt.Fprint(w, "error no iso received")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Company: %s, ISO: %s, AWB: %s, Secret: %s", html.EscapeString(d.Company), html.EscapeString(d.IsoCode), html.EscapeString(d.AWB), html.EscapeString(secret))
}

func extractAPIKey(r *http.Request) string {
	return r.Header.Get("X-API-KEY")

}

func validateApiKey(r *http.Request) bool {
	return extractAPIKey(r) == "12345"
}

// secret key: 1295cbba-0c2e-4021-9626-c4f2ef516a30
func accessSecret(w io.Writer, name string) string {
	// Create the client.
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		// return fmt.Errorf("failed to create secretmanager client: %v", err)
		fmt.Fprintf(w, "failed to create secretmanager client: %v", err)
	}
	defer client.Close()
	// Build the request.
	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: name,
	}

	// Call the API.
	result, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		// return fmt.Errorf("failed to access secret version: %v", err)
		fmt.Fprintf(w, "failed to access secret version: %v", err)
	}

	// WARNING: Do not print the secret in a production environment - this snippet
	// is showing how to access the secret material.
	fmt.Fprintf(w, "Plaintext: %s\n", string(result.Payload.Data))
	// return nil
	return string(result.Payload.Data)
}