package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

// QRCode is the domain model for a QR code
type QRCode struct {
	// DestinationURL is the URL that the QR code will redirect to
	DestinationURL string `json:"destination_url" bson:"destination_url"`
	// ShortURL is the URL that the QR code will redirect to
	ShortURL string `json:"short_url" bson:"short_url"`
	// ForegroundColor is the color of the QR code
	ForegroundColor string `json:"foreground_color" bson:"foreground_color"`
	// BackgroundColor is the color of the QR code
	BackgroundColor string `json:"background_color" bson:"background_color"`
	// Logo is the URL of the logo that will be displayed in the center of the QR code
	Logo string `json:"logo" bson:"logo"`
	// FrameColor is the color of the frame
	FrameColor string `json:"frame_color" bson:"frame_color"`
	// FrameText is the text that will be displayed below the frame
	FrameText string `json:"frame_text" bson:"frame_text"`
	// Branded is a boolean when true removes the watermark
	Branded bool `json:"branded" bson:"branded"`
	// ID is the unique identifier for the QR code
	ID primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
}
