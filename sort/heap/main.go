
// 最小/大堆排序，以及入堆出堆与删除        //删除空白行：^[ \t]*\n
//根结点为0，左子节点角标为2*k+1,右子节点角标为2*k+2,父节点角标为（k-1）/2
//如果只需要排序输出的话，只需要down()和Pop即可
// 最大堆只需要修改less函数即可
//主要堆的method包括: 入堆Push()、出堆Pop()、向上搜索定位的up()、向下搜索定位的down()、删除元素的Remove()
//主要函数：方便大小堆转换的less()、和两值交换的swap()、以及main
// 注意go中所有参数转递都是值传递// 所以要让h的变化在函数外也起作用，此处得传指针
package main

import "fmt"
type Heap []int

func (h Heap) swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
//单独定义出来时为了方便大小堆排序的转换，当前小堆根排序转换成大堆根，只需要将"<"改成">"
func (h Heap) less(i, j int) bool {
	return h[i] < h[j]
}

//h.Push()入堆，从树的末梢入堆，并使用up来向上搜索定位入堆值的位置
func (h *Heap) Push(x int) {
	*h = append(*h, x)
	h.up(len(*h) - 1)
}

//h.up()向上搜索定位入堆值h(i)的位置，需要与父节点比较，和他的兄弟结点无关，
func (h Heap) up(i int) {
	for {
		f := (i - 1) / 2 // 父亲结点
		if i == f || h.less(f, i) {  //循环结束条件，h(i)为根结点，或者大于父结点
			break
		}
		h.swap(f, i)  //交换当前值h(i)与父结点h(f)的值
		i = f    //交换指针，将当前值上升到父结点那层，并进入到下一次的循环，直到达到循环结束条件
	}
}

//h.down() 向下搜索定位当前值h(i)的位置，传入当前结点，
//并将其与左右子结点的较小值进行比较，当大于较小值时，交换，并进入下一层循环
func (h Heap) down(i int) {
	for {
		l := 2*i + 1    // 得到当前值的左结点
		if l >= len(h) {
			break       //循环结束条件1：h(i)没有左右结点时，即此时h（i）已经是叶子结点了
		}
		j := l        //将左结点给j，j记为左右结点的较小值
		if r := l + 1; r < len(h) && h.less(r, l) {
			j = r
		}
		if h.less(i, j) {
			break      // 环结束条件2:如果父结点比孩子结点小，则不交换
		}
		h.swap(i, j) // 如果父结点比子结点大，则交换父结点和最小的子结点
		i = j        //继续向下比较，进行下一个循环，直到满足结束条件
	}
}

// h.Remove() 删除堆中位置为i的元素
//思路：用最后一个值将待删除值替换，再对删除位置的最后一个值从新进行搜索定位排序
func (h *Heap) Remove(i int) (int, bool) {
	if i < 0 || i > len(*h)-1 {   //判断删除值下标的有效性，下标在（0，len(*h)-1)之间
		return 0, false
	}
	n := len(*h) - 1  //用于替换后丢弃删除的值
	h.swap(i, n)    // 用最后的元素值替换被删除元素
	// 删除最后的元素
	x := (*h)[n]   //x为待删除值
	*h = (*h)[0:n]
	// 如果当前元素大于父结点，满足小堆根排序，则向下筛选
	if (*h)[i] > (*h)[(i-1)/2] {
		h.down(i)
	} else { // 当前元素小于父结点，不满足小堆根排序，向上筛选
		h.up(i)
	}
	return x, true
}

// 弹出堆顶的元素，并返回其值,
func (h *Heap) Pop() int {
	n := len(*h) - 1
	h.swap(0, n)
	x := (*h)[n]
	*h = (*h)[0:n]
	h.down(0)    //弹出堆顶元素后需要从根部重新排序
	return x
}

func main() {
	var h = Heap{20, 7, 3, 10, 15, 25,30, 1, 17, 19}  //9个
	n := len(h)    //通过h.down()将h按照小根堆排序
	//这里不能直接h.down(0),而是需要从最后一个父结点开始，依次循环直到h.down(0)
	for i := (n - 1)/2; i >= 0; i-- {
		h.down(i)
	}
	//h.down(0)       //从下标为0的根结点开始
	fmt.Println(h)     // [1 7 3 10 15 25 30 20 17 19]
	h.Push(6)     //入堆
	fmt.Println(h)    // [1 6 3 10 7 25 30 20 17 19 15]
	x, ok := h.Remove(5)    //删除下标是5的值，返回删除的值x,和删除true/false?
	fmt.Println(x, ok, h)  // 25 true [1 6 3 10 7 15 30 20 17 19]
	y, ok := h.Remove(1)
	fmt.Println(y, ok, h)      // 6 true [1 7 3 10 19 15 30 20 17]
	z := h.Pop()     //出堆
	fmt.Println(z, h)       // 1 [3 7 15 10 19 17 30 20]
	//升序输出---将最小的值依次pop出来存入数组并输出
	var fh []int
	for len(h) > 0 {
		z := h.Pop()
		fh = append(fh, z)
	}
	fmt.Println(fh)    // [3 7 10 15 17 19 20 30]
}

