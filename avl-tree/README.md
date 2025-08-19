# AVL Tree 

This is a tiny implementation of an AVL Tree made in golang. 
This AVL Tree for now is append only. Another necessary implementations for study 
reasons will be made in the future.

## Supported operations

1. Insertion
2. Ordered (DFS) traversal 

## Principles 

The overall principle of an AVL Tree is to provide an optimized balanced insertion/retrieval of data.
A general Binary Search Tree normally is not balanced properly. So, when you insert sequentially ordered 
data, you get a linked list:

1 -> 2 -> 3 -> 4 -> 5 -> 6 

In self balanced trees, when you insert the data, the tree will find means to self balance itself 
to make it's overall structure consistent, providing a better amortized performance for reading,
at cost of a slight worse performance for writing:

       -> 6
  -> 5 
4
       -> 3
  -> 2 
       -> 1

Normally, for indexing and search data structures, AVL and Red-Black trees are used, more commonly 
Red-Black trees, because they have less strict balance rules.

AVL Trees define they balance rules based on the height of it's subtrees.
A balance factor is defined as the difference between the height of the left subtree and the right subtree, 
defined as:

BF = LSH - RSH 

When the balance factor goes bellow -1 or higher than 1, then the tree is disbalanced, and need to be 
balanced back by doing rotations of the nodes.

Rotations are simple operations with low overhead at all. 
The disbalance occurs generally when nodes are inserted/deleted.
Then rotations only are necessary on write operations. Read operations don't chnage the tree.
