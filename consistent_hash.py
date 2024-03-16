import hashlib
import logging
# config basic logging info inclue line number, module, timestamp
logging.basicConfig(level=logging.INFO, format="%(asctime)s:  %(module)s:%(lineno)d: %(message)s")


class ConsistentHashing:
    def __init__(self, nodes=None, replicas=3):
        self.replicas = replicas
        self.ring = {}
        self.sorted_keys = []

        if nodes:
            for node in nodes:
                self.add_node(node)

    def add_node(self, node):
        for i in range(self.replicas):
            replica_key = self.get_replica_key(node, i)
            self.ring[replica_key] = node
            self.sorted_keys.append(replica_key)

        self.sorted_keys.sort()

    def remove_node(self, node):
        for i in range(self.replicas):
            replica_key = self.get_replica_key(node, i)
            del self.ring[replica_key]
            self.sorted_keys.remove(replica_key)

    def get_node(self, key):
        if not self.ring:
            return None

        hash_key = self.hash_key(key)
        for replica_key in self.sorted_keys:
            if hash_key <= replica_key:
                return self.ring[replica_key]

        return self.ring[self.sorted_keys[0]]

    def get_replica_key(self, node, replica_index):
        return self.hash_key(f"{node}-{replica_index}")

    def hash_key(self, key):
        return int(hashlib.md5(key.encode()).hexdigest(), 16)


"""
    Rewrite ConsistentHashing class, where I can choose ring size and number of replicas.
"""

if __name__ == '__main__':
    # Example usage
    nodes = ["Node1", "Node2", "Node3"]
    ch = ConsistentHashing(nodes)

    logging.info(ch.get_node("Key1"))  # Output: Node1
    logging.info(ch.get_node("Key2"))  # Output: Node2
    logging.info(ch.get_node("Key3"))  # Output: Node3

    ch.remove_node("Node2")
    logging.info("Node2 removed")
    logging.info(ch.get_node("Key1"))  # Output: Node1
    logging.info(ch.get_node("Key2"))  # Output: Node3
    logging.info(ch.get_node("Key3"))  # Output: Node1

    ch.add_node("Node4")
    logging.info("Node4 added")
    ch.add_node("Node5")
    logging.info("Node5 added")
    logging.info(ch.get_node("Key1"))  # Output: Node1
    logging.info(ch.get_node("Key2"))  # Output: Node3
    logging.info(ch.get_node("Key3"))  # Output: Node1
    logging.info(ch.get_node("Key4"))  # Output: Node1
    logging.info(ch.get_node("Key5"))  # Output: Node3
    logging.info(ch.get_node("Key6"))  # Output: Node1

    logging.info(ch.ring)
    logging.info(ch.sorted_keys)