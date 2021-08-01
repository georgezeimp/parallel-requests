package processor

import (
	"sort"
	"testing"

	"github.com/georgezeimp/parallel-requests/hasher"
	"github.com/georgezeimp/parallel-requests/output"
	"github.com/georgezeimp/parallel-requests/request"
)

func TestProcess(t *testing.T) {
	tests := []struct {
		addresses []string
		npr       int
		want      []string
	}{
		{
			addresses: []string{"https://github.com", "http://google.com"},
			npr:       1,
			want:      []string{"http://google.com 995e88839234099ab6e045a418fcf746", "https://github.com 01569ffdd33a067b01576a4ab500dece"},
		},
		{
			addresses: []string{"https://github.com", "http://google.com"},
			npr:       2,
			want:      []string{"http://google.com 995e88839234099ab6e045a418fcf746", "https://github.com 01569ffdd33a067b01576a4ab500dece"},
		},
		{
			addresses: []string{"https://github.com", "http://google.com"},
			npr:       10,
			want:      []string{"http://google.com 995e88839234099ab6e045a418fcf746", "https://github.com 01569ffdd33a067b01576a4ab500dece"},
		},
		{
			addresses: []string{
				"http://google.com",
				"https://github.com",
				"https://stackoverflow.com",
				"https://www.linkedin.com",
				"https://www.facebook.com",
				"https://www.the-scorpions.com",
				"https://www.nightwish.com",
				"https://twitter.com",
				"http://www.astronio.gr",
				"https://felifromgermany.com",
				"https://www.linkedin.com",
			},
			npr: 5,
			want: []string{
				"http://google.com 995e88839234099ab6e045a418fcf746",
				"https://github.com 01569ffdd33a067b01576a4ab500dece",
				"https://stackoverflow.com 615f5caf15003dc248d62e9fb480aae0",
				"https://www.linkedin.com 62838fd5ca73115a115ba539bd0e6266",
				"https://www.facebook.com e29eefd027231e4744de765ea6f23143",
				"https://www.the-scorpions.com 5e378ba569625f36af93c91984f7d74f",
				"https://www.nightwish.com ab6c66c504ee81bd3187ef032bd203b7",
				"https://twitter.com ddce4de336bae7e2616a98f4d7378917",
				"http://www.astronio.gr 27ea27ba6fde21a58e65ff96b17a3f73",
				"https://felifromgermany.com 7f7efb402ca75b29fe85ad85884629b0",
				"https://www.linkedin.com 62838fd5ca73115a115ba539bd0e6266",
			},
		},
	}

	h := hasher.NewHasher()
	op := output.NewPresenter()
	rs := request.NewServiceMock()
	processor := NewProcessor(h, op, rs)

	for _, tt := range tests {
		got := processor.Process(tt.addresses, tt.npr)

		sort.Strings(got)
		sort.Strings(tt.want)

		equal := true
		if len(tt.want) != len(got) {
			equal = false
		}
		for i := range tt.want {
			if tt.want[i] != got[i] {
				equal = false
			}
		}

		if !equal {
			t.Errorf("Test failed for Input: %+v, Expected Output: %+v, Actual Output: %+v", tt.addresses, tt.want, got)
		}
	}
}
