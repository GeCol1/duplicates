package core

import (
	"os"
	"path/filepath"
	"testing"
)

func TestComputeHashOfFiles(t *testing.T) {
	var tests = []struct {
		desc string
		path string
		want map[string]string
	}{
		{desc: "Compute hash for tests/folder1",
			path: "../tests/folder1",
			want: map[string]string{
				"60b36fdde2c951af3677fad59c0dd5251450ca53": "/home/djay/Work/01_PROJECTS/42_GO/duplicates/tests/folder1/subFolder1/dave-hoefler-Nkwrpa_v0WA-unsplash.jpg",
				"990b299b95bd5e1f3a24857d3d8273d18cae6418": "/home/djay/Work/01_PROJECTS/42_GO/duplicates/tests/folder1/subFolder1/alex-houque-CC_ZnZJdM4E-unsplash.jpg",
				"eae82804dc4ed05103c260f6eefe4800cdd23dc3": "/home/djay/Work/01_PROJECTS/42_GO/duplicates/tests/folder1/aaron-burden-UyNrNfdKjwg-unsplash.jpg",
				"8041404f5e19c4620e4341274a7dcb4ee44be5b1": "/home/djay/Work/01_PROJECTS/42_GO/duplicates/tests/folder1/filipp-romanovski-EsCDHtM7K5s-unsplash.jpg",
				"0553c08570cad91dafb8254dca08d36b3a907d19": "/home/djay/Work/01_PROJECTS/42_GO/duplicates/tests/folder1/jason-hudson-y_83VDSKzyg-unsplash.jpg",
				"23ee8d6d932b22c72ef410866e796ae518e4640d": "/home/djay/Work/01_PROJECTS/42_GO/duplicates/tests/folder1/nguyen-minh-iUvxXOCKUMg-unsplash.jpg",
				"1ab96e9fd2878823877012204c24cc68b89f29d5": "/home/djay/Work/01_PROJECTS/42_GO/duplicates/tests/folder1/patrick-federi-dAOYYI4LIeE-unsplash.jpg",
				"a721a51784aba76a79b043f77c27ffd6a4f1abf7": "/home/djay/Work/01_PROJECTS/42_GO/duplicates/tests/folder1/gustavo-zambelli-MzlLPB18EFc-unsplash.jpg",
				"deccb9399893bb96fb7c7912a579b3f4d1bf4a9e": "/home/djay/Work/01_PROJECTS/42_GO/duplicates/tests/folder1/radek-kilijanek-ID3YnmQ0ia8-unsplash.jpg",
			},
		},
	}
	for _, test := range tests {
		wd, _ := os.Getwd()
		got, err := ComputeHashOfFiles(filepath.Join(wd, test.path))
		if err != nil {
			t.Errorf("error in test %s : %s", test.desc, err)
		}
		for key, val := range got {
			if expected, ok := test.want[key]; !ok || val != expected {
				t.Errorf("key not found in expected %s or val diff %s", key, val)
			}

		}
	}
}
