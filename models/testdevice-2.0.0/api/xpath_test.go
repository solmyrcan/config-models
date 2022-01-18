// Copyright 2021-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

import (
	"fmt"
	"github.com/antchfx/xpath"
	"github.com/onosproject/config-models/pkg/xpath/navigator"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

// Test_XPathSelect See https://devhints.io/xpath for an XPath Cheat sheet
func Test_XPathSelect(t *testing.T) {
	sampleConfig, err := ioutil.ReadFile("../testdata/sample-testdevice2-config.json")
	if err != nil {
		assert.NoError(t, err)
	}
	device := new(Device)

	schema, err := Schema()
	if err := schema.Unmarshal(sampleConfig, device); err != nil {
		assert.NoError(t, err)
	}
	schema.Root = device
	assert.NotNil(t, device)
	ynn := navigator.NewYangNodeNavigator(schema.RootSchema(), device)
	assert.NotNil(t, ynn)

	tests := []navigator.XpathSelect{
		{
			Name: "test leaf2a",
			Path: "/t1:cont1a/t1:cont2a/t1:leaf2a",
			Expected: []string{
				"Iter Value: leaf2a: 1",
			},
		},
		{
			Name: "test leaf2b",
			Path: "/t1:cont1a/t1:cont2a/t1:leaf2b",
			Expected: []string{
				"Iter Value: leaf2b: 0.4321",
			},
		},
		{
			Name:     "test leaf2c",
			Path:     "/t1:cont1a/t1:cont2a/t1:leaf2c",
			Expected: []string{}, // No value, so no response
		},
		{
			Name: "test leaf2g",
			Path: "/t1:cont1a/t1:cont2a/t1:leaf2g",
			Expected: []string{
				"Iter Value: leaf2g: true",
			},
		},
		{
			Name: "test descendant-or-self", // `//` is short for `descendant-or-self::` axis
			Path: "/t1:cont1a/descendant-or-self::t1:leaf2g",
			Expected: []string{
				"Iter Value: leaf2g: true",
			},
		},
		{
			Name: "test leaf2g ancestor(s)",
			Path: "/t1:cont1a/t1:cont2a/t1:leaf2g/ancestor::node()",
			Expected: []string{
				"Iter Value: cont2a: value of cont2a",
				"Iter Value: cont1a: value of cont1a",
				"Iter Value: device: value of device",
			},
		},
		{
			Name: "test leaf2g parent",
			Path: "/t1:cont1a/t1:cont2a/t1:leaf2g/parent::node()",
			Expected: []string{
				"Iter Value: cont2a: value of cont2a",
			},
		},
		{
			Name: "test leaf2g self",
			Path: "/t1:cont1a/t1:cont2a/t1:leaf2g/self::node()",
			Expected: []string{
				"Iter Value: leaf2g: true",
			},
		},
		{
			Name: "test leaf2g child",
			Path: "/t1:cont1a/t1:cont2a/child::node()",
			Expected: []string{
				"Iter Value: leaf2a: 1",
				"Iter Value: leaf2b: 0.4321",
				"Iter Value: leaf2e: [5 4 3 2 1]",
				"Iter Value: leaf2f: dGhpcyBpcyBhIHRlc3QgdGVzdAo=",
				"Iter Value: leaf2g: true",
			},
		},
		{
			Name: "test leaf2g preceding",
			Path: "/t1:cont1a/t1:cont2a/t1:leaf2g/preceding::node()",
			Expected: []string{
				"Iter Value: leaf2f: dGhpcyBpcyBhIHRlc3QgdGVzdAo=",
				"Iter Value: leaf2e: [5 4 3 2 1]",
				"Iter Value: leaf2b: 0.4321",
				"Iter Value: leaf2a: 1",
			},
		},
		{
			Name: "test leaf2g following-sibling",
			Path: "/t1:cont1a/t1:cont2a/t1:leaf2e/following-sibling::node()",
			Expected: []string{
				"Iter Value: leaf2f: dGhpcyBpcyBhIHRlc3QgdGVzdAo=",
				"Iter Value: leaf2g: true",
			},
		},
		{
			Name: "test leaf2g following", // Everything follow - all levels
			Path: "/t1:cont1a/t1:cont2a/t1:leaf2e/following::node()",
			Expected: []string{
				"Iter Value: leaf2f: dGhpcyBpcyBhIHRlc3QgdGVzdAo=",
				"Iter Value: leaf2g: true",
				"Iter Value: leaf1a: leaf1aval",
				"Iter Value: list2a: value of list2a",
				"Iter Value: name: l2a1",
				"Iter Value: rx-power: 25",
				"Iter Value: tx-power: 5",
				"Iter Value: list2a: value of list2a",
				"Iter Value: name: l2a2",
				"Iter Value: rx-power: 26",
				"Iter Value: tx-power: 6",
				"Iter Value: list2a: value of list2a",
				"Iter Value: name: l2a3",
				"Iter Value: rx-power: 27",
				"Iter Value: tx-power: 8",
				"Iter Value: cont1b-state: value of cont1b-state",
				"Iter Value: list2b: value of list2b",
				"Iter Value: index1: 10",
				"Iter Value: index2: 20",
				"Iter Value: leaf3c: 3c 10-20 test",
				"Iter Value: list2b: value of list2b",
				"Iter Value: index1: 11",
				"Iter Value: index2: 20",
				"Iter Value: leaf3c: 3c 11-20 test",
				"Iter Value: leaf3d: IDTYPE2",
				"Iter Value: list2b: value of list2b",
				"Iter Value: index1: 11",
				"Iter Value: index2: 21",
				"Iter Value: leaf3c: 3c 11-21 test",
				"Iter Value: leaf3d: IDTYPE1",
				"Iter Value: list2b: value of list2b",
				"Iter Value: index1: 12",
				"Iter Value: index2: 22",
				"Iter Value: leaf3c: 3c 12-22 test",
				"Iter Value: leaf3d: IDTYPE2",
			},
		},
		{
			Name: "test leaf2g descendant",
			Path: "/t1:cont1a/t1:cont2a/descendant::node()",
			Expected: []string{
				"Iter Value: leaf2a: 1",
				"Iter Value: leaf2b: 0.4321",
				"Iter Value: leaf2e: [5 4 3 2 1]",
				"Iter Value: leaf2f: dGhpcyBpcyBhIHRlc3QgdGVzdAo=",
				"Iter Value: leaf2g: true",
			},
		},
		{
			Name: "test leaf1a",
			Path: "/t1:cont1a/t1:leaf1a",
			Expected: []string{
				"Iter Value: leaf1a: leaf1aval",
			},
		},
		{
			Name: "test list2a all instances names",
			Path: "/t1:cont1a/t1:list2a/@t1:name", // List indices are always attributes - referred to with @
			Expected: []string{
				"Iter Value: name: l2a1",
				"Iter Value: name: l2a2",
				"Iter Value: name: l2a3",
			},
		},
		{
			Name: "test list2a select 2nd instance rx-power",
			Path: "/t1:cont1a/t1:list2a[@t1:name='l2a2']/t1:rx-power", // select with []
			Expected: []string{
				"Iter Value: rx-power: 26",
			},
		},
		{
			Name: "test list2a select 1st instance tx-power",
			Path: "/t1:cont1a/t1:list2a[@t1:name='l2a1']/t1:tx-power", // select with []
			Expected: []string{
				"Iter Value: tx-power: 5",
			},
		},
		{
			Name: "test list2a select instance where tx-power < 10",
			Path: "/t1:cont1a/t1:list2a[t1:tx-power < 8]/@t1:name", // select with []
			Expected: []string{
				"Iter Value: name: l2a1",
				"Iter Value: name: l2a2",
			},
		},
		{
			Name: "test list2a select instance where tx-power < rx-power",
			Path: "/t1:cont1a/t1:list2a[number(t1:tx-power) < number(t1:rx-power) -19]/@t1:name", // select with []
			Expected: []string{
				"Iter Value: name: l2a1",
				"Iter Value: name: l2a2",
			},
		},
		{
			Name: "test index 1 filter",
			Path: "/t1:cont1b-state/t1:list2b[@t1:index1=11 and @t1:index2=20]/t1:leaf3c",
			Expected: []string{
				"Iter Value: leaf3c: 3c 11-20 test",
			},
		},
		{
			Name: "test index 2 filter",
			Path: "/t1:cont1b-state/t1:list2b[@t1:index2=20]/@t1:index1",
			Expected: []string{
				"Iter Value: index1: 10",
				"Iter Value: index1: 11",
			},
		},
	}

	for _, test := range tests {
		expr, err := xpath.Compile(test.Path)
		assert.NoError(t, err, test.Name)
		assert.NotNil(t, expr, test.Name)

		iter := expr.Select(ynn)
		resultCount := 0
		for iter.MoveNext() {
			assert.LessOrEqual(t, resultCount, len(test.Expected)-1, test.Name, ". More results than expected")
			assert.Equal(t, test.Expected[resultCount], fmt.Sprintf("Iter Value: %s: %s",
				iter.Current().LocalName(), iter.Current().Value()), test.Name)
			resultCount++
		}
		assert.Equal(t, len(test.Expected), resultCount, "%s. Did not receive all the expected results", test.Name)
	}
}

func Test_XPathEvaluate(t *testing.T) {
	sampleConfig, err := ioutil.ReadFile("../testdata/sample-testdevice2-config.json")
	if err != nil {
		assert.NoError(t, err)
	}
	device := new(Device)

	schema, err := Schema()
	if err := schema.Unmarshal(sampleConfig, device); err != nil {
		assert.NoError(t, err)
	}
	schema.Root = device
	assert.NotNil(t, device)
	ynn := navigator.NewYangNodeNavigator(schema.RootSchema(), device)
	assert.NotNil(t, ynn)

	tests := []navigator.XpathEvaluate{
		// For leaf2a
		{
			Name:     "test check present",
			Path:     "count(/t1:cont1a/t1:cont2a[t1:leaf2a])",
			Expected: float64(1),
		},
		{
			Name:     "test equals leaf2a",
			Path:     "/t1:cont1a/t1:cont2a/t1:leaf2a = 1",
			Expected: true,
		},
		{
			Name:     "test not equals leaf2a",
			Path:     "not(/t1:cont1a/t1:cont2a/t1:leaf2a = 1)",
			Expected: false,
		},
		{
			Name:     "test number value",
			Path:     "number(/t1:cont1a/t1:cont2a/t1:leaf2a)",
			Expected: float64(1),
		},
		{
			Name:     "test string value",
			Path:     "string(/t1:cont1a/t1:cont2a/t1:leaf2a)",
			Expected: "1",
		},
		{
			Name:     "test boolean value",
			Path:     "boolean(/t1:cont1a/t1:cont2a/t1:leaf2a)",
			Expected: true,
		},
		{
			Name:     "test name()",
			Path:     "name(/t1:cont1a/t1:cont2a/t1:leaf2a)",
			Expected: "t1:leaf2a",
		},
		{
			Name:     "test concat value with name",
			Path:     "concat(/t1:cont1a/t1:cont2a/t1:leaf2a, name(/t1:cont1a/t1:cont2a/t1:leaf2a))",
			Expected: "1t1:leaf2a",
		},
		{
			Name:     "test substring of name",
			Path:     "substring(name(/t1:cont1a/t1:cont2a/t1:leaf2a), 4, 5)",
			Expected: "leaf2",
		},
		{
			Name:     "test substring-before :",
			Path:     "substring-before(name(/t1:cont1a/t1:cont2a/t1:leaf2a), ':')",
			Expected: "t1",
		},
		{
			Name:     "test substring after :",
			Path:     "substring-after(name(/t1:cont1a/t1:cont2a/t1:leaf2a), ':')",
			Expected: "leaf2a",
		},
		{
			Name:     "test string length",
			Path:     "string-length(name(/t1:cont1a/t1:cont2a/t1:leaf2a))",
			Expected: float64(9),
		},
		{
			Name:     "test translate",
			Path:     "translate(name(/t1:cont1a/t1:cont2a/t1:leaf2a), 't1','x1')",
			Expected: "x1:leaf2a",
		},
		{
			Name:     "test translate",
			Path:     "translate(name(/t1:cont1a/t1:cont2a/t1:leaf2a), 'tleaf','TLEAF')",
			Expected: "T1:LEAF2A",
		},
		{
			Name:     "test translate",
			Path:     "normalize-space('  This is a  test  ')",
			Expected: "This is a test",
		},
		{
			Name:     "test lt leaf2a",
			Path:     "/t1:cont1a/t1:cont2a/t1:leaf2a < 2",
			Expected: true,
		},
		{
			Name:     "test lte leaf2a",
			Path:     "/t1:cont1a/t1:cont2a/t1:leaf2a <= 1",
			Expected: true,
		},
		{
			Name:     "test gte leaf2a",
			Path:     "/t1:cont1a/t1:cont2a/t1:leaf2a >= 1",
			Expected: true,
		},
		{
			Name:     "test gt leaf2a",
			Path:     "/t1:cont1a/t1:cont2a/t1:leaf2a > 0",
			Expected: true,
		},
		{
			Name:     "test gt 1 false leaf2a",
			Path:     "/t1:cont1a/t1:cont2a/t1:leaf2a > 1",
			Expected: false,
		},
		{
			Name:     "test count leaf2a",
			Path:     "count(/t1:cont1a/t1:cont2a/t1:leaf2a)",
			Expected: float64(1),
		},
		{
			Name:     "test count leaf2a alternate syntax",
			Path:     "count(/t1:cont1a/t1:cont2a[t1:leaf2a])",
			Expected: float64(1),
		},
		{
			Name:     "test leaf2b",
			Path:     "number(/t1:cont1a/t1:cont2a/t1:leaf2b)",
			Expected: 0.4321,
		},
		//{ // product() not yet supported
		//	Name:     "test leaf2b product 10",
		//	Path:     "product(/t1:cont1a/t1:cont2a/t1:leaf2b, 10)",
		//	Expected: 4.321,
		//},
		{
			Name:     "test last leaf2g",
			Path:     "boolean(/t1:cont1a/t1:cont2a[last()])",
			Expected: true,
		},
		{
			Name:     "test leaf2g or leaf2a",
			Path:     "string(/t1:cont1a/t1:cont2a/t1:leaf2g) = 'true' or number(/t1:cont1a/t1:cont2a/t1:leaf2a) < 4",
			Expected: true,
		},
		{
			Name:     "test sum leaf2a",
			Path:     "sum(/t1:cont1a/t1:cont2a/t1:leaf2a)",
			Expected: float64(1),
		},
		// For leaf2b
		{
			Name:     "test lt leaf2b",
			Path:     "/t1:cont1a/t1:cont2a/t1:leaf2b < 0.5",
			Expected: true,
		},
		{
			Name:     "test eq leaf2b",
			Path:     "/t1:cont1a/t1:cont2a/t1:leaf2b = 0.4321",
			Expected: true,
		},
		{
			Name:     "test eq false leaf2b",
			Path:     "/t1:cont1a/t1:cont2a/t1:leaf2b = 0.432",
			Expected: false,
		},
		{
			Name:     "test gt leaf2b",
			Path:     "/t1:cont1a/t1:cont2a/t1:leaf2b > 0.4",
			Expected: true,
		},
		// For leaf2c - not present
		{
			Name:     "test eq leaf2c",
			Path:     "/t1:cont1a/t1:cont2a/t1:leaf2c = 1",
			Expected: false,
		},
		{
			Name:     "test count leaf2c", // Checking that node is not present
			Path:     "count(/t1:cont1a/t1:cont2a/t1:leaf2c)",
			Expected: float64(0),
		},
		{
			Name:     "test count leaf2c alt syntax", // Checking that node is not present
			Path:     "count(/t1:cont1a/t1:cont2a[t1:leaf2c])",
			Expected: float64(0),
		},
		// For Leaf1a
		{
			Name:     "test eq leaf1a",
			Path:     "/t1:cont1a/t1:leaf1a = 'leaf1aval'",
			Expected: true,
		},
		{
			Name:     "test neq leaf1a",
			Path:     "/t1:cont1a/t1:leaf1a != 'leaf1aval'",
			Expected: false,
		},
		{
			Name:     "test eq leaf1a",
			Path:     "/t1:cont1a/t1:leaf1a != 'leaf1avalWrong'",
			Expected: true,
		},
		{
			Name:     "test count children of cont2a star",
			Path:     "count(/t1:cont1a/*)",
			Expected: float64(5),
		},
		{
			Name:     "test count children of cont2a child",
			Path:     "count(/t1:cont1a/t1:cont2a/child::node())",
			Expected: float64(5),
		},
		// For List2a
		{
			Name:     "test count list2a",
			Path:     "count(/t1:cont1a/t1:list2a)",
			Expected: float64(3),
		},
		{
			Name:     "test first item list2a",
			Path:     "string(/t1:cont1a/t1:list2a[1]/@t1:name)",
			Expected: "l2a1",
		},
		{
			Name:     "test second item list2a",
			Path:     "string(/t1:cont1a/t1:list2a[2]/@t1:name)",
			Expected: "l2a2",
		},
		{
			Name:     "test last item list2a",
			Path:     "string(/t1:cont1a/t1:list2a[last()]/@t1:name)",
			Expected: "l2a3",
		},
		{
			Name:     "test sum list2a tx-power",
			Path:     "sum(/t1:cont1a/t1:list2a/t1:tx-power)",
			Expected: float64(19),
		},
		{
			Name:     "test sum list2a rx-power",
			Path:     "sum(/t1:cont1a/t1:list2a/t1:rx-power)",
			Expected: float64(78),
		},
		{
			Name:     "test sum list2a rx-power part a", // Selecting only 1
			Path:     "sum(/t1:cont1a/t1:list2a[@t1:name='l2a1']/t1:rx-power)",
			Expected: float64(25),
		},
		{
			Name:     "test sum list2a rx-power part a or b", // Selecting only 2
			Path:     "sum(/t1:cont1a/t1:list2a[@t1:name='l2a1' or @t1:name='l2a3']/t1:rx-power)",
			Expected: float64(52),
		},
		{
			Name:     "test sum list2a rx-power starts-with", // Selects all 3
			Path:     "sum(/t1:cont1a/t1:list2a[starts-with(@t1:name,'l2a')]/t1:rx-power)",
			Expected: float64(78),
		},
		{
			Name:     "test sum list2a rx-power contains", // Selects all 3
			Path:     "sum(/t1:cont1a/t1:list2a[contains(@t1:name,'l2')]/t1:rx-power)",
			Expected: float64(78),
		},
		{
			Name:     "test sum list2a rx-power ends-with", // Selects all 3
			Path:     "sum(/t1:cont1a/t1:list2a[ends-with(@t1:name,'a2')]/t1:rx-power)",
			Expected: float64(26),
		},
		{
			Name:     "test position list2a part b only", // Selecting only 1
			Path:     "position(/t1:cont1a/t1:list2a[@t1:name='l2a2'])",
			Expected: float64(1),
		},
		//{ // max not supported
		//	Name:     "test max list2a rx-power",
		//	Path:     "max(/t1:cont1a/t1:list2a/t1:tx-power)",
		//	Expected: float64(28),
		//},
		{
			Name:     "test sum list2a tx-power < rx-power",
			Path:     "sum(/t1:cont1a/t1:list2a/t1:tx-power) < sum(/t1:cont1a/t1:list2a/t1:rx-power)",
			Expected: true,
		},
		{ // It will do a string comparison by default - at least 1 side has to be converted to number
			Name:     "test individual list2a tx-power < rx-power",
			Path:     "count(/t1:cont1a/t1:list2a[number(t1:tx-power) < number(t1:rx-power)])",
			Expected: float64(3),
		},
		// cont1b-state List 2b
		{
			Name:     "test count list2b",
			Path:     "count(/t1:cont1b-state/t1:list2b)",
			Expected: float64(4),
		},
		{
			Name:     "test count list2b list3c",
			Path:     "count(/t1:cont1b-state/t1:list2b/t1:leaf3c)",
			Expected: float64(4),
		},
		{
			Name:     "test count list2b filtered",
			Path:     "count(/t1:cont1b-state/t1:list2b[@t1:index2=20]/t1:leaf3c)",
			Expected: float64(2),
		},
		{
			Name:     "test count list2b part1",
			Path:     "count(/t1:cont1b-state/t1:list2b[@t1:index1=10]/t1:leaf3c)",
			Expected: float64(1),
		},
		{
			Name:     "test count list2b part2 only", // Selecting 11 AND 20
			Path:     "string(/t1:cont1b-state/t1:list2b[@t1:index1=11 and @t1:index2=20]/t1:leaf3c)",
			Expected: "3c 11-20 test",
		},
		{
			Name:     "test count list2b part 2 different syntax", // Selecting 11 AND 20
			Path:     "count(/t1:cont1b-state/t1:list2b[@t1:index1=11][@t1:index2=20]/t1:leaf3c)",
			Expected: float64(1),
		},
		{
			Name:     "test value list2b part2 only", // Selecting 11 AND 20
			Path:     "/t1:cont1b-state/t1:list2b[@t1:index1=11 and @t1:index2=20]/t1:leaf3c = '3c 11-20 test'",
			Expected: true,
		},
		//{ // TODO figure out why this gives invalid token
		//	Name:     "test text list2b part2 only", // Selecting 11 AND 20
		//	Path:     "text(/t1:cont1b-state/t1:list2b/t1:leaf3c)",
		//	Expected: true,
		//},
		{
			Name:     "test count list2b part1",
			Path:     "/t1:cont1b-state/t1:list2b[@t1:index1=10]/t1:leaf3c = '3c 10-20 test'",
			Expected: true,
		},
	}

	for _, test := range tests {
		expr, testErr := xpath.Compile(test.Path)
		assert.NoError(t, testErr, test.Name)
		assert.NotNil(t, expr, test.Name)

		result := expr.Evaluate(ynn)
		assert.Equal(t, test.Expected, result, test.Name)
	}
}

func Test_XPathNodeNavigation(t *testing.T) {

	sampleConfig, err := ioutil.ReadFile("../testdata/sample-testdevice2-config.json")
	if err != nil {
		assert.NoError(t, err)
	}
	device := new(Device)

	schema, err := Schema()
	if err := schema.Unmarshal(sampleConfig, device); err != nil {
		assert.NoError(t, err)
	}
	schema.Root = device
	assert.NotNil(t, device)
	ynn := navigator.NewYangNodeNavigator(schema.RootSchema(), device)
	assert.NotNil(t, ynn)
	assert.Equal(t, xpath.ElementNode, ynn.NodeType())
	assert.Equal(t, "device", ynn.LocalName())
	assert.Equal(t, "", ynn.Prefix())
	assert.Equal(t, "value of device", ynn.Value())

	assert.True(t, ynn.MoveToChild())
	assert.Equal(t, "cont1a", ynn.LocalName())
	assert.Equal(t, xpath.ElementNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())

	assert.True(t, ynn.MoveToChild())
	assert.Equal(t, "cont2a", ynn.LocalName())
	assert.Equal(t, xpath.ElementNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())

	assert.True(t, ynn.MoveToChild())
	assert.Equal(t, "leaf2a", ynn.LocalName())
	assert.Equal(t, xpath.ElementNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())
	assert.Equal(t, "1", ynn.Value())

	assert.False(t, ynn.MoveToChild())

	assert.True(t, ynn.MoveToNext())
	assert.Equal(t, "leaf2b", ynn.LocalName())
	assert.Equal(t, xpath.ElementNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())
	assert.Equal(t, "0.4321", ynn.Value())

	// Skips leaf2c and leaf2d as they have no values
	assert.True(t, ynn.MoveToNext())
	assert.Equal(t, "leaf2e", ynn.LocalName())
	assert.Equal(t, xpath.ElementNode, ynn.NodeType()) // Leaf list
	assert.Equal(t, "t1", ynn.Prefix())
	assert.Equal(t, "[5 4 3 2 1]", ynn.Value())

	assert.True(t, ynn.MoveToNext())
	assert.Equal(t, "leaf2f", ynn.LocalName())
	assert.Equal(t, xpath.ElementNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())
	assert.Equal(t, "dGhpcyBpcyBhIHRlc3QgdGVzdAo=", ynn.Value())

	assert.True(t, ynn.MoveToNext())
	assert.Equal(t, "leaf2g", ynn.LocalName())
	assert.Equal(t, xpath.ElementNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())
	assert.Equal(t, "true", ynn.Value())

	// no next exists so returns false
	assert.False(t, ynn.MoveToNext())

	assert.True(t, ynn.MoveToPrevious())
	assert.Equal(t, "leaf2f", ynn.LocalName())
	assert.Equal(t, xpath.ElementNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())
	assert.Equal(t, "dGhpcyBpcyBhIHRlc3QgdGVzdAo=", ynn.Value())

	assert.True(t, ynn.MoveToFirst())
	assert.Equal(t, "leaf2a", ynn.LocalName())
	assert.Equal(t, xpath.ElementNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())
	assert.Equal(t, "1", ynn.Value())

	assert.True(t, ynn.MoveToParent())
	assert.Equal(t, "cont2a", ynn.LocalName())
	assert.Equal(t, xpath.ElementNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())
	assert.Equal(t, "value of cont2a", ynn.Value())

	assert.True(t, ynn.MoveToNext())
	assert.Equal(t, "leaf1a", ynn.LocalName())
	assert.Equal(t, xpath.ElementNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())
	assert.Equal(t, "leaf1aval", ynn.Value())

	assert.True(t, ynn.MoveToNext())
	assert.Equal(t, "list2a", ynn.LocalName())
	assert.Equal(t, xpath.ElementNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())

	assert.True(t, ynn.MoveToChild())
	assert.Equal(t, "name", ynn.LocalName())
	assert.Equal(t, xpath.AttributeNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())
	assert.Equal(t, "l2a1", ynn.Value())

	assert.True(t, ynn.MoveToNext())
	assert.Equal(t, "rx-power", ynn.LocalName())
	assert.Equal(t, xpath.ElementNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())
	assert.Equal(t, "25", ynn.Value())

	assert.True(t, ynn.MoveToNext())
	assert.Equal(t, "tx-power", ynn.LocalName())
	assert.Equal(t, xpath.ElementNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())
	assert.Equal(t, "5", ynn.Value())

	assert.False(t, ynn.MoveToNext())

	assert.True(t, ynn.MoveToParent())
	assert.Equal(t, "list2a", ynn.LocalName())
	assert.Equal(t, xpath.ElementNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())

	assert.True(t, ynn.MoveToNext())
	assert.Equal(t, "list2a", ynn.LocalName())
	assert.Equal(t, xpath.ElementNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())

	assert.True(t, ynn.MoveToChild())
	assert.Equal(t, "name", ynn.LocalName())
	assert.Equal(t, xpath.AttributeNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())
	assert.Equal(t, "l2a2", ynn.Value())

	assert.True(t, ynn.MoveToNext())
	assert.Equal(t, "rx-power", ynn.LocalName())
	assert.Equal(t, xpath.ElementNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())
	assert.Equal(t, "26", ynn.Value())

	assert.True(t, ynn.MoveToNext())
	assert.Equal(t, "tx-power", ynn.LocalName())
	assert.Equal(t, xpath.ElementNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())
	assert.Equal(t, "6", ynn.Value())

	assert.False(t, ynn.MoveToNext())

	assert.True(t, ynn.MoveToParent())
	assert.Equal(t, "list2a", ynn.LocalName())
	assert.Equal(t, xpath.ElementNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())

	assert.True(t, ynn.MoveToNext())
	assert.False(t, ynn.MoveToNext())

	assert.True(t, ynn.MoveToParent())
	assert.Equal(t, "cont1a", ynn.LocalName())
	assert.Equal(t, xpath.ElementNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())

	assert.True(t, ynn.MoveToNext())
	assert.Equal(t, "cont1b-state", ynn.LocalName())
	assert.Equal(t, xpath.ElementNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())

	assert.True(t, ynn.MoveToChild())
	assert.Equal(t, "list2b", ynn.LocalName())
	assert.Equal(t, xpath.ElementNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())

	assert.True(t, ynn.MoveToChild())
	assert.Equal(t, "index1", ynn.LocalName())
	assert.Equal(t, xpath.AttributeNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())
	assert.Equal(t, "10", ynn.Value())

	assert.False(t, ynn.MoveToChild())

	assert.True(t, ynn.MoveToNext())
	assert.Equal(t, "index2", ynn.LocalName())
	assert.Equal(t, xpath.AttributeNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())
	assert.Equal(t, "20", ynn.Value())

	assert.True(t, ynn.MoveToNext())
	assert.Equal(t, "leaf3c", ynn.LocalName())
	assert.Equal(t, xpath.ElementNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())
	assert.Equal(t, "3c 10-20 test", ynn.Value())

	assert.False(t, ynn.MoveToNext()) // There's no leaf3d present

	assert.True(t, ynn.MoveToParent())
	assert.Equal(t, "list2b", ynn.LocalName())
	assert.Equal(t, xpath.ElementNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())

	assert.True(t, ynn.MoveToNext()) // the next list entry

	assert.True(t, ynn.MoveToChild())
	assert.Equal(t, "index1", ynn.LocalName())
	assert.Equal(t, xpath.AttributeNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())
	assert.Equal(t, "11", ynn.Value())

	assert.True(t, ynn.MoveToNext())
	assert.Equal(t, "index2", ynn.LocalName())
	assert.Equal(t, xpath.AttributeNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())
	assert.Equal(t, "20", ynn.Value())

	assert.True(t, ynn.MoveToNext())
	assert.Equal(t, "leaf3c", ynn.LocalName())
	assert.Equal(t, xpath.ElementNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())
	assert.Equal(t, "3c 11-20 test", ynn.Value())

	assert.True(t, ynn.MoveToNext())
	assert.Equal(t, "leaf3d", ynn.LocalName())
	assert.Equal(t, xpath.ElementNode, ynn.NodeType())
	assert.Equal(t, "t1", ynn.Prefix())
	assert.Equal(t, "IDTYPE2", ynn.Value())

	assert.False(t, ynn.MoveToNext()) // No further leaves

}