package main

// PlayersDatabase holds
type PlayersDatabase interface {
	Add(players ...Player) error
	Delete(IDs ...int) error
	Get(IDs ...int) []Player
	ListAll() []Player
	Update(ID int, player Player) error
}

type InMemoryPlayersDatabase struct {
	allPlayers map[int]*Player
	nextID     int
}

func NewInMemoryPlayersDatabase() *InMemoryPlayersDatabase {
	return &InMemoryPlayersDatabase{
		allPlayers: make(map[int]*Player),
		nextID:     1,
	}
}

func (m *InMemoryPlayersDatabase) Add(players ...Player) error {
	for _, player := range players {
		player.ID = m.nextID
		m.allPlayers[player.ID] = &player
		m.nextID++
	}
	return nil
}

func (m *InMemoryPlayersDatabase) Delete(IDs ...int) error {
	for _, id := range IDs {
		delete(m.allPlayers, id)
	}
	return nil
}

func (m *InMemoryPlayersDatabase) Get(IDs ...int) []Player {
	players := make([]Player, len(IDs))
	for i, id := range IDs {
		players[i] = *m.allPlayers[id]
	}
	return players
}

func (m *InMemoryPlayersDatabase) ListAll() []Player {
	players := make([]Player, 0)
	for _, player := range m.allPlayers {
		players = append(players, *player)
	}
	return players
}

func (m *InMemoryPlayersDatabase) Update(ID int, player Player) error {
	m.allPlayers[ID] = &player
	return nil
}
