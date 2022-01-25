# go-libwebp 

This is a ccgo translation from [libwebp](https://github.com/webmproject/libwebp/)
in the hope to bring webp codec into Go space. 

`go run . megopher.png` 

Can consume JPEG and PNG images. 

## Notes

- iterations were manually cleaned up to reflect C idiom of treating integers like booleans
        - the cleanup may have been incorrect
        - ccgo issue is open to make the translation step correct to avoid needing manual work
- libwebp depends on libc bsearch, which modernc.org/libc doesn't provide; as such an experimental Go implementation is provided (that actually cheats by doing a simple linear search) 
        - again, this should ideally be upstreamed to ccgo (providing a bsearch)

