// Copyright 2017 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

//go:generate go run generated/gen.go

// DO NOT EDIT THIS FILE.
// It is generated from the source in generated/sample-template.go

package main

import (
	"fmt"
	"io"
	"os"

	"cloud.google.com/go/vision"
	"golang.org/x/net/context"
)

func init() {
	// Refer to these functions so that goimports is happy before boilerplate is inserted.
	_ = context.Background()
	_ = vision.NewClient
	_ = os.Open
}

// detectFaces gets faces from the Vision API for an image at the given file path.
func detectFaces(w io.Writer, file string) error {
	ctx := context.Background()

	client, err := vision.NewClient(ctx)
	if err != nil {
		return err
	}

	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	image, err := vision.NewImageFromReader(f)
	if err != nil {
		return err
	}
	annotations, err := client.DetectFaces(ctx, image, 10)
	if err != nil {
		return err
	}

	if len(annotations) == 0 {
		fmt.Fprintln(w, "No faces found.")
	} else {
		fmt.Fprintln(w, "Faces:")
		for i, annotation := range annotations {
			fmt.Fprintln(w, "  Face", i)
			fmt.Fprintln(w, "    Anger:", annotation.Likelihoods.Anger)
			fmt.Fprintln(w, "    Joy:", annotation.Likelihoods.Joy)
			fmt.Fprintln(w, "    Surprise:", annotation.Likelihoods.Surprise)
		}
	}

	return nil
}

// detectLabels gets labels from the Vision API for an image at the given file path.
func detectLabels(w io.Writer, file string) error {
	ctx := context.Background()

	client, err := vision.NewClient(ctx)
	if err != nil {
		return err
	}

	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	image, err := vision.NewImageFromReader(f)
	if err != nil {
		return err
	}
	annotations, err := client.DetectLabels(ctx, image, 10)
	if err != nil {
		return err
	}

	if len(annotations) == 0 {
		fmt.Fprintln(w, "No labels found.")
	} else {
		fmt.Fprintln(w, "Labels:")
		for _, annotation := range annotations {
			fmt.Fprintln(w, annotation.Description)
		}
	}

	return nil
}

// detectLandmarks gets landmarks from the Vision API for an image at the given file path.
func detectLandmarks(w io.Writer, file string) error {
	ctx := context.Background()

	client, err := vision.NewClient(ctx)
	if err != nil {
		return err
	}

	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	image, err := vision.NewImageFromReader(f)
	if err != nil {
		return err
	}
	annotations, err := client.DetectLandmarks(ctx, image, 10)
	if err != nil {
		return err
	}

	if len(annotations) == 0 {
		fmt.Fprintln(w, "No landmarks found.")
	} else {
		fmt.Fprintln(w, "Landmarks:")
		for _, annotation := range annotations {
			fmt.Fprintln(w, annotation.Description)
		}
	}

	return nil
}

// detectText gets text from the Vision API for an image at the given file path.
func detectText(w io.Writer, file string) error {
	ctx := context.Background()

	client, err := vision.NewClient(ctx)
	if err != nil {
		return err
	}

	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	image, err := vision.NewImageFromReader(f)
	if err != nil {
		return err
	}
	annotations, err := client.DetectTexts(ctx, image, 10)
	if err != nil {
		return err
	}

	if len(annotations) == 0 {
		fmt.Fprintln(w, "No text found.")
	} else {
		fmt.Fprintln(w, "Text:")
		for _, annotation := range annotations {
			fmt.Fprintf(w, "%q\n", annotation.Description)
		}
	}

	return nil
}

// detectDocumentText gets the full document text from the Vision API for an image at the given file path.
func detectDocumentText(w io.Writer, file string) error {
	ctx := context.Background()

	client, err := vision.NewClient(ctx)
	if err != nil {
		return err
	}

	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	image, err := vision.NewImageFromReader(f)
	if err != nil {
		return err
	}
	annotation, err := client.DetectDocumentText(ctx, image)
	if err != nil {
		return err
	}

	fmt.Fprintln(w, "Text:")
	fmt.Fprintf(w, "%q\n", annotation.Text)

	return nil
}

// detectProperties gets image properties from the Vision API for an image at the given file path.
func detectProperties(w io.Writer, file string) error {
	ctx := context.Background()

	client, err := vision.NewClient(ctx)
	if err != nil {
		return err
	}

	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	image, err := vision.NewImageFromReader(f)
	if err != nil {
		return err
	}
	props, err := client.DetectImageProps(ctx, image)
	if err != nil {
		return err
	}

	fmt.Fprintln(w, "Dominant colors:")
	for _, quantized := range props.DominantColors {
		color := quantized.Color
		fmt.Fprintf(w, "%2.1f%% - #%02x%02x%02x\n", quantized.PixelFraction*100, color.R&0xff, color.G&0xff, color.B&0xff)
	}

	return nil
}

// detectCropHints gets suggested croppings the Vision API for an image at the given file path.
func detectCropHints(w io.Writer, file string) error {
	ctx := context.Background()

	client, err := vision.NewClient(ctx)
	if err != nil {
		return err
	}

	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	image, err := vision.NewImageFromReader(f)
	if err != nil {
		return err
	}
	hints, err := client.CropHints(ctx, image, nil)
	if err != nil {
		return err
	}

	fmt.Fprintln(w, "Crop hints:")
	for _, hint := range hints {
		fmt.Fprintf(w, "%v\n", hint.BoundingPoly)
	}

	return nil
}

// detectSafeSearch gets image properties from the Vision API for an image at the given file path.
func detectSafeSearch(w io.Writer, file string) error {
	ctx := context.Background()

	client, err := vision.NewClient(ctx)
	if err != nil {
		return err
	}

	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	image, err := vision.NewImageFromReader(f)
	if err != nil {
		return err
	}
	props, err := client.DetectSafeSearch(ctx, image)
	if err != nil {
		return err
	}

	fmt.Fprintln(w, "Safe Search properties:")
	fmt.Fprintln(w, "Adult:", props.Adult)
	fmt.Fprintln(w, "Medical:", props.Medical)
	fmt.Fprintln(w, "Spoofed:", props.Spoof)
	fmt.Fprintln(w, "Violence:", props.Violence)

	return nil
}

// detectWeb gets image properties from the Vision API for an image at the given file path.
func detectWeb(w io.Writer, file string) error {
	ctx := context.Background()

	client, err := vision.NewClient(ctx)
	if err != nil {
		return err
	}

	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	image, err := vision.NewImageFromReader(f)
	if err != nil {
		return err
	}
	web, err := client.DetectWeb(ctx, image)
	if err != nil {
		return err
	}

	fmt.Fprintln(w, "Web properties:")
	if len(web.FullMatchingImages) != 0 {
		fmt.Fprintln(w, "\tFull image matches:")
		for _, full := range web.FullMatchingImages {
			fmt.Fprintf(w, "\t\t%s\n", full.URL)
		}
	}
	if len(web.PagesWithMatchingImages) != 0 {
		fmt.Fprintln(w, "\tPages with this image:")
		for _, page := range web.PagesWithMatchingImages {
			fmt.Fprintf(w, "\t\t%s\n", page.URL)
		}
	}
	if len(web.WebEntities) != 0 {
		fmt.Fprintln(w, "\tEntities:")
		for _, entity := range web.WebEntities {
			fmt.Fprintf(w, "\t\t%-12s %s\n", entity.ID, entity.Description)
		}
	}

	return nil
}

// detectLogos gets logos from the Vision API for an image at the given file path.
func detectLogos(w io.Writer, file string) error {
	ctx := context.Background()

	client, err := vision.NewClient(ctx)
	if err != nil {
		return err
	}

	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	image, err := vision.NewImageFromReader(f)
	if err != nil {
		return err
	}
	annotations, err := client.DetectLogos(ctx, image, 10)
	if err != nil {
		return err
	}

	if len(annotations) == 0 {
		fmt.Fprintln(w, "No logos found.")
	} else {
		fmt.Fprintln(w, "Logos:")
		for _, annotation := range annotations {
			fmt.Fprintln(w, annotation.Description)
		}
	}

	return nil
}

func init() {
	// Refer to these functions so that goimports is happy before boilerplate is inserted.
	_ = context.Background()
	_ = vision.NewClient
	_ = os.Open
}

// detectFaces gets faces from the Vision API for an image at the given file path.
func detectFacesURI(w io.Writer, file string) error {
	ctx := context.Background()

	client, err := vision.NewClient(ctx)
	if err != nil {
		return err
	}

	image := vision.NewImageFromURI(file)
	annotations, err := client.DetectFaces(ctx, image, 10)
	if err != nil {
		return err
	}

	if len(annotations) == 0 {
		fmt.Fprintln(w, "No faces found.")
	} else {
		fmt.Fprintln(w, "Faces:")
		for i, annotation := range annotations {
			fmt.Fprintln(w, "  Face", i)
			fmt.Fprintln(w, "    Anger:", annotation.Likelihoods.Anger)
			fmt.Fprintln(w, "    Joy:", annotation.Likelihoods.Joy)
			fmt.Fprintln(w, "    Surprise:", annotation.Likelihoods.Surprise)
		}
	}

	return nil
}

// detectLabels gets labels from the Vision API for an image at the given file path.
func detectLabelsURI(w io.Writer, file string) error {
	ctx := context.Background()

	client, err := vision.NewClient(ctx)
	if err != nil {
		return err
	}

	image := vision.NewImageFromURI(file)
	annotations, err := client.DetectLabels(ctx, image, 10)
	if err != nil {
		return err
	}

	if len(annotations) == 0 {
		fmt.Fprintln(w, "No labels found.")
	} else {
		fmt.Fprintln(w, "Labels:")
		for _, annotation := range annotations {
			fmt.Fprintln(w, annotation.Description)
		}
	}

	return nil
}

// detectLandmarks gets landmarks from the Vision API for an image at the given file path.
func detectLandmarksURI(w io.Writer, file string) error {
	ctx := context.Background()

	client, err := vision.NewClient(ctx)
	if err != nil {
		return err
	}

	image := vision.NewImageFromURI(file)
	annotations, err := client.DetectLandmarks(ctx, image, 10)
	if err != nil {
		return err
	}

	if len(annotations) == 0 {
		fmt.Fprintln(w, "No landmarks found.")
	} else {
		fmt.Fprintln(w, "Landmarks:")
		for _, annotation := range annotations {
			fmt.Fprintln(w, annotation.Description)
		}
	}

	return nil
}

// detectText gets text from the Vision API for an image at the given file path.
func detectTextURI(w io.Writer, file string) error {
	ctx := context.Background()

	client, err := vision.NewClient(ctx)
	if err != nil {
		return err
	}

	image := vision.NewImageFromURI(file)
	annotations, err := client.DetectTexts(ctx, image, 10)
	if err != nil {
		return err
	}

	if len(annotations) == 0 {
		fmt.Fprintln(w, "No text found.")
	} else {
		fmt.Fprintln(w, "Text:")
		for _, annotation := range annotations {
			fmt.Fprintf(w, "%q\n", annotation.Description)
		}
	}

	return nil
}

// detectDocumentText gets the full document text from the Vision API for an image at the given file path.
func detectDocumentTextURI(w io.Writer, file string) error {
	ctx := context.Background()

	client, err := vision.NewClient(ctx)
	if err != nil {
		return err
	}

	image := vision.NewImageFromURI(file)
	annotation, err := client.DetectDocumentText(ctx, image)
	if err != nil {
		return err
	}

	fmt.Fprintln(w, "Text:")
	fmt.Fprintf(w, "%q\n", annotation.Text)

	return nil
}

// detectProperties gets image properties from the Vision API for an image at the given file path.
func detectPropertiesURI(w io.Writer, file string) error {
	ctx := context.Background()

	client, err := vision.NewClient(ctx)
	if err != nil {
		return err
	}

	image := vision.NewImageFromURI(file)
	props, err := client.DetectImageProps(ctx, image)
	if err != nil {
		return err
	}

	fmt.Fprintln(w, "Dominant colors:")
	for _, quantized := range props.DominantColors {
		color := quantized.Color
		fmt.Fprintf(w, "%2.1f%% - #%02x%02x%02x\n", quantized.PixelFraction*100, color.R&0xff, color.G&0xff, color.B&0xff)
	}

	return nil
}

// detectCropHints gets suggested croppings the Vision API for an image at the given file path.
func detectCropHintsURI(w io.Writer, file string) error {
	ctx := context.Background()

	client, err := vision.NewClient(ctx)
	if err != nil {
		return err
	}

	image := vision.NewImageFromURI(file)
	hints, err := client.CropHints(ctx, image, nil)
	if err != nil {
		return err
	}

	fmt.Fprintln(w, "Crop hints:")
	for _, hint := range hints {
		fmt.Fprintf(w, "%v\n", hint.BoundingPoly)
	}

	return nil
}

// detectSafeSearch gets image properties from the Vision API for an image at the given file path.
func detectSafeSearchURI(w io.Writer, file string) error {
	ctx := context.Background()

	client, err := vision.NewClient(ctx)
	if err != nil {
		return err
	}

	image := vision.NewImageFromURI(file)
	props, err := client.DetectSafeSearch(ctx, image)
	if err != nil {
		return err
	}

	fmt.Fprintln(w, "Safe Search properties:")
	fmt.Fprintln(w, "Adult:", props.Adult)
	fmt.Fprintln(w, "Medical:", props.Medical)
	fmt.Fprintln(w, "Spoofed:", props.Spoof)
	fmt.Fprintln(w, "Violence:", props.Violence)

	return nil
}

// detectWeb gets image properties from the Vision API for an image at the given file path.
func detectWebURI(w io.Writer, file string) error {
	ctx := context.Background()

	client, err := vision.NewClient(ctx)
	if err != nil {
		return err
	}

	image := vision.NewImageFromURI(file)
	web, err := client.DetectWeb(ctx, image)
	if err != nil {
		return err
	}

	fmt.Fprintln(w, "Web properties:")
	if len(web.FullMatchingImages) != 0 {
		fmt.Fprintln(w, "\tFull image matches:")
		for _, full := range web.FullMatchingImages {
			fmt.Fprintf(w, "\t\t%s\n", full.URL)
		}
	}
	if len(web.PagesWithMatchingImages) != 0 {
		fmt.Fprintln(w, "\tPages with this image:")
		for _, page := range web.PagesWithMatchingImages {
			fmt.Fprintf(w, "\t\t%s\n", page.URL)
		}
	}
	if len(web.WebEntities) != 0 {
		fmt.Fprintln(w, "\tEntities:")
		for _, entity := range web.WebEntities {
			fmt.Fprintf(w, "\t\t%-12s %s\n", entity.ID, entity.Description)
		}
	}

	return nil
}

// detectLogos gets logos from the Vision API for an image at the given file path.
func detectLogosURI(w io.Writer, file string) error {
	ctx := context.Background()

	client, err := vision.NewClient(ctx)
	if err != nil {
		return err
	}

	image := vision.NewImageFromURI(file)
	annotations, err := client.DetectLogos(ctx, image, 10)
	if err != nil {
		return err
	}

	if len(annotations) == 0 {
		fmt.Fprintln(w, "No logos found.")
	} else {
		fmt.Fprintln(w, "Logos:")
		for _, annotation := range annotations {
			fmt.Fprintln(w, annotation.Description)
		}
	}

	return nil
}
