package runtimejs

// Prelude is javascript strings which setups the go runtime primitives included
// in js code compiled by gopherjs.
const Prelude = includePlain + preludePlain + numericPlain + typesPlain +
	goroutinesPlain + jsmappingPlain
