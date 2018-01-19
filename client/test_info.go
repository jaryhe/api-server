package client

import (
	"net/url"
	"golang.org/x/net/context"
	"encoding/json"
	"fmt"
)

func (cli *Client) TestInfo(ctx context.Context) (error) {
	query := url.Values{}
	query.Set("name", "jary")

	//serverResp, err := cli.post(ctx, "/testget", query, nil, nil)
	serverResp, err := cli.get(ctx, "/testget/"+"jary", query, nil)
	if err != nil {
		fmt.Println(err)
		return err
	}
	var ttype map[string]int
	err = json.NewDecoder(serverResp.body).Decode(&ttype)
	fmt.Println(ttype)
	fmt.Println(serverResp.header)
	ensureReaderClosed(serverResp)
	return err
}
