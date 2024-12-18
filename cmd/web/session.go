package main

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	"net/http"
	"sync"
	"time"
)

// Create a structure to represent a sessions and its associated data.
type Session struct {
	ID        string
	Data      map[string]any
	ExpiresAt time.Time
}

func (s *Session) GetOrSetDefault(key string, defaultValue any) any {
	if value, ok := s.Data[key]; ok {
		return value
	} else {
		s.Data[key] = defaultValue
		return defaultValue
	}
}

func (s *Session) Get(key string) any {
	return s.Data[key]
}

func (s *Session) Set(key string, value any) {
	s.Data[key] = value
}

// A session store will hold a number of sessions of type Session
type SessionStore struct {
	sync.RWMutex
	Sessions map[string]*Session
}

func NewSessionStore() *SessionStore {
	return &SessionStore{
		Sessions: make(map[string]*Session),
	}
}

func generateSessionId() string {
	b := make([]byte, 16) // 16 bytes = 128 bits

	// fill the byte slice with cryptographically random bytes.
	_, err := rand.Read(b)

	if err != nil {
		panic(err)
	}

	return hex.EncodeToString(b)
}

func (store *SessionStore) CreateSession(w http.ResponseWriter, name string) *Session {
	sessionId := generateSessionId()

	session := &Session{
		ID:        sessionId,
		Data:      make(map[string]any),
		ExpiresAt: time.Now().Add(30 * time.Minute),
	}

	store.Lock()
	store.Sessions[sessionId] = session
	store.Unlock()

	// Set session id in a cookie
	http.SetCookie(w, &http.Cookie{
		Name:     name,
		Value:    sessionId,
		Expires:  session.ExpiresAt,
		HttpOnly: true,
	})

	return session
}

func (store *SessionStore) GetSession(w http.ResponseWriter, r *http.Request, name string) *Session {
	cookie, err := r.Cookie(name)

	if err != nil {
		log.Println("Unable to find cookie, session_id")
		return store.CreateSession(w, name)
	}

	store.Lock()
	session, exists := store.Sessions[cookie.Value]
	store.Unlock()

	if !exists || session.ExpiresAt.Before(time.Now()) {
		return store.CreateSession(w, name)
	}

	return session
}

func (store *SessionStore) DestroySession(w http.ResponseWriter, r *http.Request, name string) {
	cookie, err := r.Cookie(name)
	if err == nil {
		store.Lock()
		delete(store.Sessions, cookie.Value)
		store.Unlock()
	}

	http.SetCookie(w, &http.Cookie{
		Name:   name,
		Value:  "",
		MaxAge: -1, // expires immediately
	})
}
