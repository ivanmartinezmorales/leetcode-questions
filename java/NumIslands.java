public class NumIslands {
    public int numIslands(char[][] grid) {
        // in order for us to count the number of islands
        // we need to have some error checking, immediately bounce
        if (grid == null || grid.length == 0)
            return 0;

        int numIslands = 0;
        for (int i = 0; i < grid.length; i++) {
            for (int j = 0; j < grid[i].length; j++) {
                if (grid[i][j] == '1') {
                    // this is the meat of the problem
                    // we know that we've visited an island
                    // sometimes its a big one, sometimes its a small one
                    // do the dfs here
                    numIslands += dfs(grid, i, j);
                }
            }
        }
        return numIslands;
    }

    public int dfs(char[][] grid, int i, int j) {
        // first thing that we want to do is put a base case
        // check that indicies that are valid
        if (i < 0 || i >= grid.length || j < 0 || j >= grid[i].length || grid[i][j] == '0')
            return 0;

        // Run DFS through entire matrix
        grid[i][j] = '0';
        dfs(grid, i + 1, j);
        dfs(grid, i - 1, j);
        dfs(grid, i, j + 1);
        dfs(grid, i, j - 1);

        return 1;
    }
}
