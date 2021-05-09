package resources

import "github.com/veandco/go-sdl2/sdl"

func CreateBlankTexture(r *sdl.Renderer, w, h int32) (*sdl.Texture, error) {
	/*	surface, err := sdl.CreateRGBSurface(0, w, h, 32, 0, 0, 0, 0)
		if err != nil {
			return nil, err
		}
		defer surface.Free()

		text, err := r.CreateTextureFromSurface(surface)
		if err != nil {
			return nil, err
		}*/

	return r.CreateTexture(sdl.PIXELFORMAT_RGBA8888, sdl.TEXTUREACCESS_TARGET, w, h)
}
