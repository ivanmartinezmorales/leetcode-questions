def hasSingleCycle(array):
    visited = 0
    current = 0
    while visited < len(array):
        if visited > 0 and current == 0:
            return False
        visited += 1	
        current = getNextIdx(current, array)
    return current == 0

def getNextIdx(current, array):
    # We want to see how many indcies we have to jump through	
    jump = array[current]
    nextIdx = (current + jump) % len(array)
    # Claify if we can negative integers, can we have decimasls?
     return nextIdx if nextIdx >= 0 else nextIdx + len(array)

