package progress_bar

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProgressBarCreation(t *testing.T) {
	progressBar := Create(100)
	progressBar = progressBar.Fill(10, 4)

	assert.Equal(t, Size(100), progressBar.Size, "Size of the progress bar should be 100")
	assert.True(t, progressBar.Size.Equal(100), "Size equal should work")
	assert.Equal(t, 2, len(progressBar.Fills), "We should have 2 fills")
	assert.Equal(t, Fill(10), progressBar.Fills[0], "First fill should be 10")
	assert.True(t, progressBar.Fills[0].Equal(10), "Fills equal should work")
	assert.Equal(t, Fill(4), progressBar.Fills[1], "Second fill should be 4")
}
