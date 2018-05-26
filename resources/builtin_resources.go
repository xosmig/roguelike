package resources

// BuiltinLoader is used to load builtin resources
var BuiltinLoader Loader = NewInMemoryLoader()

var exampleMap = "6 11\n" +
	"###########\n" +
	"#@.#.$##.O#\n" +
	"#.#..###.##\n" +
	"#..#..##..#\n" +
	"#.....z...#\n" +
	"###########\n" +
	""

func init() {
	BuiltinLoader.(*InMemoryLoader).AddString("builtin/maps/example", exampleMap)
}
