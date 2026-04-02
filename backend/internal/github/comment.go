package github

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) PostComment(ctx context.Context, owner, repo string, prNumber int, body string) error {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues/%d/comments", owner, repo, prNumber)

	payload := map[string]string{"body": body}

	jsonBody, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("JSONエンコードエラー: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(jsonBody))
	if err != nil {
		return fmt.Errorf("リクエスト生成エラー: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Content-Type", "application/json")

	// 送信
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("APIリクエストエラー: %w", err)
	}
	defer resp.Body.Close()

	// ステータス確認（今回は201）
	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("コメント投稿エラー: status=%d", resp.StatusCode)
	}

	return nil
}
