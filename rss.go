package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Items       []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {

	client := &http.Client{
		Timeout: 10 * http.DefaultClient.Timeout,
	}
	// Create a new request
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", "curl/7.81.0")
	req.Header.Add("Accept", "application/rss+xml, application/xml; q=0.9, */*; q=0.8")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Pragma", "no-cache")
	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	contentType := resp.Header.Get("Content-Type")
	fmt.Println("Content-Type:", contentType)
	fmt.Println("Status:", resp.Status)
	fmt.Println("Location:", resp.Header.Get("Location"))
	// if contentType != "application/rss+xml" {
	// 	return nil, fmt.Errorf("unexpected content type: %s", contentType)
	// }

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unable to fetch feed: %s", resp.Status)
	}

	// Read the response body
	respBodyData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(respBodyData[:200]))

	// Parse the response body
	feed := &RSSFeed{}
	err = xml.Unmarshal(respBodyData, feed)
	if err != nil {
		fmt.Println("UNMARSHAL ERROR")
		return nil, err
	}

	// Trim the description and title
	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)
	feed.Channel.Link = html.UnescapeString(feed.Channel.Link)
	for i, item := range feed.Channel.Items {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
		feed.Channel.Items[i] = item
	}

	return feed, nil
}
