package ads

import (
	"context"
	"encoding/xml"
	"net/http"
)

type Result struct {
	Title    string   `xml:"title"`
	Authors  []string `xml:"author"`
	Abstract string   `xml:"abstract"`
}

type Results []Result

// Query sends query to ADS and returns the results.
func Query(ctx context.Context, url string) (Results, error) {
	// Prepare the ADS Query API request.
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Set("data_type", "XML")
	req.URL.RawQuery = q.Encode()

	// Issue the HTTP request and handle the response. The httpDo function
	// cancels the request if ctx.Done is closed.
	var results Results
	err = httpDo(ctx, req, func(resp *http.Response, err error) error {
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		// Parse the XML query result.
		var data struct {
			Records []Result `xml:"record"`
		}

		if err := xml.NewDecoder(resp.Body).Decode(&data); err != nil {
			return err
		}
		for _, res := range data.Records {
			results = append(results, res)
		}
		return nil
	})
	// httpDo waits for the closure we provided to return, so it's safe to
	// read results here.
	return results, err
}

// httpDo issues the HTTP request and calls f with the response. If ctx.Done is
// closed while the request or f is running, httpDo cancels the request, waits
// for f to exit, and returns ctx.Err. Otherwise, httpDo returns f's error.
func httpDo(ctx context.Context, req *http.Request, f func(*http.Response, error) error) error {
	// Run the HTTP request in a goroutine and pass the response to f.
	c := make(chan error, 1)
	req = req.WithContext(ctx)
	go func() { c <- f(http.DefaultClient.Do(req)) }()
	select {
	case <-ctx.Done():
		<-c // Wait for f to return.
		return ctx.Err()
	case err := <-c:
		return err
	}
}
