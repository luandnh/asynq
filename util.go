package asynq

func prepareQueues(config map[string]int, priority bool) (queues map[string]int, orderedQueues []string) {
	queues = normalizeQueues(config)
	orderedQueues = []string(nil)
	if priority {
		orderedQueues = sortByPriority(queues)
	}
	return queues, orderedQueues
}
