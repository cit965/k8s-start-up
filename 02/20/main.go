// 实现 GetPodsInfo 方法
func (p *PodCache) GetPodsInfo() ([]interface{}, error) {
	p.podLock.Lock()
	defer p.podLock.Unlock()

	if len(p.podInfo) == 0 {
		return []interface{}{}, fmt.Errorf("No pod information available")
	}

	podsInfo := make([]interface{}, 0, len(p.podInfo))
	for _, info := range p.podInfo {
		podsInfo = append(podsInfo, info)
	}

	return podsInfo, nil
}
