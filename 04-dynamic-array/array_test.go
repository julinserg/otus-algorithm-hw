package p04dynamicarray

import (
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func NewSingleArrayString(size int) IArray[string] {
	sa := &SingleArray[string]{}
	sa.Create(size)
	return sa
}

func NewVectorArrayString(size int) IArray[string] {
	sa := &VectorArray[string]{}
	sa.Create(size)
	return sa
}

func NewFactorArrayString(size int) IArray[string] {
	sa := &FactorArray[string]{}
	sa.Create(size)
	return sa
}

func NewGoArrayString(size int) IArray[string] {
	sa := &GoArray[string]{}
	sa.Create(size)
	return sa
}

func testItemAddAndRemove(t *testing.T, array IArray[string]) {
	array.Set(0, "123")
	array.Set(1, "456")
	require.Equal(t, 2, array.Size())
	require.Equal(t, "123", array.Get(0))
	require.Equal(t, "456", array.Get(1))
	array.Add(1, "789")
	require.Equal(t, 3, array.Size())
	require.Equal(t, "123", array.Get(0))
	require.Equal(t, "789", array.Get(1))
	require.Equal(t, "456", array.Get(2))
	array.Add(0, "000")
	require.Equal(t, "000", array.Get(0))
	require.Equal(t, "123", array.Get(1))
	require.Equal(t, "789", array.Get(2))
	require.Equal(t, "456", array.Get(3))
	array.Add(4, "999")
	require.Equal(t, "000", array.Get(0))
	require.Equal(t, "123", array.Get(1))
	require.Equal(t, "789", array.Get(2))
	require.Equal(t, "456", array.Get(3))
	require.Equal(t, "999", array.Get(4))
	array.Add(3, "555")
	require.Equal(t, "000", array.Get(0))
	require.Equal(t, "123", array.Get(1))
	require.Equal(t, "789", array.Get(2))
	require.Equal(t, "555", array.Get(3))
	require.Equal(t, "456", array.Get(4))
	require.Equal(t, "999", array.Get(5))
	array.Add(100, "777")
	require.Equal(t, "000", array.Get(0))
	require.Equal(t, "123", array.Get(1))
	require.Equal(t, "789", array.Get(2))
	require.Equal(t, "555", array.Get(3))
	require.Equal(t, "456", array.Get(4))
	require.Equal(t, "999", array.Get(5))
	require.Equal(t, "777", array.Get(6))
	require.Equal(t, 7, array.Size())
	require.Equal(t, "000", array.Remove(0))
	require.Equal(t, 6, array.Size())
	require.Equal(t, "777", array.Remove(5))
	require.Equal(t, 5, array.Size())
	require.Equal(t, "555", array.Remove(2))
	require.Equal(t, 4, array.Size())

	require.Equal(t, "123", array.Get(0))
	require.Equal(t, "789", array.Get(1))
	require.Equal(t, "456", array.Get(2))
	require.Equal(t, "999", array.Get(3))

	array.Add(2, "777")

	require.Equal(t, "123", array.Get(0))
	require.Equal(t, "789", array.Get(1))
	require.Equal(t, "777", array.Get(2))
	require.Equal(t, "456", array.Get(3))
	require.Equal(t, "999", array.Get(4))

}

func testItemAddAndRemoveFrontElement(t *testing.T, array IArray[string]) {
	startAdd := time.Now()
	for j := 0; j < 10000; j++ {
		array.Add(0, strconv.Itoa(j))
	}
	elapsedAdd := time.Since(startAdd)
	log.Printf("Add to front time %s \n", elapsedAdd)
	startRemove := time.Now()
	for j := 0; j < 10000; j++ {
		array.Remove(0)
	}
	elapsedRemove := time.Since(startRemove)
	log.Printf("Remove from front time %s \n", elapsedRemove)
}

func testItemAddAndRemoveBackElement(t *testing.T, array IArray[string]) {
	startAdd := time.Now()
	for j := 0; j < 10000; j++ {
		array.Add(array.Size(), strconv.Itoa(j))
	}
	elapsedAdd := time.Since(startAdd)
	log.Printf("Add to back time %s \n", elapsedAdd)
	startRemove := time.Now()
	for j := 0; j < 10000; j++ {
		backIndex := array.Size() - 1
		if backIndex < 0 {
			backIndex = 0
		}
		array.Remove(backIndex)
	}
	elapsedRemove := time.Since(startRemove)
	log.Printf("Remove from back time %s \n", elapsedRemove)
}

func testItemAddAndRemoveMiddleElement(t *testing.T, array IArray[string]) {
	startAdd := time.Now()
	for j := 0; j < 10000; j++ {
		array.Add(array.Size()/2, strconv.Itoa(j))
	}
	elapsedAdd := time.Since(startAdd)
	log.Printf("Add to middle time %s \n", elapsedAdd)
	startRemove := time.Now()
	for j := 0; j < 10000; j++ {
		array.Remove(array.Size() / 2)
	}
	elapsedRemove := time.Since(startRemove)
	log.Printf("Remove from middle time %s \n", elapsedRemove)
}

func TestSingleArray(t *testing.T) {
	sa := NewSingleArrayString(2)
	testItemAddAndRemove(t, sa)
}

func TestSingleArrayTime(t *testing.T) {
	sa := NewSingleArrayString(0)
	testItemAddAndRemoveFrontElement(t, sa)
	testItemAddAndRemoveBackElement(t, sa)
	testItemAddAndRemoveMiddleElement(t, sa)
}

func TestVectorArray(t *testing.T) {
	sa := NewVectorArrayString(2)
	testItemAddAndRemove(t, sa)
}

func TestVectorArrayTime(t *testing.T) {
	sa := NewVectorArrayString(0)
	testItemAddAndRemoveFrontElement(t, sa)
	testItemAddAndRemoveBackElement(t, sa)
	testItemAddAndRemoveMiddleElement(t, sa)
}

func TestFactorArray(t *testing.T) {
	sa := NewFactorArrayString(2)
	testItemAddAndRemove(t, sa)
}

func TestFactorArrayTime(t *testing.T) {
	sa := NewFactorArrayString(0)
	testItemAddAndRemoveFrontElement(t, sa)
	testItemAddAndRemoveBackElement(t, sa)
	testItemAddAndRemoveMiddleElement(t, sa)
}

func TestGoArray(t *testing.T) {
	sa := NewGoArrayString(2)
	testItemAddAndRemove(t, sa)
}

func TestGoArrayTime(t *testing.T) {
	sa := NewGoArrayString(0)
	testItemAddAndRemoveFrontElement(t, sa)
	testItemAddAndRemoveBackElement(t, sa)
	testItemAddAndRemoveMiddleElement(t, sa)
}
