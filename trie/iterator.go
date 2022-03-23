}

func newNodeIterator(trie *Trie, start []byte) NodeIterator {
	if trie.Hash() == emptyState {
		return new(nodeIterator)
	if trie.Hash() == emptyRoot {
		return &nodeIterator{
			trie: trie,
			err:  errIteratorEnd,
		}
	}
	it := &nodeIterator{trie: trie}
	it.err = it.seek(start)
@@ -425,7 +428,7 @@ func findChild(n *fullNode, index int, path []byte, ancestor common.Hash) (node,
func (it *nodeIterator) nextChild(parent *nodeIteratorState, ancestor common.Hash) (*nodeIteratorState, []byte, bool) {
	switch node := parent.node.(type) {
	case *fullNode:
		//Full node, move to the first non-nil child.
		// Full node, move to the first non-nil child.
		if child, state, path, index := findChild(node, parent.index+1, it.path, ancestor); child != nil {
			parent.index = index - 1
			return state, path, true
