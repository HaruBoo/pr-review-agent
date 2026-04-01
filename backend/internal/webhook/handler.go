package webhook

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Processor interface {
	ProcessPullRequest(event PullRequestEvent) error
}

type Handler struct {
	secret    string
	processor Processor
}

func NewHandler(secret string, processor Processor) *Handler {
	return &Handler{
		secret:    secret,
		processor: processor,
	}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POSTのみ受け付けます", http.StatusMethodNotAllowed)
		return // ← if の中
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "読み取りエラー", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	signature := r.Header.Get("X-Hub-Signature-256")
	if err := h.verifySignature(signature, body); err != nil {
		http.Error(w, "署名が無効です", http.StatusUnauthorized)
		return
	}

	eventType := r.Header.Get("X-GitHub-Event")
	if eventType != "pull_request" {
		w.WriteHeader(http.StatusOK)
		return
	}
}

func (h *Handler) verifySignature(signature string, body []byte) error {
	if !strings.HasPrefix(signature, "sha256=") {
		return fmt.Errorf("署名の形式が不正です")
	}
	received := strings.TrimPrefix(signature, "sha256=")
	mac := hmac.New(sha256.New, []byte(h.secret))
	mac.Write(body)
	expected := hex.EncodeToString(mac.Sum(nil))

	if !hmac.Equal([]byte(received), []byte(expected)) {
		return fmt.Errorf("署名が一致しません")
	}
	return nil
}
