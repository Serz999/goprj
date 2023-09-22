package list

import (
	"reflect"
	"testing"
)

func (list *List) fill(nodesValues []int64) {
    for _, value := range nodesValues {
        list.Add(value)
    }
}

func (list *List) getNodesContain() (values []int64, idxs []int64) {
    
    node := list.firstNode 
    for node != nil {
        values = append(values, node.value)
        idxs = append(idxs, node.index)
        node = node.next
    } 
    
    return
}

func TestAdd(t *testing.T) {
    type TestCase struct {
        nodesValues []int64
        arg []int64
        expectedValues []int64
        expectedIdxs []int64
    }
    
    testCases := []TestCase{
        TestCase{[]int64{5, 4, 3, 2}, []int64{2}, []int64{5, 4, 3, 2, 2}, []int64{0, 1, 2, 3, 4}},
        TestCase{[]int64{5, 4, 3, 2}, []int64{44, 9}, []int64{5, 4, 3, 2, 44, 9}, []int64{0, 1, 2, 3, 4, 5}},
        TestCase{[]int64{5, 4, 3, 2}, []int64{}, []int64{5, 4, 3, 2}, []int64{0, 1, 2, 3}},
    }

    for _, tc := range testCases {
        list := NewList() 

        list.fill(tc.nodesValues)

        for _, value := range tc.arg {
            list.Add(value)
        }

        gotValues, gotIdxs := list.getNodesContain() 
        
        if !reflect.DeepEqual(gotValues, tc.expectedValues) {
            t.Errorf("got values %v, expected values %v", gotValues, tc.expectedValues)
        }

        if !reflect.DeepEqual(gotIdxs, tc.expectedIdxs) {
            t.Errorf("got idxs %v, expected idxs %v", gotIdxs, tc.expectedIdxs)
        }
    }
}

func TestLen(t *testing.T) {
    type TestCase struct {
        nodesValues []int64
        expectedLen int64
    }
    
    testCases := []TestCase{
        TestCase{[]int64{5, 4, 3, 2}, 4},
        TestCase{[]int64{}, 0},
    }

    for _, tc := range testCases {
        list := NewList() 

        for i := 0; i < int(tc.expectedLen); i++ {
            list.Add(int64(i))
        }

        got := list.Len() 
    
        if  got != tc.expectedLen {
            t.Errorf("got %v, expected %v", got, tc.expectedLen)
        }
    }
}

func TestRemoveByIndex(t *testing.T) {
    type TestCase struct {
        nodesValues []int64
        arg int64
        expectedValues []int64
        expectedIdxs []int64
    }
    
    testCases := []TestCase{
        TestCase{[]int64{5, 4, 3, 2}, 2, []int64{5, 4, 2}, []int64{0, 1, 2}},
        TestCase{[]int64{5, 4, 3, 2}, 3, []int64{5, 4, 3}, []int64{0, 1, 2}},
    }

    for _, tc := range testCases {
        list := NewList()

        list.fill(tc.nodesValues)

        list.RemoveByIndex(tc.arg)     

        gotValues, gotIdxs := list.getNodesContain() 
        
        if !reflect.DeepEqual(gotValues, tc.expectedValues) {
            t.Errorf("got values %v, expected values %v", gotValues, tc.expectedValues)
        }

        if !reflect.DeepEqual(gotIdxs, tc.expectedIdxs) {
            t.Errorf("got idxs %v, expected idxs %v", gotIdxs, tc.expectedIdxs)
        }
    }
}

func TestRemoveByValue(t *testing.T) {
    type TestCase struct {
        nodesValues []int64
        arg int64
        expectedValues []int64
        expectedIdxs []int64
    }
    
    testCases := []TestCase{
        TestCase{[]int64{5, 4, 3, 4, 4, 2}, 3, []int64{5, 4, 4, 4, 2}, []int64{0, 1, 2, 3, 4}},
        TestCase{[]int64{5, 4, 3, 4, 4, 2}, 4, []int64{5, 3, 4, 4, 2}, []int64{0, 1, 2, 3, 4}},
    }

    for _, testCase := range testCases {

        list := NewList()

        list.fill(testCase.nodesValues)

        list.RemoveByValue(testCase.arg)     

        gotValues, gotIdxs := list.getNodesContain() 
        
        if !reflect.DeepEqual(gotValues, testCase.expectedValues) {
            t.Errorf("got values %v, expected values %v", gotValues, testCase.expectedValues)
        }

        if !reflect.DeepEqual(gotIdxs, testCase.expectedIdxs) {
            t.Errorf("got idxs %v, expected idxs %v", gotIdxs, testCase.expectedIdxs)
        }
    }
}

func TestRemoveAllByValue(t *testing.T) {
     type TestCase struct {
        nodesValues []int64
        arg int64
        expectedValues []int64
        expectedIdxs []int64
    }
    
    testCases := []TestCase{
        TestCase{[]int64{5, 4, 3, 4, 4, 2}, 2, []int64{5, 4, 3, 4, 4}, []int64{0, 1, 2, 3, 4}},
        TestCase{[]int64{5, 4, 3, 4, 4, 2}, 4, []int64{5, 3, 2}, []int64{0, 1, 2}},
    }

    for _, testCase := range testCases {
        list := NewList()
    
        list.fill(testCase.nodesValues)

        list.RemoveAllByValue(testCase.arg)     

        gotValues, gotIdxs := list.getNodesContain() 
        
        if !reflect.DeepEqual(gotValues, testCase.expectedValues) {
            t.Errorf("got values %v, expected values %v", gotValues, testCase.expectedValues)
        }

        if !reflect.DeepEqual(gotIdxs, testCase.expectedIdxs) {
            t.Errorf("got idxs %v, expected idxs %v", gotIdxs, testCase.expectedIdxs)
        }
    }
}

func TestGetByIndex(t *testing.T) {

    type TestCase struct {
        nodesValues []int64
        arg int64
        expectedValue int64
        expectedStatus bool
    }
    
    testCases := []TestCase{
        TestCase{[]int64{5, 4, 3, 4, 4, 2}, 2, 3, true},
        TestCase{[]int64{5, 4, 3, 4, 4, 2}, 100, 0, false},
    }

    for _, testCase := range testCases {
        list := NewList()

        list.fill(testCase.nodesValues)

        value, status := list.GetByIndex(testCase.arg) 
    
        if value != int64(testCase.expectedValue) {
            t.Errorf("got values %v, expected values %v", value, testCase.expectedValue)
        }

        if status != testCase.expectedStatus {
            t.Errorf("got status %v, expected status %v", value, testCase.expectedStatus)
        }
    }
}

func TestGetByValue(t *testing.T) {

    type TestCase struct {
        nodesValues []int64
        arg int64
        expectedValue int64
        expectedStatus bool
    }
    
    testCases := []TestCase{
        TestCase{[]int64{5, 4, 3, 4, 4, 2}, 3, 2, true},
        TestCase{[]int64{5, 4, 3, 4, 4, 2}, 4, 1, true},
        TestCase{[]int64{5, 4, 3, 4, 4, 2}, 444, 0, false},
    }

    for _, testCase := range testCases {
        list := NewList()

        list.fill(testCase.nodesValues) 

        value, status := list.GetByValue(testCase.arg) 
    
        if value != int64(testCase.expectedValue) {
            t.Errorf("got idx %v, expected idx %v", value, testCase.expectedValue)
        }

        if status != testCase.expectedStatus {
            t.Errorf("got status %v, expected status %v", value, testCase.expectedStatus)
        }
    }
}

func TestGetAllByValue(t *testing.T) {

    type TestCase struct {
        nodesValues []int64
        arg int64
        expectedValue []int64
        expectedStatus bool
    }
    
    testCases := []TestCase{
        TestCase{[]int64{5, 4, 3, 4, 4, 2}, 3, []int64{2}, true},
        TestCase{[]int64{5, 4, 3, 4, 4, 2}, 4, []int64{1, 3, 4}, true},
        TestCase{[]int64{5, 4, 3, 4, 4, 2}, 444, nil, false},
    }

    for _, testCase := range testCases {
        list := NewList()

        list.fill(testCase.nodesValues) 

        idxs, status := list.GetAllByValue(testCase.arg) 
    
        if !reflect.DeepEqual(idxs, testCase.expectedValue) {
            t.Errorf("got idxs %v, expected idxs %v", idxs, testCase.expectedValue)
        }

        if status != testCase.expectedStatus {
            t.Errorf("got status %v, expected status %v", status, testCase.expectedStatus)
        }
    }
}

func TestGetAll(t *testing.T) {

    type TestCase struct {
        nodesValues []int64
        expectedValue []int64
        expectedStatus bool
    }
    
    testCases := []TestCase{
        TestCase{[]int64{5, 4, 3, 4, 4, 2}, []int64{5, 4, 3, 4, 4, 2}, true},
        TestCase{[]int64{1, 3, 4}, []int64{1, 3, 4}, true},
        TestCase{[]int64{},  nil, false},
    }

    for _, testCase := range testCases {
        list := NewList()

        list.fill(testCase.nodesValues) 

        idxs, status := list.GetAll() 
    
        if !reflect.DeepEqual(idxs, testCase.expectedValue) {
            t.Errorf("got idxs %v, expected idxs %v", idxs, testCase.expectedValue)
        }

        if status != testCase.expectedStatus {
            t.Errorf("got status %v, expected status %v", status, testCase.expectedStatus)
        }
    }
}

func TestClear(t *testing.T) {

    type TestCase struct {
        nodesValues []int64
    }
    
    testCases := []TestCase{
        TestCase{[]int64{5, 4, 3, 4, 4, 2}},
        TestCase{[]int64{1, 3, 4}},
        TestCase{[]int64{}},
    }

    for _, testCase := range testCases {
        list := NewList()

        list.fill(testCase.nodesValues) 

        list.Clear() 
 
        if list.firstNode != nil {
            t.Errorf("got %v, expected %v", list.firstNode, nil)
        }         
        
        if list.len != 0 {
            t.Errorf("got %v, expected %v", list.len, 0)
        }
    }
}

func ExamplePrint() {
    list := NewList()
    list.fill([]int64{5, 4, 3, 4, 4, 2})
    list.Print()
}

