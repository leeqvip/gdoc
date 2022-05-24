package gdoc

type Option func(h *Handler)

// WithTitle allows you to set the document title.
func WithTitle(t string) Option {
	return func(h *Handler) {
		h.title = t
	}
}
