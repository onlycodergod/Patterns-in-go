package main

import (
	"os"
	"testing"
)

func TestSortUtil(t *testing.T) {
	testTable := []struct {
		sc          SortConfig
		name        string
		out         string
		haveError   bool
		errorString string
	}{
		{
			name: "sort without parameters",
			sc: SortConfig{
				filename: "testing/1_sort.txt",
			},
			out: "1\n1\n1\n55\n56\n6\n6\n7\n7\n8\n8\n80",
		},
		{
			name: "reverse sort",
			sc: SortConfig{
				filename:    "1_sort.txt",
				reverseSort: true,
			},
			out: "80\n8\n8\n7\n7\n6\n6\n56\n55\n1\n1\n1",
		},
		{
			name: "check, if already sorted -c",
			sc: SortConfig{
				isAlreadySorted: true,
				filename:        "testing/1_sort.txt",
			},
			out: "false",
		},
		{
			name: "check, if already sorted -c",
			sc: SortConfig{
				isAlreadySorted: true,
				filename:        "testing/2_sort.txt",
			},
			out: "true",
		},
		{
			name: "error: file not found",
			sc: SortConfig{
				filename: "testing/100_sort.txt",
			},
			haveError:   true,
			errorString: "can not read file 'testing/100_sort.txt': open testing/100_sort.txt: no such file or directory",
		},
	}

	for _, testingCase := range testTable {
		t.Run(testingCase.name, func(t *testing.T) {
			result, err := Start(&testingCase.sc)
			if !testingCase.haveError {
				if err != nil {
					t.Errorf("expected err == nil; got '%s'", err.Error())
				}

				if result != testingCase.out {
					t.Errorf("expected result \n'%s';\n\ngot\n'%s'", testingCase.out, result)
				}
			} else {
				if err != nil {
					if err.Error() != testingCase.errorString {
						t.Errorf("expected err.Error() == '%s'; got '%s'", testingCase.errorString, err.Error())
					}
				} else {
					t.Errorf("expected err != nil, but errs is nil")
				}
			}
		})
	}

}

// exists returns whether the given file or directory exists
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
