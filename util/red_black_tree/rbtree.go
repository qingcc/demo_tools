package RBTree

type node struct {
	color               bool
	key                 int
	left, right, parent *node
}

type RBTree struct {
	root *node
}

var (
	Red   = false
	Black = true
)

func NewRBTree() *RBTree {
	return &RBTree{root: nil}
}

func (n *node) compare(key int) int {
	if n.key > key {
		return -1 //key 比该节点值小，左子树查找
	} else if n.key < key {
		return 1 //key 比该节点值大， 右字树查找
	} else {
		return 0 //相等，返回
	}
}

func (tree *RBTree) Search(key int) *node {
	return tree.root.find(key)
}

//search
func (n *node) find(key int) *node {
	for n != nil {
		switch n.compare(key) {
		case -1:
			n = n.left
		case 1:
			n = n.right
		default:
			return n
		}
	}
	return n
}

func (n *node) getParent() (parent *node) {
	if n != nil {
		parent = n.parent
	}
	return
}

func (n *node) setParent(parent *node) {
	if n != nil {
		n.parent = parent
	}
}

func (n *node) isRed() (isRed bool) {
	if n != nil {
		isRed = !n.color
	}
	return
}

func (n *node) setRed() {
	if n != nil {
		n.color = false
	}
	return
}

func (n *node) isBlack() (isBlack bool) {
	return !n.isRed()
}

func (n *node) setBlack() {
	if n != nil {
		n.color = true
	}
	return
}

func (n *node) getColor() (color bool) {
	if n != nil {
		color = n.color
	} else {
		color = Black
	}
	return
}

func (n *node) setColor(color bool) {
	if n != nil {
		n.color = color
	}
	return
}

//前序遍历
func (n *node) PreOrder() {
	if n != nil {
		print(n.key, " ")
		n.left.PreOrder()
		n.right.PreOrder()
	}
}

//中序遍历
func (n *node) inOrder() {
	if n != nil {
		n.left.inOrder()
		print(n.key, " ")
		n.right.inOrder()
	}
}

//后序遍历
func (n *node) postOrder() {
	if n != nil {
		n.left.postOrder()
		n.right.postOrder()
		print(n.key, " ")
	}
}

func (n *node) minValue() (key int) {
	if min := n.minNode(); min != nil {
		key = min.key
	}
	return key
}

func (n *node) minNode() *node {
	for n != nil {
		n = n.left
	}
	return n
}

func (n *node) maxValue() (key int) {
	if min := n.maxNode(); min != nil {
		key = min.key
	}
	return key
}

func (n *node) maxNode() *node {
	for n != nil {
		n = n.right
	}
	return n
}

/********* 查找节点x的后继节点,即大于节点x的最小节点 ***********/
func (n *node) successor() *node {
	//有右子节点
	if n.right != nil {
		return n.right.minNode()
	}
	//1.x是其父节点的左子节点，则x的后继节点为它的父节点
	//2. x是其父节点的右子节点，则先查找x的父节点p，然后对p再次进行这两个条件的判断
	p := n.parent
	for p != nil && (n == p.right) { //对应情况2
		n = p
		p = n.parent
	}
	return p //对应情况1
}

/********* 查找节点x的前驱节点，即小于节点x的最大节点 ************/
func (n *node) predecessor() *node {
	//有左子节点
	if n.left != nil {
		return n.left.maxNode()
	}
	//1. x是其父节点的右子节点，则x的前驱节点是它的父节点
	//2. x是其父节点的左子节点，则先查找x的父节点p，然后对p再次进行这两个条件的判断
	p := n.parent
	for p != nil && (n == p.left) { //对应情况2
		n = p
		p = n.parent
	}
	return p //对应情况1
}

/*************对红黑树节点x进行左旋操作 ******************/
/*
 * 左旋示意图：对节点x进行左旋
 *     p                       p
 *    /                       /
 *   x                       y
 *  / \                     / \
 * lx  y      ----->       x  ry
 *    / \                 / \
 *   ly ry               lx ly
 * 左旋做了三件事：
 * 1. 将y的左子节点赋给x的右子节点,并将x赋给y左子节点的父节点(y左子节点非空时)
 * 2. 将x的父节点p(非空时)赋给y的父节点，同时更新p的子节点为y(左或右)
 * 3. 将y的左子节点设为x，将x的父节点设为y
 */

func (x *node) leftRotate(tree *RBTree) {
	if x == nil || x.right == nil { //x的右子节点为nil，无法左旋
		return
	}
	p, y, ly := x.parent, x.right, x.right.left
	y.parent, y.left, x.parent, x.right = p, x, y, ly
	if p == nil {
		tree.root = y
	} else if p.left == x {
		p.left = y
	} else {
		p.right = y
	}
	//ly := y.left
	//p := x.parent
	//x.right = ly
	//y.left = x
	//x.parent = y
	//y.parent = p
	return
}

/*************对红黑树节点y进行右旋操作 ******************/
/*
 * 右旋示意图：对节点y进行右旋
 *        p                   p
 *       /                   /
 *      y                   x
 *     / \                 / \
 *    x  ry   ----->      lx  y
 *   / \                     / \
 * lx  rx                   rx ry
 * 右旋做了三件事：
 * 1. 将x的右子节点赋给y的左子节点,并将y赋给x右子节点的父节点(x右子节点非空时)
 * 2. 将y的父节点p(非空时)赋给x的父节点，同时更新p的子节点为x(左或右)
 * 3. 将x的右子节点设为y，将y的父节点设为x
 */
func (y *node) rightRotate(tree *RBTree) {
	if y == nil || y.left == nil {
		return
	}
	x, rx, p := y.left, y.left.right, y.parent
	x.right, x.parent, y.left, y.parent = y, p, rx, x
	if p == nil {
		tree.root = x
	} else if p.left == y {
		p.left = x
	} else {
		p.right = x
	}
	return
}

/*********************** 向红黑树中插入节点 **********************/
func (tree *RBTree) Insert(key int) {
	newNode := &node{key: key, color: Red}
	tree.insert(newNode)
	return
}

func (tree *RBTree) insert(n *node) {
	current := tree.root
	var p *node
	//1. 找到插入的位置
	//树中已经有key，直接跳出，不做任何操作
	//p 为current的父节点，current为需要插入的node， 跳出循环时，即找到需要插入到p节点下
	for current != nil {
		p = current
		if current.compare(n.key) == 1 {
			current = current.right
		} else if current.compare(n.key) == -1 {
			current = current.left
		} else {
			return
		}
	}
	n.parent = p

	//2. 判断，是插入到p的左节点还是右节点
	if p != nil {
		if p.compare(n.key) == 1 {
			p.right = n
		} else {
			p.left = n
		}
	} else {
		tree.root = n
	}
	//3. 将它重新修整为一颗红黑树
	n.insertFixUp(tree)
	return
}

func (n *node) insertFixUp(tree *RBTree) {
	var parent, gparent *node
	for parent = n.parent; parent != nil && parent.isRed(); {

		gparent = n.parent.parent
		if gparent == nil {
			return
		}
		//若父节点为祖父节点左子节点
		if parent == gparent.left {
			uncle := gparent.right //叔叔节点

			//case1:叔叔节点也是红色
			if uncle != nil && uncle.isRed() {
				uncle.setBlack()
				parent.setBlack()
				gparent.setRed()
				n = gparent
				continue
			}

			//case2:叔叔节点是黑色，且当前节点是右子节点
			if n == parent.right {
				parent.leftRotate(tree)
				n, parent = parent, n
			}
			//case3: 叔叔节点是黑色，且当前节点是左子节点
			parent.setBlack()
			gparent.setRed()
			gparent.rightRotate(tree)
		} else { //若父节点为祖父节点右子节点
			uncle := gparent.left

			//case1:叔叔节点也是红色的
			if uncle != nil && uncle.isRed() {
				uncle.setBlack()
				parent.setBlack()
				gparent.setRed()
				n = gparent
				continue
			}

			//case2: 叔叔节点是黑色的，且当前节点是左子节点
			if n == parent.left {
				parent.rightRotate(tree)
				n, parent = parent, n
			}

			//叔叔节点是黑色的
			parent.setBlack()
			gparent.setRed()
			gparent.leftRotate(tree)
		}
	}
	//将根节点设置为黑色
	tree.root.setBlack()
}

/*********************** 删除红黑树中的节点 **********************/
func (r *node) Remove(x *node) {

	return
}
