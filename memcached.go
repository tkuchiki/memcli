package memcli

import "github.com/bradfitz/gomemcache/memcache"

type Client struct {
	client *memcache.Client
}

func NewClient(servers ...string) *Client {
	return &Client{
		client: memcache.New(servers...),
	}
}

func newItem(key string, value []byte, flags uint32, expiration int32) *memcache.Item {
	return &memcache.Item{
		Key:        key,
		Value:      value,
		Flags:      flags,
		Expiration: expiration,
	}
}

func (c *Client) Get(key string) (*memcache.Item, error) {
	return c.client.Get(key)
}

func (c *Client) Delete(key string) error {
	return c.client.Delete(key)
}

func (c *Client) DeleteAll() error {
	return c.client.DeleteAll()
}

func (c *Client) FlushAll() error {
	return c.client.FlushAll()
}

func (c *Client) Set(key string, value []byte, flags uint32, expiration int32) error {
	return c.client.Set(newItem(key, value, flags, expiration))
}

func (c *Client) Close() error {
	c.client = nil
	return nil
}
