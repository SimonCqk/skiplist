## What is SkipList

`SkipList` is a randomized data structure, whose performance can be
compared with **Balanced Binary Search Tree**.

The idea is simple, we skip some nodes to boost searching by creating
multiply layers,

The time complexity of skip lists can be reduced further by adding
more layers. In fact, the time complexity of search, insert and delete
can become O(Log n) in average case with O(n) extra space.

HUNT FOR MORE ON WIKI: [SkipList](https://en.wikipedia.org/wiki/Skip_list)