package types

type Event struct {
	Type                     string    `json:"type"`
	StartAnnouncementContent string    `json:"announcementContent"`
	EndAnnouncementContent   string    `json:"endAnnouncementContent"`
	RestartContent           string    `json:"restartContent"`
	Value                    string    `json:"value"`
	CurrentValue             string    `json:"currentValue"`
	LastSender               string    `json:"lastSender"`
	Channel                  string    `json:"channel"`
	SlowModeChannel          bool      `json:"slowModeChannel"`
	MessageProgression       []Message `json:"messageProgression"`
	RequiredSequence         bool      `json:"requiredSequence"`
}

type Message struct {
	ID           string `json:"id"`
	Content      string `json:"content"`
	RequiredRole string `json:"requiredRole"`
}
