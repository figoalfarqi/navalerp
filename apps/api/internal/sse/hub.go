package sse

import (
	"encoding/json"
	"net/http"
	"slices"
	"sync"
	"time"

	"github.com/figoalfarqi/navalerp/config"
	"github.com/figoalfarqi/navalerp/internal/auth"
	"github.com/figoalfarqi/navalerp/internal/helper"
	"github.com/figoalfarqi/navalerp/pkg/response"
)

type Client struct {
	ch        chan []byte
	AppUserID int
	AppRoleID int
	topic     string
}

type Hub struct {
	mu      sync.Mutex
	clients map[*Client]bool
}

func NewHub() *Hub {
	return &Hub{clients: make(map[*Client]bool)}
}

func (h *Hub) Add(c *Client) {
	h.mu.Lock()
	h.clients[c] = true
	h.mu.Unlock()
}

func (h *Hub) Remove(c *Client) {
	h.mu.Lock()
	delete(h.clients, c)
	close(c.ch)
	h.mu.Unlock()
}

func (h *Hub) Broadcast(topic, event string, data any) {
	payload, _ := json.Marshal(data)

	msg := []byte(
		"event: " + event + "\n" +
			"data: " + string(payload) + "\n\n",
	)

	h.mu.Lock()
	defer h.mu.Unlock()

	for c := range h.clients {
		if c.topic == topic {
			select {
			case c.ch <- msg:
			default:
				delete(h.clients, c)
			}
		}
	}
}

func (h *Hub) ServeHTTP(role string, cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		flusher, ok := w.(http.Flusher)
		if !ok {
			response.JSON(w, http.StatusInternalServerError, "no flusher", nil, nil)

			return
		}

		token := r.URL.Query().Get("token")
		topic := r.URL.Query().Get("topic")

		// claims, err := ValidateJWT(token)
		claims, err := auth.ParseToken(cfg.JWTKey, token)
		if err != nil {
			response.JSON(w, http.StatusUnauthorized, "invalid token", nil, map[string]string{"error": err.Error()})
			return
		}

		if !slices.Contains(helper.GetAppRoleIDsByRoleTypeName(role), claims.AppRoleID) {
			response.JSON(w, http.StatusUnauthorized, "unauthorized user", nil, nil)
			return
		}

		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		client := &Client{
			ch:        make(chan []byte, 10),
			AppUserID: claims.AppUserID,
			AppRoleID: claims.AppRoleID,
			topic:     topic,
		}

		h.Add(client)
		defer h.Remove(client)

		// connected event
		w.Write([]byte("event: connected\ndata: {}\n\n"))
		flusher.Flush()

		ticker := time.NewTicker(20 * time.Second)
		defer ticker.Stop()

		ctx := r.Context()

		for {
			select {
			case msg := <-client.ch:
				w.Write(msg)
				flusher.Flush()
			case <-ticker.C:
				w.Write([]byte("event: ping\ndata: {}\n\n"))
				flusher.Flush()
			case <-ctx.Done():
				return
			}
		}
	}
}
