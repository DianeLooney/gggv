package gql

func (q *Schema) Sources() []Source {
	var out []Source
	for _, src := range q.d.Scene.Sources {
		out = append(out, Source{src})
	}
	return out
}
