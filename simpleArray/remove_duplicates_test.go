/**
 * @Author : admin
 * @File : remove_duplicates_test
 * @Date: 2021/7/22 0:03
 * @Description:
 */
package simpleArray

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{1,1,2,3,3,4,4,4,4,4,5,7,8,8,9}
	numsLen := RemoveDuplicates(nums)
	fmt.Println(numsLen)
}
