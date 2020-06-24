package RBTree

type node struct {
	color               bool
	key                 int
	left, right, parent *node
}

type RBTree struct {
	root *node
}

/*
										      |
				     		b
				    / 				\
				   r 				 b
				  / \				/ \
				b    b  		   b   b
			  / \   / \
			r   r  r  b
           /\   /\ /\
           b b  b b b b






*/

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
func (n *node) getBrother() (brother *node) {
	if n != nil && n.parent != nil {
		if n.parent.left == n {
			brother = n.parent.right
		} else {
			brother = n.parent.left
		}
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
func (n *node) isLeftSon() (isLeftSon bool) {
	if n != nil && n.parent != nil && n.parent.left == n {
		isLeftSon = true
	}
	return
}
func (n *node) isRightSon() (isRightSon bool) {
	isRightSon = !n.isLeftSon()
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
	for n != nil && n.left != nil {
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
	for n != nil && n.right != nil {
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

//
///*********************** 删除红黑树中的节点 **********************/
//func (tree *RBTree) Remove(key int) {
//	if node := tree.Search(key); node != nil {
//		node.remove(tree)
//	}
//	return
//}
//
//func (n *node)remove(tree *RBTree)  {
//	//n为叶子节点
//	if n.left == nil && n.right == nil {
//		if n.isRed() {
//			n.clear()
//		}else {
//			//最复杂的，n为黑色子节点（n的兄弟节点一定存在有2种情况：）
//			/* n节点可能为左/右子节点，这里只展示右子节点（左子节点左右对称）
//			 *
//			 * 			      |                     |
//			 * 			     pb                     pr/b
//			 * 			    / \                    / \
//			 * 			   br  nb       or       bb  nb
//			 * 			  / \                    / \
//			 * 			blb  brb                 	(r/nil)
//			 * 		   / \   / \
//			 * 	            	  (r/nil)
//			 */
//			if n.parent.isRed() {
//				n.complexRemoveParentIsRed(tree)
//			}else if n.getBrother().isRed() {
//				n.complexRemoveParentIsBlackAndbrotherIsRed(tree)
//			}else {
//				n.complexRemoveParentAndbrotherIsBlack(tree)
//			}
//			if parent := n.parent; parent.isRed() { //如果parent是 red， 则parent.parent.Color=Black, parent可设置为r/b
//				if n.isRightSon() {
//					n.clear()
//					if brother := n.left; brother.left != nil {					//     pr								    ub
//						parent.rightRotate(tree)								//    / \								   / \
//						if brother.right == nil {								//  bb  nb								 blr  pr
//							brother.setRed()			//可以不调色				//  /
//							brother.left.setBlack()		//可以不调色				// blr
//							brother.right.setBlack()	//可以不调色
//						}else {												//     pr								    ub
//							brother.setRed()								//    / \								   / \
//							brother.left.setBlack()							//  ub  nb								 blr  pr
//							parent.setBlack()								//  / \										  /
//						}													// blr brr									brr
//					}else {
//						if brother.right != nil {
//							//     pr					   pr					       	brr
//							//    / \		u ll		  / 			p rr		    / \
//							//  ub  nb		--->		brr  			--->		   ub  pr
//							//   \						/
//							//    brr				   ub
//							brother.leftRotate(tree)
//							parent.rightRotate(tree)
//							parent.setBlack()
//						}else {
//							//     pr
//							//    / \
//							//  ub  nb
//							parent.setBlack()
//							brother.setRed()
//						}
//					}
//				}else {  //n is left son
//					n.clear()
//					if brother := n.left; brother.right != nil {				//     pr								    ub
//						brother.leftRotate(tree)								//    / \								   / \
//						if brother.left == nil {								//  nb  ub								  pr  brr
//							brother.setRed()			//可以不调色				//     	  \
//							brother.left.setBlack()		//可以不调色				//        brr
//							brother.right.setBlack()	//可以不调色
//						}else {												//     pr								    ub
//							brother.setRed()								//    / \								   / \
//							brother.left.setBlack()							//  nb  ub								 pr  brr
//							brother.right.setBlack()						//     / \								   \
//						}													//    blr brr							   blr
//					}else {
//						if brother.left != nil {
//							//     pr					   pr					       	blr
//							//    / \		u rr		    \			p rr		    / \
//							//  nb  ub		--->		    blr	    	--->		   pr  ub
//							//      / 					      \
//							//    blr  				          ub
//							brother.rightRotate(tree)
//							parent.leftRotate(tree)
//							parent.setBlack()
//						}else {
//							//     pr
//							//    / \
//							//  nb  ub
//							parent.setBlack()
//							brother.setRed()
//						}
//					}
//				}
//			}else {
//				//parent is black
//				if brother := n.getBrother(); brother.isBlack() {
//					if n.isRightSon() {
//						n.clear()
//						if brother.left != nil {
//							//		 pb					ub
//							//		/ \      p rr	    / \
//							//	   ub  nb    --->      blr pb
//							//	  / \                      /
//							//	 blr (brr/nil)           (brr/nil)
//							parent.rightRotate(tree)
//							brother.left.setBlack()
//						}else if brother.right != nil {
//							//		 pb					pb							    brr
//							//		/ \      u ll		/		p rr				    / \
//							//	   ub  nb    --->	  brr       --->				  ub   pb
//							//	    \            	   /
//							//	   brr           	  ub
//							brother.leftRotate(tree)
//							parent.rightRotate(tree)
//							brother.parent.setBlack()
//						}else {
//							//		 pb
//							//		/ \
//							//	   ub  nb
//							//todo 待完成
//						}
//					}else { // parent is black, brother is black, n is left son
//						if brother.right != nil {
//							//		 pb					ub
//							//		/ \      p ll	    / \
//							//	   nb  ub    --->      pb brr
//							//	      / \               \
//							//(blr/nil) brr           (blr/nil)
//							parent.leftRotate(tree)
//							brother.right.setBlack()
//						}else if brother.left != nil {
//							//		 pb					pb							    blr
//							//		/ \      u rr		 \	       p ll				    / \
//							//	   nb  ub    --->	      blr      --->				  pb   ub
//							//	       /            	   \
//							//	      blr           	    ub
//							brother.rightRotate(tree)
//							parent.leftRotate(tree)
//							brother.parent.setBlack()
//						} else {
//							//		 pb
//							//		/ \
//							//	   ub  nb
//							//todo 待完成
//						}
//					}
//				}else { //parent, n, bl, br is black, brother is red
//					if n.isRightSon() {
//						if brother.left.left != nil || brother.left.right != nil {
//							//		      |
//							//		     pb    						br					br
//							//		    / \      p rr              / \      p rr	   / \
//							//		   br  nb    --->           blb   pb    --->	blb   brb
//							//		  / \                       /\    /				/\    / \
//							//		blb  brb   						 brb				 brlr pb
//							//	   / \   / \						/ \						  /
//							//		   brlr                       brlr				  		(brrr/nil)
//							if brother.right.left != nil  {
//								parent.rightRotate(tree)
//								parent.rightRotate(tree)
//								brother.setBlack()
//								brother.right.setRed()
//								brother.right.left.setBlack()
//							}else if brother.right.right != nil {
//								//		      |
//								//		     pb    						br						br								br
//								//		    / \      p rr              / \      br ll		   / \    				p rr	   / \
//								//		   br  nb    --->           blb   pb    --->		blb   pb  				--->	blb   brlr
//								//		  / \                       /\    /					/\    /							/\    / \
//								//		blb  brb   						 brb					 brlr							 brb pb
//								//	   / \    \							  \						 /
//								//		   	  brrr                        brlr				    brb
//								parent.rightRotate(tree)
//								parent.left.leftRotate(tree)
//								parent.rightRotate(tree)
//								brother.setBlack()
//							}else {
//								//		      |
//								//		     pb    						br
//								//		    / \      p rr              / \
//								//		   br  nb    --->           blb   pb
//								//		  / \                       /\    /
//								//		blb  brb   						 brb
//								//	   / \
//								parent.rightRotate(tree)
//								brother.setBlack()
//								parent.left.setRed()
//							}
//						}else {
//							if brother.right.left != nil || brother.right.right != nil {
//								//		      |
//								//		     pb    						br
//								//		    / \      p rr              / \
//								//		   br  nb    --->           blb   pb
//								//		  / \                             /
//								//		blb  brb   						 brb
//								//	   		/ \    					     /\
//								parent.rightRotate(tree)			//							br
//								if parent.left.left != nil {		//		p ll			   / \
//									parent.leftRotate(tree)			//		--->			blb   brb
//									brother.right.setRed()			//						      / \
//									brother.right.left.setBlack()	//							brlr pb
//									brother.setBlack()
//								}else { //brl is nil, brr is red
//									//		      |
//									//		     pb    						br						br							br
//									//		    / \      p rr              / \		br ll	       / \		p rr		       / \
//									//		   br  nb    --->           blb   pb     --->	    blb   pb    --->		    blb   brrr
//									//		  / \                             /				          /					          /	 \
//									//		blb  brb   						 brb					 brrr						 brb  pb
//									//	   		 \    					       \				      /
//									//	   		 brrr    					    brrr				  brb
//									parent.left.leftRotate(tree)
//									parent.rightRotate(tree)
//									brother.setBlack()
//								}
//							}else {
//								//		      |
//								//		     pb    						br
//								//		    / \      p rr              / \
//								//		   br  nb    --->           blb   pb
//								//		  / \                             /
//								//		blb  brb   						 brb
//								parent.rightRotate(tree)
//								brother.setBlack()
//								parent.left.setRed()
//							}
//						}
//					}else { //n is left son
//						if brother.right.left != nil || brother.right.right != nil {
//							//		      |
//							//		     pb    						br					br						br
//							//		    / \      p ll              / \      bl rr	   /   \       p ll		   /   \
//							//		   nb  br    --->            pb   brb    --->	  pb    brb    --->		 bllr    brb
//							//		  	  / \                     \    /\			   \     / \			 /  \     / \
//							//			blb  brb   				  blb				 	bllr 				pb 	blb
//							//	   		/ \   / \				  /	\				  	 \	 				  	 \
//							//		   bllr                      bllr				    blb
//							//															  \
//							//
//							if brother.left.left != nil  {
//								parent.leftRotate(tree)
//								parent.right.rightRotate(tree)
//								parent.leftRotate(tree)
//								brother.setBlack()
//							}else if brother.left.right != nil {
//								//		      |
//								//		     pb    						br					br
//								//		    / \      p ll              / \       p ll	   /   \
//								//		   nb  br    --->            pb   brb    --->	  blb    brb
//								//		  	  / \                     \    /\			  / \     / \
//								//			blb  brb   				  blb				 pb	blrr
//								//	   		 \   / \				  	\
//								//		    blrr                      blrr
//								//
//								//
//								parent.leftRotate(tree)
//								parent.leftRotate(tree)
//								brother.setBlack()
//								parent.setRed()
//							}else {
//								//		      |
//								//		     pb    						br
//								//		    / \      p ll              / \
//								//		   nb  br    --->            pb   brb
//								//		  	  / \                     \   / \
//								//			blb  brb   				  blb
//								//	   			/ \
//								parent.leftRotate(tree)
//								brother.setBlack()
//								parent.right.setRed()
//							}
//						}else { // brother.right has no son
//							if brother.right.left != nil || brother.right.right != nil {
//								//		      |
//								//		     pb    						br
//								//		    / \      p ll              / \
//								//		   nb  br    --->           pb   brb
//								//		  	  / \                     \
//								//			blb  brb   				  blb
//								//	   		/ \    					  /\
//								parent.leftRotate(tree)				//							br
//								if parent.right.right != nil {		//		p ll			   / \
//									parent.leftRotate(tree)			//		--->			blb   brb
//									brother.left.setRed()			//				        / \
//									brother.left.right.setBlack()	//					  pb  blrr
//									brother.setBlack()				//						\
//								}else { //blr is nil, bll is red
//									//		      |
//									//		     pb    						br						br							br
//									//		    / \      p ll              / \		bl rr	       / \		p ll		       / \
//									//		   nb  br    --->           pb   brb     --->	    pb   brb    --->		    bllr  brb
//									//		  	  / \                     \        				 \      					/ \
//									//			blb  brb   				  blb					 bllr						pb blb
//									//	   		/    					  /     				    \
//									//	      bllr	    				bllr	    				 blb
//									parent.right.rightRotate(tree)
//									parent.leftRotate(tree)
//									brother.setBlack()
//								}
//							}else {
//								//		      |
//								//		     pb    						br
//								//		    / \      p ll              / \
//								//		   nb  br    --->            pb   brb
//								//		  	  / \                     \
//								//			blb  brb   				  blb
//								parent.leftRotate(tree)
//								brother.setBlack()
//								parent.right.setRed()
//							}
//						}
//					}
//
//				}
//			}
//		}
//	}
//	//n只有一个子节点(n只能是黑色节点，且子节点为红色节点，否则不满足红黑树性质）,只需把子节点的值赋给n节点，删除子节点即可
//	if n.left == nil || n.right == nil {
//		if n.left == nil {
//			n.key = n.right.key
//			n.right.destroy()
//			n.right = nil
//		}else {
//			n.key = n.left.key
//			n.left.destroy()
//			n.left = nil
//		}
//	}
//	//n有2个子节点， 查找n的前驱节点或后继节点，把前驱节点的值，或后继节点的值赋给n节点，删除前驱或后继节点，把问题转为删除只有一个子节点 或 无子节点。
//	pre := n.predecessor()
//	n.key = pre.key
//	pre.remove(tree)
//}

/*********************** 删除红黑树中的节点 **********************/
func (tree *RBTree) Remove(key int) {
	if node := tree.Search(key); node != nil {
		tree.RemoveNode(node)
	}
}

func (tree *RBTree) RemoveNode(n *node) {
	var replace *node
	//node 有双子节点
	if n.left != nil && n.right != nil {
		for {
			succ := n.successor()
			n.key = succ.key
			tree.RemoveNode(succ)
			return
		}
	} else { //node 有单子节点或无子节点
		//node 为根节点
		if n.parent == nil {
			if n.left != nil {
				tree.root = n.left
			} else {
				tree.root = n.right
			}
			replace = tree.root
			if tree.root != nil {
				tree.root.parent = nil
			}
		} else { //非跟节点
			var child *node
			if n.left != nil {
				child = n.left
			} else {
				child = n.right
			}
			if n.parent.left == n {
				n.parent.left = child
			} else {
				n.parent.right = child
			}
			if child != nil {
				child.parent = n.parent
			}
			replace = child
		}
	}

	//如果待删除节点为红色，直接删除
	if n.isBlack() {
		replace.deleteFixUp(tree, n.parent)
	}
	n.clear()
}

func (replace *node) deleteFixUp(tree *RBTree, parent *node) { //im. replace is nil todo need fix
	for (replace == nil || replace.isBlack()) && replace != tree.root {
		//parent := replace.parent
		brother := replace.getBrother()
		if replace.isLeftSon() {
			//case1, brother is red
			/*
			 *   p                       bb
			 *  / \       p ll          / \
			 * n   br      ----->      pr  brb
			 *    / \                 / \
			 *   blb brb             n  blb
			 */
			if brother.isRed() { // brother is red
				brother.setBlack()
				parent.setRed()
				parent.leftRotate(tree)
				brother = parent.right
			}
			//case 2 brother and brother's two son are black
			if (brother.left == nil || brother.left.isBlack()) && (brother.right == nil || brother.right.isBlack()) {
				if parent.isRed() {
					parent.setBlack()
					brother.setRed()
					break
					// 如果此时parent为黑，即此时全黑了，则把brother涂红，导致brother分支少一个黑，使整个分支都少了一个黑，需要对parent又进行一轮调整
				} else {
					brother.setRed()
					replace = parent
					parent = replace.parent
				}
			} else {
				//case3 brother is black, brother's left son is red
				if brother.left != nil && brother.left.isRed() {
					brother.left.color = parent.color
					parent.setBlack()
					brother.rightRotate(tree)
					parent.leftRotate(tree)
					//case4 brother.right is red
				} else if brother.right != nil && brother.right.isRed() {
					brother.color = parent.color
					parent.setBlack()
					brother.right.setBlack()
					parent.leftRotate(tree)
				}
				break
			}
		} else { //
			//case1 brother is red
			if brother.isRed() {
				brother.setBlack()
				parent.setRed()
				parent.rightRotate(tree)
				brother = parent.left
			}
			//case2 brother, brother.left, brother.right are black
			if (brother.left == nil || brother.left.isBlack()) && (brother.right == nil || brother.right.isBlack()) {
				if parent.isRed() {
					parent.setBlack()
					brother.setRed()
					break
				} else {
					brother.setRed()
					replace = parent
					parent = replace.parent
				}
			} else {
				//case3 brother is black, brother.left is red
				if brother.right != nil && brother.right.isRed() {
					brother.right.color = parent.color
					parent.setBlack()
					brother.leftRotate(tree)
					parent.rightRotate(tree)
					//case4 brother is black, brother.right is red
				} else if brother.left != nil && brother.left.isRed() {
					brother.color = parent.color
					parent.setBlack()
					brother.left.setBlack()
					parent.rightRotate(tree)
				}
				break
			}
		}
	}
	if replace != nil {
		replace.setBlack()
	}
}

/* n节点可能为左/右子节点，这里只展示右子节点（左子节点左右对称）
 * 	        |
 * 	        pr
 * 	       / \
 * 	     ub  nb
 * 	     / \
 * 	     	(r/nil)
 */
func (n *node) complexRemoveParentIsRed(tree *RBTree) {
	parent, brother, isRightSon := n.parent, n.getBrother(), n.isRightSon()
	n.clear()
	if isRightSon {
		if brother.left != nil { //     pr								    ub
			//parent.rightRotate(tree)							//    / \								   / \
			//if brother.right == nil {							//  ub  nb								 blr  pr
			//	brother.setRed()			//可以不调色			//  /
			//	brother.left.setBlack()	//可以不调色				// blr
			//	brother.right.setBlack()	//可以不调色
			//}else {											//     pr								    ub
			//	brother.setRed()								//    / \								   / \
			//	brother.left.setBlack()							//  ub  nb								 blr  pr
			//	parent.setBlack()								//  / \										  /
			//}													// blr brr									brr
			if brother.right != nil { //     pr								    ub
				brother.setRed()        //    / \								   / \
				brother.left.setBlack() //  ub  nb								 blr  pr
				parent.setBlack()       //  / \										  /
			}
		} else { //brother.left is nil
			if brother.right != nil {
				//     pr					   pr					       	brr
				//    / \		u ll		  / 			p rr		    / \
				//  ub  nb		--->		brr  			--->		   ub  pr
				//   \						/
				//    brr				   ub
				brother.leftRotate(tree)
				parent.rightRotate(tree)
				parent.setBlack()
			} else {
				//     pr
				//    / \
				//  ub  nb
				parent.setBlack()
				brother.setRed()
			}
		}
	} else { //n is left son
		if brother.right != nil { //     pr								    ub
			//if brother.left == nil {							//    / \								   / \
			//brother.leftRotate(tree)							//  nb  ub								  pr  brr
			//	brother.setRed()			//可以不调色			//     	  \
			//	brother.left.setBlack()	//可以不调色				//        brr
			//	brother.right.setBlack()	//可以不调色
			//}else {											//     pr								    ub
			//	brother.setRed()								//    / \								   / \
			//	brother.left.setBlack()							//  nb  ub								 pr  brr
			//	brother.right.setBlack()						//     / \								   \
			//}													//    blr brr							   blr

			//brother.left, brother.right is not nil
			if brother.left != nil { //     pr								    ub
				parent.leftRotate(tree)  //    / \			p ll				   / \
				brother.setRed()         //  nb  ub			--->				 pr  brr
				brother.left.setBlack()  //     / \								   \
				brother.right.setBlack() //    blr brr							   blr
			}
		} else { //brother.right is nil
			if brother.left != nil {
				//     pr					   pr					       	blr
				//    / \		u rr		    \			p ll		    / \
				//  nb  ub		--->		    blr	    	--->		   pr  ub
				//      / 					      \
				//    blr  				          ub
				brother.rightRotate(tree)
				parent.leftRotate(tree)
				parent.setBlack()
			} else {
				//     pr
				//    / \
				//  nb  ub
				parent.setBlack()
				brother.setRed()
			}
		}
	}
}

/* n节点可能为左/右子节点，这里只展示右子节点（左子节点左右对称）
 * 			      |
 * 			     pb
 * 			    / \
 * 			   br  nb
 * 			  / \
 * 			blb  brb
 * 		   / \   / \
 * 	            	  (r/nil)
 */
func (n *node) complexRemoveParentIsBlackAndbrotherIsRed(tree *RBTree) {
	parent, brother, isRightSon := n.parent, n.getBrother(), n.isRightSon()
	n.clear()
	if isRightSon {
		if brother.left.left != nil || brother.left.right != nil {
			//		      |
			//		     pb    						br					br
			//		    / \      p rr              / \      p rr	   / \
			//		   br  nb    --->           blb   pb    --->	blb   brb
			//		  / \                       /\    /				/\    / \
			//		blb  brb   						 brb				 brlr pb
			//	   / \   / \						/ \						  /
			//		   brlr                       brlr				  		(brrr/nil)
			if brother.right.left != nil {
				parent.rightRotate(tree)
				parent.rightRotate(tree)
				brother.setBlack()
				brother.right.setRed()
				brother.right.left.setBlack()
			} else if brother.right.right != nil {
				//		      |
				//		     pb    						br						br								br
				//		    / \      p rr              / \      br ll		   / \    				p rr	   / \
				//		   br  nb    --->           blb   pb    --->		blb   pb  				--->	blb   brlr
				//		  / \                       /\    /					/\    /							/\    / \
				//		blb  brb   						 brb					 brlr							 brb pb
				//	   / \    \							  \						 /
				//		   	  brrr                        brlr				    brb
				parent.rightRotate(tree)
				parent.left.leftRotate(tree)
				parent.rightRotate(tree)
				brother.setBlack()
			} else {
				//		      |
				//		     pb    						br
				//		    / \      p rr              / \
				//		   br  nb    --->           blb   pb
				//		  / \                       /\    /
				//		blb  brb   						 brb
				//	   / \
				parent.rightRotate(tree)
				brother.setBlack()
				parent.left.setRed()
			}
		} else {
			if brother.right.left != nil || brother.right.right != nil {
				//		      |
				//		     pb    						br
				//		    / \      p rr              / \
				//		   br  nb    --->           blb   pb
				//		  / \                             /
				//		blb  brb   						 brb
				//	   		/ \    					     /\
				parent.rightRotate(tree)     //							br
				if parent.left.left != nil { //		p ll			   / \
					parent.leftRotate(tree)       //		--->			blb   brb
					brother.right.setRed()        //						      / \
					brother.right.left.setBlack() //							brlr pb
					brother.setBlack()
				} else { //brl is nil, brr is red
					//		      |
					//		     pb    						br						br							br
					//		    / \      p rr              / \		br ll	       / \		p rr		       / \
					//		   br  nb    --->           blb   pb     --->	    blb   pb    --->		    blb   brrr
					//		  / \                             /				          /					          /	 \
					//		blb  brb   						 brb					 brrr						 brb  pb
					//	   		 \    					       \				      /
					//	   		 brrr    					    brrr				  brb
					parent.left.leftRotate(tree)
					parent.rightRotate(tree)
					brother.setBlack()
				}
			} else {
				//		      |
				//		     pb    						br
				//		    / \      p rr              / \
				//		   br  nb    --->           blb   pb
				//		  / \                             /
				//		blb  brb   						 brb
				parent.rightRotate(tree)
				brother.setBlack()
				parent.left.setRed()
			}
		}
	} else { //n is left son
		if brother.right.left != nil || brother.right.right != nil {
			//		      |
			//		     pb    						br					br						br
			//		    / \      p ll              / \      bl rr	   /   \       p ll		   /   \
			//		   nb  br    --->            pb   brb    --->	  pb    brb    --->		 bllr    brb
			//		  	  / \                     \    /\			   \     / \			 /  \     / \
			//			blb  brb   				  blb				 	bllr 				pb 	blb
			//	   		/ \   / \				  /	\				  	 \	 				  	 \
			//		   bllr                      bllr				    blb
			//															  \
			//
			if brother.left.left != nil {
				parent.leftRotate(tree)
				parent.right.rightRotate(tree)
				parent.leftRotate(tree)
				brother.setBlack()
			} else if brother.left.right != nil {
				//		      |
				//		     pb    						br					br
				//		    / \      p ll              / \       p ll	   /   \
				//		   nb  br    --->            pb   brb    --->	  blb    brb
				//		  	  / \                     \    /\			  / \     / \
				//			blb  brb   				  blb				 pb	blrr
				//	   		 \   / \				  	\
				//		    blrr                      blrr
				//
				//
				parent.leftRotate(tree)
				parent.leftRotate(tree)
				brother.setBlack()
				parent.setRed()
			} else {
				//		      |
				//		     pb    						br
				//		    / \      p ll              / \
				//		   nb  br    --->            pb   brb
				//		  	  / \                     \   / \
				//			blb  brb   				  blb
				//	   			/ \
				parent.leftRotate(tree)
				brother.setBlack()
				parent.right.setRed()
			}
		} else { // brother.right has no son
			if brother.right.left != nil || brother.right.right != nil {
				//		      |
				//		     pb    						br
				//		    / \      p ll              / \
				//		   nb  br    --->           pb   brb
				//		  	  / \                     \
				//			blb  brb   				  blb
				//	   		/ \    					  /\
				parent.leftRotate(tree)        //							br
				if parent.right.right != nil { //		p ll			   / \
					parent.leftRotate(tree)       //		--->			blb   brb
					brother.left.setRed()         //				        / \
					brother.left.right.setBlack() //					  pb  blrr
					brother.setBlack()            //						\
				} else { //blr is nil, bll is red
					//		      |
					//		     pb    						br						br							br
					//		    / \      p ll              / \		bl rr	       / \		p ll		       / \
					//		   nb  br    --->           pb   brb     --->	    pb   brb    --->		    bllr  brb
					//		  	  / \                     \        				 \      					/ \
					//			blb  brb   				  blb					 bllr						pb blb
					//	   		/    					  /     				    \
					//	      bllr	    				bllr	    				 blb
					parent.right.rightRotate(tree)
					parent.leftRotate(tree)
					brother.setBlack()
				}
			} else {
				//		      |
				//		     pb    						br
				//		    / \      p ll              / \
				//		   nb  br    --->            pb   brb
				//		  	  / \                     \
				//			blb  brb   				  blb
				parent.leftRotate(tree)
				brother.setBlack()
				parent.right.setRed()
			}
		}
	}
}

/* n节点可能为左/右子节点，这里只展示右子节点（左子节点左右对称）
 * 	        |
 * 	        pb
 * 	       / \
 * 	     ub  nb
 * 	     / \
 * 	     	(r/nil)
 */
func (n *node) complexRemoveParentAndbrotherIsBlack(tree *RBTree) {
	parent, brother, isRightSon := n.parent, n.getBrother(), n.isRightSon()
	n.clear()
	if isRightSon {
		if brother.left != nil {
			//		 pb					ub
			//		/ \      p rr	    / \
			//	   ub  nb    --->      blr pb
			//	  / \                      /
			//	 blr (brr/nil)           (brr/nil)
			parent.rightRotate(tree)
			brother.left.setBlack()
		} else if brother.right != nil {
			//		 pb					pb							    brr
			//		/ \      u ll		/		p rr				    / \
			//	   ub  nb    --->	  brr       --->				  ub   pb
			//	    \            	   /
			//	   brr           	  ub
			brother.leftRotate(tree)
			parent.rightRotate(tree)
			brother.parent.setBlack()
		} else {
			//		 pb
			//		/ \
			//	   ub  nb
			//todo 待完成
		}
	} else { // parent is black, brother is black, n is left son
		if brother.right != nil {
			//		 pb					ub
			//		/ \      p ll	    / \
			//	   nb  ub    --->      pb brr
			//	      / \               \
			//(blr/nil) brr           (blr/nil)
			parent.leftRotate(tree)
			brother.right.setBlack()
		} else if brother.left != nil {
			//		 pb					pb							    blr
			//		/ \      u rr		 \	       p ll				    / \
			//	   nb  ub    --->	      blr      --->				  pb   ub
			//	       /            	   \
			//	      blr           	    ub
			brother.rightRotate(tree)
			parent.leftRotate(tree)
			brother.parent.setBlack()
		} else {
			//		 pb
			//		/ \
			//	   ub  nb
			//todo 待完成
		}
	}
}

func (tree *RBTree) Clear() {
	tree.root.destroy()
}

func (n *node) cut() {
	if n == nil || n.parent == nil {
		return
	}
	if n.parent.left == n {
		n.parent.left = nil
	} else {
		n.parent.right = nil
	}
}

//删除 以n节点为跟节点的书，且删除n的父节点指向n
func (n *node) clear() {
	if n == nil || n.parent == nil {
		n = nil
		return
	}
	if n.parent.left == n {
		n.parent.left = nil
	} else {
		n.parent.right = nil
	}
	n.destroy()
}
func (n *node) destroy() {
	if n == nil {
		return
	}
	if n.left != nil {
		n.left.destroy()
	}
	if n.right != nil {
		n.right.destroy()
	}
	n = nil
}
