package id3conv

import (
	"log"

	"github.com/bogem/id3v2"
)

func fieldConvert(str string) (b []byte, err error) {
	latin, err := Utf8ToLatin([]byte(str))
	if err != nil {
		return nil, err
	}
	utf8, err := GbkToUtf8(latin)
	if err != nil {
		return nil, err
	}
	return utf8, err
}

// Convert convert the tag of the given music file to UTF8 encoding
func Convert(filename string) {
	tag, err := id3v2.Open(filename, id3v2.Options{Parse: true})
	if err != nil {
		log.Fatal(err)
	}
	tag.SetDefaultEncoding(id3v2.EncodingUTF8) // destination encoding
	for key, f := range tag.AllFrames() {
		if len(f) == 0 {
			// empty frame
			continue
		}
		if tf, ok := f[0].(id3v2.TextFrame); ok {
			// only deal with TextFrame
			if !tf.Encoding.Equals(id3v2.EncodingISO) {
				// only deal with wrong latin-1 aka ISO-8859-1 encoding
				continue
			}
			if b, err := fieldConvert(tf.Text); err == nil {
				converted := string(b)
				log.Printf("convert key (%s) from \"%s\" to \"%s\"", key, tf.Text, converted)
				tag.AddTextFrame(tag.CommonID(key), tag.DefaultEncoding(), converted)
			} else {
				log.Printf("warn: failed to convert key (%s) - \"%s\" with err \"%v\"", key, tf.Text, err)
			}
		}
	}
	if err := tag.Save(); err != nil {
		log.Printf("warn: failed to save to %s with err \"%v\"", filename, err)
	}
}
