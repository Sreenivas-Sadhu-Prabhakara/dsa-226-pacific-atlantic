# Pacific Atlantic Water Flow

> NeetCode 150 · Graphs · Blind 75 — trace the water backwards: climb uphill from each shoreline, then intersect the two reachable sets. Go, stdlib only, tested.

## Problem

Heights on a grid; water flows to equal-or-lower neighbors. The Pacific laps the top/left edges, the Atlantic the bottom/right. Which cells can drain to **both** oceans?

## Approach

Simulating flow from every cell re-explores the grid per cell. The inversion that collapses it: instead of asking "where can this cell's water go?", ask "which cells can reach this ocean?" — start at each ocean's shoreline and walk **uphill** (reverse flow: neighbor height ≥ current), marking reachability. Two DFS sweeps produce two boolean grids; the answer is their intersection. Each sweep touches every cell at most once because the reach-grid doubles as the visited set. This direction flip — compute the inverse relation from the targets — is the same move as multi-source BFS, in DFS clothing.

**Complexity:** O(R·C) time, O(R·C) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

Part of a 275-repo DSA series in Go — Striver's SDE Sheet (dsa-001–191) plus the NeetCode 150 / Blind 75 additions (dsa-192–275), stdlib only, every solution tested. Browse: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
