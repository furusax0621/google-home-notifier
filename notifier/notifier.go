package notifier

import (
	"context"
	"net"
	"net/url"

	"github.com/barnybug/go-cast"
	"github.com/barnybug/go-cast/controllers"
)

// Notifier is a google-home-notifier-go client
type Notifier struct {
	client *cast.Client
}

// NewClient makes a connection and create a client
func NewClient(ctx context.Context, host string, port int) (*Notifier, error) {
	ips, err := net.LookupIP(host)
	if err != nil {
		return nil, err
	}
	client := cast.NewClient(ips[0], port)
	if err := client.Connect(ctx); err != nil {
		return nil, err
	}

	return &Notifier{
		client: client,
	}, nil
}

// Notify sends a message to google home
func (n *Notifier) Notify(ctx context.Context, text, language string) error {
	u := &url.URL{
		Scheme: "https",
		Host:   "translate.google.com",
		Path:   "translate_tts",
	}

	q := u.Query()
	q.Add("ie", "UTF-8")
	q.Add("q", text)
	q.Add("tl", language)
	q.Add("client", "tw-ob")
	u.RawQuery = q.Encode()

	return n.Play(ctx, u.String())
}

//Play sound via URL
func (n *Notifier) Play(ctx context.Context, url string) error {
	media, err := n.client.Media(ctx)
	if err != nil {
		return err
	}

	item := controllers.MediaItem{
		ContentId:   url,
		StreamType:  "BUFFERED",
		ContentType: "audio/mpeg",
	}
	_, err = media.LoadMedia(ctx, item, 0, true, map[string]interface{}{})
	return err
}

//Stop sound
func (n *Notifier) Stop(ctx context.Context) error {
	if !n.client.IsPlaying(ctx) {
		return nil
	}
	media, err := n.client.Media(ctx)
	if err != nil {
		return err
	}
	_, err = media.Stop(ctx)
	return err
}

// Quit
func (n *Notifier) Quit(ctx context.Context) error {
	receiver := n.client.Receiver()
	_, err := receiver.QuitApp(ctx)
	return err
}

// Close connection
func (n *Notifier) Close() {
	n.client.Close()
}
