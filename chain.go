package imaging

import (
	"image"
	"image/jpeg"
	"image/png"
	"io"
)

type Chain struct {
	img image.Image
	err error
}

func DecodeToChain(r io.Reader, opts ...DecodeOption) (*Chain, error) {
	m, err := Decode(r, opts...)
	if err != nil {
		return nil, err
	}
	return &Chain{
		img: m,
		err: nil,
	}, nil
}

func OpenToChain(s string, opts ...DecodeOption) (*Chain, error) {
	m, err := Open(s, opts...)
	if err != nil {
		return nil, err
	}
	return &Chain{
		img: m,
		err: nil,
	}, nil
}

func (c *Chain) Resize(width, height int, filter ResampleFilter) *Chain {
	if c.err != nil {
		return c
	}
	c.img = Resize(c.img, width, height, filter)
	return c
}

func (c *Chain) CropAnchor(width, height int, anchor Anchor) *Chain {
	if c.err != nil {
		return c
	}
	c.img = CropAnchor(c.img, width, height, anchor)
	return c
}

func (c *Chain) CropCenter(width, height int, anchor Anchor) *Chain {
	if c.err != nil {
		return c
	}
	c.img = CropCenter(c.img, width, height)
	return c
}

func (c *Chain) GaussianBlur(sigma float64, radius int) *Chain {
	if c.err != nil {
		return c
	}
	c.img = GaussianBlur(c.img, sigma, radius)
	return c
}

func (c *Chain) StackBlur(sigma float64, radius int) *Chain {
	if c.err != nil {
		return c
	}
	c.img, c.err = StackBlur(c.img, uint32(radius))
	if c.err != nil {
		return c
	}
	return c
}

// PNGWriter encode image to jpeg format
func (c *Chain) PNGWriter(w io.Writer) error {
	if c.err != nil {
		return c.err
	}
	return png.Encode(w, c.img)
}

// JPEGWriter encode image to jpeg format , quality ranges from 1 to 100 inclusive, higher is better.
func (c *Chain) JPEGWriter(w io.Writer, quality int) error {
	if c.err != nil {
		return c.err
	}
	return jpeg.Encode(w, c.img, &jpeg.Options{
		Quality: quality,
	})
}