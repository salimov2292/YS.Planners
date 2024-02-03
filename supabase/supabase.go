package supabase

import (
	"planners/types"

	"github.com/nedpals/supabase-go"
)

type Client struct {
	*supabase.Client
}

func NewClient(supabaseUrl, supabaseKey string) *Client {

	return &Client{
		Client: supabase.CreateClient(supabaseUrl, supabaseKey),
	}
}

func (c *Client) InsertTask(task types.Tab) error {
	var results []types.Tab
	err := c.Client.DB.From("tasks").Insert(task).Execute(&results)
	return err
}
