package emoji

// Map returns the emoji at the provided position.
// This list is from 0-255
func Map(b byte) Words {
	return mapGen.EmojiWords[int(b)]
}

// Version returns the version of the emojisum document currently compiled
// against
func Version() string {
	return mapGen.Version
}

var mapGen VersionedMap

// VersionedMap is the structure used for the `emojimap.json` document
type VersionedMap struct {
	Description string `json:"description"`
	Version     string `json:"version"`
	// these are an ordered list, referened by a byte (each byte of a checksum digest)
	EmojiWords []Words `json:"emojiwords"`
}

// Words are a set of options to represent an emoji.
// Possible options could be the ":colon_notion:" or a "U+26CF" style codepoint.
type Words []string
