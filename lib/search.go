package lib

import "container/list"

// AStar uses the A* algorithm to find the minimum number of steps from the initial state(s)
// to a state for which the done function returns true. The next function should fill the supplied
// map with all states reachable in a single step from the state passed to it along with the
// corresponding cost, and the estimate function should return a lower bound on the remaining
// cost to go from the supplied state to a target state.
// See https://www.redblobgames.com/pathfinding/a-star/introduction.html.
func AStar(
	initial []interface{},
	done func(interface{}) bool,
	next func(interface{}, map[interface{}]int),
	estimate func(interface{}) int) int {
	// TODO: Add some way to track the path if needed.
	frontier := NewHeap(func(a, b interface{}) bool { return a.(asNode).pri < b.(asNode).pri })
	costs := make(map[interface{}]int)
	for _, init := range initial {
		frontier.Insert(asNode{init, 0})
		costs[init] = 0
	}

	for frontier.Len() != 0 {
		cur := frontier.Pop().(asNode).state
		cost := costs[cur]

		// Check if we're done.
		if done(cur) {
			return cost
		}

		nmap := make(map[interface{}]int)
		next(cur, nmap)
		for ns, nc := range nmap {
			newCost := cost + nc
			if oldCost, ok := costs[ns]; !ok || newCost < oldCost {
				costs[ns] = newCost
				pri := newCost + estimate(ns)
				frontier.Insert(asNode{ns, pri})
			}
		}
	}
	panic("No paths found")
}

type asNode struct {
	state interface{}
	pri   int // lower is better
}

// BFS performs a breadth-first search to discover paths to states reachable from initial.
// If opts is non-nil, it is used to configure the search.
// The returned steps map contains the minimum number of steps to each state.
// The returned from map contains the state (value) preceding each destination state (key).
func BFS(initial []interface{}, next func(interface{}, map[interface{}]struct{}), opts *BFSOptions) (
	steps map[interface{}]int, from map[interface{}]interface{}) {
	queue := list.New() // next states to check
	steps = make(map[interface{}]int)
	from = make(map[interface{}]interface{})
	for _, s := range initial {
		queue.PushBack(s)
		steps[s] = 0
		from[s] = s
	}

	var remain map[interface{}]struct{}
	if opts != nil && len(opts.AllDests) > 0 {
		remain = make(map[interface{}]struct{})
		for _, d := range opts.AllDests {
			remain[d] = struct{}{}
		}
	}

Loop:
	for queue.Len() > 0 {
		cur := queue.Remove(queue.Front())
		cost := steps[cur] + 1

		// Early exit if we've exceeded the maximum number of steps.
		if opts != nil && opts.MaxSteps > 0 && cost > opts.MaxSteps {
			break Loop
		}

		nmap := make(map[interface{}]struct{})
		next(cur, nmap)
		for n := range nmap {
			// Skip already-visited states.
			if _, ok := from[n]; ok {
				continue
			}

			queue.PushBack(n)
			from[n] = cur
			steps[n] = cost

			// Early exit if we've reached one of the "any" destinations.
			if opts != nil && MapHasKey(opts.AnyDests, n) {
				break Loop
			}

			// Early exit if we've reached all required destinations.
			if remain != nil {
				delete(remain, n)
				if len(remain) == 0 {
					break Loop
				}
			}
		}
	}

	return steps, from
}

// BFSOptions specifies optional configuration for BFS.
type BFSOptions struct {
	// AllDests contains states that must all be reached before exiting.
	AllDests []interface{}
	// AnyDests contains states of which just one must be reached before exiting.
	AnyDests map[interface{}]struct{}
	// MaxSteps contains the maximum number of steps before exiting.
	MaxSteps int
}
