package main

var topics = []Topic{
	{
		ID:          "topic-1",
		Name:        "Self-development",
		Description: "Content to help you grow, stay focused, and build discipline.",
		Content: []Content{
			{
				ID:      "content-1",
				Title:   "Focus your mind on what matters",
				TopicID: "topic-1",
				Type:    "audio_clip",
				Source:  "original",
				URI:     "https://storage.example.com/audio/self_dev/focus_clip_1.mp3",
			},
			{
				ID:      "content-2",
				Title:   "Consistency over motivation",
				TopicID: "topic-1",
				Type:    "audio_clip",
				Source:  "original",
				URI:     "https://storage.example.com/audio/self_dev/consistency_clip.mp3",
			},
			{
				ID:      "content-3",
				Title:   "Discipline is freedom (short reading)",
				TopicID: "topic-1",
				Type:    "audiobook_excerpt",
				Source:  "public_domain",
				URI:     "https://storage.example.com/audio/self_dev/discipline_excerpt.mp3",
			},
		},
	},
	{
		ID:          "topic-2",
		Name:        "Spiritual focus",
		Description: "Content to help you reconnect with faith, gratitude, and inner trust.",
		Content: []Content{
			{
				ID:      "content-4",
				Title:   "A moment of gratitude",
				TopicID: "topic-2",
				Type:    "audio_clip",
				Source:  "original",
				URI:     "https://storage.example.com/audio/spiritual/gratitude_clip.mp3",
			},
			{
				ID:      "content-5",
				Title:   "Psalm reading (public domain)",
				TopicID: "topic-2",
				Type:    "audiobook_excerpt",
				Source:  "public_domain",
				URI:     "https://storage.example.com/audio/spiritual/psalm_reading.mp3",
			},
			{
				ID:      "content-6",
				Title:   "Instrumental worship background",
				TopicID: "topic-2",
				Type:    "music",
				Source:  "licensed",
				URI:     "https://storage.example.com/audio/spiritual/instrumental_worship.mp3",
			},
		},
	},
}
