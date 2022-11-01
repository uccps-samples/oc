package catalog

import (
	"fmt"
	"hash/fnv"
	"strings"

	"k8s.io/apimachinery/pkg/util/errors"

	"github.com/uccps-samples/oc/pkg/cli/image/imagesource"
)

type Mirrorer interface {
	Mirror() (map[imagesource.TypedImageReference]imagesource.TypedImageReference, error)
}

// IndexExtractor knows how to pull an index image and extract its index file(s)
type IndexExtractor interface {
	Extract(from imagesource.TypedImageReference) (string, error)
}

type IndexExtractorFunc func(from imagesource.TypedImageReference) (string, error)

func (f IndexExtractorFunc) Extract(from imagesource.TypedImageReference) (string, error) {
	return f(from)
}

type RelatedImagesParser interface {
	Parse(path string) (map[string]struct{}, error)
}

type RelatedImagesParserFunc func(path string) (map[string]struct{}, error)

func (f RelatedImagesParserFunc) Parse(path string) (map[string]struct{}, error) {
	return f(path)
}

// ImageMirrorer knows how to mirror an image from one registry to another
type ImageMirrorer interface {
	Mirror(mapping map[imagesource.TypedImageReference]imagesource.TypedImageReference) error
}

type ImageMirrorerFunc func(mapping map[imagesource.TypedImageReference]imagesource.TypedImageReference) error

func (f ImageMirrorerFunc) Mirror(mapping map[imagesource.TypedImageReference]imagesource.TypedImageReference) error {
	return f(mapping)
}

type IndexImageMirrorer struct {
	ImageMirrorer       ImageMirrorer
	IndexExtractor      IndexExtractor
	RelatedImagesParser RelatedImagesParser

	// options
	Source, Dest      imagesource.TypedImageReference
	MaxPathComponents int
}

var _ Mirrorer = &IndexImageMirrorer{}

func NewIndexImageMirror(options ...ImageIndexMirrorOption) (*IndexImageMirrorer, error) {
	config := DefaultImageIndexMirrorerOptions()
	config.Apply(options)
	if err := config.Complete(); err != nil {
		return nil, err
	}
	if err := config.Validate(); err != nil {
		return nil, err
	}
	return &IndexImageMirrorer{
		ImageMirrorer:       config.ImageMirrorer,
		IndexExtractor:      config.IndexExtractor,
		RelatedImagesParser: config.RelatedImagesParser,
		Source:              config.Source,
		Dest:                config.Dest,
		MaxPathComponents:   config.MaxPathComponents,
	}, nil
}

func (b *IndexImageMirrorer) Mirror() (map[imagesource.TypedImageReference]imagesource.TypedImageReference, error) {
	catalogPath, err := b.IndexExtractor.Extract(b.Source)
	if err != nil {
		return nil, fmt.Errorf("extract catalog files: %v", err)
	}

	images, err := b.RelatedImagesParser.Parse(catalogPath)
	if err != nil {
		return nil, fmt.Errorf("parse related images: %v", err)
	}

	var errs = make([]error, 0)
	mapping, mapErrs := mappingForImages(images, b.Source, b.Dest, b.MaxPathComponents)
	if len(mapErrs) > 0 {
		errs = append(errs, mapErrs...)
	}

	mappedIndex, err := mount(b.Source, b.Dest, b.MaxPathComponents)
	if err != nil {
		errs = append(errs, fmt.Errorf("unable to map index image to new location in dest"))
	}
	mapping[b.Source] = mappedIndex

	if err := b.ImageMirrorer.Mirror(mapping); err != nil {
		errs = append(errs, fmt.Errorf("mirroring failed: %s", err.Error()))
	}

	return mapping, errors.NewAggregate(errs)
}

func mappingForImages(images map[string]struct{}, src, dest imagesource.TypedImageReference, maxComponents int) (mapping map[imagesource.TypedImageReference]imagesource.TypedImageReference, errs []error) {
	if dest.Type != imagesource.DestinationRegistry {
		// don't do any name mangling when not mirroring to a real registry
		// this allows us to assume the names are preserved when doing multi-hop mirrors that use a file or s3 as an
		// intermediate step
		maxComponents = 0

		// if mirroring a source (like quay.io/my/index:1) to a file location like file://local/store
		// we will remount all of the content in the file store under the catalog name
		// i.e. file://local/store/my/index
		var err error
		dest, err = mount(src, dest, 0)
		if err != nil {
			errs = []error{err}
			return
		}
	}

	mapping = map[imagesource.TypedImageReference]imagesource.TypedImageReference{}
	for img := range images {
		if img == "" {
			continue
		}

		parsed, err := imagesource.ParseReference(img)
		if err != nil {
			errs = append(errs, fmt.Errorf("couldn't parse image for mirroring (%s), skipping mirror: %v", img, err))
			continue
		}

		targetRef, err := mount(parsed, dest, maxComponents)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		// set docker defaults, but don't set default tag for digest refs
		s := parsed
		parsed.Ref = parsed.Ref.DockerClientDefaults()
		if len(s.Ref.Tag) == 0 && len(s.Ref.ID) > 0 {
			parsed.Ref.Tag = ""
		}

		// if src is a file store, assume all other references are in the same location on disk
		if src.Type != imagesource.DestinationRegistry {
			srcRef, err := mount(parsed, src, 0)
			if err != nil {
				errs = append(errs, err)
				continue
			}
			if len(parsed.Ref.Tag) == 0 {
				srcRef.Ref.Tag = ""
			}
			mapping[srcRef] = targetRef
			continue
		}

		mapping[parsed] = targetRef
	}
	return
}

func mount(in, dest imagesource.TypedImageReference, maxComponents int) (out imagesource.TypedImageReference, err error) {
	out = in
	out.Type = dest.Type

	hasher := fnv.New32a()
	// tag with hash of source ref if no tag given
	if len(out.Ref.Tag) == 0 && len(out.Ref.ID) > 0 {
		hasher.Reset()
		_, err = hasher.Write([]byte(out.Ref.String()))
		if err != nil {
			err = fmt.Errorf("couldn't generate tag for image (%s), skipping mirror", in.String())
		}
		out.Ref.Tag = fmt.Sprintf("%x", hasher.Sum32())
	}

	// fill in default registry / tag if missing
	out.Ref = out.Ref.DockerClientDefaults()

	components := []string{}
	if len(dest.Ref.Namespace) > 0 {
		components = append(components, dest.Ref.Namespace)
	}
	if len(dest.Ref.Name) > 0 {
		components = append(components, strings.Split(dest.Ref.Name, "/")...)
	}
	if len(out.Ref.Namespace) > 0 {
		components = append(components, out.Ref.Namespace)
	}
	if len(out.Ref.Name) > 0 {
		components = append(components, strings.Split(out.Ref.Name, "/")...)
	}

	out.Ref.Registry = dest.Ref.Registry
	out.Ref.Namespace = components[0]
	if maxComponents > 1 && len(components) > maxComponents {
		out.Ref.Name = strings.Join(components[1:maxComponents-1], "/") + "/" + strings.Join(components[maxComponents-1:], "-")
	} else if maxComponents == 0 {
		out.Ref.Name = strings.Join(components[1:], "/")
	} else if len(components) > 1 {
		out.Ref.Name = strings.Join(components[1:maxComponents], "/")
	} else {
		// only one component, make it the name, not the namespace
		out.Ref.Name = in.Ref.Name
		out.Ref.Namespace = ""
	}
	out.Ref.Name = strings.TrimPrefix(out.Ref.Name, "/")
	return
}
