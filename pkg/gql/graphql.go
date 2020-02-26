package gql

import "github.com/dianelooney/gggv/pkg/daemon"

func NewSchema(d *daemon.D) *Schema {
	return &Schema{d: d}
}

const S = `
schema {
	query: Query
}
type Query {
	sources(): [Source!]!
}
type Source {
	id: ID!
	programID: ID
	sourceIDs: [ID!]!
	uniforms: [Uniform!]
}
type Uniform {
	name: String!
	value1f: Float
}
`
