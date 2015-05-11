package models

import (
	"gopkg.in/mgo.v2"
)

const (
	MgoHost    = "localhost"
	BlogDB     = "blog"
	AdminC     = "admin"
	BlogC      = "blogs"
	TagC       = "tags"
	BlogCacheC = "blog_cache"
	BlogTagC   = "blog_tag"
	HistoryC   = "historys"
)

type Mgo struct {
	session *mgo.Session
}

func NewDB() (*Mgo, error) {
	session, err := mgo.Dial(MgoHost)
	if err != nil {
		return nil, err
	}
	return &Mgo{session}, nil
}
func (m *Mgo) Close() {
	m.session.Close()
}
