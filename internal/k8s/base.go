package k8s

type base struct {
	fieldSelector string
	labelSelector string
}

// SetFieldSelector refines query results via selector.
func (b *base) SetFieldSelector(s string) {
	b.fieldSelector = s
}

// SetLabelSelector refines query results via labels.
func (b *base) SetLabelSelector(s string) {
	b.labelSelector = s
}

// GetFieldSelector returns field selector.
func (b *base) GetFieldSelector() string {
	return b.fieldSelector
}

// GetLabelSelector returns label selector.
func (b *base) GetLabelSelector() string {
	return b.labelSelector
}

func (b *base) HasSelectors() bool {
	return b.labelSelector != "" || b.fieldSelector != ""
}

func (b *base) Kill(ns, n string) error {
	return nil
}
