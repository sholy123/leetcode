package main

import (
	"golang.org/x/crypto/cryptobyte/asn1"
)

func findMiddleOfTwoArrays(nums1 []int,nums2 []int) float64{
	l1 := len(nums1)
	l2 := len(nums2)
	if l1 > l2{
		return findMiddleOfTwoArrays(nums2, nums1)
	}
	var L1,L2,R1,R2,c1,c2,l0 int
	l0 = 0
	h2 := 2 * l1
	for l0 <= h2{
		c1 = (l0+h2) / 2
		c2 = l1 + l2 - c1
		if c1 == 0{
			L1 =
		}
	}

}
