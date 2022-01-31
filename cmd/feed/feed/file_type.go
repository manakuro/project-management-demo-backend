package feed

import (
	"context"
	"log"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/filetype"
)

var fileTypeFeed = struct {
	image ent.CreateFileTypeInput
	pdf   ent.CreateFileTypeInput
	text  ent.CreateFileTypeInput
}{
	image: ent.CreateFileTypeInput{Name: "Image", TypeCode: filetype.TypeCodeImage},
	pdf:   ent.CreateFileTypeInput{Name: "PDF", TypeCode: filetype.TypeCodePDF},
	text:  ent.CreateFileTypeInput{Name: "Text", TypeCode: filetype.TypeCodeText},
}

// FileType generates color data
func FileType(ctx context.Context, client *ent.Client) {
	_, err := client.FileType.Delete().Exec(ctx)
	if err != nil {
		log.Fatalf("FileType failed to delete data: %v", err)
	}

	data := []ent.CreateFileTypeInput{
		fileTypeFeed.image,
		fileTypeFeed.pdf,
		fileTypeFeed.text,
	}
	bulk := make([]*ent.FileTypeCreate, len(data))
	for i, t := range data {
		bulk[i] = client.FileType.Create().SetInput(t)
	}
	if _, err = client.FileType.CreateBulk(bulk...).Save(ctx); err != nil {
		log.Fatalf("FileType failed to feed data: %v", err)
	}
}
