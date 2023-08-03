# Finding and Counting Number of Rectangles in a 2D Array

![Golang Logo](https://golang.org/doc/gopher/frontpage.png)

This Go package provides a function `CountRectangles` to find and count the number of rectangles formed by adjacent 1s in a 2D grid represented by a 2D slice of integers.

## Problem Statement

Given a 2D array of 0s and 1s, we want to determine the number of rectangles that can be formed by adjacent 1s. Two 1s are considered adjacent if they are horizontally or vertically adjacent (not diagonally). A rectangle is formed by a group of adjacent 1s.

Example:
Consider the following 2D array:

```go
arr := [][]int{
    {1, 0, 0, 0, 0, 0, 0},
    {0, 0, 0, 0, 0, 0, 0},
    {1, 0, 0, 1, 1, 1, 0},
    {0, 1, 0, 1, 1, 1, 0},
    {0, 1, 0, 0, 0, 0, 0},
    {0, 1, 0, 1, 1, 0, 0},
    {0, 0, 0, 1, 1, 0, 0},
    {0, 0, 0, 0, 0, 0, 1},
}
```

There are four rectangles in this 2D array, and the `CountRectangles` function will return 6.

## Step-by-Step Explanation of `CountRectangles` Function

The `CountRectangles` function works as follows:

1. It takes a 2D slice of integers, `rectangles`, as input.

2. It initializes some variables, `rows` and `cols`, to store the dimensions of the grid (number of rows and columns).

3. It sets `rectangleCount` to 0, which will be used to count the number of rectangles in the grid.

4. The function defines a nested DFS function named `dfs`. This DFS function is used to explore each rectangle starting from a specific cell (row, col).

5. The DFS function takes two parameters, `row` and `col`, representing the starting cell for exploration.

6. Inside the DFS function, it checks if the current cell (row, col) is out of bounds (i.e., outside the grid) or if the cell's value is 0 (i.e., it's not part of a rectangle). If either of these conditions is true, the function returns, indicating that the current cell cannot be part of a rectangle.

7. If the current cell's value is 1 (i.e., it's part of a rectangle), the function marks it as visited by setting its value to 0.

8. The DFS function then recursively calls itself for each adjacent cell (up, down, left, and right) from the current cell (row, col). This recursive exploration continues until it reaches the boundary or encounters a 0, effectively covering the entire rectangle that contains the initial cell.

9. The main function uses a nested loop to iterate over each cell in the grid. When it finds a cell with a value of 1, it calls the DFS function with the coordinates of that cell. This DFS exploration covers all cells forming a rectangle, and each time it encounters a new 1 (part of a new rectangle), it increments the `rectangleCount`.

10. After iterating over all cells, the function returns the final `rectangleCount`, which represents the number of rectangles in the grid.

In the provided example grid, there are four rectangles, so the `CountRectangles` function will return 6.

---

**Note:** This is a sample implementation for educational purposes. For production use, it is recommended to handle edge cases and error conditions appropriately.

Please feel free to modify and use this [Readme](README.md) file or navigate back to this [Readme](../../README.md) file to see the overview Golang Learning repository as needed.
