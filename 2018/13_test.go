package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTrack(t *testing.T) {
	testCases := []struct {
		In string
	}{
		{`
/-----\   
|     |   
|  /--+--\
|  |  |  |
\--+--/  |
   |     |
   \-----/`,
		},
		{`
/->-\        
|   |  /----\
| /-+--+-\  |
| | |  | v  |
\-+-/  \-+--/
  \------/   `,
		},
	}

	for i, c := range testCases {
		tr := NewTracks(c.In)
		assert.Equal(t, c.In, tr.String(false), "(case=%d)", i)
	}
}

func TestRunUntilCrash(t *testing.T) {
	testCases := []struct {
		In               string
		OutX, OutY, OutI int
	}{
		{`
/->-\        
|   |  /----\
| /-+--+-\  |
| | |  | v  |
\-+-/  \-+--/
  \------/   `, 7, 3, 14,
		},
	}

	for i, c := range testCases {
		tr := NewTracks(c.In)
		x, y, n := tr.RunUntilCrash()
		assert.Equal(t, fmt.Sprintf("%d,%d (%d)", c.OutX, c.OutY, c.OutI), fmt.Sprintf("%d,%d (%d)", x, y, n), "(case=%d)", i)
	}
}

func TestRunUntilOneCartLeft(t *testing.T) {
	testCases := []struct {
		In               string
		OutX, OutY, OutI int
	}{
		{`/>-<\  
|   |  
| /<+-\
| | | v
\>+</ |
  |   ^
  \<->/`, 6, 4, 3,
		},
	}

	for i, c := range testCases {
		tr := NewTracks(c.In)
		x, y, n := tr.RunUntilOneCartLeft()
		assert.Equal(t, fmt.Sprintf("%d,%d (%d)", c.OutX, c.OutY, c.OutI), fmt.Sprintf("%d,%d (%d)", x, y, n), "(case=%d)", i)
	}
}
