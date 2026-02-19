#class ShoreSpec(object):
#
#    def __init__(self,X,Y):
#        self.x = X
#        self.y = Y

class Solution(object):

#    def isIsland(self,grid,row,col):
#
#        # An island is made of land
#        if grid[row][col] == '1':
#            ret=True
#        else:
#            return False
#
#        #North
#        if (( (row > 0) && (grid[row-1][col]==)

#    def isConnected(self, grid, SS, R, C):
#        adsf

    # bounds check
    def outOfScope(self,grid,x,y):

        # too far west or east
        if (x == -1) or (x > len(grid[0]) - 1):
            return True
        # too far north or south
        if (y == -1) or (y > len(grid) - 1):
            return True
        # already processed
        if self.visited[(y,x)]:
            return True

        # passed all bounds checks, return false
        return False

    # DFS using recursion
    def walkIsland(self, grid, x, y):
        if self.outOfScope(grid,x,y):
            return

        # in bounds, mark node as visited so we can skip in walk
        self.visited[(y,x)] = True

        # not land, so back out of recursion
        if grid[y][x] == '0':
            return

        # walk N, S, E, W recursively
        self.walkIsland(grid, x, y-1) # North
        self.walkIsland(grid, x, y+1) # South
        self.walkIsland(grid, x+1, y) # East
        self.walkIsland(grid, x-1, y) # West

    def numIslands(self, grid):
        """
        :type grid: List[List[str]]
        :rtype: int
        """

        # return value
        islands=0

        # initialize visited hash map
        self.visited={}
        for r in range(len(grid)):
            for c in range(len(grid[r])):
                self.visited[(r,c)] = False

#        foundLand=False
        for r in range(len(grid)):
            for c in range(len(grid[r])):

#                # mark node as visited so we can skip in walk
#                visited[(r,c)] = True

                # if we've seend this node by walking the island already, skip
                if self.visited[(r,c)]:
                    continue

                if grid[r][c] == '1':
                    islands += 1
                    self.walkIsland(grid, c, r)

        return islands
                #search until we find our first land
#                if !foundLand:
#                    if grid[r][c] == '0':
#                        continue
                    #X=c
                    #W=1 # min 1 cell wide
                    #islands = islands+1
#                    shores.apend(ShoreSpec(c,r))
                # we found land, now find the width
                #else:
                #    if grid[r][c] == '1':
                #        W = W+1
                #        continue