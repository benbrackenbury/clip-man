package store

type Store interface {
    GetLatestContent() string
    SetLatestContent(content string)
    Close()
}
