package fastimage

func (f *FastImage) getWEBPImageSize() (*ImageSize, error) {
	slice, err := f.getBytes(26, 4)
	if err != nil {
		return nil, err
	}

	imageSize := ImageSize{}

	imageSize.Width = uint32(slice[1]&0x3f)<<8 | uint32(slice[0])
	imageSize.Height = uint32(slice[3]&0x3f)<<8 | uint32(slice[2])

	return &imageSize, nil
}
