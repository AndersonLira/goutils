package str

import "testing"

func TestSimilarity(t *testing.T) {

	per := Similarity("Anderson", "anderson")

	if 100 != int(per) {
		t.Errorf("Similarity should be 100, but %.2f", per)
	}

	per = Similarity("andré", "andre")

	if 80 > int(per) {
		t.Errorf("Similarity should be greater than 80, but %.2f", per)
	}

}

func TestSimilarityNormalized(t *testing.T) {

	per := SimilarityNormalized("andré", "andre")

	if 100 != int(per) {
		t.Errorf("Similarity should be 100, but %.2f", per)
	}

}
