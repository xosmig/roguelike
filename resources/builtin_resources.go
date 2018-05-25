package resources

var BuiltinLoader Loader = NewInMemoryLoader()

var exampleMap = "6 11\n" +
	"###########\n" +
	"#@.#*.##.O#\n" +
	"#.#..###.##\n" +
	"#..#..##..#\n" +
	"#.........#\n" +
	"###########\n" +
	""

func init() {
	BuiltinLoader.(*InMemoryLoader).AddString("builtin/maps/example", exampleMap)
}
