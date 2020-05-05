
func killProcess(pid []int, ppid []int, kill int) []int {
	ans := []int{}

	ppidMap := getPPIDMap(pid, ppid)
	dfs(ppidMap, kill, &ans)

	return ans
}

func dfs(ppidMap map[int][]int, kill int, killedPids *[]int) {
	*killedPids = append(*killedPids, kill)

	for _, pid := range ppidMap[kill] {
		dfs(ppidMap, pid, killedPids)
	}
}

func getPPIDMap(pids, ppids []int) map[int][]int {
	m := map[int][]int{}
	for i, ppid := range ppids {
		m[ppid] = append(m[ppid], pids[i])
	}
	return m
}