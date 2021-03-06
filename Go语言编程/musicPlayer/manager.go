package library

import (
	"errors"
)

type MusicEntry struct {
	Id string
	Name string
	Artist string
	Source string
	Type string
}

type MusicManager struct {
	musics []MusicEntry
}

func NewMusicManager() *MusicManager {
	return &MusicManager{make([]MusicEntry, 0)}
}

func (m *MusicManager) Len() int {
	return len(m.musics)
}

func (m *MusicManager) Get(index int) (music *MusicEntry, err error) {
	if index < 0 || index >= len(m.musics) {
		return nil, errors.New("Index out of range.")
	}

	return &m.musics[index], nil
}

func (m *MusicManager) Find(name string) (music *MusicEntry, err error) {
	if len(m.musics) == 0 {
		return nil, errors.New("not found.")
	}

	for _, val := range m.musics {
		if val.Name == name {
			return &val, nil
		}
	}

	return nil, errors.New("not found.")
}

func (m *MusicManager) Add(music *MusicEntry) {
	m.musics = append(m.musics, *music)
}

func (m *MusicManager) Remove(index int) (music *MusicEntry, err error) {
	if index < 0 || index >= len(m.musics) {
		return nil, errors.New("not found.")
	}

	removedMusic := &m.musics[index]

	if index < len(m.musics) - 1 {
		m.musics = append(m.musics[:index - 1], m.musics[index + 1:]...)
		
	}else if index == 0 {
		m.musics = make([]MusicEntry, 0)
	} else {
		m.musics = m.musics[:index - 1]
	}
	return removedMusic, nil
}
