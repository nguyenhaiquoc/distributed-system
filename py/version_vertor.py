class VersionVector:
    def __init__(self, process_id):
        self.process_id = process_id
        self.vector = {process_id: 0}

    def increment(self):
        self.vector[self.process_id] += 1

    def update(self, other_vector):
        for process_id, version in other_vector.items():
            if process_id not in self.vector or version > self.vector[process_id]:
                self.vector[process_id] = version

    def compare(self, other_vector):
        for process_id, version in self.vector.items():
            if process_id not in other_vector or version > other_vector[process_id]:
                return 1
            elif version < other_vector[process_id]:
                return -1
        return 0