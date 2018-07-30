package str

import (
	"testing"

	"github.com/andersonlira/goutils/str"
)

func TestSimilarity(t *testing.T) {

	per := str.Similarity("Anderson", "anderson")

	if 100 != int(per) {
		t.Errorf("Similarity should be 100, but %.2f", per)
	}

	per = str.Similarity("andré", "andre")

	if 80 > int(per) {
		t.Errorf("Similarity should be greater than 80, but %.2f", per)
	}

}

func TestSimilarityNormalized(t *testing.T) {

	per := str.SimilarityNormalized("andré", "andre")

	if 100 != int(per) {
		t.Errorf("Similarity should be 100, but %.2f", per)
	}

}

func TestSimilarityTolerance(t *testing.T) {

	per := str.Similarity("es", "espanhola, lingua")

	if 75 < int(per) {
		t.Errorf("Similarity should less than 75, but %.2f", per)
	}

}
