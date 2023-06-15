[![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.17-blue)](https://img.shields.io/badge/Go-%3E%3D%201.20-blue)

# Lexorank implementation in Go

Lexorank is an algorithm introduced by Atlassian company to manage "rank/order system" efficiently.

For example, you can use integer numbers to store the rank of items in DB like below:

```azure
Item    |  Display order
--------|-----------------
item-1  |  1
item-2  |  2  
item-3  |  3
item-4  |  4  
item-5  |  5
item-6  |  6  
item-7  |  7
item-8  |  8  
item-9  |  9
item-10 |  10  
```

If you want to insert a new item with rank 5, you need to update the ranks of (item-5 to item-10) which is inefficiently.

Lexorank uses string to store the rank values.