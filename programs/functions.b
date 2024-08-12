fn main() {
    fn hello() string {
        return "Hello, World";
    }

    print(hello());

    fn helloName(name string) string {
        return "Hello, " + name;
    }

    fn add(x int, y int) int {
        return x + y;
    }

    add(2,3);


    let nameFunc = fn(name) string {
        return "Hello, " + name;
    }

    nameFunc("pspiagicw")

    fn arithmetic(operation fn(int,int), x int , y int) int {
        return operation(x, y);
    }

    arithmetic(add, 3, 4);

    arithmetic(fn(x int, y int) { return x + y; }, 3, 4);
}
